#include <iostream>

using namespace std;
#define maximum(a,b) ( (a) > (b) ? (a) : (b))

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
    if (head == NULL)
    {
        head = temp;
        return;
    }
    while (true)
    {
        if (current->key >= newKey)
        {
            if (current->left == NULL)
            {
                current->left = temp;
                break;
            }
            current = current->left;
        } else {
            if (current->right == NULL)
            {
                current->right = temp;
                break;
            }
            current = current->right;
        }
    }
}

// An iteration from low to high in tree
void tree_walk(Node* start)
{
    if (start->left != NULL)
    {
        tree_walk(start->left);
    }
    cout << start->key << " ";
    if (start->right != NULL)
    {
        tree_walk(start->right);
    }
    return;
}

// Returns the size of the linked list
int size(Node* start, int x)
{
    if (start == NULL)
        return x; 
    return 1 + size(start->right, x) + size(start->left, x);
}

// Gets the minimum number in tree
int min(Node* root)
{
    while (root->left != NULL)
    {
        root = root->left;
    }
    return root->key;
}

// Gets the maximum number in tree
int max(Node* root)
{
    while (root->right != NULL)
    {
        root = root->right;
    }
    return root->key;
}

// Returns the hight of the tree
int hight(Node * root, int x)
{
    if (root == 0)
        return x;
    return 1 + maximum(hight(root->right, x), hight(root->left, x));
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
    cout << "\nTree Size = " << size(head, 0) << endl;
    cout << "Tree min = " << min(head) << " max = " << max(head) << endl;
    cout << "Tree hight = " << hight(head, 0) << endl;
    tree_walk(head);
    return 0;
}