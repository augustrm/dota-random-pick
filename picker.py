import networkx as nx
import matplotlib.pyplot as plt
import random as r

with open("heroes.txt", 'r') as f:
    heroes = f.read().splitlines()

def cantor_pair(x,y):
    return int(0.5*((x + y)*(x + y + 1))+y)

cantor_map = []
for i in range(0,4): #includes 0-3
    for j in range(0,31): # includes 0-30
        cantor_map.append(cantor_pair(i,j))
cantor_map.sort()


def get_hero(t:tuple):
    cantor_index = cantor_pair(t[0], t[1])
    hero_index = cantor_map.index(cantor_index)
    return heroes[hero_index]

G = nx.grid_2d_graph(4,31, periodic=True)

def do_picks(j):
    picks = {}
    for i in range(j):
        m = r.randint(0,3)
        n = r.randint(0,30)

        hero_node = (m,n)
        hero = get_hero(hero_node)
        # [(top), bottom, left, right] #adjacency library-free
        #_adj_nodes = [((m-1)%4, n), ((m+1)%4, n), (m, (n-1)%31), (m, (n+1)%31)]

        adj_nodes = list(G.neighbors(hero_node))
        adj_heroes = [get_hero(x) for x in adj_nodes]
        picks[i] = {"hero":hero, "alts": adj_heroes}
    return picks

p = do_picks(5)


print(p)
