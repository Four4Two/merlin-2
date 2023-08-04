#!/bin/bash

CHAIN_ID=blackfury-1
VAL=$(merlin keys show -a validator --keyring-backend test)

# We can make this a loop in the future, hard with bash, so I copy it twice

CONTRACT=cw20_base
# we cannot really do this progamatically, find this from the events, so we hardcode
PROPOSAL=1

merlin tx gov submit-proposal wasm-store $CONTRACT.wasm --title "Add $CONTRACT" \
  --description "Let's upload this contract" --run-as $VAL \
  --from validator --keyring-backend test --chain-id $CHAIN_ID -y -b block \
  --gas 9000000 --gas-prices 0.025stake

merlin query gov proposal $PROPOSAL

merlin tx gov deposit $PROPOSAL 10000000stake --from validator --keyring-backend test \
    --chain-id $CHAIN_ID -y -b block --gas 5000000 --gas-prices 0.025stake

merlin tx gov vote $PROPOSAL yes --from validator --keyring-backend test \
    --chain-id $CHAIN_ID -y -b block --gas 400000 --gas-prices 0.025stake


# repeat with new variables
CONTRACT=cw1_whitelist
PROPOSAL=2

merlin tx gov submit-proposal wasm-store $CONTRACT.wasm --title "Add $CONTRACT" \
  --description "Let's upload this contract" --run-as $VAL \
  --from validator --keyring-backend test --chain-id $CHAIN_ID -y -b block \
  --gas 9000000 --gas-prices 0.025stake

merlin query gov proposal $PROPOSAL

merlin tx gov deposit $PROPOSAL 10000000stake --from validator --keyring-backend test \
    --chain-id $CHAIN_ID -y -b block --gas 5000000 --gas-prices 0.025stake

merlin tx gov vote $PROPOSAL yes --from validator --keyring-backend test \
    --chain-id $CHAIN_ID -y -b block --gas 400000 --gas-prices 0.025stake


# now check the results

merlin query wasm list-code

echo "Waiting for voting periods to finish..."
sleep 120

merlin query wasm list-code