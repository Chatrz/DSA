#include <iostream>
#include <queue>
#include <list>

using namespace std;

struct edge
{
    int index;
    int u;
};

struct node
{
    node() : visited(false){};
    bool visited;
    list<edge> edges;
};

void print(int **matrix, int m, int n);
list<edge> bfs(node *graph, int vertecies);

int main()
{
    int m, n;
    cin >> n >> m;
    node *graph = new node[n];
    int u, v;
    for (int i = 1; i <= m; i++)
    {
        cin >> u >> v;
        u--;
        v--;
        graph[u].edges.push_back(edge{i, v});
        graph[v].edges.push_back(edge{i, u});
    }
    list<edge> res = bfs(graph, n);
    cout << res.size() << endl;
    for (edge curr : res)
        cout << curr.index << " ";
    return 0;
}

list<edge> bfs(node *graph, int vertecies)
{
    list<edge> edges;
    queue<node> q;
    graph[0].visited = true;
    q.push(graph[0]);
    while (!q.empty())
    {
        node u = q.front();
        q.pop();
        for (edge curr : u.edges)
        {
            if (!graph[curr.u].visited)
            {
                graph[curr.u].visited = true;
                edges.push_back(curr);
                q.push(graph[curr.u]);
            }
        }
    }
    return edges;
}