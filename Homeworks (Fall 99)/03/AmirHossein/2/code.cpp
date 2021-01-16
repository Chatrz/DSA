#include <iostream>
#include <string>

using namespace std;
#define max(X, Y) ((X) > (Y) ? (X) : (Y))

// ----------------- Classes ------------------
// The box to hold the pointers and the money of list
class Slot_Data
{
    public:
        long money;
        Slot_Data* next_in;
        Slot_Data* next_out;
        Slot_Data()
        {
            next_in = NULL;
            next_out = NULL;
        };
};

class Slot
{
    // Each slot is a representor of a minute, that has two lists of guests comming-in and guests leaving
    public:
        long slot_number;
        Slot* next;
        Slot_Data* in_linkedlist;
        Slot_Data* out_linkedlist;
        Slot()
        {
            next = NULL;
            in_linkedlist = NULL;
            out_linkedlist = NULL;
        };
        void add_to_in(Slot_Data* pointer)
        {
            pointer->next_in = in_linkedlist;
            in_linkedlist = pointer;
        };
        void add_to_out(Slot_Data* pointer)
        {
            pointer->next_out = out_linkedlist;
            out_linkedlist = pointer;
        };
};

// Head of slot linked list
Slot* head = NULL;

Slot* create_slot(long slot_number)
{
    Slot* current = head;
    Slot* prev = NULL;

    Slot* temp = new Slot();
    temp->slot_number = slot_number;

    // Empty list
    if (head == NULL)
    {
        head = temp;
        return temp;
    }

    if (head->slot_number > slot_number)
    {
        temp->next = head;
        head = temp;
        return temp;
    }

    while (current->next != NULL && current->next->slot_number < slot_number)
    {
        current = current->next;
    }

    temp->next = current->next;
    current->next = temp;

    return temp;
}

// Search for slot
Slot* get_slot(long slot_number)
{
    Slot* current = head;
    while (current != NULL && current->slot_number != slot_number)
    {
        current = current->next;
    }
    return current;
}

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
    while (y != nil && temp == y->right)
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
void deletion(long key)
{
    Node* z = search(key);
    Node* y = NULL;
    Node* x = NULL;
    if (z->left == nil || z->right == nil)
    {
        y = z;
    } else
    {
        y = tree_successor(z);
    }
    if (z->left != nil)
    {
        x = y->left;
    } else
    {
        x = y->right;
    }
    x->parent = y->parent;
    if (y->parent == nil)
    {
        root = x;
    } else if (y == y->parent->left)
    {
        y->parent->left = x;
    } else
    {
        y->parent->right = x;
    }
    if (y != z)
    {
        z->data = y->data;
    }
    if (y->color == 'b')
    {
        rb_delete_fixup(x);
    }
}
// ------------ Finish RBT implement ----------------

void see_all()
{
    Slot* current = head;
    while (current != NULL)
    {
        cout << current->slot_number << " ";
        current = current->next;
    }
    cout << endl;
}

// ------------ Finish data structure and classes --------------
// ------------ Program functions --------------
void initialize(long number)
{
    for (long i = 0; i < number; i++)
    {
        // Getting input
        long in_time, out_time, money;
        cin >> in_time;
        cin >> out_time;
        cin >> money;
        // Adding to lists
        Slot_Data* temp = new Slot_Data();
        temp->money = money;

        Slot* in_slot = get_slot(in_time);
        Slot* out_slot = get_slot(out_time);

        if (in_slot == NULL)
        {
            in_slot = create_slot(in_time);
        } 
        in_slot->add_to_in(temp);
        if (out_slot == NULL)
        {
            out_slot = create_slot(out_time);
        } 
        out_slot->add_to_out(temp);
        //see_all();
    }
}

void lets_party(long number)
{
    Slot* current = head;
    long value = 0;
    for (long i = 1; i <= number; i++)
    {
        Node* max_tree = maximum();
        if (current != NULL && i == current->slot_number)
        {
            Slot_Data* begin_out = current->out_linkedlist;
            while (begin_out != NULL)
            {
                deletion(begin_out->money);
                begin_out = begin_out->next_out;
            }
            Slot_Data* begin_in = current->in_linkedlist;
            while (begin_in != NULL)
            {
                insert(begin_in->money);
                begin_in = begin_in->next_in;
            }
            current = current->next;
        }
        Node* new_max = maximum();
        value = max(max_tree->data, new_max->data);
        cout << value << " ";
    }
}

// Execute
int main() 
{
    long number, minute;
    cin >> number;
    cin >> minute;
    
    initialize(number);
    lets_party(minute);

    return 0;
}