# LocalMerlin

LocalMerlin is a complete Merlin testnet containerized with Docker and orchestrated with a simple docker-compose file. LocalMerlin comes preconfigured with opinionated, sensible defaults for a standard testing environment.

LocalMerlin comes in two flavors:

1. No initial state: brand new testnet with no initial state. 
2. With mainnet state: creates a testnet from a mainnet state export

## Prerequisites

Ensure you have docker and docker-compose installed:

```sh
# Docker
sudo apt-get remove docker docker-engine docker.io
sudo apt-get update
sudo apt install docker.io -y

# Docker compose
sudo apt install docker-compose -y
```

## 1. LocalMerlin - No Initial State

The following commands must be executed from the root folder of the Merlin repository.

1. Make any change to the merlin code that you want to test

2. Initialize LocalMerlin:

```bash
make localnet-init
```

The command:

- Builds a local docker image with the latest changes
- Cleans the `$HOME/.merlin-local` folder

3. Start LocalMerlin:

```bash
make localnet-start
```

> Note
>
> You can also start LocalMerlin in detach mode with:
>
> `make localnet-startd`

4. (optional) Add your validator wallet and 9 other preloaded wallets automatically:

```bash
make localnet-keys
```

- These keys are added to your `--keyring-backend test`
- If the keys are already on your keyring, you will get an `"Error: aborted"`
- Ensure you use the name of the account as listed in the table below, as well as ensure you append the `--keyring-backend test` to your txs
- Example: `merlin tx bank send lo-test2 fury1cyyzpxplxdzkeea7kwsydadg87357qnahakaks --keyring-backend test --chain-id LocalMerlin`

5. You can stop chain, keeping the state with

```bash
make localnet-stop
```

6. When you are done you can clean up the environment with:

```bash
make localnet-clean
```

## 2. LocalMerlin - With Mainnet State

Running LocalMerlin with mainnet state is resource intensive and can take a bit of time.
It is recommended to only use this method if you are testing a new feature that must be thoroughly tested before pushing to production.

A few things to note before getting started. The below method will only work if you are using the same version as mainnet. In other words,
if mainnet is on v8.0.0 and you try to do this on a v9.0.0 tag or on main, you will run into an error when initializing the genesis.
(yes, it is possible to create a state exported testnet on a upcoming release, but that is out of the scope of this tutorial)

Additionally, this process requires 64GB of RAM. If you do not have 64GB of RAM, you will get an OOM error.

### Create a mainnet state export

