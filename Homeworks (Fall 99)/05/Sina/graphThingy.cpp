#include <iostream>
#include <list>
#include <utility>
#include <bits/stdc++.h>

using namespace std;
class Graph{
    private:
    int v;
    int keptNum;
    list<pair<int,int>>* adj;
    list<int>* keptEdges;
    public:
    void addEdge(int index,int u, int v){
        adj[u].push_back(make_pair(index,v));
        adj[v].push_back(make_pair(index,u));
    }
    Graph(int v){
        this->v=v;
        adj= new list<pair<int,int>>[v+1];
        // for (int i=0;i<=v;i++){
        //     adj[i]=new list<pair<int,int>>();
        // }
        keptEdges= new list<int>();
        keptNum=0;
    }
    void solve(){
        bool* visited = new bool[v];
        list <int> queue ;
        visited[1]=true;
        queue.push_back(1);


        list<pair<int,int>>::iterator iter;
        while(!queue.empty()){
            int s=queue.front();
            queue.pop_front();
            for(iter = adj[s].begin();iter!=adj[s].end();iter++){
                int index,node;
                tie(index,node)=*iter;
                if(!visited[node]){
                    keptEdges->push_back(index);
                    keptNum++;
                    visited[node]=true;
                    queue.push_back(node);
                }
            }
        }
    }
    void print_result(){
        cout<<keptNum<<"\n";
        list<int>::iterator iter;
        for(iter = keptEdges->begin();iter !=keptEdges->end();iter++){
            cout<<*iter<<" ";
        }

    }
};
int main (){
    int v, e;
    cin>>v>>e;
    Graph* graph= new Graph(v);
    for(int i=1;i<=e;i++){
        int u,v;
        cin>>u>>v;
        graph->addEdge(i,u,v);
    }
    graph->solve();
    graph->print_result();

}