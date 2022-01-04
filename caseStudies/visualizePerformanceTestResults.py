import pandas
import sys
import matplotlib.pyplot as plt

def main(file_path):
    print(file_path)
    df = pandas.read_csv(file_path)
    df.boxplot(column="throughput", by="parallel events")
    plt.xticks(rotation=45)
    plt.title("Boxplot of the throughput of concurrent events")
    plt.suptitle("")
    plt.savefig(file_path[:-4]+"Throughput.png")
    plt.xlabel("Amount of concurrent events")
    plt.ylabel("Throughput in events/second")

    df.boxplot(column="latency", by="parallel events")
    plt.title("Boxplot of the latency of concurrent events")
    plt.suptitle("")
    plt.savefig(file_path[:-4]+"Latency.png")
    plt.xlabel("Amount of concurrent events")
    plt.ylabel("Latency in second/event")

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Please pass the path to the csv file like: python visualizePerformanceTestResults.py <path/to/file.csv>")
    main(sys.argv[1])
