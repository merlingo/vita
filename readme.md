# vita
**vita** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Additionally, Ignite CLI offers both Vue and React options for frontend scaffolding:

For a Vue frontend, use: `ignite scaffold vue`
For a React frontend, use: `ignite scaffold react`
These commands can be run within your scaffolded blockchain project. 


For more information see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/vita@latest! | sudo bash
```
`username/vita` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)


# On Docker 
build
$ docker build -f Dockerfile-ubuntu . -t vita_i

check version of ignite
$ docker run --rm -it vita_i ignite version

create persistent container
$ docker create --name vita -i \
    -v $(pwd):/vita -w /vita \
    -p 1317:1317 -p 3000:3000 -p 4500:4500 -p 5000:5000 -p 26657:26657 \
    vita_i
$ docker start vita

$ docker exec -it vita ignite chain serve


docker run --rm -it \
    -v $(pwd):/vita -w /vita \
    -p 1317:1317 -p 3000:3000 -p 4500:4500 -p 5000:5000 -p 26657:26657 \
    --name vita vita_i \
    ignite chain serve


## Testnet Preparation

build with makefile

docker run --rm -it \
    -v $(pwd):/vita \
    -w /vita \
    vita_i \
    make build-with-checksum

Build the image of vitad node
docker build -f prod-sim/Dockerfile-vitad-debian . -t vitad_tstnet

test everything is ok
docker run --rm -it vitad_tstnet help

Build the image of KMS key management system 
docker build -f prod-sim/Dockerfile-tmkms-debian . -t tmkms_i:v0.12.2

Nodes dir creation
mkdir -p prod-sim/kms-mert
mkdir -p prod-sim/node-yasin
mkdir -p prod-sim/sentry-mert
mkdir -p prod-sim/sentry-yigit
mkdir -p prod-sim/val-mert
mkdir -p prod-sim/val-yigit

mkdir -p prod-sim/desk-mert
mkdir -p prod-sim/desk-yigit

basic initialization
echo -e desk-mert'\n'desk-yigit'\n'node-yasin'\n'sentry-mert'\n'sentry-yigit'\n'val-mert'\n'val-yigit \
    | xargs -I {} \
    docker run --rm -i \
    -v $(pwd)/prod-sim/{}:/root/.vita \
    vitad_tstnet \
    init vita

In the authoritative config/genesis.json (desk-mert's):
docker run --rm -it \
    -v $(pwd)/prod-sim/desk-mert:/root/.vita \
    --entrypoint sed \
    vitad_tstnet \
    -i 's/"stake"/"uvita"/g' /root/.vita/config/genesis.json

In all seven config/app.toml:
echo -e desk-mert'\n'desk-yigit'\n'node-yasin'\n'sentry-mert'\n'sentry-yigit'\n'val-mert'\n'val-yigit \
    | xargs -I {} \
    docker run --rm -i \
    -v $(pwd)/prod-sim/{}:/root/.vita \
    --entrypoint sed \
    vitad_tstnet \
    -Ei 's/([0-9]+)stake/\1uvita/g' /root/.vita/config/app.toml

Make sure that config/client.toml mentions vita-1, the chain's name:
echo -e desk-mert'\n'desk-yigit'\n'node-yasin'\n'sentry-mert'\n'sentry-yigit'\n'val-mert'\n'val-yigit \
    | xargs -I {} \
    docker run --rm -i \
    -v $(pwd)/prod-sim/{}:/root/.vita \
    --entrypoint sed \
    vitad_tstnet \
    -Ei 's/^chain-id = .*$/chain-id = "vita-1"/g' \
    /root/.vita/config/client.toml

### Keys
The validator operator keys for Mert and Yigit.
The consensus keys, whether they stay on Yigit's node or are kept inside Mert's KMS.
docker run --rm -it \
    -v $(pwd)/prod-sim/desk-mert:/root/.vita \
    vitad_tstnet \
    keys \
    --keyring-backend file --keyring-dir /root/.vita/keys \
    add mert

Do the same for val-yigit:
docker run --rm -it \
    -v $(pwd)/prod-sim/desk-yigit:/root/.vita \
    vitad_tstnet \
    keys \
    --keyring-backend file --keyring-dir /root/.vita/keys \
    add yigit
echo -n Intra135vis! > prod-sim/desk-yigit/keys/passphrase.txt

Prepare the KMS
docker run --rm -it \
    -v $(pwd)/prod-sim/kms-mert:/root/tmkms \
    tmkms_i:v0.12.2 \
    init /root/tmkms

Import the consensus key
docker run --rm -t \
    -v $(pwd)/prod-sim/val-mert:/root/.vita \
    vitad_tstnet \
    tendermint show-validator \
    | tr -d '\n' | tr -d '\r' \
    > prod-sim/desk-mert/config/pub_validator_key-val-mert.json

copy and move key json files
cp prod-sim/val-mert/config/priv_validator_key.json \
  prod-sim/desk-mert/config/priv_validator_key-val-mert.json
mv prod-sim/val-mert/config/priv_validator_key.json \
  prod-sim/kms-mert/secrets/priv_validator_key-val-mert.json

Import it into the softsign "device" as defined in tmkms.toml:
docker run --rm -i \
    -v $(pwd)/prod-sim/kms-mert:/root/tmkms \
    -w /root/tmkms \
    tmkms_i:v0.12.2 \
    softsign import secrets/priv_validator_key-val-mert.json \
    secrets/val-mert-consensus.key

default private key prevention on val-mert:
cp prod-sim/sentry-mert/config/priv_validator_key.json \
    prod-sim/val-mert/config/

cp prod-sim/sentry-mert/config/priv_validator_key.json \
    prod-sim/val-mert/config

Initial balances
Mert account
MERT=$(echo Intra135vis! | docker run --rm -i \
    -v $(pwd)/prod-sim/desk-mert:/root/.vita \
    vitad_tstnet \
    keys \
    --keyring-backend file --keyring-dir /root/.vita/keys \
    show mert --address)

Have Mert add her initial balance in the genesis:
docker run --rm -it \
    -v $(pwd)/prod-sim/desk-mert:/root/.vita \
    vitad_tstnet \
    genesis add-genesis-account $MERT 1000000000uvita

Now move the genesis file to desk-yigit:
mv prod-sim/desk-mert/config/genesis.json \
    prod-sim/desk-yigit/config/

Have Yigit add his own initial balance:
YIGIT=$(echo Intra135vis! | docker run --rm -i \
    -v $(pwd)/prod-sim/desk-yigit:/root/.vita \
    vitad_tstnet \
    keys \
    --keyring-backend file --keyring-dir /root/.vita/keys \
    show yigit --address)
docker run --rm -it \
    -v $(pwd)/prod-sim/desk-yigit:/root/.vita \
    vitad_tstnet \
    genesis add-genesis-account $YIGIT 500000000uvita

Yigit's stake
cp prod-sim/val-yigit/config/priv_validator_key.json \
    prod-sim/desk-yigit/config/priv_validator_key.json

echo Intra135vis! | docker run --rm -i \
    -v $(pwd)/prod-sim/desk-yigit:/root/.vita \
    vitad_tstnet \
    genesis gentx yigit 40000000uvita \
    --keyring-backend file --keyring-dir /root/.vita/keys \
    --account-number 0 --sequence 0 \
    --chain-id vita-1 \
    --gas 1000000 \
    --gas-prices 0.1uvita

mv prod-sim/desk-yigit/config/genesis.json \
    prod-sim/desk-mert/config/

Mert's stake
echo Intra135vis! | docker run --rm -i \
    -v $(pwd)/prod-sim/desk-mert:/root/.vita \
    vitad_tstnet \
    genesis gentx mert 60000000uvita \
    --keyring-backend file --keyring-dir /root/.vita/keys \
    --account-number 0 --sequence 0 \
    --pubkey $(cat prod-sim/desk-mert/config/pub_validator_key-val-mert.json) \
    --chain-id vita-1 \
    --gas 1000000 \
    --gas-prices 0.1uvita

Genesis assembly
cp prod-sim/desk-yigit/config/gentx/gentx-* \
    prod-sim/desk-mert/config/gentx
docker run --rm -it \
    -v $(pwd)/prod-sim/desk-mert:/root/.vita \
    vitad_tstnet genesis collect-gentxs

validate:
docker run --rm -it \
    -v $(pwd)/prod-sim/desk-mert:/root/.vita \
    vitad_tstnet \
    genesis validate-genesis

Genesis distribution
echo -e desk-yigit'\n'node-yasin'\n'sentry-mert'\n'sentry-yigit'\n'val-mert'\n'val-yigit \
    | xargs -I {} \
    cp prod-sim/desk-mert/config/genesis.json prod-sim/{}/config

Network preparation
docker run --rm -i \
    -v $(pwd)/prod-sim/val-mert:/root/.vita \
    vitad_tstnet \
    tendermint show-node-id

seed node connection setting:
docker run --rm -i \
    -v $(pwd)/prod-sim/sentry-yigit:/root/.vita \
    vitad_tstnet \
    tendermint show-node-id
docker run --rm -i \
    -v $(pwd)/prod-sim/node-yasin:/root/.vita \
    vitad_tstnet \
    tendermint show-node-id

