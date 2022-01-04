network_name=pub-sub-oracle-network
current_dir = $(shell pwd)

TEST_SMART_CONTRACT=${smart_contract}

all: docker-network eth-testnet eth-test-contract n8n

docker:
	docker build -t "paulsbecks/pub-sub-oracle" .

docker-start:
	docker run -p 8080:8080 -d --name pub-sub-oracle --network=$(network_name) --add-host=host.docker.internal:host-gateway -v /var/run/docker.sock:/var/run/docker.sock paulsbecks/pub-sub-oracle

docker-test-start:
	docker run -p 8080:8080 -d --name pub-sub-oracle --network=$(network_name) --add-host=host.docker.internal:host-gateway --env ENV=PERFORMANCE_TEST -v /var/run/docker.sock:/var/run/docker.sock paulsbecks/pub-sub-oracle

docker-stop:
	docker rm $$(docker stop $$(docker ps -a -q --filter ancestor="paulsbecks/pub-sub-oracle" --format="{{.ID}}"))

docker-update: docker-stop docker docker-start

oracle-blueprint:
	docker build -t "paulsbecks/blf-outbound-oracle" ./oracleBlueprint

eth-testnet:
	docker run --detach -p 8545:8545 -p 7545:7545 --network=$(network_name) --name eth-test-net trufflesuite/ganache-cli:latest --accounts 10  --blockTime 2 --seed OracleFramework
	sleep 20

eth-testnet-stop:
	docker rm $$(docker stop $$(docker ps -a -q --filter ancestor="truffelsuite/ganache-cli:latest" --format="{{.ID}}"))

install-eth-contract:
	cd ethereumOnChainOracle; truffle migrate; cd ../caseStudies/testContract; truffle migrate;

docker-network:
	docker network create -d bridge $(network_name) || true

fmt:
	go fmt ./...

frontend-build:
	cd ./frontend; docker build -t "paulsbecks/pub-sub-oracle-frontend" .; cd ..

frontend-start:
	docker run --detach -p 3000:3000 --network=$(network_name) --name pub-sub-oracle-frontend paulsbecks/pub-sub-oracle-frontend

frontend-stop:
	docker rm $$(docker stop $$(docker ps -a -q --filter ancestor="pub-sub-oracle-frontend" --format="{{.ID}}"))

frontend-update: frontend-stop frontend-build frontend-start

n8n:
	docker run --detach --rm --name n8n -p 5678:5678 -v ${current_dir}/.n8n:/home/node/.n8n --network=$(network_name) n8nio/n8n

init-test-setup: docker-network eth-testnet hyperledger-testnet oracle-blueprint docker docker-test-start

init-visual-test-setup: init-test-setup frontend-build frontend-start n8n

prune-test-setup:
	docker stop $$(docker ps -aq) || $$(true)
	docker network prune -f
	docker container prune -f

register:
	selenium-side-runner "caseStudies/${TEST_SMART_CONTRACT}/register.side"

create-oracle-template:
	selenium-side-runner "caseStudies/${TEST_SMART_CONTRACT}/create_smart_contract.side"

create-oracle:
	echo "Not implemented yet"

use-case: register create-oracle-template create-oracle

evaluation:
	echo "Evaluation not implemented yet!"

case-study: init-test-setup install-eth-contract use-case evaluation prune-test-setup

hyperledger-testnet:
	curl -sSL https://bit.ly/2ysbOFE | bash -s
	cd fabric-samples/test-network; ./network.sh down; COMPOSE_PROJECT_NAME=docker ./network.sh up createChannel -c mychannel -ca; ./network.sh deployCC -ccs 1  -ccv 1 -ccep "OR('Org1MSP.peer','Org2MSP.peer')" -ccl go -ccp ../../hyperledgerTestContract -ccn test-contract; ./network.sh deployCC -ccs 1  -ccv 1 -ccep "OR('Org1MSP.peer','Org2MSP.peer')" -ccl go -ccp ../../hyperledgerTestContract -ccn test-contract2; ./network.sh deployCC -ccs 1  -ccv 1 -ccep "OR('Org1MSP.peer','Org2MSP.peer')" -ccl go -ccp ../../hyperledgerTestContract -ccn test-contract3; ./network.sh deployCC -ccs 1  -ccv 1 -ccep "OR('Org1MSP.peer','Org2MSP.peer')" -ccl go -ccp ../../hyperledgerOnChainOracle -ccn on-chain-oracle; cd ../..
	cp fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/connection-org1.yaml . 
	cp fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/keystore/* hyperledger_key
	cp fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/signcerts/cert.pem hyperledger_cert
	docker network connect pub-sub-oracle-network peer0.org1.example.com
	docker network connect pub-sub-oracle-network peer0.org2.example.com
	docker network connect pub-sub-oracle-network ca_org1
	docker network connect pub-sub-oracle-network ca_org2
	docker network connect pub-sub-oracle-network orderer.example.com

test-setup: prune-test-setup init-test-setup

performance-test:
	sleep 3
	cd caseStudies; sh ./executePerformanceTests.sh
	echo "Performance test started in background"

setup-and-test: test-setup performance-test

hyperledger-oracle-test: docker-network hyperledger-testnet 
	
hyperledger-install-contracts:
	sh ./install-oracle-on-hyperledger.sh
	sh ./install-test-contract.sh

make swagger:
	swag init --parseDependency --parseDepth 2