#include <iostream>
#include <string>
#include <vector>

using namespace std;
#define max(X, Y) ((X) > (Y) ? (X) : (Y))


// ----------------- Classes ------------------

// In this class I keep the information about a guest
// also with two pointers for their Slot linkedlist.
class Guest 
{
    // Used for storing a single guest information 
    public:
        long money;
        // To use in slot linked lists
        Guest* next_in;
        Guest* next_out;
        Guest()
        {
            next_in = NULL;
            next_out = NULL;
        }
};

// Each slot is a representor of a minute, that has two lists of guests comming-in and guests leaving
class Slot
{
    public:
        // Pointers for two in list and out list
        Guest* in_linkedlist;
        Guest* out_linkedlist;
        Slot()
        {
            in_linkedlist = NULL;
            out_linkedlist = NULL;
        }
        void add_to_in(Guest* guest_pointer)
        {
            guest_pointer->next_in = in_linkedlist;
            in_linkedlist = guest_pointer;
        }
        void add_to_out(Guest* guest_pointer)
        {
            guest_pointer->next_out = out_linkedlist;
            out_linkedlist = guest_pointer;
        }
};


// ----------------- Red-Black-Tree ------------------
class Node
{
    public:
        long data; // Which is the money of guest
        char color; // "b" for black and "r" for red
        // Pointers to other nodes
        Node* parent; 
        Node* left;
        Node* right;
        Node(long input_data, char input_color);
};
// Node constructor
Node::Node(long input_data, char input_color)
{
    data = input_data;
    color = input_color;
    left = NULL;
    right = NULL;
    parent = NULL;
}

// Creating two important part of red black trees
Node* nil = new Node(0, 'b');
Node* root = nil;

// RBT methods

// Search
Node* search(long key)
{
    Node* x = root;
    while (x != nil)
    {
        if (key < x->data)
        {
            x = x->left;
        } else if (key > x->data)
        {
            x = x->right;
        } else
        {
            break;
        }
    }
    return x;
}

// Maximum finder
Node* maximum()
{
    Node* current = root;
    if (root == nil)
    {
        return nil;
    }
    while (current->right != nil)
    {
        current = current->right;
    }
    return current;
}

// Minimum finder
Node* minimum(Node* subtree_root)
{
    if (subtree_root == nil)
    {
        return subtree_root;
    }
    while (subtree_root->left != nil)
    {
        subtree_root = subtree_root->left;
    }
    return subtree_root;
}

// Successor
Node* tree_successor(Node* temp)
{
    if (temp->right != nil)
    {
        return minimum(temp->right);
    }
    Node* y = temp->parent;
    while (y != root && temp == y->right)
    {
        temp = y;
        y = y->parent;
    }
    return y;
}

// -- Rotations --
void left_rotate(Node* x)
{
    Node* y = x->right;
    x->right = y->left;
    if (y->left != nil)
    {
        y->left->parent = x;
    }
    y->parent = x->parent;
    if (x->parent == nil)
    {
        root = y;
    } else if (x == x->parent->left)
    {
        x->parent->left = y;
    } else
    {
        x->parent->right = y;
    }
    y->left = x;
    x->parent = y;
}

void right_rotate(Node* x)
{
    Node* y = x->left;
    x->left = y->right;
    if (y->right != nil)
    {
        y->right->parent = x;
    }
    y->parent = x->parent;
    if (x->parent == nil)
    {
        root = y;
    } else if (x == x->parent->right)
    {
        x->parent->right = y;
    } else
    {
        x->parent->left = y;
    }
    y->right = x;
    x->parent = y;
}

// -- Insertion fix --
void rb_insert_fixup(Node* newnode)
{
    while (newnode->parent->color == 'r')
    {
        if (newnode->parent == newnode->parent->parent->left)
        {
            Node* y = newnode->parent->parent->right;
            if (y->color == 'r')
            {
                newnode->parent->color = 'b';
                y->color = 'b';
                newnode->parent->parent->color = 'r';
                newnode = newnode->parent->parent;
            } else
            {
                if (newnode == newnode->parent->right)
                {
                    newnode = newnode->parent;
                    left_rotate(newnode);
                }
                newnode->parent->color = 'b';
                newnode->parent->parent->color = 'r';
                right_rotate(newnode->parent->parent);
            }
        } else
        {
            Node* y = newnode->parent->parent->left;
            if (y->color == 'r')
            {
                newnode->parent->color = 'b';
                y->color = 'b';
                newnode->parent->parent->color = 'r';
                newnode = newnode->parent->parent;
            } else
            {
                if (newnode == newnode->parent->left)
                {
                    newnode = newnode->parent;
                    right_rotate(newnode);
                }
                newnode->parent->color = 'b';
                newnode->parent->parent->color = 'r';
                left_rotate(newnode->parent->parent);
            }
        }
    }
    root->color = 'b';
}

