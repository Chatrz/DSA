#include <iostream>
#include <vector>
using namespace std;

class Graph{
    private:
        int v;
        vector<int>* adjList;
    public:
        Graph(int vertex_num){
            this->v= vertex_num;
            adjList = new vector<int>[vertex_num];
        }
        void add_edge(int u, int v){
            adjList[u].push_back(v);
        }
        void print_adj_list(){
            for(int i=0; i<v ;i++){
                vector<int>::iterator iter;
                cout<< i <<"- " ;
                for(iter = adjList[i].begin(); iter!=adjList[i].end();iter++){
                    cout<< *iter;
                }
                cout<<endl;
            }
        }
        void print_adj_matrix(){
            for (int i=0;i<v;i++){
                vector<int> temp_vect(v,0);
                vector<int>::iterator iter;
                cout<< i <<"- ";
                for(iter = adjList[i].begin(); iter!=adjList[i].end();iter++){
                    temp_vect[*iter]=1;
                }
                for(iter = adjList[i].begin(); iter!=adjList[i].end();iter++){
                    cout<<*iter;
                }
                cout<<endl;
            }
        }
};