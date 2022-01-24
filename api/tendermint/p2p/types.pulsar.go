package p2p

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_ProtocolVersion       protoreflect.MessageDescriptor
	fd_ProtocolVersion_p2p   protoreflect.FieldDescriptor
	fd_ProtocolVersion_block protoreflect.FieldDescriptor
	fd_ProtocolVersion_app   protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_p2p_types_proto_init()
	md_ProtocolVersion = File_tendermint_p2p_types_proto.Messages().ByName("ProtocolVersion")
	fd_ProtocolVersion_p2p = md_ProtocolVersion.Fields().ByName("p2p")
	fd_ProtocolVersion_block = md_ProtocolVersion.Fields().ByName("block")
	fd_ProtocolVersion_app = md_ProtocolVersion.Fields().ByName("app")
}

var _ protoreflect.Message = (*fastReflection_ProtocolVersion)(nil)

type fastReflection_ProtocolVersion ProtocolVersion

func (x *ProtocolVersion) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ProtocolVersion)(x)
}

func (x *ProtocolVersion) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_p2p_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ProtocolVersion_messageType fastReflection_ProtocolVersion_messageType
var _ protoreflect.MessageType = fastReflection_ProtocolVersion_messageType{}

type fastReflection_ProtocolVersion_messageType struct{}

func (x fastReflection_ProtocolVersion_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ProtocolVersion)(nil)
}
func (x fastReflection_ProtocolVersion_messageType) New() protoreflect.Message {
	return new(fastReflection_ProtocolVersion)
}
func (x fastReflection_ProtocolVersion_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ProtocolVersion
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ProtocolVersion) Descriptor() protoreflect.MessageDescriptor {
	return md_ProtocolVersion
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ProtocolVersion) Type() protoreflect.MessageType {
	return _fastReflection_ProtocolVersion_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ProtocolVersion) New() protoreflect.Message {
	return new(fastReflection_ProtocolVersion)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ProtocolVersion) Interface() protoreflect.ProtoMessage {
	return (*ProtocolVersion)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ProtocolVersion) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.P2P != uint64(0) {
		value := protoreflect.ValueOfUint64(x.P2P)
		if !f(fd_ProtocolVersion_p2p, value) {
			return
		}
	}
	if x.Block != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Block)
		if !f(fd_ProtocolVersion_block, value) {
			return
		}
	}
	if x.App != uint64(0) {
		value := protoreflect.ValueOfUint64(x.App)
		if !f(fd_ProtocolVersion_app, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_ProtocolVersion) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.p2p.ProtocolVersion.p2p":
		return x.P2P != uint64(0)
	case "tendermint.p2p.ProtocolVersion.block":
		return x.Block != uint64(0)
	case "tendermint.p2p.ProtocolVersion.app":
		return x.App != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.ProtocolVersion"))
		}
		panic(fmt.Errorf("message tendermint.p2p.ProtocolVersion does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ProtocolVersion) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.p2p.ProtocolVersion.p2p":
		x.P2P = uint64(0)
	case "tendermint.p2p.ProtocolVersion.block":
		x.Block = uint64(0)
	case "tendermint.p2p.ProtocolVersion.app":
		x.App = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.ProtocolVersion"))
		}
		panic(fmt.Errorf("message tendermint.p2p.ProtocolVersion does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ProtocolVersion) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.p2p.ProtocolVersion.p2p":
		value := x.P2P
		return protoreflect.ValueOfUint64(value)
	case "tendermint.p2p.ProtocolVersion.block":
		value := x.Block
		return protoreflect.ValueOfUint64(value)
	case "tendermint.p2p.ProtocolVersion.app":
		value := x.App
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.ProtocolVersion"))
		}
		panic(fmt.Errorf("message tendermint.p2p.ProtocolVersion does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ProtocolVersion) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.p2p.ProtocolVersion.p2p":
		x.P2P = value.Uint()
	case "tendermint.p2p.ProtocolVersion.block":
		x.Block = value.Uint()
	case "tendermint.p2p.ProtocolVersion.app":
		x.App = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.ProtocolVersion"))
		}
		panic(fmt.Errorf("message tendermint.p2p.ProtocolVersion does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ProtocolVersion) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.p2p.ProtocolVersion.p2p":
		panic(fmt.Errorf("field p2p of message tendermint.p2p.ProtocolVersion is not mutable"))
	case "tendermint.p2p.ProtocolVersion.block":
		panic(fmt.Errorf("field block of message tendermint.p2p.ProtocolVersion is not mutable"))
	case "tendermint.p2p.ProtocolVersion.app":
		panic(fmt.Errorf("field app of message tendermint.p2p.ProtocolVersion is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.ProtocolVersion"))
		}
		panic(fmt.Errorf("message tendermint.p2p.ProtocolVersion does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ProtocolVersion) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.p2p.ProtocolVersion.p2p":
		return protoreflect.ValueOfUint64(uint64(0))
	case "tendermint.p2p.ProtocolVersion.block":
		return protoreflect.ValueOfUint64(uint64(0))
	case "tendermint.p2p.ProtocolVersion.app":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.ProtocolVersion"))
		}
		panic(fmt.Errorf("message tendermint.p2p.ProtocolVersion does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ProtocolVersion) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.p2p.ProtocolVersion", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ProtocolVersion) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ProtocolVersion) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_ProtocolVersion) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ProtocolVersion) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ProtocolVersion)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.P2P != 0 {
			n += 1 + runtime.Sov(uint64(x.P2P))
		}
		if x.Block != 0 {
			n += 1 + runtime.Sov(uint64(x.Block))
		}
		if x.App != 0 {
			n += 1 + runtime.Sov(uint64(x.App))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*ProtocolVersion)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.App != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.App))
			i--
			dAtA[i] = 0x18
		}
		if x.Block != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Block))
			i--
			dAtA[i] = 0x10
		}
		if x.P2P != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.P2P))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*ProtocolVersion)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ProtocolVersion: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ProtocolVersion: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field P2P", wireType)
				}
				x.P2P = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.P2P |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
				}
				x.Block = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Block |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field App", wireType)
				}
				x.App = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.App |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_NodeInfo                  protoreflect.MessageDescriptor
	fd_NodeInfo_protocol_version protoreflect.FieldDescriptor
	fd_NodeInfo_node_id          protoreflect.FieldDescriptor
	fd_NodeInfo_listen_addr      protoreflect.FieldDescriptor
	fd_NodeInfo_network          protoreflect.FieldDescriptor
	fd_NodeInfo_version          protoreflect.FieldDescriptor
	fd_NodeInfo_channels         protoreflect.FieldDescriptor
	fd_NodeInfo_moniker          protoreflect.FieldDescriptor
	fd_NodeInfo_other            protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_p2p_types_proto_init()
	md_NodeInfo = File_tendermint_p2p_types_proto.Messages().ByName("NodeInfo")
	fd_NodeInfo_protocol_version = md_NodeInfo.Fields().ByName("protocol_version")
	fd_NodeInfo_node_id = md_NodeInfo.Fields().ByName("node_id")
	fd_NodeInfo_listen_addr = md_NodeInfo.Fields().ByName("listen_addr")
	fd_NodeInfo_network = md_NodeInfo.Fields().ByName("network")
	fd_NodeInfo_version = md_NodeInfo.Fields().ByName("version")
	fd_NodeInfo_channels = md_NodeInfo.Fields().ByName("channels")
	fd_NodeInfo_moniker = md_NodeInfo.Fields().ByName("moniker")
	fd_NodeInfo_other = md_NodeInfo.Fields().ByName("other")
}

