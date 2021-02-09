import java.util.*;

public class Main {

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int vertices = in.nextInt();
        int edges = in.nextInt();

        Vertex[] adjList = new Vertex[vertices + 1];

        for (int i = 0; i <= vertices; i++) {
            adjList[i] = new Vertex();
        }

        for (int k = 0; k < edges; k++) {
            int start = in.nextInt();
            int finish = in.nextInt();
            int cost = in.nextInt();
            adjList[start].neighbors.add(new Edge(start, finish, cost, k + 1));
            adjList[finish].neighbors.add(new Edge(finish, start, cost, k + 1));
        }
        LinkedList<Edge> mst = prime_edge(adjList, 1);
        int main_mst_cost = 0;
//        System.out.println("printing mst:");
        for (Edge e : mst) {
//            System.out.println(e.name);
            main_mst_cost += e.cost;
        }
//        System.out.println("main_mst_cost:" + main_mst_cost);
        for (int i = 0; i < mst.size(); i++) {
            Edge edge = mst.get(i);
//            System.out.println("removing :" + edge.name);
            remove_edge(adjList, edge);
            int new_mst_cost = prime_cost(adjList, 1);
//            System.out.println("new_mst_cost:" + i + " " + new_mst_cost);

            if (new_mst_cost == -1 || main_mst_cost + 1 <= new_mst_cost) {
                System.out.println(edge.name);
                return;
            }
            add_edge(adjList, edge);

        }
        System.out.println(-1);

    }

    public static void remove_edge(Vertex[] adjList, Edge edgeToRemove) {
        int v_start = edgeToRemove.v_start;
        int v_end = edgeToRemove.v_end;
        adjList[v_start].removeNeighbor(v_end);
        adjList[v_end].removeNeighbor(v_start);
    }

    public static void add_edge(Vertex[] adjList, Edge edgeToAdd) {
        int v_start = edgeToAdd.v_start;
        int v_end = edgeToAdd.v_end;
        int cost = edgeToAdd.cost;
        int name = edgeToAdd.name;
        adjList[v_start].addNeighbor(v_start, v_end, cost, name);
        adjList[v_end].addNeighbor(v_end, v_start, cost, name);
    }

    public static int prime_cost(Vertex[] adjList, int s) {
//        System.out.println("in prime_cost:\nadjList size:" + adjList.length);
        int cost = 0;
        boolean[] visited = new boolean[adjList.length];
        visited[s] = true;
        PriorityQueue<Edge> priorityQueue = new PriorityQueue<>(adjList[s].neighbors);
        while (!priorityQueue.isEmpty()) {
            Edge e = priorityQueue.remove();
            if (!visited[e.v_end]) {
                visited[e.v_end] = true;
//                System.out.println(e.name + " is added");
                cost += e.cost;
                priorityQueue.addAll(adjList[e.v_end].neighbors);
            }
        }

        for (int i = 1; i < adjList.length; i++) {
            if (!visited[i]) {
                return -1;
            }
        }
        return cost;
    }

    public static LinkedList<Edge> prime_edge(Vertex[] adjList, int s) {
        LinkedList<Edge> mst = new LinkedList<>();
        boolean[] visited = new boolean[adjList.length];
        visited[s] = true;
        PriorityQueue<Edge> priorityQueue = new PriorityQueue<>(adjList[s].neighbors);
        while (!priorityQueue.isEmpty()) {
            Edge e = priorityQueue.remove();
            if (!visited[e.v_end]) {
                visited[e.v_end] = true;
                mst.add(e);
                priorityQueue.addAll(adjList[e.v_end].neighbors);
            }
        }
        return mst;
    }

    static class Vertex {
        LinkedList<Edge> neighbors;

        public Vertex() {
            neighbors = new LinkedList<>();
        }

        public void addNeighbor(int start, int finish, int cost, int name) {
            neighbors.add(new Edge(start, finish, cost, name));
        }

        public void removeNeighbor(int v_end) {
            for (int i = 0; i < neighbors.size(); i++) {
                Edge edge_to_remove = neighbors.get(i);
                if (edge_to_remove.v_end == v_end) {
                    neighbors.remove(edge_to_remove);
                    return;
                }
            }
        }
    }
    static class Edge implements Comparable<Edge> {
        int v_start, v_end, cost, name;

        public Edge(int v_start, int v_end, int cost, int name) {
            this.v_start = v_start;
            this.v_end = v_end;
            this.cost = cost;
            this.name = name;
        }

        @Override
        public int compareTo(Edge edge) {
            return this.cost - edge.cost;
        }
    }

}


