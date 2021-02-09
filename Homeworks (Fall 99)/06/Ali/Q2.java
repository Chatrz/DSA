import java.util.*;

public class Main2 {
    static Scanner in = new Scanner(System.in);
    static int inf = Integer.MAX_VALUE;

    public static void main(String[] args) {
        int vertices = in.nextInt();
        int edges = in.nextInt();
        int questions = in.nextInt();

        Vertex[] adjList = new Vertex[vertices + 1];

        for (int i = 0; i <= vertices; i++) {
            adjList[i] = new Vertex(i);
        }


        for (int k = 0; k < edges; k++) {
            int start = in.nextInt();
            int finish = in.nextInt();
            int cost = in.nextInt();
            adjList[start].neighbors.add(new Edge(start, finish, cost, k + 1));
            adjList[finish].neighbors.add(new Edge(finish, start, cost, k + 1));
        }

        LinkedList<Integer> answers = new LinkedList<>();
        for (int i = 0; i < questions; i++) {
            String c = in.next();
            if (c.equals("?")) {
                int s = in.nextInt();
                int d = in.nextInt();
                int ans = Dijkstras(adjList, s, d);
                answers.add(ans);
            } else if (c.equals("+")) {
                int vertex = in.nextInt();
                adjList[vertex].isValid = true;
            }
        }
        for (int ans : answers) {
            System.out.println(ans);
        }
    }

    public static void initSinguleSource(Vertex[] graph) {
        for (Vertex v :
                graph) {
            v.minDistance = inf;
        }
    }

    public static int Dijkstras(Vertex[] graph, int source, int destination) {
        initSinguleSource(graph);
        Vertex current = graph[source];
        current.minDistance = 0;

        PriorityQueue<Vertex> priorityQueue = new PriorityQueue<Vertex>();
        if (current.isValid) {
            priorityQueue.add(current);

            while (!priorityQueue.isEmpty()) {
                current = priorityQueue.remove();

                if (current.id == destination) {
                    break;
                }

                for (int i = 0; i < current.neighbors.size(); i++) {
                    Edge edge = current.neighbors.get(i);
                    Vertex end = graph[edge.end];
                    if (!current.isValid || !end.isValid) {
                        continue;
                    }
                    int distance = edge.cost;
                    int newDistance = current.minDistance + distance;

                    if (newDistance < end.minDistance) {
                        priorityQueue.remove(end);
                        end.minDistance = newDistance;
                        priorityQueue.add(end);
                    }
                }
            }
        }

        if (graph[destination].minDistance == Integer.MAX_VALUE) {
            return -1;
        } else {
            return graph[destination].minDistance;
        }
    }

    static class Vertex implements Comparable<Vertex> {
        int id, minDistance;
        List<Edge> neighbors;
        boolean isValid;

        public Vertex(int id) {
            this.id = id;
            this.minDistance = inf;
            this.neighbors = new ArrayList<Edge>();
            this.isValid = false;
        }

        public int compareTo(Vertex v) {
            return (this.minDistance > v.minDistance) ? 1 : -1;
        }
    }

    static class Edge {
        int start, end, cost, name;

        public Edge(int start, int end, int cost, int name) {
            this.start = start;
            this.end = end;
            this.cost = cost;
            this.name = name;
        }
    }

}