var _ protoreflect.Message = (*fastReflection_NodeInfo)(nil)

type fastReflection_NodeInfo NodeInfo

func (x *NodeInfo) ProtoReflect() protoreflect.Message {
	return (*fastReflection_NodeInfo)(x)
}

func (x *NodeInfo) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_p2p_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_NodeInfo_messageType fastReflection_NodeInfo_messageType
var _ protoreflect.MessageType = fastReflection_NodeInfo_messageType{}

type fastReflection_NodeInfo_messageType struct{}

func (x fastReflection_NodeInfo_messageType) Zero() protoreflect.Message {
	return (*fastReflection_NodeInfo)(nil)
}
func (x fastReflection_NodeInfo_messageType) New() protoreflect.Message {
	return new(fastReflection_NodeInfo)
}
func (x fastReflection_NodeInfo_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_NodeInfo
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_NodeInfo) Descriptor() protoreflect.MessageDescriptor {
	return md_NodeInfo
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_NodeInfo) Type() protoreflect.MessageType {
	return _fastReflection_NodeInfo_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_NodeInfo) New() protoreflect.Message {
	return new(fastReflection_NodeInfo)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_NodeInfo) Interface() protoreflect.ProtoMessage {
	return (*NodeInfo)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_NodeInfo) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ProtocolVersion != nil {
		value := protoreflect.ValueOfMessage(x.ProtocolVersion.ProtoReflect())
		if !f(fd_NodeInfo_protocol_version, value) {
			return
		}
	}
	if x.NodeId != "" {
		value := protoreflect.ValueOfString(x.NodeId)
		if !f(fd_NodeInfo_node_id, value) {
			return
		}
	}
	if x.ListenAddr != "" {
		value := protoreflect.ValueOfString(x.ListenAddr)
		if !f(fd_NodeInfo_listen_addr, value) {
			return
		}
	}
	if x.Network != "" {
		value := protoreflect.ValueOfString(x.Network)
		if !f(fd_NodeInfo_network, value) {
			return
		}
	}
	if x.Version != "" {
		value := protoreflect.ValueOfString(x.Version)
		if !f(fd_NodeInfo_version, value) {
			return
		}
	}
	if len(x.Channels) != 0 {
		value := protoreflect.ValueOfBytes(x.Channels)
		if !f(fd_NodeInfo_channels, value) {
			return
		}
	}
	if x.Moniker != "" {
		value := protoreflect.ValueOfString(x.Moniker)
		if !f(fd_NodeInfo_moniker, value) {
			return
		}
	}
	if x.Other != nil {
		value := protoreflect.ValueOfMessage(x.Other.ProtoReflect())
		if !f(fd_NodeInfo_other, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_NodeInfo) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.p2p.NodeInfo.protocol_version":
		return x.ProtocolVersion != nil
	case "tendermint.p2p.NodeInfo.node_id":
		return x.NodeId != ""
	case "tendermint.p2p.NodeInfo.listen_addr":
		return x.ListenAddr != ""
	case "tendermint.p2p.NodeInfo.network":
		return x.Network != ""
	case "tendermint.p2p.NodeInfo.version":
		return x.Version != ""
	case "tendermint.p2p.NodeInfo.channels":
		return len(x.Channels) != 0
	case "tendermint.p2p.NodeInfo.moniker":
		return x.Moniker != ""
	case "tendermint.p2p.NodeInfo.other":
		return x.Other != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.NodeInfo"))
		}
		panic(fmt.Errorf("message tendermint.p2p.NodeInfo does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_NodeInfo) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.p2p.NodeInfo.protocol_version":
		x.ProtocolVersion = nil
	case "tendermint.p2p.NodeInfo.node_id":
		x.NodeId = ""
	case "tendermint.p2p.NodeInfo.listen_addr":
		x.ListenAddr = ""
	case "tendermint.p2p.NodeInfo.network":
		x.Network = ""
	case "tendermint.p2p.NodeInfo.version":
		x.Version = ""
	case "tendermint.p2p.NodeInfo.channels":
		x.Channels = nil
	case "tendermint.p2p.NodeInfo.moniker":
		x.Moniker = ""
	case "tendermint.p2p.NodeInfo.other":
		x.Other = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.NodeInfo"))
		}
		panic(fmt.Errorf("message tendermint.p2p.NodeInfo does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_NodeInfo) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.p2p.NodeInfo.protocol_version":
		value := x.ProtocolVersion
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.p2p.NodeInfo.node_id":
		value := x.NodeId
		return protoreflect.ValueOfString(value)
	case "tendermint.p2p.NodeInfo.listen_addr":
		value := x.ListenAddr
		return protoreflect.ValueOfString(value)
	case "tendermint.p2p.NodeInfo.network":
		value := x.Network
		return protoreflect.ValueOfString(value)
	case "tendermint.p2p.NodeInfo.version":
		value := x.Version
		return protoreflect.ValueOfString(value)
	case "tendermint.p2p.NodeInfo.channels":
		value := x.Channels
		return protoreflect.ValueOfBytes(value)
	case "tendermint.p2p.NodeInfo.moniker":
		value := x.Moniker
		return protoreflect.ValueOfString(value)
	case "tendermint.p2p.NodeInfo.other":
		value := x.Other
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.NodeInfo"))
		}
		panic(fmt.Errorf("message tendermint.p2p.NodeInfo does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_NodeInfo) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.p2p.NodeInfo.protocol_version":
		x.ProtocolVersion = value.Message().Interface().(*ProtocolVersion)
	case "tendermint.p2p.NodeInfo.node_id":
		x.NodeId = value.Interface().(string)
	case "tendermint.p2p.NodeInfo.listen_addr":
		x.ListenAddr = value.Interface().(string)
	case "tendermint.p2p.NodeInfo.network":
		x.Network = value.Interface().(string)
	case "tendermint.p2p.NodeInfo.version":
		x.Version = value.Interface().(string)
	case "tendermint.p2p.NodeInfo.channels":
		x.Channels = value.Bytes()
	case "tendermint.p2p.NodeInfo.moniker":
		x.Moniker = value.Interface().(string)
	case "tendermint.p2p.NodeInfo.other":
		x.Other = value.Message().Interface().(*NodeInfoOther)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.NodeInfo"))
		}
		panic(fmt.Errorf("message tendermint.p2p.NodeInfo does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_NodeInfo) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.p2p.NodeInfo.protocol_version":
		if x.ProtocolVersion == nil {
			x.ProtocolVersion = new(ProtocolVersion)
		}
		return protoreflect.ValueOfMessage(x.ProtocolVersion.ProtoReflect())
	case "tendermint.p2p.NodeInfo.other":
		if x.Other == nil {
			x.Other = new(NodeInfoOther)
		}
		return protoreflect.ValueOfMessage(x.Other.ProtoReflect())
	case "tendermint.p2p.NodeInfo.node_id":
		panic(fmt.Errorf("field node_id of message tendermint.p2p.NodeInfo is not mutable"))
	case "tendermint.p2p.NodeInfo.listen_addr":
		panic(fmt.Errorf("field listen_addr of message tendermint.p2p.NodeInfo is not mutable"))
	case "tendermint.p2p.NodeInfo.network":
		panic(fmt.Errorf("field network of message tendermint.p2p.NodeInfo is not mutable"))
	case "tendermint.p2p.NodeInfo.version":
		panic(fmt.Errorf("field version of message tendermint.p2p.NodeInfo is not mutable"))
	case "tendermint.p2p.NodeInfo.channels":
		panic(fmt.Errorf("field channels of message tendermint.p2p.NodeInfo is not mutable"))
	case "tendermint.p2p.NodeInfo.moniker":
		panic(fmt.Errorf("field moniker of message tendermint.p2p.NodeInfo is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.NodeInfo"))
		}
		panic(fmt.Errorf("message tendermint.p2p.NodeInfo does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_NodeInfo) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.p2p.NodeInfo.protocol_version":
		m := new(ProtocolVersion)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.p2p.NodeInfo.node_id":
		return protoreflect.ValueOfString("")
	case "tendermint.p2p.NodeInfo.listen_addr":
		return protoreflect.ValueOfString("")
	case "tendermint.p2p.NodeInfo.network":
		return protoreflect.ValueOfString("")
	case "tendermint.p2p.NodeInfo.version":
		return protoreflect.ValueOfString("")
	case "tendermint.p2p.NodeInfo.channels":
		return protoreflect.ValueOfBytes(nil)
	case "tendermint.p2p.NodeInfo.moniker":
		return protoreflect.ValueOfString("")
	case "tendermint.p2p.NodeInfo.other":
		m := new(NodeInfoOther)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.NodeInfo"))
		}
		panic(fmt.Errorf("message tendermint.p2p.NodeInfo does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_NodeInfo) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.p2p.NodeInfo", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_NodeInfo) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_NodeInfo) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_NodeInfo) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_NodeInfo) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*NodeInfo)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.ProtocolVersion != nil {
			l = options.Size(x.ProtocolVersion)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.NodeId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ListenAddr)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Network)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Version)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Channels)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Moniker)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Other != nil {
			l = options.Size(x.Other)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*NodeInfo)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Other != nil {
			encoded, err := options.Marshal(x.Other)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x42
		}
		if len(x.Moniker) > 0 {
			i -= len(x.Moniker)
			copy(dAtA[i:], x.Moniker)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Moniker)))
			i--
			dAtA[i] = 0x3a
		}
		if len(x.Channels) > 0 {
			i -= len(x.Channels)
			copy(dAtA[i:], x.Channels)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Channels)))
			i--
			dAtA[i] = 0x32
		}
		if len(x.Version) > 0 {
			i -= len(x.Version)
			copy(dAtA[i:], x.Version)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Version)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.Network) > 0 {
			i -= len(x.Network)
			copy(dAtA[i:], x.Network)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Network)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.ListenAddr) > 0 {
			i -= len(x.ListenAddr)
			copy(dAtA[i:], x.ListenAddr)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ListenAddr)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.NodeId) > 0 {
			i -= len(x.NodeId)
			copy(dAtA[i:], x.NodeId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.NodeId)))
			i--
			dAtA[i] = 0x12
		}
		if x.ProtocolVersion != nil {
			encoded, err := options.Marshal(x.ProtocolVersion)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*NodeInfo)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: NodeInfo: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: NodeInfo: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProtocolVersion", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.ProtocolVersion == nil {
					x.ProtocolVersion = &ProtocolVersion{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.ProtocolVersion); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NodeId", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.NodeId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ListenAddr", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ListenAddr = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Network", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Network = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Version = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Channels", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Channels = append(x.Channels[:0], dAtA[iNdEx:postIndex]...)
				if x.Channels == nil {
					x.Channels = []byte{}
				}
				iNdEx = postIndex
			case 7:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Moniker", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Moniker = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 8:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Other", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Other == nil {
					x.Other = &NodeInfoOther{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Other); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_NodeInfoOther             protoreflect.MessageDescriptor
	fd_NodeInfoOther_tx_index    protoreflect.FieldDescriptor
	fd_NodeInfoOther_rpc_address protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_p2p_types_proto_init()
	md_NodeInfoOther = File_tendermint_p2p_types_proto.Messages().ByName("NodeInfoOther")
	fd_NodeInfoOther_tx_index = md_NodeInfoOther.Fields().ByName("tx_index")
	fd_NodeInfoOther_rpc_address = md_NodeInfoOther.Fields().ByName("rpc_address")
}

