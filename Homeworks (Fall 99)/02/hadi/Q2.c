#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct box
{
    int number;
    struct box *next;
};

struct dock
{
    struct box *boxes;
};

int main()
{
    int n, m, i, j, d;
    scanf("%d %d", &n, &m);
    struct dock *docks = (struct dock *)malloc(n * sizeof(struct dock));
    for (int i = 0; i < n; i++)
    {
        docks[i].boxes = (struct box *)malloc(sizeof(struct box));
        docks[i].boxes->number = i + 1;
        docks[i].boxes->next = NULL;
    }
    for (m; m > 0; m--)
    {
        scanf(" %d %d", &i, &j);
        i--;
        j--;
        struct box *curr = docks[j].boxes;
        if (curr != NULL)
        {
            for (curr; curr->next != NULL; curr = curr->next);
            curr->next = docks[i].boxes;
        } else {
            docks[j].boxes = docks[i].boxes;
        }
        docks[i].boxes = NULL;
    }
    scanf(" %d", &d);
    for (struct box *curr = docks[d-1].boxes; curr != NULL; curr= curr->next, m++);
    printf("%d", m);
    for (struct box *curr = docks[d-1].boxes; curr != NULL; curr= curr->next)
        printf(" %d", curr->number);
    
}
