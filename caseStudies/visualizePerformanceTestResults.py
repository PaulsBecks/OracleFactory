import pandas
import sys
import matplotlib.pyplot as plt

def main(file_path):
    df = pandas.read_csv(file_path)
    df.boxplot(column="Throughput", by="ParallelWorkers")
    plt.xticks(rotation=45)
    plt.savefig(file_path[:-4]+"Throughput.png")

    df.boxplot(column="Latency", by="ParallelWorkers")
    plt.savefig(file_path[:-4]+"Latency.png")

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Please pass the path to the csv file like: python visualizePerformanceTestResults.py <path/to/file.csv>")
    main(sys.argv[1])