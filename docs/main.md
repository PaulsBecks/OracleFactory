# Oracle Factory

The oracle factory shall provide an artifact to create publish-subscribe oracles. It's based on the theories of this [paper](https://www.researchgate.net/profile/Jan-Ladleif/publication/344080716_External_Data_Monitoring_Using_Oracles_in_Blockchain-Based_Process_Execution/links/6062bb4ba6fdccbfea15da38/External-Data-Monitoring-Using-Oracles-in-Blockchain-Based-Process-Execution.pdf) by Ladleif, Weber and Weske. We understand a publish-subscribe oracle as a piece of software, that allows smart contracts to subscribe events, web services to publish events and to match published events to subscribers. 

:danger: This piece of software is not yet ready for production.

## Startup

To run this project on your local machine make sure docker is installed. You can then choose to run it by compiling this repo and a couple of others or by downloading the docker images and running them. 

:danger: right now sudo right are necessary to run this project on you machine.

```
docker pull paulsbecks/pub-sub-oracle 
docker pull paulsbecks/pub-sub-oracle-frontend 
docker pull paulsbecks/blf-outbound-oracle

docker network create -d bridge pub-sub-oracle-network
docker run -p 8080:8080 -d --name pub-sub-oracle --network=pub-sub-oracle-network --add-host=host.docker.internal:host-gateway -v /var/run/docker.sock:/var/run/docker.sock paulsbecks/pub-sub-oracle
docker run -p 3000:3000 -d --name pub-sub-oracle-frontend --network=pub-sub-oracle-network paulsbecks/pub-sub-oracle-frontend

```

## Usage

After the application was started you have to create an account to use. Next, you create a new blockchain connector, by providing all necessary information to connect and authenticate against your blockchain and account. 

### Creating Oracles
Finally you can start the blockchain oracle for this connection. You can create pub-sub oracles for ethereum and hyperledger. There are two ways to run a pub-sub oracle - either on or off the blockchain. This means the code, that manages subscriptions and events runs on  or off the blockchain. Both can have advantages and disadvantages on latency, throughput and cost as summorized in section performance. From now on the oracle will be displayed on the dashboard.

### Creating Publishers
You can also create publishing endpoints, for the events that shall be received. You can decide whether others can see the endpoint as well.

### Subscribing to Off-Chain Oracle
To subscribe a smart contract to an off-chain oracle you need to add the following event emission to your smart contract code. 

#### Ethereum
#### Hyperledger

### Subscribing to On-Chain Oracle

#### Ethereum
#### Hyperledger

## Performance

Performance of the artifact Ethereum one subscription per event.
|  | On Chain | Off Chain |
| --- | --- | --- |
| Cost | - | + |
| Latency | o | o |
| Throughput | o | o |


Performance of the artifact Ethereum many subscriptions per event.

|  | On Chain | Off Chain |
| --- | --- | --- |
| Cost | + | - |
| Latency | + | - |
| Throughput | + | - |

Hyperledger one subscriptions per event.

| Add Table here

Hyperledger many subscriptions per event.

| Add Table here

## Motivation

When implementing processes on a blockchain that contain monitoring points, the current oracle solutions are not sufficient, as data may be missed or fetched in the wrong order. To tackle this problem Weber et al designed pub-sub oracles which are implemented in this project.

## Design
