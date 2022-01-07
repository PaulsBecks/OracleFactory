import pandas
import sys
import matplotlib.pyplot as plt

file_paths = [
    "inboundOraclePerformanceTests/hyperledger1SubscriptionSameChoice.csv",
    "inboundOraclePerformanceTests/hyperledger2SubscriptionSameChoice.csv",
    "inboundOraclePerformanceTests/hyperledger3SubscriptionSameChoice.csv",
    "inboundOraclePerformanceTests/hyperledger1SubscriptionDifferentChoice.csv",
    "inboundOraclePerformanceTests/hyperledger2SubscriptionDifferentChoice.csv",
    "inboundOraclePerformanceTests/hyperledger3SubscriptionDifferentChoice.csv",
    "inboundOraclePerformanceTests/ethereum1subscription.csv",
    "inboundOraclePerformanceTests/ethereum2subscription.csv",
    "inboundOraclePerformanceTests/ethereum3subscription.csv",
    "inboundOraclePerformanceTests/ethereum1subscriptionDifferentChoice.csv",
    "inboundOraclePerformanceTests/ethereum2subscriptionDifferentChoice.csv",
    "inboundOraclePerformanceTests/ethereum3subscriptionDifferentChoice.csv",
]

def read_merged_files(all_files):
    li = []
    for filename in all_files:
        df = pandas.read_csv(filename)
        li.append(df)
    return pandas.concat(li)


def main():
    file_path = file_paths[0]
    df = pandas.read_csv(file_path)
    df.boxplot(column="throughput", by="parallel events")
    plt.xticks(rotation=45)
    plt.title("Boxplot of the artifacts event publishing throughput \n with one Hyperledger Fabric subscription")
    plt.suptitle("")
    plt.xlabel("Amount of concurrent events")
    plt.ylabel("Throughput (events/second)")
    plt.savefig(file_path[:-4]+"Throughput.png")

    df.boxplot(column="latency", by="parallel events")
    plt.title("Boxplot of the artifacts event publishing latency \n with one Hyperledger Fabric subscription")
    plt.suptitle("")
    plt.ylabel("Latency in seconds/event")
    plt.xlabel("Amount of concurrent events")
    plt.ylim(ymin=0)
    plt.savefig(file_path[:-4]+"Latency.png")

    df = read_merged_files(file_paths[:4])
    df = df[df["parallel events"] == 1]
    df.boxplot(column="latency", by="subscriptions")
    plt.title("Boxplot of the artifacts event publishing latency \n for Hyperledger Fabric subscriptions with one concurrent event")
    plt.suptitle("")
    plt.ylabel("Latency in seconds/event")
    plt.xlabel("Amount of oracle subscriptions")
    plt.ylim(ymin=0)
    plt.savefig("inboundOraclePerformanceTests/hyperledgerSubscriptionsLatency.png")

    plt.xticks(rotation=45)
    df.boxplot(column="throughput", by="subscriptions")
    plt.title("Boxplot of the artifacts event publishing throughput \n for Hyperledger Fabric subscriptions with one concurrent event ")
    plt.suptitle("")
    plt.xlabel("Amount of concurrent events")
    plt.ylabel("Throughput (events/second)")
    plt.ylim(ymin=0)
    plt.savefig("inboundOraclePerformanceTests/hyperledgerSubscriptionsThroughput.png")

    file_path = file_paths[3]
    df = pandas.read_csv(file_path)
    df.boxplot(column="throughput", by="parallel events")
    plt.xticks(rotation=45)
    plt.title("Boxplot of the artifacts event publishing throughput \n with one Hyperledger Fabric subscription")
    plt.suptitle("")
    plt.xlabel("Amount of concurrent events")
    plt.ylabel("Throughput (events/second)")
    plt.ylim(ymin=0)
    plt.savefig(file_path[:-4]+"Throughput.png")

    df.boxplot(column="latency", by="parallel events")
    plt.title("Boxplot of the artifacts event publishing latency \n with one Hyperledger Fabric subscription")
    plt.suptitle("")
    plt.ylabel("Latency in seconds/event")
    plt.xlabel("Amount of concurrent events")
    plt.ylim(ymin=0)
    plt.savefig(file_path[:-4]+"Latency.png")

    df = read_merged_files(file_paths[3:6])
    df = df[df["parallel events"] == 1]
    df.boxplot(column="latency", by="subscriptions")
    plt.title("Boxplot of the artifacts event publishing latency \n for Hyperledger Fabric subscriptions with one concurrent event")
    plt.suptitle("")
    plt.ylabel("Latency in seconds/event")
    plt.xlabel("Amount of oracle subscriptions")
    plt.ylim(ymin=0)
    plt.savefig("inboundOraclePerformanceTests/hyperledgerSubscriptionsLatency.png")

    plt.xticks(rotation=45)
    df.boxplot(column="throughput", by="subscriptions")
    plt.title("Boxplot of the artifacts event publishing throughput \n for Hyperledger Fabric subscriptions with one concurrent event ")
    plt.suptitle("")
    plt.xlabel("Amount of concurrent events")
    plt.ylabel("Throughput (events/second)")
    plt.ylim(ymin=0)
    plt.savefig("inboundOraclePerformanceTests/hyperledgerSubscriptionsThroughput.png")

    file_path = file_paths[6]
    df = pandas.read_csv(file_path)
    df.boxplot(column="throughput", by="parallel events")
    plt.xticks(rotation=45)
    plt.title("Boxplot of the artifacts event publishing throughput \n with one Hyperledger Fabric subscription")
    plt.suptitle("")
    plt.xlabel("Amount of concurrent events")
    plt.ylabel("Throughput (events/second)")
    plt.ylim(ymin=0)
    plt.savefig(file_path[:-4]+"Throughput.png")

    df.boxplot(column="latency", by="parallel events")
    plt.title("Boxplot of the artifacts event publishing latency \n  with one Ethereum subscription")
    plt.suptitle("")
    plt.ylabel("Latency in seconds/event")
    plt.xlabel("Amount of concurrent events")
    plt.ylim(ymin=0)
    plt.savefig(file_path[:-4]+"Latency.png")
    
    df = read_merged_files(file_paths[6:9])
    df = df[df["parallel events"] == 1]
    df.boxplot(column="latency", by="subscriptions")
    plt.title("Boxplot of the artifacts event publishing latency \n for Ethereum subscriptions with one concurrent event")
    plt.suptitle("")
    plt.ylabel("Latency in seconds/event")
    plt.xlabel("Amount of oracle subscriptions")
    plt.ylim(ymin=0)
    plt.savefig("inboundOraclePerformanceTests/ethereumSubscriptionsLatency.png")

    df.boxplot(column="throughput", by="subscriptions")
    plt.xticks(rotation=45)
    plt.title("Boxplot of the artifacts event publishing throughput \n for Ethereum subscriptions with one concurrent event ")
    plt.suptitle("")
    plt.xlabel("Amount of concurrent events")
    plt.ylabel("Throughput (events/second)")
    plt.ylim(ymin=0)
    plt.savefig("inboundOraclePerformanceTests/ethereumSubscriptionsThroughput.png")

if __name__ == "__main__":
    main()
