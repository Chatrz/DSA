#include <stdio.h>

#define MAX(x, y) (((x) > (y)) ? (x) : (y))
#define MIN(x, y) (((x) < (y)) ? (x) : (y))
#define ABS(x) (((x) > 0) ? (x) : (-1 * x))
#define NELEMS(x) (sizeof(x) / sizeof((x)[0]))

const int LIMIT = 10;
int memory[LIMIT];
int head = 1, tail = 1;

void enqueue(int a)
{
    if ((tail+1)%LIMIT == head)
        printf("Queue overflow.\n");
    else
    {
        memory[tail] = a;
        (++tail) %= LIMIT;
    }
    
}

void dequeue()
{
    if ((head+1)%LIMIT == tail)
        printf("Queue underflow.\n");
    else
    {
        (++head) %= LIMIT;
    }
        
}

int size()
{
    int size = ABS(head - tail) + 1;
    return (tail >= head) ? size : LIMIT - size + 2;
}    

void recover()
{
    if ((tail+1)%LIMIT == head)
        printf("Queue overflow.\n");
    else
    {
        (--head) %= LIMIT;
    }
}

void print()
{
    for(int i = 0; i < NELEMS(memory); i++)
    {
        printf(" ---");
    }
    printf("\n|");
    for(int i = 0; i < NELEMS(memory); i++)
    {
        printf(" %d |", memory[i]);
    }
    printf("\n");
    for(int i = 0; i < NELEMS(memory); i++)
    {
        printf(" ---");
    }
    printf("\n ");
    for(int i = 0; i <= MAX(head, tail); i++)
    {
        if(i == tail) 
        {
            printf(" T ");
            if(i < MAX(head, tail))
                printf(" ");
        }
        else if(i == head) 
        { 
            printf(" H "); 
            if(i < MAX(head, tail))
                printf(" "); 
        }
        else
            printf("    ");
    }
    printf("\n\n");
}

void auto_enqueue(int number)
{
    for(int i = 1; i <= number; i++)
    {
        printf("Enqueue:\n");
        enqueue(i % 10);
        print();
    }
}

void auto_dequeue(int number)
{
    for(int i = 0; i < number; i++)
    {
        printf("Dequeue:\n");
        dequeue();
        print();
    }
}

void auto_recover(int number)
{
    for(int i = 0; i < number; i++)
    {
        printf("Recover:\n");
        recover();
        print();
    }
}

int main() 
{
    printf("This is how queue works >> \n");
    printf("Start:\n");
    print();
    auto_enqueue(3);
    auto_dequeue(2);
    auto_enqueue(5);
    auto_recover(2);
    auto_enqueue(2);
    printf("Finish.");
    return 0;
}