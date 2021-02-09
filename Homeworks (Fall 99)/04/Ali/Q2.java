import java.util.*;

public class Main {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int n = scanner.nextInt();
        int m = scanner.nextInt();

        ArrayList<ArrayList<Integer>> adjList = new ArrayList<>(n);
        HashMap<List<Integer>, Integer> edges = new HashMap<>();

        for (int i = 0; i < n; i++) {
            adjList.add(new ArrayList<>());
        }
        for (int i = 0; i < m; i++) {
            int src = scanner.nextInt();
            int des = scanner.nextInt();
            addEdge(adjList, edges, i + 1, src - 1, des - 1);
        }

        int[] predecessors = new int[n];
        int[] distances = new int[n];

        for (int i = 0; i < n; i++) {
            distances[i] = Integer.MAX_VALUE;
            predecessors[i] = -1;
        }

        BFS(adjList, 0, n, predecessors, distances);

        printResult(n, predecessors, edges);

    }

    private static void addEdge(ArrayList<ArrayList<Integer>> adj, HashMap<List<Integer>, Integer> edges,
                                int edgeName, int src, int des) {
        adj.get(src).add(des);
        adj.get(des).add(src);
        List<Integer> newEdgeSD =new ArrayList<>();
        List<Integer> newEdgeDS = new ArrayList<>();
        newEdgeDS.add(des);
        newEdgeDS.add(src);
        newEdgeSD.add(src);
        newEdgeSD.add(des);

        edges.put(newEdgeDS, edgeName);
        edges.put(newEdgeSD, edgeName);
    }

    private static void printResult(int n, int[] predecessors, HashMap<List<Integer>, Integer> edges) {
//        System.out.println(edges.keySet().size());
        int finalEdgesNumber = 0;
        ArrayList<Integer> finalEdges = new ArrayList<>();
        for (int i = 0; i < n; i++) {
            int k = i;
            while (predecessors[k] != -1) {
                int j = k;
                k = predecessors[k];
                List<Integer> edgeToCheck = new ArrayList<>();
                edgeToCheck.add(k);
                edgeToCheck.add(j);
//                System.out.println(edgeToCheck.get(0) + " " + edgeToCheck.get(1));
                Integer edgeNumber = edges.get(edgeToCheck);
                if (edgeNumber != null && !finalEdges.contains(edgeNumber)) {
                    finalEdgesNumber++;
                    finalEdges.add(edgeNumber);
                }
            }
        }
        System.out.println(finalEdgesNumber);
        for (Integer finalEdge : finalEdges) {
            System.out.print(finalEdge + " ");
        }
    }

    private static void BFS(ArrayList<ArrayList<Integer>> adjList, int src,
                            int n, int[] predecessors, int[] distances) {

        LinkedList<Integer> queue = new LinkedList<>();
        boolean[] visited = new boolean[n];

        for (int i = 0; i < n; i++)
            visited[i] = false;

        visited[src] = true;
        distances[src] = 0;
        queue.add(src);

        while (!queue.isEmpty()) {
            int u = queue.remove();
            for (int i = 0; i < adjList.get(u).size(); i++) {
                if (!visited[adjList.get(u).get(i)]) {
                    visited[adjList.get(u).get(i)] = true;
                    distances[adjList.get(u).get(i)] = distances[u] + 1;
                    predecessors[adjList.get(u).get(i)] = u;
                    queue.add(adjList.get(u).get(i));
                }
            }
        }

    }
}

