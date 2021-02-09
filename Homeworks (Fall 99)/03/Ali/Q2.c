#include <stdio.h>
#include <stdlib.h>
#define MAX(a, b) (((a) > (b)) ? (a) : (b))

int main()
{
    int m, n;
    scanf("%d", &m);
    scanf("%d", &n);
    int maxMoney[n];
    for (int i = 0; i < n; i++)
        maxMoney[i] = 0;
    for (int i = 0; i < m; i++)
    {
        int enter, exit, money;
        scanf("%d", &enter);
        scanf("%d", &exit);
        scanf("%d", &money);
        for (int j = enter-1; j < exit; j++)
        {
            maxMoney[j] = MAX(maxMoney[j],money);
        }
        
    }
    for (int i = 0; i < n; i++)
        printf("%d ",maxMoney[i]);    

    return 0;
}