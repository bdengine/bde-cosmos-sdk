package middleware

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/proto"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/tmhash"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

type runMsgsTxHandler struct {
	legacyRouter     sdk.Router        // router for redirecting legacy Msgs
	msgServiceRouter *MsgServiceRouter // router for redirecting Msg service messages
}

func NewRunMsgsTxHandler(msr *MsgServiceRouter, legacyRouter sdk.Router) tx.TxHandler {
	return runMsgsTxHandler{
		legacyRouter:     legacyRouter,
		msgServiceRouter: msr,
	}
}

var _ tx.TxHandler = runMsgsTxHandler{}

func (txh runMsgsTxHandler) CheckTx(ctx context.Context, tx sdk.Tx, req abci.RequestCheckTx) (res abci.ResponseCheckTx, err error) {
	// Don't run Msgs during CheckTx.
	return abci.ResponseCheckTx{}, nil
}

func (txh runMsgsTxHandler) DeliverTx(ctx context.Context, tx sdk.Tx, req abci.RequestDeliverTx) (res abci.ResponseDeliverTx, err error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Create a new Context based off of the existing Context with a MultiStore branch
	// in case message processing fails. At this point, the MultiStore
	// is a branch of a branch.
	runMsgCtx, msCache := cacheTxContext(sdkCtx, req.Tx)

	// Attempt to execute all messages and only update state if all messages pass
	// and we're in DeliverTx. Note, runMsgs will never return a reference to a
	// Result if any single message fails or does not have a registered Handler.
	msgLogs := make(sdk.ABCIMessageLogs, 0, len(tx.GetMsgs()))
	events := sdkCtx.EventManager().Events()
	txMsgData := &sdk.TxMsgData{
		Data: make([]*sdk.MsgData, 0, len(tx.GetMsgs())),
	}

	// NOTE: GasWanted is determined by the AnteHandler and GasUsed by the GasMeter.
	for i, msg := range tx.GetMsgs() {
		var (
			msgResult    *sdk.Result
			eventMsgName string // name to use as value in event `message.action`
			err          error
		)

		if handler := txh.msgServiceRouter.Handler(msg); handler != nil {
			// ADR 031 request type routing
			msgResult, err = handler(runMsgCtx, msg)
			eventMsgName = sdk.MsgTypeURL(msg)
		} else if legacyMsg, ok := msg.(legacytx.LegacyMsg); ok {
			// legacy sdk.Msg routing
			// Assuming that the app developer has migrated all their Msgs to
			// proto messages and has registered all `Msg services`, then this
			// path should never be called, because all those Msgs should be
			// registered within the `MsgServiceRouter` already.
			msgRoute := legacyMsg.Route()
			eventMsgName = legacyMsg.Type()
			handler := txh.legacyRouter.Route(sdkCtx, msgRoute)
			if handler == nil {
				return abci.ResponseDeliverTx{}, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized message route: %s; message index: %d", msgRoute, i)
			}

			msgResult, err = handler(sdkCtx, msg)
		} else {
			return abci.ResponseDeliverTx{}, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "can't route message %+v", msg)
		}

		if err != nil {
			return abci.ResponseDeliverTx{}, sdkerrors.Wrapf(err, "failed to execute message; message index: %d", i)
		}

		msgEvents := sdk.Events{
			sdk.NewEvent(sdk.EventTypeMessage, sdk.NewAttribute(sdk.AttributeKeyAction, eventMsgName)),
		}
		msgEvents = msgEvents.AppendEvents(msgResult.GetEvents())

		// append message events, data and logs
		//
		// Note: Each message result's data must be length-prefixed in order to
		// separate each result.
		events = events.AppendEvents(msgEvents)

		txMsgData.Data = append(txMsgData.Data, &sdk.MsgData{MsgType: sdk.MsgTypeURL(msg), Data: msgResult.Data})
		msgLogs = append(msgLogs, sdk.NewABCIMessageLog(uint32(i), msgResult.Log, msgEvents))
	}

	msCache.Write()
	data, err := proto.Marshal(txMsgData)
	if err != nil {
		return abci.ResponseDeliverTx{}, sdkerrors.Wrap(err, "failed to marshal tx data")
	}

	return abci.ResponseDeliverTx{
		Log:    msgLogs.String(),
		Data:   data,
		Events: events.ToABCIEvents(),
	}, nil
}

// cacheTxContext returns a new context based off of the provided context with
// a branched multi-store.
func cacheTxContext(sdkCtx sdk.Context, txBytes []byte) (sdk.Context, sdk.CacheMultiStore) {
	ms := sdkCtx.MultiStore()
	// TODO: https://github.com/cosmos/cosmos-sdk/issues/2824
	msCache := ms.CacheMultiStore()
	if msCache.TracingEnabled() {
		msCache = msCache.SetTracingContext(
			sdk.TraceContext(
				map[string]interface{}{
					"txHash": fmt.Sprintf("%X", tmhash.Sum(txBytes)),
				},
			),
		).(sdk.CacheMultiStore)
	}

	return sdkCtx.WithMultiStore(msCache), msCache
}
