import pandas
import sys
import matplotlib.pyplot as plt

file_paths = [
    "inboundOraclePerformanceTests/hyperledger1Subscription.csv"
    "inboundOraclePerformanceTests/hyperledger2Subscription.csv"
    "inboundOraclePerformanceTests/hyperledger3Subscription.csv"
    "inboundOraclePerformanceTests/ethereum1Subscription.csv"
    "inboundOraclePerformanceTests/ethereum2Subscription.csv"
    "inboundOraclePerformanceTests/ethereum3Subscription.csv"
]
def main():
    file_path = file_paths[0]
    df = pandas.read_csv()
    df.boxplot(column="throughput", by="parallel events")
    plt.xticks(rotation=45)
    plt.title("Boxplot of the artifacts event publishing throughput with one Hyperledger Fabric subscription")
    plt.suptitle("")
    plt.xlabel("Amount of concurrent events")
    plt.ylabel("Throughput (events/second)")
    plt.savefig(file_path[:-4]+"Throughput.png")

    df.boxplot(column="latency", by="parallel events")
    plt.title("Boxplot of the artifacts event publishing latency with one Hyperledger Fabric subscription")
    plt.suptitle("")
    plt.ylabel("Latency in second/event")
    plt.xlabel("Amount of concurrent events")
    plt.savefig(file_path[:-4]+"Latency.png")

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Please pass the path to the csv file like: python visualizePerformanceTestResults.py <path/to/file.csv>")
    main(sys.argv[1])
