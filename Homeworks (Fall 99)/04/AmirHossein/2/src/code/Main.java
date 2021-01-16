package code;

import java.util.*;

class Vertex
{
    public HashMap<Integer, Vertex> adj;
    public int dv;
    public Vertex ()
    {
        dv = -1;
        adj = new HashMap<>();
    }
    public int chose()
    {
        int chose = 0;
        int min = dv;
        for (Integer key : adj.keySet())
        {
            Vertex temp = adj.get(key);
            if (min > temp.dv)
            {
                min = temp.dv;
                chose = key;
            }
        }
        return chose;
    }
}

public class Main {

    public static void bfs (Vertex root)
    {
        root.dv = 0;
        Queue<Vertex> queue = new LinkedList<>();
        queue.add(root);
        while ( !queue.isEmpty() )
        {
            Vertex u = queue.remove();
            for (Vertex v : u.adj.values())
            {
                if (v.dv == -1)
                {
                    v.dv = u.dv + 1;
                    queue.add(v);
                }
            }
        }
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int n = scanner.nextInt();
        int m = scanner.nextInt();

        Vertex[] vertices = new Vertex[n];
        for (int i = 0; i < n; i++)
            vertices[i] = new Vertex();

        for (int i = 0; i < m; i++)
        {
            int head = scanner.nextInt();
            int tail = scanner.nextInt();
            vertices[head-1].adj.put(i+1, vertices[tail-1]);
            vertices[tail-1].adj.put(i+1, vertices[head-1]);
        }

        bfs(vertices[0]);

        SortedSet<Integer> results = new TreeSet<>();
        for(int i = 1; i < n; i++)
        {
            results.add(vertices[i].chose());
        }
        System.out.println(results.size());
        for (Integer i : results)
        {
            System.out.print(i + " ");
        }
    }
}