var _ protoreflect.Message = (*fastReflection_NodeInfoOther)(nil)

type fastReflection_NodeInfoOther NodeInfoOther

func (x *NodeInfoOther) ProtoReflect() protoreflect.Message {
	return (*fastReflection_NodeInfoOther)(x)
}

func (x *NodeInfoOther) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_p2p_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_NodeInfoOther_messageType fastReflection_NodeInfoOther_messageType
var _ protoreflect.MessageType = fastReflection_NodeInfoOther_messageType{}

type fastReflection_NodeInfoOther_messageType struct{}

func (x fastReflection_NodeInfoOther_messageType) Zero() protoreflect.Message {
	return (*fastReflection_NodeInfoOther)(nil)
}
func (x fastReflection_NodeInfoOther_messageType) New() protoreflect.Message {
	return new(fastReflection_NodeInfoOther)
}
func (x fastReflection_NodeInfoOther_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_NodeInfoOther
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_NodeInfoOther) Descriptor() protoreflect.MessageDescriptor {
	return md_NodeInfoOther
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_NodeInfoOther) Type() protoreflect.MessageType {
	return _fastReflection_NodeInfoOther_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_NodeInfoOther) New() protoreflect.Message {
	return new(fastReflection_NodeInfoOther)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_NodeInfoOther) Interface() protoreflect.ProtoMessage {
	return (*NodeInfoOther)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_NodeInfoOther) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.TxIndex != "" {
		value := protoreflect.ValueOfString(x.TxIndex)
		if !f(fd_NodeInfoOther_tx_index, value) {
			return
		}
	}
	if x.RpcAddress != "" {
		value := protoreflect.ValueOfString(x.RpcAddress)
		if !f(fd_NodeInfoOther_rpc_address, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_NodeInfoOther) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.p2p.NodeInfoOther.tx_index":
		return x.TxIndex != ""
	case "tendermint.p2p.NodeInfoOther.rpc_address":
		return x.RpcAddress != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.NodeInfoOther"))
		}
		panic(fmt.Errorf("message tendermint.p2p.NodeInfoOther does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_NodeInfoOther) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.p2p.NodeInfoOther.tx_index":
		x.TxIndex = ""
	case "tendermint.p2p.NodeInfoOther.rpc_address":
		x.RpcAddress = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.NodeInfoOther"))
		}
		panic(fmt.Errorf("message tendermint.p2p.NodeInfoOther does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_NodeInfoOther) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.p2p.NodeInfoOther.tx_index":
		value := x.TxIndex
		return protoreflect.ValueOfString(value)
	case "tendermint.p2p.NodeInfoOther.rpc_address":
		value := x.RpcAddress
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.NodeInfoOther"))
		}
		panic(fmt.Errorf("message tendermint.p2p.NodeInfoOther does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_NodeInfoOther) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.p2p.NodeInfoOther.tx_index":
		x.TxIndex = value.Interface().(string)
	case "tendermint.p2p.NodeInfoOther.rpc_address":
		x.RpcAddress = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.NodeInfoOther"))
		}
		panic(fmt.Errorf("message tendermint.p2p.NodeInfoOther does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_NodeInfoOther) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.p2p.NodeInfoOther.tx_index":
		panic(fmt.Errorf("field tx_index of message tendermint.p2p.NodeInfoOther is not mutable"))
	case "tendermint.p2p.NodeInfoOther.rpc_address":
		panic(fmt.Errorf("field rpc_address of message tendermint.p2p.NodeInfoOther is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.NodeInfoOther"))
		}
		panic(fmt.Errorf("message tendermint.p2p.NodeInfoOther does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_NodeInfoOther) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.p2p.NodeInfoOther.tx_index":
		return protoreflect.ValueOfString("")
	case "tendermint.p2p.NodeInfoOther.rpc_address":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.NodeInfoOther"))
		}
		panic(fmt.Errorf("message tendermint.p2p.NodeInfoOther does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_NodeInfoOther) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.p2p.NodeInfoOther", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_NodeInfoOther) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_NodeInfoOther) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_NodeInfoOther) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_NodeInfoOther) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*NodeInfoOther)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.TxIndex)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.RpcAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*NodeInfoOther)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.RpcAddress) > 0 {
			i -= len(x.RpcAddress)
			copy(dAtA[i:], x.RpcAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RpcAddress)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.TxIndex) > 0 {
			i -= len(x.TxIndex)
			copy(dAtA[i:], x.TxIndex)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.TxIndex)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*NodeInfoOther)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: NodeInfoOther: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: NodeInfoOther: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TxIndex", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.TxIndex = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RpcAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.RpcAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_PeerInfo_2_list)(nil)

