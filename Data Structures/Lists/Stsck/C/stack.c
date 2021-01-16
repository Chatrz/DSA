#include <stdio.h>

#define MAX(x, y) (((x) > (y)) ? (x) : (y))
#define MIN(x, y) (((x) < (y)) ? (x) : (y))
#define NELEMS(x) (sizeof(x) / sizeof((x)[0]))

const int LIMIT = 10;
int memory[LIMIT];
int pointer = -1;

void push(int a)
{
    if (pointer+1 == LIMIT)
        printf("Stack overflow.\n");
    else
    {
       pointer+=1;
       memory[pointer] = a;
    }
        
}

void pop()
{
    if (pointer == -1)
        printf("Stack underflow.\n");
    else
    {
        pointer-=1;
    }
    
}

int size()
{
    return pointer + 1;
}

void recover(int number)
{
    pointer = MAX(MIN(LIMIT - 1, number + pointer), -1);
}

void print()
{
    if (size() == 0)
        return;

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
    for(int i = 0; i < pointer; i++)
    {
        printf("    ");
    }
    printf(" | \n\n");
}

int main() 
{
    printf("This is how stack works >> \n");
    printf("Start:\n");
    print();
    printf("Push:\n");
    push(1);
    print();
    printf("Push:\n");
    push(2);
    print();
    printf("Push:\n");
    push(6);
    print();
    printf("Pop:\n");
    pop();
    print();
    printf("Pop:\n");
    pop();
    print();
    printf("Push:\n");
    push(7);
    print();
    printf("Recover:\n");
    recover(1);
    print();
    printf("Finish.");
    return 0;
}