// -- Insert --
void insert(long key)
{
    // Creation
    Node* newnode = new Node(key, 'r');
    newnode->left = nil;
    newnode->right = nil;

    Node* y = nil;
    Node* x = root;
    // Placing
    while (x != nil)
    {
        y = x;
        if ( newnode->data < x->data )
            x = x->left;
        else
            x = x->right;
    }
    newnode->parent = y;
    if (y == nil)
    {
        root = newnode;
    } else if (newnode->data < y->data)
    {
        y->left = newnode;
    } else
    {
        y->right = newnode;
    }
    // Fixing
    rb_insert_fixup(newnode);
}

// -- Deletion fix --
void rb_delete_fixup(Node* x)
{
    while (x != root && x->color == 'b')
    {
        if (x == x->parent->left)
        {
            Node* w = x->parent->right;
            if (w->color == 'r'){
                w->color = 'b';
                x->parent->color = 'r';
                left_rotate(x->parent);
                w = x->parent->right;
            }
            if (w->left->color == 'b' && w->right->color == 'b')
            {
                w->color = 'r';
                x = x->parent;
            } else
            {
                if (w->right->color == 'b')
                {
                    w->left->color = 'b';
                    w->color = 'r';
                    right_rotate(w);
                    w = x->parent->right;
                }
                w->color = x->parent->color;
                x->parent->color = 'b';
                w->right->color = 'b';
                left_rotate(x->parent);
                x = root;
            }
        } else
        {
            Node* w = x->parent->left;
            if (w->color == 'r'){
                w->color = 'b';
                x->parent->color = 'r';
                right_rotate(x->parent);
                w = x->parent->left;
            }
            if (w->right->color == 'b' && w->left->color == 'b')
            {
                w->color = 'r';
                x = x->parent;
            } else
            {
                if (w->left->color == 'b')
                {
                    w->right->color = 'b';
                    w->color = 'r';
                    left_rotate(w);
                    w = x->parent->left;
                }
                w->color = x->parent->color;
                x->parent->color = 'b';
                w->left->color = 'b';
                right_rotate(x->parent);
                x = root;
            }
        }
    }
    x->color = 'b';
}

// -- Deletion --
void rb_transplant(Node* u, Node* v)
{
    if (u->parent == nil)
    {
        root = v;
    } else if (u == u->parent->left)
    {
        u->parent->left = v;
    } else
    {
        u->parent->right = v;
    }
    v->parent = u->parent;
}

void deletion(long key)
{
    Node* z = search(key);
    Node* y = z;
    Node* x = nil;
    char y_orginal_color = y->color;

    if (z->left == nil)
    {
        x = z->right;
        rb_transplant(z, z->right);
    } else if (z->right == nil)
    {   
        x = z->left;
        rb_transplant(z, z->left);
    } else
    {
        y = minimum(z->right);
        y_orginal_color = y->color;
        x = y->right;
        if (y->parent == z)
        {
            x->parent = y;
        } else
        {
            rb_transplant(y, y->right);
            y->right = z->right;
            y->right->parent = y;
        }
        rb_transplant(z, y);
        y->left = z->left;
        y->left->parent = y;
        y->color = z->color;
    }
    
    if (y_orginal_color == 'b')
        rb_delete_fixup(x);
    
}

// ------------ Finish RBT implemention ----------------
// ------------ Finish data structure and classes --------------

// ------------ Program functions --------------

// Then we update the tree in each minute and give the maximum.
void lets_party(long time, vector<Slot*> minutes)
{
    // Each round max money
    long value = 0;
    for (long i = 0; i < time; i++)
    {
        // Create a pointer for finding max
        Node* max_data = nil;

        // We update the Tree then
        Slot* s = minutes[i];
        // Inserts
        Guest* begin_in = s->in_linkedlist;
        while (begin_in != NULL)
        {
            insert(begin_in->money);
            begin_in = begin_in->next_in;
        }
        // Getting maximum
        max_data = maximum();
        // Deletions
        Guest* begin_out = s->out_linkedlist;
        while (begin_out != NULL)
        {
            deletion(begin_out->money);
            begin_out = begin_out->next_out;
        }
        // Setting the new max if needed
        if (max_data != nil)
        {
            value = max_data->data;
        } else
        {
            value = 0;
        }
        cout << value << " ";
    }
}

// In initialize we add the guests to the list of minutes.
void initialize(long people, long time, vector<Slot*> minutes)
{
    for (long i = 0; i < people; i++)
    {
        // Getting input
        long in_time, out_time, money;
        cin >> in_time;
        cin >> out_time;
        cin >> money;
        // Adding to lists
        Guest* temp = new Guest();
        temp->money = money;
        // Checking for comming in list
        minutes[in_time - 1]->add_to_in(temp);
        // Checking for comming out list
        minutes[out_time - 1]->add_to_out(temp);
    }
    lets_party(time, minutes);
}

// clearup for set the minutes to null at beginning
void define(long people, long time, vector<Slot*> minutes)
{
    for (long i = 0; i < time; i++)
    {
        Slot* temp = new Slot();
        minutes.push_back(temp);
    }
    initialize(people, time, minutes);
}

// Execute
int main() 
{
    long people, time;
    cin >> people;
    cin >> time;

    vector<Slot*> minutes;

    define(people, time, minutes);

    return 0;
}