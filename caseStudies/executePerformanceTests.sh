#!/bin/bash

# run tests
cd ./inboundOraclePerformanceTests; go build; ./inboundOraclePerformanceTests; cd ..
cd ./outboundOraclePerformanceTests; go build;./outboundOraclePerformanceTests;cd ..
# outbound oracle test visualization
python3 visualizePerformanceTestResults.py inboundOraclePerformanceTests/ethereumMintTokenTest.csv
python3 visualizePerformanceTestResults.py inboundOraclePerformanceTests/hyperledgerCreateAssetTest.csv
# outbound oracle test visualization
python3 visualizePerformanceTestResults.py outboundOraclePerformanceTests/test.csv