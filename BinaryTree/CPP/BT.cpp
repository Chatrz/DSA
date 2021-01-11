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

// Returns the size of the linked list
int size(Node* start, int x)
{
    if (start == NULL)
        return x; 
    return 1 + size(start->right, x) + size(start->left, x);
}

// Returns the minimum number in tree
int min(Node* root)
{
    return root->key;
}

// Returns the maximum number in tree
int max(Node* root)
{
    if (root->left == NULL && root->right == NULL)
        return root->key; 
    if (root->left != NULL)
        return max(root->left);
    else
        return max(root->right);     
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
    if( node != NULL )
    {
        cout << prefix;
        cout << (isLeft ? "|__" : "|__" );
        // print the value of the node
        cout << node->key << "\n" <<endl;
        // enter the next tree level - left and right branch
        printBT( prefix + (isLeft ? "|   " : "    "), node->right, true);
        printBT( prefix + (isLeft ? "|   " : "    "), node->left, false);
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
    cout << "Tree hight = " << hight(head, 0);
    return 0;
}