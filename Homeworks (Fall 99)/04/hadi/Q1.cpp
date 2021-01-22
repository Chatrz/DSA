#include <iostream>
#include <stack>
#include <string>

using namespace std;
const int wall = -1;
const int unvisited = -2;

typedef struct
{
    int x;
    int y;
} coordinate;

coordinate getUnvisitedNode(int **graph, int i, int j, int m, int n);
void dfs(int **graph, int i, int j, int mark, int m, int n);

int main()
{
    int m, n;
    string row;
    cin >> m >> n;
    int **graph = new int *[m];
    for (int i = 0; i < m; i++)
    {
        graph[i] = new int[n];
        cin >> row;
        for (int j = 0; j < n; j++)
        {
            if (row[j] == '.')
                graph[i][j] = unvisited;
            else
                graph[i][j] = wall;
        }
    }
    int c1, c2, r1, r2, q;
    cin >> q;
    for (int i = 0; i < q; i++)
    {
        cin >> c1 >> r1 >> c2 >> r2;
        if (graph[c1 - 1][r1 - 1] == unvisited && graph[c2 - 1][r2 - 1] == unvisited)
            dfs(graph, c1 - 1, r1 - 1, i, m, n);
        if (graph[c1 - 1][r1 - 1] != graph[c2 - 1][r2 - 1] || graph[c1 - 1][r1 - 1] == wall)
            cout << "NO" << endl;
        else
            cout << "YES" << endl;
    }
}

void dfs(int **graph, int i, int j, int mark, int m, int n)
{
    stack<coordinate> s;
    coordinate c = coordinate{i, j};
    graph[c.x][c.y] = mark;
    s.push(c);
    while (!s.empty())
    {
        c = s.top();
        graph[c.x][c.y] = mark;
        s.pop();
        for (coordinate curr = getUnvisitedNode(graph, c.x, c.y, m, n); curr.x != -1 && curr.y != -1; curr = getUnvisitedNode(graph, c.x, c.y, m, n))
        {
            if (graph[curr.x][curr.y] == unvisited)
            {
                graph[curr.x][curr.y] = mark;
                s.push(curr);
            }
        }
    }
}

coordinate getUnvisitedNode(int **graph, int i, int j, int m, int n)
{
    if (i > 0 && graph[i - 1][j] == unvisited)
    {
        //case up
        return coordinate{i - 1, j};
    }
    else if (i < m - 1 && graph[i + 1][j] == unvisited)
    {
        //case down
        return coordinate{i + 1, j};
    }
    else if (j > 0 && graph[i][j - 1] == unvisited)
    {
        //case left
        return coordinate{i, j - 1};
    }
    else if (j < n - 1 && graph[i][j + 1] == unvisited)
    {
        //case right
        return coordinate{i, j + 1};
    }
    //case nothing
    return coordinate{-1, -1};
}
