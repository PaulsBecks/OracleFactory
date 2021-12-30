npm --silent install 
cp -f truffle-config.blueprint.js truffle-config.js
sed -i.bak "s#PRIVATE_KEY#$1#g" "truffle-config.js"
sed -i.bak "s#ETHEREUM_BLOCKCHAIN_URL#$2#g" "truffle-config.js"
truffle migrate --network prod  grep "PUBSUB_ORACLE_ADDRESS" | awk '{print $2}'
rm truffle-config.js.bak