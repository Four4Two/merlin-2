# this script runs under the assumption that a three-validator environment is running on your local machine(multinode-local-testnet.sh)
# this script would do basic setup that has to be achieved to actual superfluid staking
# prior to running this script, have the following json file in the directory running this script
#
# jinx-ufury.json
# {
# 	"weights": "5jinx,5ufury",
# 	"initial-deposit": "1000000jinx,1000000ufury",
# 	"swap-fee": "0.01",
# 	"exit-fee": "0.01",
# 	"future-governor": "168h"
# }

# create pool
merlin tx gamm create-pool --pool-file=./stake-ufury.json --from=validator1 --keyring-backend=test --chain-id="blackfury-1"--yes --fees 200000ufury
sleep 7

# test swap in pool created
merlin tx gamm swap-exact-amount-in 100000ufury 50000 --swap-route-pool-ids=1 --swap-route-denoms=jinx --from=validator1 --keyring-backend=test --chain-id="blackfury-1" --yes --fees 200000ufury
sleep 7

# create a lock up with lockable duration 360h
merlin tx lockup lock-tokens 10000000000000000000gamm/pool/1 --duration=360h --from=validator1 --keyring-backend=test --chain-id="blackfury-1" --broadcast-mode=block --yes --fees 200000ufury
sleep 7

# submit and pass proposal for superfluid
merlin tx gov submit-proposal set-superfluid-assets-proposal --title="set superfluid assets" --description="set superfluid assets description" --superfluid-assets="gamm/pool/1" --deposit=10000000ufury --from=validator1 --chain-id="blackfury-1" --keyring-backend=test --broadcast-mode=block --yes --fees 200000ufury
sleep 7

merlin tx gov deposit 1 10000000jinx --from=validator1 --keyring-backend=test --chain-id="blackfury-1" --broadcast-mode=block --yes --fees 200000ufury
sleep 7

merlin tx gov vote 1 yes --from=validator1 --keyring-backend=test --chain-id="blackfury-1" --yes --fees 200000ufury
sleep 7
merlin tx gov vote 1 yes --from=validator2 --keyring-backend=test --chain-id="blackfury-1" --yes --home=$HOME/.merlin/validator2 --fees 200000ufury
sleep 7
merlin tx gov vote 1 yes --from=validator3 --keyring-backend=test --chain-id="blackfury-1" --yes --home=$HOME/.merlin/validator3 --fees 200000ufury
sleep 7
merlin tx gov vote 1 yes --from=validator4 --keyring-backend=test --chain-id="blackfury-1" --yes --home=$HOME/.merlin/validator4 --fees 200000ufury
sleep 7
merlin tx gov vote 1 yes --from=validator5 --keyring-backend=test --chain-id="blackfury-1" --yes --home=$HOME/.merlin/validator5 --fees 200000ufury
sleep 7
merlin tx gov vote 1 yes --from=validator6 --keyring-backend=test --chain-id="blackfury-1" --yes --home=$HOME/.merlin/validator6 --fees 200000ufury
sleep 7
merlin tx gov vote 1 yes --from=validator7 --keyring-backend=test --chain-id="blackfury-1" --yes --home=$HOME/.merlin/validator7 --fees 200000ufury
sleep 7
merlin tx gov vote 1 yes --from=validator8 --keyring-backend=test --chain-id="blackfury-1" --yes --home=$HOME/.merlin/validator8 --fees 200000ufury
sleep 7
merlin tx gov vote 1 yes --from=validator9 --keyring-backend=test --chain-id="blackfury-1" --yes --home=$HOME/.merlin/validator9 --fees 200000ufury
sleep 7
