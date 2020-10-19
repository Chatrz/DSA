#include <iostream>

using namespace std;

// Each link object is a node that we define in this class
class Node
{
    public:
        int key;
        Node* next;
        Node* prev;
        Node(int key);
};
// Node class constructor
Node::Node(int newKey) 
{
    key = newKey;
}

// Head address as a global variable
Node* head = NULL;
Node* tail = NULL;


// Methods of list :
// Create a node in memory and return the address
Node* create_node(int newkey)
{
    Node* temp = new Node(newkey);
    temp->next = NULL;
    temp->prev = NULL;
    return temp;
}

// Returns the size of the linked list
int size()
{
    int x = 0;
    Node* current = head;
    while (current != NULL)
    {
        current = current->next;
        x++;
    }
    return x;
}

// Inserts a new node at the head
void insert_at_head(int newKey)
{
    Node* temp = create_node(newKey);
    if (head == NULL)
    {
        head = temp;
        tail = temp;
        return;
    }
    temp->next = head;
    temp->next->prev = temp;
    head = temp;
    temp->prev = NULL;
    return;
}
// Adds a new node at tail
void insert_at_tail(int newKey)
{
    Node* temp = create_node(newKey);
    if (tail == NULL)
    {
        tail = temp;
        head = temp;
        return;
    }
    temp->prev = tail;
    temp->prev->next = temp;
    tail = temp;
    temp->next = NULL;
    return;
}

// Removes the node at the head
void remove_from_head()
{
    Node* temp = head;
    head = temp->next;
    head->prev = NULL;
    delete temp;
}
// Deletes the node at the tail
void remove_from_tail()
{
    Node* temp = tail;
    tail = temp->prev;
    tail->next = NULL;
    delete temp;
}

// Adds a new node anywhere in the list
void insert(int newKey, int index)
{
    Node* temp = head;
    if (index == 1)
    {
        insert_at_head(newKey);
        return;
    }
    if (index == size()+1)
    {
        insert_at_tail(newKey);
        return;
    }
    for (int i = 1; i < index-1; i++)
    {
        temp = temp->next;
        if (temp == NULL)
            return;
    }
    Node* newNode = create_node(newKey);
    newNode->next = temp->next;
    temp->next->prev = newNode;
    newNode->prev = temp;
    temp->next = newNode;
    return;
}

// Removes a node at any given place from the list
void remove(int index)
{
    Node* temp = head;
    if (index == 1)
    {
        remove_from_head();
        return;
    }
    if (index == size())
    {
        remove_from_tail();
        return;
    }
    for (int i = 1; i < index-1; i++)
    {
        temp = temp->next;
        if (temp == NULL)
            return;
    }
    Node* temp1 = temp->next;
    temp->next = temp1->next;
    temp->next->prev = temp;
    delete temp1;
}

// Reverses the link list
void reverse(Node* current)
{
    if (current->next == NULL)
    {
        Node* holder = head;
        head = tail;
        tail = holder;
        current->prev = NULL;
        return;
    }
    reverse(current->next);
    current->next->next = current;
    current->prev = current->next;
    current->next = NULL;
    return;
}

// Sortes the linked list based on the keys with insertion sort
void sort()
{
    int len = size();
    if (len == 0)
        return;
    Node* current = head->next;    
    for (int i = 1; i < len; i++)
    {
        int key = current->key;
        Node* temp = current->prev;
        while ((temp != NULL) && (key <= temp->key))
        {
            temp->next->key = temp->key;
            temp = temp->prev;
        }
        if (temp == NULL)
        {
            head->key = key;
        } else {
            temp->next->key = key;
        }
        current = current->next;
    }    
}

// Prints the linked list objects
void print_list()
{
    Node* temp = head;
    while (temp != NULL)
    {
        cout << temp->key << " ";
        temp = temp->next;
    }
    cout << endl;
}

// A test case
int main() 
{
    insert(2, 1); // Expect 2
    insert(20, 1); // Expect 20 2
    insert(23, 1); // Expect 23 20 2
    insert(7, 1); // Expect 7 23 20 2
    insert(4, 2); // Expect 7 4 23 20 2
    insert(10, 2); // Expect 7 10 4 23 20 2
    print_list();
    sort(); // Expect 2 4 7 10 20 23
    print_list();
    remove(3); // Expeect 2 4 10 20 23
    remove(2); // Expect 2 10 20 23
    print_list();
    remove(size()); // Expect 2 10 20 
    print_list();
    insert(2, 1); // Expect 2 2 10 20 
    reverse(head);
    print_list(); // Expect 20 10 2 2
    return 0;
}