type _PeerInfo_2_list struct {
	list *[]*PeerAddressInfo
}

func (x *_PeerInfo_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_PeerInfo_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_PeerInfo_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*PeerAddressInfo)
	(*x.list)[i] = concreteValue
}

func (x *_PeerInfo_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*PeerAddressInfo)
	*x.list = append(*x.list, concreteValue)
}

func (x *_PeerInfo_2_list) AppendMutable() protoreflect.Value {
	v := new(PeerAddressInfo)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_PeerInfo_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_PeerInfo_2_list) NewElement() protoreflect.Value {
	v := new(PeerAddressInfo)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_PeerInfo_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_PeerInfo                protoreflect.MessageDescriptor
	fd_PeerInfo_id             protoreflect.FieldDescriptor
	fd_PeerInfo_address_info   protoreflect.FieldDescriptor
	fd_PeerInfo_last_connected protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_p2p_types_proto_init()
	md_PeerInfo = File_tendermint_p2p_types_proto.Messages().ByName("PeerInfo")
	fd_PeerInfo_id = md_PeerInfo.Fields().ByName("id")
	fd_PeerInfo_address_info = md_PeerInfo.Fields().ByName("address_info")
	fd_PeerInfo_last_connected = md_PeerInfo.Fields().ByName("last_connected")
}

var _ protoreflect.Message = (*fastReflection_PeerInfo)(nil)

type fastReflection_PeerInfo PeerInfo

