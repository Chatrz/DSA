
# this code is a part of udacity DS course (https://www.udacity.com/course/data-structures-and-algorithms-in-python--ud513)



class Node(object):
    def __init__(self, value):
        self.value = value
        self.edges = []
    def __repr__(self):
        return str(self.value)
class Edge(object):
    def __init__(self, value, node_from, node_to):
        self.value = value
        self.node_from = node_from
        self.node_to = node_to

class Graph(object):
    def __init__(self, nodes=[], edges=[]):
        self.nodes = nodes
        self.edges = edges

    def insert_node(self, new_node_val):
        new_node = Node(new_node_val)
        self.nodes.append(new_node)
        
    def insert_edge(self, new_edge_val, node_from_val, node_to_val):
        from_found = None
        to_found = None
        for node in self.nodes:
            if node_from_val == node.value:
                from_found = node
            if node_to_val == node.value:
                to_found = node
        if from_found == None:
            from_found = Node(node_from_val)
            self.nodes.append(from_found)
        if to_found == None:
            to_found = Node(node_to_val)
            self.nodes.append(to_found)
        new_edge = Edge(new_edge_val, from_found, to_found)
        from_found.edges.append(new_edge)
        to_found.edges.append(new_edge)
        self.edges.append(new_edge)

    def get_edge_list(self):
        edge_list=[]
        for edge in self.edges:
            edge_list.append((edge.value,edge.node_from,edge.node_to))
        return edge_list

    def get_adjacency_list(self):
        adjencency_list=[None for i in range(len(self.nodes)+1)]
        for edge in self.edges:
            from_node=edge.node_from
            info_list=[]
            for edge in from_node.edges:
                if(from_node!=edge.node_to):
                    info_list.append((edge.node_to,edge.value))
            adjencency_list[from_node.value]=info_list
        return adjencency_list
    
    def get_adjacency_matrix(self):
        matrix=[[0,0,0,0,0] for i in range(len(self.nodes)+1)]
        for node in self.nodes:
            info_list=[0 for i in range(5)]
            for edge in node.edges:
                if(node==edge.node_from):
                    info_list[edge.node_to.value]=edge.value
            matrix[node.value]=info_list
        return matrix
graph = Graph()
graph.insert_edge(100, 1, 2)
graph.insert_edge(101, 1, 3)
graph.insert_edge(102, 1, 4)
graph.insert_edge(103, 3, 4)
# Should be [(100, 1, 2), (101, 1, 3), (102, 1, 4), (103, 3, 4)]
print(graph.get_edge_list())
# Should be [None, [(2, 100), (3, 101), (4, 102)], None, [(4, 103)], None]
print(graph.get_adjacency_list())
# Should be [[0, 0, 0, 0, 0] [0, 0, 100, 101, 102], [0, 0, 0, 0, 0], [0, 0, 0, 0, 103], [0, 0, 0, 0, 0]]
print(graph.get_adjacency_matrix())