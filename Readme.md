# OracleFactory

To run this project on your local machine make sure docker is installed. You can then choose to run it by compiling this repo and a couple of others or by downloading the docker images and running them.

```
docker pull paulsbecks/pub-sub-oracle 
docker pull paulsbecks/pub-sub-oracle-frontend 
docker pull paulsbecks/blf-outbound-oracle

docker network create -d bridge pub-sub-oracle-network
docker run -p 8080:8080 -d --name pub-sub-oracle --network=pub-sub-oracle-network --add-host=host.docker.internal:host-gateway -v /var/run/docker.sock:/var/run/docker.sock paulsbecks/pub-sub-oracle
docker run -p 3000:3000 -d --name pub-sub-oracle-frontend --network=pub-sub-oracle-network paulsbecks/pub-sub-oracle-frontend

```


# Ethereum Testnet

You can start an Ethereum testnetwork locally that is connected to the same network as follows.

```
docker run --detach -p 8545:8545 -p 7545:7545 --network=pub-sub-oracle-network --name eth-test-net trufflesuite/ganache-cli:latest --accounts 10  --blockTime 2 --seed OracleFramework
```

If you want to deploy the on-chain publish-subscribe oracle manually, you can switch to the `ethereumOnChainOracle` folder and execute `truffle migrate`.

# Prerequirements

To develop and run the different components locally a bunch of thing need to be installed

* node 16.5
* golang 1.6
* npm
* Make
* Docker

Install requirements on ubuntu by running

```
sudo sh ./install.sh
```

# Interface documentation

You can find more information about the interfaces by starting the oracle factory locally and visiting http://localhost:8080/swagger/index.html