# KiiChainV3 Network

![Docker!](https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white)
![Go!](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Apache License](https://img.shields.io/badge/Apache%20License-D22128?style=for-the-badge&logo=Apache&logoColor=white)

## Kiichaind Install

Binary can be Installed for Linux or Mac (check releases for Windows)

Specify a version to install if desired. 

```sh
git clone -b <latest-release-tag> https://github.com/KiiChain/kiichainV3.git
cd kiichainV3 && make install
```

Note: Depending on your `go` setup you may need to add `$GOPATH/bin` to your `$PATH`.

```
export PATH=$PATH:$(go env GOPATH)/bin
```

## Run a Local Network
To run a local node for testing purposes, execute the following commands:
```
make install
kiichaind start
```

## Run a Local Network using scripts

```
./docker/localnode/scripts/deploy.sh
./docker/localnode/scripts/step0_build.sh
./docker/localnode/scripts/step1_configure_init.sh
./docker/localnode/scripts/step2_genesis.sh
./docker/localnode/scripts/step3_add_validator_to_genesis.sh
./docker/localnode/scripts/step4_config_override.sh
./docker/localnode/scripts/step5_start_kiichaind.sh
```

## Run a node using docker

### Run
```
docker compose pull
docker compose up
```

run `docker compose up -d` to run detached.

*NOTE:* Don't forget to pull the images first, to ensure that you're using the latest images.

### See logs
`docker compose logs -f`
