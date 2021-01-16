package code;
// Libraries used
import java.util.ArrayList;
import java.util.Scanner;

// Each vertex have a mode and color
class Vertex
{
    public char mode, color;
    public ArrayList<Vertex> adj;
    public int group;
    public Vertex()
    {
        group = -1;
        color = 'W';
        adj = new ArrayList<>();
    }
}

// Main class of program
public class Main
{

    /**
     * The dfs visit method which colors the visited vertices.
     * @param u the vertex we visit
     * @param hash_key add it to the hash-map
     */
    public static void dfs_visit(Vertex u, int hash_key)
    {
        u.color = 'G';
        for(Vertex v : u.adj)
        {
            if (v.color == 'W')
                dfs_visit(v, hash_key);
        }
        u.group = hash_key;
    }

    /**
     * Performing a dfs on the graph
     * @param nodes graph nodes
     * @param n the rows
     * @param m the cols
     */
    public static void dfs(Vertex[][] nodes, int n, int m)
    {
        int hash_key = 0;
        for(int y = 0; y < n; y++)
        {
            for(int x = 0; x < m; x++)
            {
                Vertex temp = nodes[y][x];
                if (temp.mode == '.' && temp.color == 'W')
                {
                    dfs_visit(temp, hash_key);
                    hash_key++;
                }
            }
        }
    }


    public static void main(String[] args) {
        ArrayList<ArrayList<Vertex>> table_result = new ArrayList<>();

        Scanner scanner = new Scanner(System.in);
        // The rows and cols
        int n = scanner.nextInt();
        int m = scanner.nextInt();

        Vertex[][] nodes = new Vertex[n][m];
        for (int i = 0; i < n; i++)
        {
            String line = scanner.next();
            for (int j = 0; j < m; j++)
            {
                nodes[i][j] = new Vertex();
                nodes[i][j].mode = line.charAt(j);
                if (line.charAt(j) == '.')
                {
                    if (i - 1 > -1 && nodes[i-1][j] != null && nodes[i-1][j].mode == '.')
                    {
                        nodes[i-1][j].adj.add(nodes[i][j]);
                        nodes[i][j].adj.add(nodes[i-1][j]);
                    }
                    if (j - 1 > -1 && nodes[i][j-1] != null && nodes[i][j-1].mode == '.')
                    {
                        nodes[i][j-1].adj.add(nodes[i][j]);
                        nodes[i][j].adj.add(nodes[i][j-1]);
                    }
                    if (i + 1 < n && nodes[i+1][j] != null && nodes[i+1][j].mode == '.')
                    {
                        nodes[i+1][j].adj.add(nodes[i][j]);
                        nodes[i][j].adj.add(nodes[i+1][j]);
                    }
                    if (j + 1 < m && nodes[i][j+1] != null && nodes[i][j+1].mode == '.')
                    {
                        nodes[i][j+1].adj.add(nodes[i][j]);
                        nodes[i][j].adj.add(nodes[i][j+1]);
                    }
                }
            }
        }

        // Dfs to manage the nodes
        dfs(nodes, n, m);

        // Finding the answers
        int q = scanner.nextInt();
        for (int i = 0; i < q; i++)
        {
            int y1 = scanner.nextInt()-1;
            int x1 = scanner.nextInt()-1;
            int y2 = scanner.nextInt()-1;
            int x2 = scanner.nextInt()-1;
            String result = ( nodes[y1][x1].group == nodes[y2][x2].group ) ? "YES" : "NO";
            System.out.println(result);
        }
    }
}
