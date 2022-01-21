network_name=oracle-factory-network
current_dir = $(shell pwd)

TEST_SMART_CONTRACT=${smart_contract}

all: docker-network eth-testnet eth-test-contract n8n

docker:
	docker build -t "paulsbecks/pub-sub-oracle" .

docker-start:
	docker run -p 8080:8080 -d --name oracle-factory --network=$(network_name) --add-host=host.docker.internal:host-gateway -v /var/run/docker.sock:/var/run/docker.sock paulsbecks/pub-sub-oracle

docker-test-start:
	docker run -p 8080:8080 -d --name oracle-factory --network=$(network_name) --add-host=host.docker.internal:host-gateway --env ENV=PERFORMANCE_TEST -v /var/run/docker.sock:/var/run/docker.sock paulsbecks/pub-sub-oracle

docker-stop:
	docker rm $$(docker stop $$(docker ps -a -q --filter ancestor="paulsbecks/pub-sub-oracle" --format="{{.ID}}"))

docker-update: docker-stop docker docker-start

oracle-blueprint:
	docker build --no-cache -t "oracle_blueprint" ./oracleBlueprint

eth-testnet:
	docker run --detach -p 8545:8545 -p 7545:7545 --network=$(network_name) --name eth-test-net trufflesuite/ganache-cli:latest --accounts 10 --blockTime 2 --seed OracleFramework
	sleep 20
	cd caseStudies/token; truffle migrate; cd ../..

eth-testnet-stop:
	docker rm $$(docker stop $$(docker ps -a -q --filter ancestor="truffelsuite/ganache-cli:latest" --format="{{.ID}}"))

install-eth-contract:
	cd "caseStudies/${TEST_SMART_CONTRACT}"; truffle compile; truffle migrate

docker-network:
	docker network create -d bridge $(network_name) || true

fmt:
	go fmt ./...

frontend-build:
	cd ./frontend; docker build -t "paulsbecks/pub-sub-oracle_frontend" .; cd ..

frontend-start:
	docker run --detach -p 3000:3000 --network=$(network_name) --name oracle-factory-frontend paulsbecks/pub-sub-oracle_frontend

frontend-stop:
	docker rm $$(docker stop $$(docker ps -a -q --filter ancestor="paulsbecks/pub-sub-oracle_frontend" --format="{{.ID}}"))

frontend-update: frontend-stop frontend-build frontend-start

n8n:
	docker run --detach --rm --name n8n -p 5678:5678 -v ${current_dir}/.n8n:/home/node/.n8n --network=$(network_name) n8nio/n8n

init-test-setup: docker-network eth-testnet oracle-blueprint docker docker-test-start

init-visual-test-setup: init-test-setup frontend-build frontend-start n8n

prune-test-setup:
	docker stop $$(docker ps -aq) || $$(true)
	docker network prune -f
	docker container prune -f

hyperledger-testnet:
	curl -sSL https://bit.ly/2ysbOFE | bash -s
	cd fabric-samples/test-network; ./network.sh down; ./network.sh up createChannel -c mychannel -ca; ./network.sh deployCC -ccs 1  -ccv 1 -ccep "OR('Org1MSP.peer','Org2MSP.peer')" -ccl javascript -ccp ./../asset-transfer-events/chaincode-javascript/ -ccn events; cd ../..
	cp fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/connection-org1.yaml . 
	cp fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/keystore/* hyperledger_key
	cp fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/signcerts/cert.pem hyperledger_cert
	docker network connect oracle-factory-network peer0.org1.example.com
	docker network connect oracle-factory-network peer0.org2.example.com
	docker network connect oracle-factory-network ca_org1
	docker network connect oracle-factory-network ca_org2
	docker network connect oracle-factory-network orderer.example.com

test-setup: prune-test-setup init-test-setup

performance-test:
	sleep 3
	cd caseStudies; sh ./executePerformanceTests.sh
	echo "Performance test started in background"

setup-and-test: test-setup performance-test
