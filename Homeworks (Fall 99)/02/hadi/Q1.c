#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct node
{
    int value;
    int count;
    struct node *next;
};

struct list
{
    struct node *head;
    struct node *tail;
};

void insert(struct list *lst, int value, int count)
{
    struct node *new = (struct node *)malloc(sizeof(struct node));
    new->count = count;
    new->value = value;
    new->next = NULL;
    if (lst->head != NULL)
    {
        lst->tail->next = new;
        lst->tail = new;
        return;
    }
    lst->head = new;
    lst->tail = new;
    return;
}

void rem(struct list *lst, int count)
{
    while (count > 0)
    {
        if (lst->head->count - count > 0)
        {
            lst->head->count -= count;
            break;
        }
        count -= lst->head->count;
        lst->head = lst->head->next;
    }
}

void print_list(struct list *lst)
{
    if (lst->head != NULL)
    {
        printf("%d\n", lst->head->value);
        return;
    }
    printf("empty\n");
}

void scanIn(struct list *lst)
{
    char opr;
    int d, t, j;
    scanf("%d", &j);
    for (j; j > 0; j--)
    {
        scanf(" %c", &opr);
        //printf("here: %c\n", opr);
        switch (opr)
        {
        case '+':
            scanf("%d %d", &d, &t);
            insert(lst, d, t);
            break;
        case '-':
            scanf("%d", &t);
            rem(lst, t);
            break;
        default:
            print_list(lst);
            break;
        }
    }
}

int main()
{
    struct list *lst = (struct list *)malloc(sizeof(struct list));
    lst->head = lst->tail = NULL;
    scanIn(lst);
    return 0;
}