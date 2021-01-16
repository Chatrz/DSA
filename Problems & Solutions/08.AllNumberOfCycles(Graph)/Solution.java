import java.util.*;


public class Solution{
    // Fields
    public static int total_number_edges;
    static int count = 0; // Each round we count the cycles and store in here
    static int[][] graph = new int[1000][1000]; // Our graph matrix
    static HashMap<Integer, Integer> edges = new HashMap<>(); // Used hash map to convert graph into matrix with index
    static int index = 0;

    // Using dfs to count the cycles
    static void DFS(boolean marked[], int n, int vert, int start) {
        marked[vert] = true;
        if (n == 0) {
            marked[vert] = false;
            if (graph[vert][start] == 1) {
                count++;
                return;
            } else
                return;
        }
        for (int i = 0; i < total_number_edges; i++)
            if (!marked[i] && graph[vert][i] == 1)
                DFS(marked, n-1, i, start);
        marked[vert] = false;
    }

    static int countCycles(int n) {
        count = 0;
        boolean[] marked = new boolean[total_number_edges];
        for (int i = 0; i < total_number_edges - (n - 1); i++) {
            DFS(marked, n-1, i, i);
            marked[i] = true;
        }
        return count / 2; // Half of the counts are similar
    }

    // Adding a new edge to the matrix
    static void addEdge(int u, int v)
    {
        int index_u = edges.get(u);
        int index_v = edges.get(v);
        graph[index_u][index_v] = 1;
        graph[index_v][index_u] = 1;
    }

    // A method for adding the input types into matrix type
    static void add_to_map(int key)
    {
        if ( ! edges.containsKey(key) )
        {
            edges.put(key, index);
            index++;
        }
    }

    // Adding a new number and all numbers that counts it to the map
    static void new_number(int number)
    {
        for (int i = 1; i <= number / 2; i++)
        {
            if (number % i == 0)
            {
                add_to_map(i);
                addEdge(i, number);
                new_number(i);
            }
        }
    }

    // To convert the graph given model into a matrix
    static int create_matrix()
    {
        Scanner scanner = new Scanner(System.in);
        int n = scanner.nextInt();

        for (int i = 0; i < n; i++)
        {
            int temp = scanner.nextInt();
            add_to_map(temp);
            new_number(temp);
        }

        total_number_edges = edges.size();
        return edges.size();
    }

    // Program starts
    public static void main(String[] args) {
        int total_edges = create_matrix();
        int cycles = 0;
        for (int n = 3; n < total_edges+1; n++)
        {
            cycles += countCycles(n);
        }
        System.out.println(cycles);
    }
}
