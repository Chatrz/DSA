#include <iostream>
#include <string>

using namespace std;

class Node
{
    public:
        int key; // Value
        int space; // Number of values
        Node* link;
        Node(int key, int space);
};

Node::Node(int newKey, int totalSpace) 
{
    key = newKey;
    space = totalSpace;
    link = NULL;
}

Node* head = NULL;
Node* tail = NULL;

// ************* Structure Methods ****************
Node* create_node(int newkey, int totalSpace)
{
    return new Node(newkey, totalSpace);
}

void remove(int number)
{
    Node* temp = head;
    if (temp == NULL)
    {
        return;
    }
    if (number >= temp->space)
    {
        head = temp->link;
        if (head == NULL)
        {
            tail = NULL;
        }
        remove(number - temp->space);
        delete temp;
    } else
    {
        temp->space -= number;
    }
}

void insert(int newKey, int totalSpace)
{
    Node* newNode = create_node(newKey, totalSpace);
    if (head == NULL)
    {
        head = newNode;
        tail = newNode;
    } else {
        tail->link = newNode;
        tail = newNode;
    }
    return;
}

// ************* Program Methods ****************
void print()
{
    (head == NULL) ? cout << "empty" << endl : cout << head->key << endl;
}

void command_input()
{
    char order;
    int number, key;
    cin >> order;
    switch (order)
    {
    case '?':
        print();
        break;
    case '+':
        cin >> key;
        cin >> number;
        insert(key, number);
        break;
    case '-':
        cin >> number;
        remove(number);
        break;
    default:
        break;
    }
}

int main() 
{
    int orders_number;
    cin >> orders_number;
    for (int i = 0; i < orders_number; i++)
    {
        command_input();
    }
    return 0;
}