#include <iostream>
#include <string>

using namespace std;
#define max(X, Y)  ((X) > (Y) ? (X) : (Y))

// The main function for placing
int replace_losts(long array[], int size)
{
    int leaf = int(size / 2);
    for (int index = size-1; index > -1; index--)
    {
        int parent = int(index / 2) - (1 - (index % 2));
        if (array[index] == -1)
        {
            if (index >= leaf)
            {
                array[index] = 1;  
            } else
            {
                int left = 2 * index + 1, right = left + 1;
                array[index] = (right >= size) ? array[left] + 1 : max(array[left], array[right]) + 1;
            } 
        }
        if (parent > -1 && array[parent] != -1 && array[index] >= array[parent])
        {
            return -1;
        }
    }
    return 0;
}
// The main function for printing
void print(long array[], int size)
{
    for (int i = 0; i < size; i++)
    {
        cout << array[i] << " ";
    }
}

// Execute
int main() 
{
    int size;
    cin >> size;
    long array[size];
    for (int i = 0; i < size; i++)
    {
        cin >> array[i];
    }
    int result = replace_losts(array, size);
    if (result == -1)
    {
        cout << -1;
    } else
    {
        print(array, size);
    }
    return 0;
}