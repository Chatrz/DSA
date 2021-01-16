#include<iostream>
#include <queue>
using namespace std;
class Block{
    public:
    int state;
    bool visited;
    int parent;
    Block(){
        visited=false;
    }
};
void dfs (Block** grid,int row, int col);
void rec_dfs(Block** grid,int prow, int pcol,int row, int col);
int check_parent(Block**grid,int sr,int sc,int dr,int dc);

int main (){
    int row, col;
    cin>>row;
    cin>>col;

    Block** grid=new Block*[row];
    for(int i=0;i<row;i++){
        grid[i]=new Block[col];
        for(int j=0;j<col;j++){
            char state;
            cin>>state;
            grid[i][j].state=state=='.'?0:1;
            grid[i][j].parent=i*col+j;
        }
    }
    dfs(grid,row,col);
    int numOfCheck;
    cin>>numOfCheck;
    for (int i=0;i<numOfCheck;i++){
        int sr,sc,dr,dc;
        cin>>sr>>sc>>dr>>dc;
        int result=check_parent(grid,sr-1,sc-1,dr-1,dc-1);
        string answer=result==1?"YES":"NO";
        cout<<answer<<"\n";
    }



    return 1;
}


void rec_dfs(Block** grid,int prow, int pcol,int row, int col){
    grid[prow][pcol].visited=true;
    int dir_row[4]={-1,+1,0,0};
    int dir_col[4]={0,0,+1,-1};

    for (int i=0;i<4;i++){
            int nr=prow+dir_row[i];
            int nc=pcol+dir_col[i];

            if(nr<0 || nc<0) continue;
            if(nr>=row || nc>= col) continue;
            if(grid[nr][nc].visited)continue;
            if(grid[nr][nc].state==1)continue;

            grid[nr][nc].parent=grid[prow][pcol].parent;
            rec_dfs(grid,nr,nc,row,col);
        }

}
void dfs (Block** grid,int row, int col){

    for (int r=0;r<row;r++){
        for(int c=0;c<col;c++){
            if(grid[r][c].state==0 && grid[r][c].visited==false)
                rec_dfs(grid,r,c,row,col);
        }
    }

}
int check_parent(Block**grid,int sr,int sc,int dr,int dc){
    return grid[sr][sc].parent==grid[dr][dc].parent;
}