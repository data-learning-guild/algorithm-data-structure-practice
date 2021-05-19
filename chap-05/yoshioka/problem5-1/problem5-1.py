import networkx as nx
import numpy as np
from numpy import random as rd
from itertools import permutations


def vacation(num_of_days, value_dict: dict):  # -> np.array:
    G = nx.DiGraph()

    nodes_name = list(permutations(value_dict.keys(), 2))

    for i in range(num_of_days):
        if i == 0:
            for k, v in value_dict.items():
                G.add_edge("-1", f"{k}{i}", weight=v[i])
        if i > 0:
            for p1, p2 in nodes_name:
                G.add_edge(f"{p1}{i-1}", f"{p2}{i}",
                           weight=value_dict[p2][i])

    for k, v in value_dict.items():
        G.add_edge(f"{k}{num_of_days-1}", "-2", weight=100)

    result = nx.dag_longest_path(G, weight="weight")
    return result


if __name__ == '__main__':
    num_of_days = 3

    a_array = [1, 1, 1]
    b_array = [2, 2, 2]
    c_array = [3, 3, 3]

    # a_array = rd.randint(0, 100, [1, num_of_days])[0]
    # b_array = rd.randint(0, 100, [1, num_of_days])[0]
    # c_array = rd.randint(0, 100, [1, num_of_days])[0]

    value_dict = {
        "a": a_array,
        "b": b_array,
        "c": c_array
    }

    result = vacation(num_of_days, value_dict)

    def get_value(x): return value_dict[x[0]][int(x[-1])]
    result_value = sum([get_value(x) for x in result[1:-1]])
    print(result_value)
