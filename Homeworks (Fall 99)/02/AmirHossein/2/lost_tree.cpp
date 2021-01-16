#include <iostream>

using namespace std;

class Node
{
    public:
        int key; // Value
        int l_index; // Left child
        int r_index; // Right child
};

int head = -1;
int current_value = 1;

// ************* Structure Methods ****************
void insert(Node* array, int index, int l, int r)
{
    if (head == l - 1 || head == r - 1)
    {
        head = index - 1;
    }
    array[index - 1].l_index = l - 1;
    array[index - 1].r_index = r - 1;
}

void value(Node* array, int index)
{
    if (array[index].l_index != -2)
    {
        value(array, array[index].l_index);
    }
    array[index].key = current_value;
    current_value++;
    if (array[index].r_index != -2)
    {
        value(array, array[index].r_index);
    }
}

void print(Node* array, int size)
{
    for (int i = 0; i < size; i++)
    {
        cout << array[i].key << " ";
    }
}

int main() 
{
    int size;
    cin >> size;
    Node* array = new Node[size];
    for (int i = 0; i < size; i++)
    {
        int index, l, r;
        cin >> index;
        cin >> l;
        cin >> r;
        insert(array, index, l, r);
        if (i == 0)
        {
            head = index - 1;
        }
    }
    value(array, head);
    print(array, size);
    return 0;
}