func (x *PeerInfo) ProtoReflect() protoreflect.Message {
	return (*fastReflection_PeerInfo)(x)
}

func (x *PeerInfo) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_p2p_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_PeerInfo_messageType fastReflection_PeerInfo_messageType
var _ protoreflect.MessageType = fastReflection_PeerInfo_messageType{}

type fastReflection_PeerInfo_messageType struct{}

func (x fastReflection_PeerInfo_messageType) Zero() protoreflect.Message {
	return (*fastReflection_PeerInfo)(nil)
}
func (x fastReflection_PeerInfo_messageType) New() protoreflect.Message {
	return new(fastReflection_PeerInfo)
}
func (x fastReflection_PeerInfo_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_PeerInfo
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_PeerInfo) Descriptor() protoreflect.MessageDescriptor {
	return md_PeerInfo
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_PeerInfo) Type() protoreflect.MessageType {
	return _fastReflection_PeerInfo_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_PeerInfo) New() protoreflect.Message {
	return new(fastReflection_PeerInfo)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_PeerInfo) Interface() protoreflect.ProtoMessage {
	return (*PeerInfo)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_PeerInfo) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != "" {
		value := protoreflect.ValueOfString(x.Id)
		if !f(fd_PeerInfo_id, value) {
			return
		}
	}
	if len(x.AddressInfo) != 0 {
		value := protoreflect.ValueOfList(&_PeerInfo_2_list{list: &x.AddressInfo})
		if !f(fd_PeerInfo_address_info, value) {
			return
		}
	}
	if x.LastConnected != nil {
		value := protoreflect.ValueOfMessage(x.LastConnected.ProtoReflect())
		if !f(fd_PeerInfo_last_connected, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_PeerInfo) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.p2p.PeerInfo.id":
		return x.Id != ""
	case "tendermint.p2p.PeerInfo.address_info":
		return len(x.AddressInfo) != 0
	case "tendermint.p2p.PeerInfo.last_connected":
		return x.LastConnected != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.PeerInfo"))
		}
		panic(fmt.Errorf("message tendermint.p2p.PeerInfo does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PeerInfo) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.p2p.PeerInfo.id":
		x.Id = ""
	case "tendermint.p2p.PeerInfo.address_info":
		x.AddressInfo = nil
	case "tendermint.p2p.PeerInfo.last_connected":
		x.LastConnected = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.PeerInfo"))
		}
		panic(fmt.Errorf("message tendermint.p2p.PeerInfo does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_PeerInfo) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.p2p.PeerInfo.id":
		value := x.Id
		return protoreflect.ValueOfString(value)
	case "tendermint.p2p.PeerInfo.address_info":
		if len(x.AddressInfo) == 0 {
			return protoreflect.ValueOfList(&_PeerInfo_2_list{})
		}
		listValue := &_PeerInfo_2_list{list: &x.AddressInfo}
		return protoreflect.ValueOfList(listValue)
	case "tendermint.p2p.PeerInfo.last_connected":
		value := x.LastConnected
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.PeerInfo"))
		}
		panic(fmt.Errorf("message tendermint.p2p.PeerInfo does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PeerInfo) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.p2p.PeerInfo.id":
		x.Id = value.Interface().(string)
	case "tendermint.p2p.PeerInfo.address_info":
		lv := value.List()
		clv := lv.(*_PeerInfo_2_list)
		x.AddressInfo = *clv.list
	case "tendermint.p2p.PeerInfo.last_connected":
		x.LastConnected = value.Message().Interface().(*timestamppb.Timestamp)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.PeerInfo"))
		}
		panic(fmt.Errorf("message tendermint.p2p.PeerInfo does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PeerInfo) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.p2p.PeerInfo.address_info":
		if x.AddressInfo == nil {
			x.AddressInfo = []*PeerAddressInfo{}
		}
		value := &_PeerInfo_2_list{list: &x.AddressInfo}
		return protoreflect.ValueOfList(value)
	case "tendermint.p2p.PeerInfo.last_connected":
		if x.LastConnected == nil {
			x.LastConnected = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.LastConnected.ProtoReflect())
	case "tendermint.p2p.PeerInfo.id":
		panic(fmt.Errorf("field id of message tendermint.p2p.PeerInfo is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.PeerInfo"))
		}
		panic(fmt.Errorf("message tendermint.p2p.PeerInfo does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_PeerInfo) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.p2p.PeerInfo.id":
		return protoreflect.ValueOfString("")
	case "tendermint.p2p.PeerInfo.address_info":
		list := []*PeerAddressInfo{}
		return protoreflect.ValueOfList(&_PeerInfo_2_list{list: &list})
	case "tendermint.p2p.PeerInfo.last_connected":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.PeerInfo"))
		}
		panic(fmt.Errorf("message tendermint.p2p.PeerInfo does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_PeerInfo) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.p2p.PeerInfo", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_PeerInfo) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PeerInfo) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_PeerInfo) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_PeerInfo) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*PeerInfo)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Id)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.AddressInfo) > 0 {
			for _, e := range x.AddressInfo {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.LastConnected != nil {
			l = options.Size(x.LastConnected)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*PeerInfo)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.LastConnected != nil {
			encoded, err := options.Marshal(x.LastConnected)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.AddressInfo) > 0 {
			for iNdEx := len(x.AddressInfo) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.AddressInfo[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x12
			}
		}
		if len(x.Id) > 0 {
			i -= len(x.Id)
			copy(dAtA[i:], x.Id)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Id)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*PeerInfo)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PeerInfo: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PeerInfo: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Id = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AddressInfo", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.AddressInfo = append(x.AddressInfo, &PeerAddressInfo{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.AddressInfo[len(x.AddressInfo)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field LastConnected", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.LastConnected == nil {
					x.LastConnected = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.LastConnected); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_PeerAddressInfo                   protoreflect.MessageDescriptor
	fd_PeerAddressInfo_address           protoreflect.FieldDescriptor
	fd_PeerAddressInfo_last_dial_success protoreflect.FieldDescriptor
	fd_PeerAddressInfo_last_dial_failure protoreflect.FieldDescriptor
	fd_PeerAddressInfo_dial_failures     protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_p2p_types_proto_init()
	md_PeerAddressInfo = File_tendermint_p2p_types_proto.Messages().ByName("PeerAddressInfo")
	fd_PeerAddressInfo_address = md_PeerAddressInfo.Fields().ByName("address")
	fd_PeerAddressInfo_last_dial_success = md_PeerAddressInfo.Fields().ByName("last_dial_success")
	fd_PeerAddressInfo_last_dial_failure = md_PeerAddressInfo.Fields().ByName("last_dial_failure")
	fd_PeerAddressInfo_dial_failures = md_PeerAddressInfo.Fields().ByName("dial_failures")
}

