#!/bin/bash

# run tests
cd ./inboundSubscriptionPerformanceTests; go build; ./inboundSubscriptionPerformanceTests; cd ..
cd ./outboundSubscriptionPerformanceTests; go build;./outboundSubscriptionPerformanceTests;cd ..
# outbound oracle test visualization
python3 visualizePerformanceTestResults.py inboundSubscriptionPerformanceTests/ethereumMintTokenTest.csv
python3 visualizePerformanceTestResults.py inboundSubscriptionPerformanceTests/hyperledgerCreateAssetTest.csv
# outbound oracle test visualization
python3 visualizePerformanceTestResults.py outboundSubscriptionPerformanceTests/test.csv