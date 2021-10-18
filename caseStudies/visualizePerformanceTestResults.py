import pandas
import sys
import matplotlib.pyplot as plt

def main(file_path):
    df = pandas.read_csv(file_path)
    grouped = df.groupby("ParallelWorkers")
    latencyFig = grouped.boxplot("Latency")
    latencyFig.savefig(file_path[:-3]+"png")

if __name__ == "__main__":
    if len(sys.argv) != 1:
        print("Please pass the path to the csv file like: python visualizePerformanceTestResults.py <path/to/file.csv>")
    main(sys.argv[1])