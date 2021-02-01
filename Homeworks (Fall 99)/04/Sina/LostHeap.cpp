#include<iostream>
// #include <cmath>
using namespace std;
int fill_heap(int heap[],int size);
int rec_fill_heap(int heap[],int size,int index,int parent);
int main (){
    int heap_size;
    cin >> heap_size;
    int heap [heap_size]={0};
    for (int i=0;i<heap_size;i++){
        cin>>heap[i];
    }
    int result=rec_fill_heap(heap,heap_size,0,-1);
    if(result==-1)
        cout<<-1;
    else{
        for(int el:heap){
            cout<<el<<" ";
        }
    }
}

int fill_heap(int heap[],int size){
    for(int i=1;i<size;i++){
        int parent=(i-1)/2;
        // if (heap[parent]==-1){
        //     int p_left_child=heap[parent*2+1];
        //     int p_right_child=heap[parent*2+2];
        //     int max=p_left_child>p_right_child?p_left_child:p_right_child;
        //     heap[parent]=max+1;
        // }
        if(2*i+1<size)
            if(heap[2*i+1]>heap[i])return -1;
        if(2*i+2<size)
            if(heap[2*i+2]>heap[i])return -1;
        if(heap[parent]<heap[i])return -1;
        if(heap[i]==-1){
            int left_child=2*i+1>size?-2:heap[2*i+1];
            int right_child=2*i+2>size?-2:heap[2*i+2];
            int upper_bound=heap[parent];
            if(left_child==-2 && right_child==-2){
                heap[i]=upper_bound-1;
            }else{
                int lower_bound=left_child>right_child?left_child:right_child;
                if(lower_bound>upper_bound)return -1;
                if (lower_bound==upper_bound)heap[i]=lower_bound;
                else{
                    heap[i]=upper_bound-1;
                }
            }

        }
        

    }
    return 1;
}
int rec_fill_heap(int heap[],int size,int index,int parent_key){
    if(parent_key>-1)
        if(heap[index]>parent_key)return -1;
    int max_left=-2;
    int max_right=-2;
    if(2*index+1<size){
        if(heap[2*index+1]>heap[index]&&heap[index]!=-1)return -1;
        max_left=rec_fill_heap(heap,size,2*index+1,heap[index]==-1?parent_key-1:heap[index]);
    }
    if(2*index+2<size){
        if(heap[2*index+2]>heap[index]&&heap[index]!=-1)return -1;
        max_right=rec_fill_heap(heap,size,2*index+2,heap[index]==-1?parent_key-1:heap[index]);
    }
    if(max_left==-2 && max_right==-2){
        heap[index]=heap[index]==-1?parent_key<0?1:parent_key-1:heap[index];
        return heap[index];
    }
    if(max_left==-1 || max_right==-1)return -1;//signal indicating its not a heap or cant be a heap
    int max_both=max_left>max_right?max_left:max_right;
    if(heap[index]==-1){
        if(parent_key<max_both&&parent_key>-1)return -1;
        heap[index]=max_both+1;

    }
    return heap[index];
}