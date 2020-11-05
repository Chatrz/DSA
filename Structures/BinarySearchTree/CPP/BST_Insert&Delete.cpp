// This is just BST but only with two usages of inserting and deleting
#include <iostream>

using namespace std;

// Each tree object is a node that we define in this class
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

// Methods of tree :
// Create a node in memory and return the address
Node* create_node(int newkey)
{
    Node* temp = new Node(newkey);
    temp->parrent = NULL;
    temp->left = NULL;
    temp->right = NULL;
    return temp;
}

// Inserts a new node 
void insert(int newKey)
{
    Node* temp = create_node(newKey);
    Node* current = head;
    Node* parrent = NULL;
    while (current != NULL)
    {
        parrent = current;
        if (newKey <= current->key)
        {
            current = current->left;
        } else {
            current = current->right;
        }
    }
    temp->parrent = parrent;
    if (parrent == NULL)
    {
        head = temp;
    } else {
        if (newKey <= parrent->key)
        {
            parrent->left = temp;
        } else {
            parrent->right = temp;
        }
    }
}

// To find a special node
Node* search(int target)
{
    Node* current = head;
    while (current != NULL)
    {
        if (current->key == target)
            break;
        if (current-> key > target)
        {
            current = current->left;
        } else
        {
            current = current->right;
        }   
    }
    return current;
}

// Gets the minimum number in tree
Node* min(Node* root)
{
    while (root->left != NULL)
    {
        root = root->left;
    }
    return root;
}

// Succesors of a node
Node* succesore(int target)
{
    Node* current = search(target);
    if (current == NULL)
    {
        return NULL;
    }
    if (current->right != NULL)
    {
        return min(current->right);
    } else {
        Node* x = current->parrent;
        Node* y = current;
        while (x->parrent != NULL && y == x->right)
        {
            y = x;
            x = x->parrent;
        }
        return x;
    }
}

// Removing a node from binary search tree
void deletion(int target)
{
    Node* temp = search(target);
    if (temp == NULL)
    {
        return;
    }
    if (temp->left == NULL && temp->right == NULL)
    {
        if (temp == temp->parrent->left)
        {
            temp->parrent->left = NULL;
        } else {
            temp->parrent->right = NULL;
        }
        delete temp;
        return;
    } else {
        if (temp->right == NULL && temp->left != NULL)
        {
            if (temp == temp->parrent->left)
            {
                temp->parrent->left = temp->left;
                temp->left->parrent = temp->parrent;
            } else {
                temp->parrent->right = temp->left;
                temp->left->parrent = temp->parrent;
            }
            delete temp;
        } else if (temp->left == NULL && temp->right != NULL) 
        {
            if (temp == temp->parrent->left)
            {
                temp->parrent->left = temp->right;
                temp->right->parrent = temp->parrent;
            } else {
                temp->parrent->right = temp->right;
                temp->right->parrent = temp->parrent;
            }
            delete temp;
        } else {
            Node* current_succesore = succesore(target);
            int holder = current_succesore->key;
            deletion(holder);
            temp->key = holder;
        }
    }
}

// Prints the linked list objects
void printBT(const std::string& prefix, Node* node, bool isLeft)
{
    if ( node != NULL )
    {
        cout << prefix;
        cout << (isLeft ? "|__" : "|__" );
        // print the value of the node
        cout << node->key << endl;
        // enter the next tree level - left and right branch
        printBT( prefix + (isLeft ? "|   " : "    "), node->left, true);
        printBT( prefix + (isLeft ? "|   " : "    "), node->right, false);
    } else
    {
        cout << prefix;
        cout << (isLeft ? "|__" : "|__" );
        cout << "NULL" << endl;
    }
}


// A test case
int main() 
{
    insert(2);
    insert(10);
    insert(5);
    insert(8);
    insert(5);
    insert(20);
    insert(12);
    insert(16);
    insert(100);
    printBT("", head, false);
    deletion(20);
    printBT("", head, false);
    deletion(100);
    printBT("", head, false);
}