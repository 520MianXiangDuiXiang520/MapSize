import matplotlib.pyplot as plt
import csv


class MapSizeStatistic:
    """
    A statistic of map storage usage in go where key-value pairs are all int64
    """

    def __init__(self):
        self.utilization_list = []
        self.rand_utilization_list = []
        self.str_utilization_list = []

        with open("./int64.csv") as fp:
            reader = csv.reader(fp)
            self.utilization_list = [float(i[-1]) for i in reader]

        with open("./int64_rand.csv") as fp:
            reader = csv.reader(fp)
            self.rand_utilization_list = [float(i[-1]) for i in reader]

        with open("./string_rand.csv") as fp:
            reader = csv.reader(fp)
            self.str_utilization_list = [float(i[-1]) for i in reader]

    def draw_utilization(self):
        x = [i*100 for i in range(len(self.utilization_list))]
        x2 = [i*100 for i in range(len(self.rand_utilization_list))]
        x3 = [i*100 for i in range(len(self.str_utilization_list))]
        average_x = sum(self.utilization_list)/len(self.utilization_list)
        average_x2 = sum(self.rand_utilization_list)/len(self.rand_utilization_list)
        average_x3 = sum(self.str_utilization_list)/len(self.str_utilization_list)
        plt.plot(x, self.utilization_list, color="red")
        plt.plot(x2, self.rand_utilization_list, color="green")
        plt.plot(x3, self.str_utilization_list, color="blue")
        plt.plot(x, [average_x]*len(self.utilization_list), color="red")
        plt.plot(x2, [average_x2]*len(self.rand_utilization_list), color="green")
        plt.plot(x3, [average_x3]*len(self.str_utilization_list), color="blue")
        plt.legend(['rand', 'order', 'str', 'average_order', 'average_rand', 'average_str'])
        plt.show()


if __name__ == '__main__':
    mss = MapSizeStatistic()
    mss.draw_utilization()
