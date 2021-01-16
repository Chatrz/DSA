#include <iostream>
#include <string>

using namespace std;

class Node
{
    public:
        int key; // Box label
        int size; // Number of values
        Node* link; // Points to next node in front
        Node(int key);
};

Node::Node(int newKey) 
{
    key = newKey;
    size = 1;
    link = NULL;
}

// ************* Structure Methods ****************
Node* create_node(int newkey)
{
    return new Node(newkey);
}

void move_boxes(Node* pointers[], int i, int j)
{
    Node* destination = pointers[j];
    Node* source = pointers[i];

    if (source->link == NULL)
    {
        return;
    }

    if (destination->link == NULL)
    {
        destination->link = source->link;
        source->link = NULL;
        destination->size = source->size;
        source->size = 0;
    } else 
    {
        Node* current = destination->link;

        while (current->link != NULL)
        {
            current = current->link;
        }

        current->link = source->link;
        source->link = NULL;
        destination->size += source->size;
        source->size = 0;
    }
}

void init(Node* pointers[], int size)
{
    for (int i = 0; i < size; i++)
    {
        pointers[i] = create_node(-1); // Create a sentinel
        pointers[i]->link = create_node(i + 1); // Create the first box
    }
}

// ************* Program Methods ***************
void print(Node* pointers[], int d)
{
    Node* head = pointers[d];
    cout << head->size << " ";

    Node* current = head->link;
    while (current != NULL)
    {
        cout << current->key << " ";
        current = current->link;
    }
}

void command_input(Node* pointers[])
{
    int j, i;
    cin >> i;
    cin >> j;
    move_boxes(pointers, i - 1, j - 1);
}

int main() 
{
    int total_size, orders_number;

    cin >> total_size;
    Node* pointers[total_size];
    init(pointers, total_size);

    cin >> orders_number;
    for (int i = 0; i < orders_number; i++)
    {
        command_input(pointers);
    }

    int d;
    cin >> d;
    print(pointers, d - 1);
    return 0;
}