var _ protoreflect.Message = (*fastReflection_PeerAddressInfo)(nil)

type fastReflection_PeerAddressInfo PeerAddressInfo

func (x *PeerAddressInfo) ProtoReflect() protoreflect.Message {
	return (*fastReflection_PeerAddressInfo)(x)
}

func (x *PeerAddressInfo) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_p2p_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_PeerAddressInfo_messageType fastReflection_PeerAddressInfo_messageType
var _ protoreflect.MessageType = fastReflection_PeerAddressInfo_messageType{}

type fastReflection_PeerAddressInfo_messageType struct{}

func (x fastReflection_PeerAddressInfo_messageType) Zero() protoreflect.Message {
	return (*fastReflection_PeerAddressInfo)(nil)
}
func (x fastReflection_PeerAddressInfo_messageType) New() protoreflect.Message {
	return new(fastReflection_PeerAddressInfo)
}
func (x fastReflection_PeerAddressInfo_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_PeerAddressInfo
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_PeerAddressInfo) Descriptor() protoreflect.MessageDescriptor {
	return md_PeerAddressInfo
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_PeerAddressInfo) Type() protoreflect.MessageType {
	return _fastReflection_PeerAddressInfo_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_PeerAddressInfo) New() protoreflect.Message {
	return new(fastReflection_PeerAddressInfo)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_PeerAddressInfo) Interface() protoreflect.ProtoMessage {
	return (*PeerAddressInfo)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_PeerAddressInfo) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Address != "" {
		value := protoreflect.ValueOfString(x.Address)
		if !f(fd_PeerAddressInfo_address, value) {
			return
		}
	}
	if x.LastDialSuccess != nil {
		value := protoreflect.ValueOfMessage(x.LastDialSuccess.ProtoReflect())
		if !f(fd_PeerAddressInfo_last_dial_success, value) {
			return
		}
	}
	if x.LastDialFailure != nil {
		value := protoreflect.ValueOfMessage(x.LastDialFailure.ProtoReflect())
		if !f(fd_PeerAddressInfo_last_dial_failure, value) {
			return
		}
	}
	if x.DialFailures != uint32(0) {
		value := protoreflect.ValueOfUint32(x.DialFailures)
		if !f(fd_PeerAddressInfo_dial_failures, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_PeerAddressInfo) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.p2p.PeerAddressInfo.address":
		return x.Address != ""
	case "tendermint.p2p.PeerAddressInfo.last_dial_success":
		return x.LastDialSuccess != nil
	case "tendermint.p2p.PeerAddressInfo.last_dial_failure":
		return x.LastDialFailure != nil
	case "tendermint.p2p.PeerAddressInfo.dial_failures":
		return x.DialFailures != uint32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.PeerAddressInfo"))
		}
		panic(fmt.Errorf("message tendermint.p2p.PeerAddressInfo does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PeerAddressInfo) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.p2p.PeerAddressInfo.address":
		x.Address = ""
	case "tendermint.p2p.PeerAddressInfo.last_dial_success":
		x.LastDialSuccess = nil
	case "tendermint.p2p.PeerAddressInfo.last_dial_failure":
		x.LastDialFailure = nil
	case "tendermint.p2p.PeerAddressInfo.dial_failures":
		x.DialFailures = uint32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.PeerAddressInfo"))
		}
		panic(fmt.Errorf("message tendermint.p2p.PeerAddressInfo does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_PeerAddressInfo) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.p2p.PeerAddressInfo.address":
		value := x.Address
		return protoreflect.ValueOfString(value)
	case "tendermint.p2p.PeerAddressInfo.last_dial_success":
		value := x.LastDialSuccess
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.p2p.PeerAddressInfo.last_dial_failure":
		value := x.LastDialFailure
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.p2p.PeerAddressInfo.dial_failures":
		value := x.DialFailures
		return protoreflect.ValueOfUint32(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.PeerAddressInfo"))
		}
		panic(fmt.Errorf("message tendermint.p2p.PeerAddressInfo does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PeerAddressInfo) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.p2p.PeerAddressInfo.address":
		x.Address = value.Interface().(string)
	case "tendermint.p2p.PeerAddressInfo.last_dial_success":
		x.LastDialSuccess = value.Message().Interface().(*timestamppb.Timestamp)
	case "tendermint.p2p.PeerAddressInfo.last_dial_failure":
		x.LastDialFailure = value.Message().Interface().(*timestamppb.Timestamp)
	case "tendermint.p2p.PeerAddressInfo.dial_failures":
		x.DialFailures = uint32(value.Uint())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.PeerAddressInfo"))
		}
		panic(fmt.Errorf("message tendermint.p2p.PeerAddressInfo does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PeerAddressInfo) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.p2p.PeerAddressInfo.last_dial_success":
		if x.LastDialSuccess == nil {
			x.LastDialSuccess = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.LastDialSuccess.ProtoReflect())
	case "tendermint.p2p.PeerAddressInfo.last_dial_failure":
		if x.LastDialFailure == nil {
			x.LastDialFailure = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.LastDialFailure.ProtoReflect())
	case "tendermint.p2p.PeerAddressInfo.address":
		panic(fmt.Errorf("field address of message tendermint.p2p.PeerAddressInfo is not mutable"))
	case "tendermint.p2p.PeerAddressInfo.dial_failures":
		panic(fmt.Errorf("field dial_failures of message tendermint.p2p.PeerAddressInfo is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.PeerAddressInfo"))
		}
		panic(fmt.Errorf("message tendermint.p2p.PeerAddressInfo does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_PeerAddressInfo) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.p2p.PeerAddressInfo.address":
		return protoreflect.ValueOfString("")
	case "tendermint.p2p.PeerAddressInfo.last_dial_success":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.p2p.PeerAddressInfo.last_dial_failure":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.p2p.PeerAddressInfo.dial_failures":
		return protoreflect.ValueOfUint32(uint32(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.p2p.PeerAddressInfo"))
		}
		panic(fmt.Errorf("message tendermint.p2p.PeerAddressInfo does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_PeerAddressInfo) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.p2p.PeerAddressInfo", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_PeerAddressInfo) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PeerAddressInfo) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_PeerAddressInfo) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_PeerAddressInfo) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*PeerAddressInfo)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.LastDialSuccess != nil {
			l = options.Size(x.LastDialSuccess)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.LastDialFailure != nil {
			l = options.Size(x.LastDialFailure)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.DialFailures != 0 {
			n += 1 + runtime.Sov(uint64(x.DialFailures))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*PeerAddressInfo)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.DialFailures != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.DialFailures))
			i--
			dAtA[i] = 0x20
		}
		if x.LastDialFailure != nil {
			encoded, err := options.Marshal(x.LastDialFailure)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x1a
		}
		if x.LastDialSuccess != nil {
			encoded, err := options.Marshal(x.LastDialSuccess)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*PeerAddressInfo)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PeerAddressInfo: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PeerAddressInfo: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Address = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field LastDialSuccess", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.LastDialSuccess == nil {
					x.LastDialSuccess = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.LastDialSuccess); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field LastDialFailure", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.LastDialFailure == nil {
					x.LastDialFailure = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.LastDialFailure); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DialFailures", wireType)
				}
				x.DialFailures = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.DialFailures |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: tendermint/p2p/types.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProtocolVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	P2P   uint64 `protobuf:"varint,1,opt,name=p2p,proto3" json:"p2p,omitempty"`
	Block uint64 `protobuf:"varint,2,opt,name=block,proto3" json:"block,omitempty"`
	App   uint64 `protobuf:"varint,3,opt,name=app,proto3" json:"app,omitempty"`
}

