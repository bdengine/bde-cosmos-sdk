#!/bin/sh

set -e
addr=$(simd keys show fd -a --keyring-backend=test)
echo "12345678" | simd tx bank send "$addr" "$1" 100stake --chain-id="testing" --node tcp://cosmos:26657 --grpc=cosmos:9090 --yes --keyring-backend=test
