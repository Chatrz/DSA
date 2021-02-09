#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <math.h>
bool isMaxHeap(int *nodes, int n)
{
    for (int i = 0; i < n; i++)
    {
        if ( nodes[i] == -1 || nodes[(i - 1) / 2] < nodes[i])
            return false;
    }
    return true;
}

void compeleteMaxHeap(int *nodes, int n)
{
    for (int i = 0; i < n; i++)
    {
        if (nodes[i] == -1)
        {
            if (i==0)
            {
                nodes[i] = pow(10,9);
            }
            else
            {
                nodes[i] = nodes[(i-1)/2]- 1;
            }
        }
    }
}

int main()
{
    int n;
    scanf("%d", &n);

    int nodes[n];
    for (int i = 0; i < n; i++)
    {
        int newNodeData;
        scanf("%d", &newNodeData);
        nodes[i] = newNodeData;
    }
    
    compeleteMaxHeap(nodes,n);
    if (isMaxHeap(nodes,n)){
        for (int i = 0; i < n; i++){
            printf("%d ",nodes[i]);
        }
    }else{
        printf("-1");
    }
    return 0;
}