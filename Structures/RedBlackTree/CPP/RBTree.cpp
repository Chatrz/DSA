#include <iostream>
using namespace std;
class Node{
    public :
    Node* parent;
    Node* left;
    Node* right;
    int key;
    Node(int key){
        this->key=key;
        parent=left=right=NULL;
    }
};
class RBTree{
    Node* root;
    public:
    Node* RB_insert(int key){

    }
    void RB_insert_fixup(Node* x){

    }
    Node* RB_delete(int key){

    }
    void RB_delete_fixup(Node* x){

    }
    Node* BST_search(int key){

    }
    Node* create_node(int key){
        return new Node(key);
    }
};



