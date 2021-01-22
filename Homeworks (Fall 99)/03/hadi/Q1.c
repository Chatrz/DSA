#include <stdio.h>
#include <stdlib.h>

int left(int n)
{
    return 2 * n + 1;
}

int right(int n)
{
    return 2 * n + 2;
}

int parent(int n)
{
    return (n - 1) / 2;
}
void tfix(int index, int *arr, int size)
{
    if (index < 0)
        return;
    if (*(arr + index) == -1)
    {
        int leftChild = left(index);
        int rightChild = right(index);
        int max = -1;
        if (leftChild < size)
            max = *(arr + leftChild);
        if (rightChild<size &&*(arr + rightChild)> max)
            max = *(arr + rightChild);
        max = max == -1 ? 1 : max + 1;
        *(arr + index) = max;
    }
    tfix(index - 1, arr, size);
}
int main()
{
    int n;
    int *arr;
    scanf("%d", &n);
    arr = (int *)malloc(n * sizeof(int));
    for (int i = 0; i < n; i++)
    {
        scanf("%d", arr + i);
    }
    tfix(n, arr, n);
    for (int i = 0; i <= n / 2; i++)
    {
        int lefti, righti;
        lefti = left(i);
        righti = right(i);
        if (righti < n && arr[righti] > arr[i])
        {
            printf("-1");
            return 0;
        }
        if (lefti < n && arr[lefti] > arr[i])
        {
            printf("-1");
            return 0;
        }
    }

    for (int i = 0; i < n; i++)
        printf("%d ", *(arr + i));
}
