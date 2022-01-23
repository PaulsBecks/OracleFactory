docker pull paulsbecks/pub-sub-oracle 
docker pull paulsbecks/pub-sub-oracle-frontend 
docker pull paulsbecks/blf-outbound-oracle

docker network create -d bridge pub-sub-oracle-network
docker run -p 8080:8080 -d --name pub-sub-oracle --network=pub-sub-oracle-network --add-host=host.docker.internal:host-gateway -v /var/run/docker.sock:/var/run/docker.sock paulsbecks/pub-sub-oracle
docker run -p 3000:3000 -d --name pub-sub-oracle-frontend --network=pub-sub-oracle-network paulsbecks/pub-sub-oracle-frontend