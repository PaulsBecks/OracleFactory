network_name=oracle-factory-network

TEST_SMART_CONTRACT=${smart_contract}

all: docker-network eth-testnet eth-test-contract n8n

docker:
	docker build -t "oracle_factory" .

docker-start:
	docker run -p 8080:8080 -d --name oracle-factory --network=$(network_name) -v /var/run/docker.sock:/var/run/docker.sock oracle_factory

docker-stop:
	docker rm $$(docker stop $$(docker ps -a -q --filter ancestor="oracle_factory" --format="{{.ID}}"))

docker-update: docker-stop docker docker-start

oracle-blueprint:
	docker build -t "oracle_blueprint" ./oracleBlueprint

eth-testnet:
	docker run --detach -p 8545:8545 -p 7545:7545 --network=$(network_name) --name eth-test-net trufflesuite/ganache-cli:latest --accounts 10 --seed OracleFramework

eth-testnet-stop:
	docker rm $$(docker stop $$(docker ps -a -q --filter ancestor="truffelsuite/ganache-cli" --format="{{.ID}}"))

install-eth-contract:
	cd "caseStudies/${TEST_SMART_CONTRACT}"; truffle compile; truffle migrate

docker-network:
	docker network create -d bridge $(network_name) || true

fmt:
	go fmt ./...

frontend-build:
	pushd ./frontend; yarn build; docker build -t "oracle_factory_frontend" .; popd

frontend-start:
	docker run --detach -p 3000:3000 --network=$(network_name) --name oracle-factory-frontend oracle_factory_frontend

frontend-stop:
	docker rm $$(docker stop $$(docker ps -a -q --filter ancestor="oracle_factory_frontend" --format="{{.ID}}"))

frontend-update: frontend-stop frontend-build frontend-start

n8n:
	docker run --detach --rm --name n8n -p 5678:5678 -v ~/.n8n:/home/node/.n8n --network=$(network_name) n8nio/n8n

init-test-setup: docker-network eth-testnet hyperledger-testnet oracle-blueprint docker docker-start frontend-build frontend-start n8n

prune-test-setup:
	docker stop $$(docker ps -aq)
	docker network prune -f
	docker container prune -f

register:
	selenium-side-runner "caseStudies/${TEST_SMART_CONTRACT}/register.side"

create-oracle-template:
	selenium-side-runner "caseStudies/${TEST_SMART_CONTRACT}/create_oracle_template.side"

create-oracle:
	echo "Not implemented yet"

use-case: register create-oracle-template create-oracle

evaluation:
	echo "Evaluation not implemented yet!"

case-study: init-test-setup install-eth-contract use-case evaluation prune-test-setup

hyperledger-testnet:
	curl -sSL https://bit.ly/2ysbOFE | bash -s -- 2.2.2 1.4.9
	cd fabric-samples/test-network; ./network.sh down; ./network.sh up; ./network.sh deployCC -ccep "OR('Org1MSP.peer','Org2MSP.peer')"  -ccl java -ccp ./../asset-transfer-events/chaincode-java/ -ccn asset-transfer-events-java; cd ../..