func (x *ProtocolVersion) Reset() {
	*x = ProtocolVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_p2p_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolVersion) ProtoMessage() {}

// Deprecated: Use ProtocolVersion.ProtoReflect.Descriptor instead.
func (*ProtocolVersion) Descriptor() ([]byte, []int) {
	return file_tendermint_p2p_types_proto_rawDescGZIP(), []int{0}
}

func (x *ProtocolVersion) GetP2P() uint64 {
	if x != nil {
		return x.P2P
	}
	return 0
}

func (x *ProtocolVersion) GetBlock() uint64 {
	if x != nil {
		return x.Block
	}
	return 0
}

func (x *ProtocolVersion) GetApp() uint64 {
	if x != nil {
		return x.App
	}
	return 0
}

type NodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtocolVersion *ProtocolVersion `protobuf:"bytes,1,opt,name=protocol_version,json=protocolVersion,proto3" json:"protocol_version,omitempty"`
	NodeId          string           `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	ListenAddr      string           `protobuf:"bytes,3,opt,name=listen_addr,json=listenAddr,proto3" json:"listen_addr,omitempty"`
	Network         string           `protobuf:"bytes,4,opt,name=network,proto3" json:"network,omitempty"`
	Version         string           `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	Channels        []byte           `protobuf:"bytes,6,opt,name=channels,proto3" json:"channels,omitempty"`
	Moniker         string           `protobuf:"bytes,7,opt,name=moniker,proto3" json:"moniker,omitempty"`
	Other           *NodeInfoOther   `protobuf:"bytes,8,opt,name=other,proto3" json:"other,omitempty"`
}

func (x *NodeInfo) Reset() {
	*x = NodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_p2p_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfo) ProtoMessage() {}

// Deprecated: Use NodeInfo.ProtoReflect.Descriptor instead.
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return file_tendermint_p2p_types_proto_rawDescGZIP(), []int{1}
}

func (x *NodeInfo) GetProtocolVersion() *ProtocolVersion {
	if x != nil {
		return x.ProtocolVersion
	}
	return nil
}

func (x *NodeInfo) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *NodeInfo) GetListenAddr() string {
	if x != nil {
		return x.ListenAddr
	}
	return ""
}

func (x *NodeInfo) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *NodeInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *NodeInfo) GetChannels() []byte {
	if x != nil {
		return x.Channels
	}
	return nil
}

func (x *NodeInfo) GetMoniker() string {
	if x != nil {
		return x.Moniker
	}
	return ""
}

func (x *NodeInfo) GetOther() *NodeInfoOther {
	if x != nil {
		return x.Other
	}
	return nil
}

type NodeInfoOther struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxIndex    string `protobuf:"bytes,1,opt,name=tx_index,json=txIndex,proto3" json:"tx_index,omitempty"`
	RpcAddress string `protobuf:"bytes,2,opt,name=rpc_address,json=rpcAddress,proto3" json:"rpc_address,omitempty"`
}

func (x *NodeInfoOther) Reset() {
	*x = NodeInfoOther{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_p2p_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInfoOther) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfoOther) ProtoMessage() {}

// Deprecated: Use NodeInfoOther.ProtoReflect.Descriptor instead.
func (*NodeInfoOther) Descriptor() ([]byte, []int) {
	return file_tendermint_p2p_types_proto_rawDescGZIP(), []int{2}
}

func (x *NodeInfoOther) GetTxIndex() string {
	if x != nil {
		return x.TxIndex
	}
	return ""
}

func (x *NodeInfoOther) GetRpcAddress() string {
	if x != nil {
		return x.RpcAddress
	}
	return ""
}

type PeerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AddressInfo   []*PeerAddressInfo     `protobuf:"bytes,2,rep,name=address_info,json=addressInfo,proto3" json:"address_info,omitempty"`
	LastConnected *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=last_connected,json=lastConnected,proto3" json:"last_connected,omitempty"`
}

func (x *PeerInfo) Reset() {
	*x = PeerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_p2p_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerInfo) ProtoMessage() {}

// Deprecated: Use PeerInfo.ProtoReflect.Descriptor instead.
func (*PeerInfo) Descriptor() ([]byte, []int) {
	return file_tendermint_p2p_types_proto_rawDescGZIP(), []int{3}
}

func (x *PeerInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PeerInfo) GetAddressInfo() []*PeerAddressInfo {
	if x != nil {
		return x.AddressInfo
	}
	return nil
}

func (x *PeerInfo) GetLastConnected() *timestamppb.Timestamp {
	if x != nil {
		return x.LastConnected
	}
	return nil
}

type PeerAddressInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address         string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	LastDialSuccess *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_dial_success,json=lastDialSuccess,proto3" json:"last_dial_success,omitempty"`
	LastDialFailure *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=last_dial_failure,json=lastDialFailure,proto3" json:"last_dial_failure,omitempty"`
	DialFailures    uint32                 `protobuf:"varint,4,opt,name=dial_failures,json=dialFailures,proto3" json:"dial_failures,omitempty"`
}

func (x *PeerAddressInfo) Reset() {
	*x = PeerAddressInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_p2p_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerAddressInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerAddressInfo) ProtoMessage() {}

// Deprecated: Use PeerAddressInfo.ProtoReflect.Descriptor instead.
func (*PeerAddressInfo) Descriptor() ([]byte, []int) {
	return file_tendermint_p2p_types_proto_rawDescGZIP(), []int{4}
}

func (x *PeerAddressInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *PeerAddressInfo) GetLastDialSuccess() *timestamppb.Timestamp {
	if x != nil {
		return x.LastDialSuccess
	}
	return nil
}

func (x *PeerAddressInfo) GetLastDialFailure() *timestamppb.Timestamp {
	if x != nil {
		return x.LastDialFailure
	}
	return nil
}

func (x *PeerAddressInfo) GetDialFailures() uint32 {
	if x != nil {
		return x.DialFailures
	}
	return 0
}

var File_tendermint_p2p_types_proto protoreflect.FileDescriptor

var file_tendermint_p2p_types_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x70, 0x32, 0x70,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x32, 0x70, 0x1a, 0x14, 0x67, 0x6f,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x03, 0x70, 0x32, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x07, 0xe2, 0xde, 0x1f, 0x03, 0x50, 0x32, 0x50, 0x52, 0x03, 0x70, 0x32,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x61, 0x70, 0x70, 0x22, 0xc7, 0x02, 0x0a, 0x08, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x50, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x32,
	0x70, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xe2, 0xde, 0x1f, 0x06, 0x4e,
	0x6f, 0x64, 0x65, 0x49, 0x44, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x6b, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x6b, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x05, 0x6f, 0x74, 0x68, 0x65,
	0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x05, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x22, 0x5b, 0x0a, 0x0d, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4f,
	0x74, 0x68, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x78, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x2f, 0x0a, 0x0b, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xe2, 0xde, 0x1f, 0x0a, 0x52, 0x50, 0x43, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x0a, 0x72, 0x70, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0xaf, 0x01, 0x0a, 0x08, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xde, 0x1f, 0x02, 0x49,
	0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x42, 0x0a, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x50, 0x65, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x47, 0x0a, 0x0e, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0x90,
	0xdf, 0x1f, 0x01, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x22, 0xec, 0x01, 0x0a, 0x0f, 0x50, 0x65, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x4c, 0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x64, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x0f, 0x6c,
	0x61, 0x73, 0x74, 0x44, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x4c,
	0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x64, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x0f, 0x6c, 0x61, 0x73,
	0x74, 0x44, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x64, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x64, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x73, 0x42, 0xaa, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x32, 0x70, 0x42, 0x0a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x74, 0x2f, 0x70, 0x32, 0x70, 0xa2, 0x02, 0x03, 0x54, 0x50, 0x58, 0xaa, 0x02, 0x0e,
	0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x50, 0x32, 0x70, 0xca, 0x02,
	0x0e, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x5c, 0x50, 0x32, 0x70, 0xe2,
	0x02, 0x1a, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x5c, 0x50, 0x32, 0x70,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x54,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x3a, 0x3a, 0x50, 0x32, 0x70, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tendermint_p2p_types_proto_rawDescOnce sync.Once
	file_tendermint_p2p_types_proto_rawDescData = file_tendermint_p2p_types_proto_rawDesc
)

func file_tendermint_p2p_types_proto_rawDescGZIP() []byte {
	file_tendermint_p2p_types_proto_rawDescOnce.Do(func() {
		file_tendermint_p2p_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_tendermint_p2p_types_proto_rawDescData)
	})
	return file_tendermint_p2p_types_proto_rawDescData
}

var file_tendermint_p2p_types_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tendermint_p2p_types_proto_goTypes = []interface{}{
	(*ProtocolVersion)(nil),       // 0: tendermint.p2p.ProtocolVersion
	(*NodeInfo)(nil),              // 1: tendermint.p2p.NodeInfo
	(*NodeInfoOther)(nil),         // 2: tendermint.p2p.NodeInfoOther
	(*PeerInfo)(nil),              // 3: tendermint.p2p.PeerInfo
	(*PeerAddressInfo)(nil),       // 4: tendermint.p2p.PeerAddressInfo
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_tendermint_p2p_types_proto_depIdxs = []int32{
	0, // 0: tendermint.p2p.NodeInfo.protocol_version:type_name -> tendermint.p2p.ProtocolVersion
	2, // 1: tendermint.p2p.NodeInfo.other:type_name -> tendermint.p2p.NodeInfoOther
	4, // 2: tendermint.p2p.PeerInfo.address_info:type_name -> tendermint.p2p.PeerAddressInfo
	5, // 3: tendermint.p2p.PeerInfo.last_connected:type_name -> google.protobuf.Timestamp
	5, // 4: tendermint.p2p.PeerAddressInfo.last_dial_success:type_name -> google.protobuf.Timestamp
	5, // 5: tendermint.p2p.PeerAddressInfo.last_dial_failure:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_tendermint_p2p_types_proto_init() }
func file_tendermint_p2p_types_proto_init() {
	if File_tendermint_p2p_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tendermint_p2p_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolVersion); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tendermint_p2p_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tendermint_p2p_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInfoOther); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tendermint_p2p_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tendermint_p2p_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerAddressInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tendermint_p2p_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tendermint_p2p_types_proto_goTypes,
		DependencyIndexes: file_tendermint_p2p_types_proto_depIdxs,
		MessageInfos:      file_tendermint_p2p_types_proto_msgTypes,
	}.Build()
	File_tendermint_p2p_types_proto = out.File
	file_tendermint_p2p_types_proto_rawDesc = nil
	file_tendermint_p2p_types_proto_goTypes = nil
	file_tendermint_p2p_types_proto_depIdxs = nil
}
