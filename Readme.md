# OracleFactory


## Quickstart
If you simple want to run the OracleFactory on you machine and Docker is already installed you can run the following command. Replace `<the-base-url-of-the-backend>` with the URL of the backend. If you run it on your local machine that would be `http://localhost:8080` if you run it on a server with IP 1.2.3.4 than `http://1.2.3.4:8080` etc.

```
curl -s "https://raw.githubusercontent.com/PaulsBecks/OracleFactory/master/run_docker_containers.sh" | sudo bash -s -- <the-base-url-of-the-backend>
```

At the current version the code requires root access, but if docker already has root access on your machine you can run:
```
curl -s "https://raw.githubusercontent.com/PaulsBecks/OracleFactory/master/run_docker_containers.sh" | bash -s -- <the-base-url-of-the-backend>
```

This will download and run the backend and frontend.

If you are interested in quickstart the the performance tests, you can execute 

## Longer Start
To run this project on your local machine make sure docker is installed. You can then choose to run it by compiling this repo and a couple of others or by downloading the docker images and running them.

```
docker pull paulsbecks/pub-sub-oracle 
docker pull paulsbecks/pub-sub-oracle-frontend 
docker pull paulsbecks/blf-outbound-oracle

docker network create -d bridge pub-sub-oracle-network
docker run -p 8080:8080 -d --name pub-sub-oracle --network=pub-sub-oracle-network --add-host=host.docker.internal:host-gateway -v /var/run/docker.sock:/var/run/docker.sock paulsbecks/pub-sub-oracle
docker run -p 3000:3000 -d --name pub-sub-oracle-frontend --network=pub-sub-oracle-network paulsbecks/pub-sub-oracle-frontend
```

## From Binary

```
make docker frontend-build oracle-blueprint docker-network docker-start frontend-start
```

## Ethereum Testnet

You can start an Ethereum testnetwork locally that is connected to the same network as follows.

```
docker run --detach -p 8545:8545 -p 7545:7545 --network=pub-sub-oracle-network --name eth-test-net trufflesuite/ganache-cli:latest --accounts 10  --blockTime 2 --seed SubscriptionFramework
```

If you want to deploy the on-chain publish-subscribe oracle manually, you can switch to the `ethereumOnChainSubscription` folder and execute `truffle migrate`.

## Developer Prerequirements

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

## Interface documentation

You can find more information about the interfaces by starting the oracle factory locally and visiting http://localhost:8080/swagger/index.html or you can visit https://oracles.work/api/swagger/index.html.

To update the swagger documentation run

```
swagg init
```

## Running Testscenarios

To run the visual test scenarios, where you need the GUI, you have to first setup the test environment:

```
make init-visual-test-setup
```

Now you have a running ethereum blockchain and hyperledger fabric blockchain with test smart contracts installed. Also an n8n server is running for you to create workflows with. At http://localhost:5678


If you want to just run the performance test, run

```
make init-test-setup
```

or if some docker already, you have to cleanup first and re run it with

```
make test-setup
```

Next, to execute the tests run

```
make performance-test
```

To prune all created docker networks, images and containers run
```
make prune-test-setup
```
