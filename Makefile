network_name=oracle-factory-network

all: docker-network oracle-blueprint

docker:
	docker build -t "oracle_factory" .

docker-start:
	docker run -p 8080:8080 -d --name oracle-factory --network=$(network_name) -v /var/run/docker.sock:/var/run/docker.sock oracle_factory

docker-stop:
	docker rm $$(docker stop $$(docker ps -a -q --filter ancestor="oracle_factory" --format="{{.ID}}"))

docker-update: docker-stop docker docker-start

oracle-blueprint:
	docker build -t "oracle_blueprint" ./oracleBlueprint

testnet:
	docker run --detach -p 8545:8545 -p 7545:7545 --network=$(network_name) --name eth-test-net trufflesuite/ganache-cli:latest --accounts 10 --seed OracleFramework

testnet-stop:
	docker rm $$(docker stop $$(docker ps -a -q --filter ancestor="truffelsuite/ganache-cli" --format="{{.ID}}"))

docker-network:
	docker network create -d bridge $(network_name)

deploy-test-contract:
	cd testContract; truffle compile; truffle migrate

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