1. Set up a node on mainnet (easiest to use the [get.merlin.zone](https://get.merlin.zone) tool). This will be the node you use to run the state exported testnet, so ensure it has at least 64GB of RAM.

```sh
curl -sL https://get.merlin.zone/install > i.py && python3 i.py
```

2. Once the installer is done, ensure your node is hitting blocks.

```sh
source ~/.profile
journalctl -u merlin.service -f
```

3. Stop your Merlin daemon

```sh
systemctl stop merlin.service
```

4. Take a state export snapshot with the following command:

```sh
cd $HOME
merlin export 2> state_export.json
```

After a while (~15 minutes), this will create a file called `state_export.json` which is a snapshot of the current mainnet state.

### Use the state export in LocalMerlin

1. Copy the `state_export.json` to the `LocalMerlin/state_export` folder within the merlin repo

```sh
cp $HOME/state_export.json $HOME/merlin/tests/LocalMerlin/state_export/
```

6. Ensure you have docker and docker-compose installed:

```sh
# Docker
sudo apt-get remove docker docker-engine docker.io
sudo apt-get update
sudo apt install docker.io -y

# Docker compose
sudo apt install docker-compose -y
```

7. Build the `local:merlin` docker image:

```bash
make localnet-state-export-init
```

The command:

- Builds a local docker image with the latest changes
- Cleans the `$HOME/.merlin` folder

3. Start LocalMerlin:

```bash
make localnet-state-export-start
```

> Note
>
> You can also start LocalMerlin in detach mode with:
>
> `make localnet-state-export-startd`

When running this command for the first time, `local:merlin` will:

- Modify the provided `state_export.json` to create a new state suitable for a testnet
- Start the chain

You will then go through the genesis initialization process. This will take ~15 minutes.
You will then hit the first block (not block 1, but the block number after your snapshot was taken), and then you will just see a bunch of p2p error logs with some KV store logs.
**This will happen for about 1 hour**, and then you will finally hit blocks at a normal pace.

9. On your host machine, add this specific wallet which holds a large amount of fury funds

```sh
MNEMONIC="bottom loan skill merry east cradle onion journey palm apology verb edit desert impose absurd oil bubble sweet glove shallow size build burst effort"
echo $MNEMONIC | merlin keys add wallet --recover --keyring-backend test
```

You now are running a validator with a majority of the voting power with the same state as mainnet state (at the time you took the snapshot)

10. On your host machine, you can now query the state-exported testnet:

```sh
merlin status
```

11. Here is an example command to ensure complete understanding:

```sh
merlin tx bank send wallet fury1nyphwl8p5yx6fxzevjwqunsfqpcxukmtk8t60m 10000000ufury --chain-id testing1 --keyring-backend test
```

12. You can stop chain, keeping the state with

```bash
make localnet-state-export-stop
```

13. When you are done you can clean up the environment with:

```bash
make localnet-state-export-clean
```

Note: At some point, all the validators (except yours) will get jailed at the same block due to them being offline.

When this happens, it may take a little bit of time to process. Once all validators are jailed, you will continue to hit blocks as you did before.
If you are only running the validator for a short time (< 24 hours) you will not experience this.

## LocalMerlin Accounts

LocalMerlin is pre-configured with one validator and 9 accounts with ION and FURY balances.

| Account   | Address                                                                                                | Mnemonic                                                                                                                                                                   |
|-----------|--------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| lo-val    | `fury1phaxpevm5wecex2jyaqty2a4v02qj7qmlmzk5a`<br/>`furyvaloper1phaxpevm5wecex2jyaqty2a4v02qj7qm9v24r6` | `satisfy adjust timber high purchase tuition stool faith fine install that you unaware feed domain license impose boss human eager hat rent enjoy dawn`                    |
| lo-test1  | `fury1cyyzpxplxdzkeea7kwsydadg87357qnahakaks`                                                          | `notice oak worry limit wrap speak medal online prefer cluster roof addict wrist behave treat actual wasp year salad speed social layer crew genius`                       |
| lo-test2  | `fury18s5lynnmx37hq4wlrw9gdn68sg2uxp5rgk26vv`                                                          | `quality vacuum heart guard buzz spike sight swarm shove special gym robust assume sudden deposit grid alcohol choice devote leader tilt noodle tide penalty`              |
| lo-test3  | `fury1qwexv7c6sm95lwhzn9027vyu2ccneaqad4w8ka`                                                          | `symbol force gallery make bulk round subway violin worry mixture penalty kingdom boring survey tool fringe patrol sausage hard admit remember broken alien absorb`        |
| lo-test4  | `fury14hcxlnwlqtq75ttaxf674vk6mafspg8xwgnn53`                                                          | `bounce success option birth apple portion aunt rural episode solution hockey pencil lend session cause hedgehog slender journey system canvas decorate razor catch empty` |
| lo-test5  | `fury12rr534cer5c0vj53eq4y32lcwguyy7nndt0u2t`                                                          | `second render cat sing soup reward cluster island bench diet lumber grocery repeat balcony perfect diesel stumble piano distance caught occur example ozone loyal`        |
| lo-test6  | `fury1nt33cjd5auzh36syym6azgc8tve0jlvklnq7jq`                                                          | `spatial forest elevator battle also spoon fun skirt flight initial nasty transfer glory palm drama gossip remove fan joke shove label dune debate quick`                  |
| lo-test7  | `fury10qfrpash5g2vk3hppvu45x0g860czur8ff5yx0`                                                          | `noble width taxi input there patrol clown public spell aunt wish punch moment will misery eight excess arena pen turtle minimum grain vague inmate`                       |
| lo-test8  | `fury1f4tvsdukfwh6s9swrc24gkuz23tp8pd3e9r5fa`                                                          | `cream sport mango believe inhale text fish rely elegant below earth april wall rug ritual blossom cherry detail length blind digital proof identify ride`                 |
| lo-test9  | `fury1myv43sqgnj5sm4zl98ftl45af9cfzk7nhjxjqh`                                                          | `index light average senior silent limit usual local involve delay update rack cause inmate wall render magnet common feature laundry exact casual resource hundred`       |
| lo-test10 | `fury14gs9zqh8m49yy9kscjqu9h72exyf295afg6kgk`                                                          | `prefer forget visit mistake mixture feel eyebrow autumn shop pair address airport diesel street pass vague innocent poem method awful require hurry unhappy shoulder`     |

## Tests

### Software-upgrade test

To test a software upgrade, you can use the `submit_upgrade_proposal.sh` script located in the `scripts/` folder. This script automatically creates a proposal to upgrade the software to the specified version and votes "yes" on the proposal. Once the proposal passes and the upgrade height is reached, you can update your localmerlin instance to use the new version.

#### Usage 

To use the script:

1. make sure you have a running LocalMerlin instance

2. run the following command:

```bash
./scripts/submit_upgrade_proposal.sh <upgrade version>
```

Replace `<upgrade version>` with the version of the software you want to upgrade to, for example. If no version is specified, the script will default to `v15` version.

The script does the following:

- Creates an upgrade proposal with the specified version and description.
- Votes "yes" on the proposal.

#### Upgrade

Once the upgrade height is reached, you need to update your `localmerlin` instance to use the new software. 

There are two ways to do this:

1. Change the image in the `docker-compose.yml` file to use the new version, and then restart LocalMerlin using `make localnet-start`. For example:

```yaml
services:
  merlin:
    image: <NEW_IMAGE_I_WANT_TO_USE>
    # All this needs to be commented to don't build the image with local changes
    # 
    # build:
    #     context: ../../
    #     dockerfile: Dockerfile
    #     args:
    #     RUNNER_IMAGE: alpine:3.17
    #     GO_VERSION: 1.20
```

2. Checkout the Merlin repository to a different `ref` that includes the new version, and then rebuild and restart LocalMerlin using `make localnet-start`. Make sure to don't delete your `~/.merlin-local` folder.

## FAQ

Q: How do I enable pprof server in localmerlin?

A: everything but the Dockerfile is already configured. Since we use a production Dockerfile in localmerlin, we don't want to expose the pprof server there by default. As a result, if you would like to use pprof, make sure to add `EXPOSE 6060` to the Dockerfile and rebuild the localmerlin image.
