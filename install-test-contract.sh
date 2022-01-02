cd fabric-samples/test-network
export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
peer version

export CHANNEL=mychannel
export CHAINCODE_NAME=test-contract
export CHAINCODE_PATH=../../hyperledgerTestContract/$CHAINCODE_NAME.tar.gz
export CA_FILE_PATH=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
export CORE_PEER_TLS_ENABLED=true

#org1
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

peer lifecycle chaincode install $CHAINCODE_PATH
export CC_PACKAGE_ID=$(peer lifecycle chaincode queryinstalled | grep "Package ID: test-contract" | sed -E 's/.*(test-contract:[0-9a-f]+),.*/\1/') 
echo $CC_PACKAGE_ID
peer lifecycle chaincode approveformyorg -o localhost:7050 --init-required --signature-policy "OR('Org1MSP.peer','Org2MSP.peer')" --ordererTLSHostnameOverride orderer.example.com --channelID $CHANNEL  --name $CHAINCODE_NAME --version 1.0 --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile $CA_FILE_PATH

#org2
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
export CORE_PEER_ADDRESS=localhost:9051

peer lifecycle chaincode install $CHAINCODE_PATH
export CC_PACKAGE_ID=$(peer lifecycle chaincode queryinstalled | grep "Package ID: test-contract" | sed -E 's/.*(test-contract:[0-9a-f]+),.*/\1/') 
echo $CC_PACKAGE_ID
peer lifecycle chaincode approveformyorg -o localhost:7050 --signature-policy "OR('Org1MSP.peer','Org2MSP.peer')" --ordererTLSHostnameOverride orderer.example.com --channelID $CHANNEL --name $CHAINCODE_NAME --version 1.0 --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile $CA_FILE_PATH

peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL --name $CHAINCODE_NAME --version 1.0 --sequence 1 --tls --cafile $CA_FILE_PATH --output json
peer lifecycle chaincode commit -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID $CHANNEL --name $CHAINCODE_NAME --version 1.0 --sequence 1 --tls --cafile $CA_FILE_PATH --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt

peer lifecycle chaincode querycommitted --channelID $CHANNEL --name $CHAINCODE_NAME --cafile $CA_FILE_PATH

export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile $CA_FILE_PATH -C $CHANNEL -n $CHAINCODE_NAME --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE -c '{"function":"InitLedger","Args":[]}'