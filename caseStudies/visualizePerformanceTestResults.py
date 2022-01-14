import pandas
import sys
import matplotlib.pyplot as plt
import matplotlib

matplotlib.rcParams.update({'font.size': 22})

def main(file_path):
    df = pandas.read_csv(file_path)
    df.boxplot(column="throughput", by="parallel events")
    plt.title("Boxplot of the artifacts event publishing throughput \n with one subscription")
    plt.suptitle("")
    plt.xlabel("Amount of concurrent events")
    plt.ylabel("Throughput (events/second)")
    plt.ylim(ymin=0)
    plt.savefig(file_path[:-4]+"Throughput.png")

    df.boxplot(column="latency", by="parallel events")
    plt.title("Boxplot of the latency by parallel events")
    plt.suptitle("")
    plt.savefig(file_path[:-4]+"Latency.png")

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Please pass the path to the csv file like: python visualizePerformanceTestResults.py <path/to/file.csv>")
    main(sys.argv[1])
