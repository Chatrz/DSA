#include <iostream>

using namespace std;

// Each link object is a node that we define in this class
class Node
{
    public:
        int key;
        Node* parrent;
        Node* left;
        Node* right;
        Node(int key);
};
// Node class constructor
Node::Node(int newKey) 
{
    key = newKey;
}

// Head address as a global variable
Node* head = NULL;

// Methods of list :
// Create a node in memory and return the address
Node* create_node(int newkey)
{
    Node* temp = new Node(newkey);
    temp->parrent = NULL;
    temp->left = NULL;
    temp->right = NULL;
    return temp;
}

// Inserts a new node at the head
void insert(int newKey)
{
    Node* temp = create_node(newKey);
    if (head == NULL)
    {
        head = temp;
        return;
    }
    Node* current = head;
    while(true)
    {
        if (newKey < current->key)
        {
            temp->parrent = current->parrent;
            temp->left = current;
            if (current == current->parrent->left)
            {
                current->parrent->left = temp;
            } else {
                current->parrent->right = temp;
            }
            current->parrent = temp;
            return;
        }
        if (current->left == NULL)
        {
            current->left = temp;
            temp->parrent = current;
            return;
        }
        if (newKey > current->left->key)
        {
            current = current->left;
            continue;
        }
        if (current->right == NULL)
        {
            current->right = temp;
            temp->parrent = current;
            return;
        }
        if (newKey <= current->left->key)
        {
            current = current->right;
            continue;
        }
        
    }
}

// Do the deleting process
void do_delete(Node* parrent_side, Node* current)
{
    if (current->right == NULL)
    {
        parrent_side = current->left;
        if (current->left != NULL)
        {
            current->left->parrent = current->parrent;
        }
    } else {
        parrent_side = current->right;
        Node* right = current->right;
        while (right->left != NULL)
        {
            right = right->left;
        }
        right->left = current->left;
        if (current->left != NULL)
        {
            current->left->parrent = right;
        }
        current->right->parrent = current->parrent;
    }
    delete current;
    return;
}

// Removes the node at the head
void remove(int target)
{
    if (head == NULL)
    {
        return;
    }
    Node* current = head;
    while (current != NULL)
    {
        if (current->key == target)
        {
            if (current->parrent->left == current)
            {
                do_delete(current->parrent->left, current);
            } else {
                do_delete(current->parrent->right, current);
            }
            return;
        }  
        if (current->key < target)
        {
            current = current->left;
        } else {
            current = current->right;
        }
    }
    
}

// Returns the size of the linked list
int size(Node* start, int x)
{
    if (start == NULL)
        return x; 
    return 1 + size(start->right, x) + size(start->left, x);
}

// Prints the linked list objects
void print(Node* current)
{
    if (current == NULL)
    {
        return;
    }
    cout << "Self " << current->key << " [ ";
    cout << "Left ";
    print(current->left);
    cout << "Right ";
    print(current->right);
    cout << "] ";
    return;
}

// A test case
int main() 
{
    insert(2);
    insert(4);
    insert(3);
    print(head);
    cout << endl;
    insert(8);
    insert(5);
    print(head);
    cout << "Size = " << size(head, 0) << endl;
    return 0;
}