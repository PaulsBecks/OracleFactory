#!/bin/bash

# run tests
#cd ./pubSubOraclePerformanceTests; go build; ./pubSubOraclePerformanceTests; cd ..
cd ./outboundOraclePerformanceTests; go build; ./outboundOraclePerformanceTests; cd ..
# outbound oracle test visualization
#python3 visualizePerformanceTestResults.py pubSubOraclePerformanceTests/ethereumMintTokenTest.csv
#python3 visualizePerformanceTestResults.py pubSubOraclePerformanceTests/ethereumTransferTokenTest.csv
#python3 visualizePerformanceTestResults.py pubSubOraclePerformanceTests/hyperledgerCreateAssetTest.csv
# outbound oracle test visualization
#python3 visualizePerformanceTestResults.py outboundOraclePerformanceTests/ethereumTransferTokenTest.csv
#python3 visualizePerformanceTestResults.py outboundOraclePerformanceTests/hyperledgerCreateAssetTest.csv