#include <iostream>
using namespace std;
enum Color {red,black};
class Node{
    public :
    Node* parent;
    Node* left;
    Node* right;
    Color color;
    int key;
    Node(int key,Color color){
        this->key=key;
        parent=left=right=NULL;
        this->color=color;
    }
};
class RBTree{
    private:
        Node* root;
        Node* nullNode;
        Node* flipColors(Node* parentOfParent){
            parentOfParent->color=red;
            parentOfParent->right->color=black;
            parentOfParent->left->color=black;
            return parentOfParent;
        }
        Node* rotateRight(Node* p,Node* x){

        }
        Node* rotateRight(Node* p,Node* x){
            
        }
    public:
        RBTree(){
            root=NULL;
            nullNode=new Node(-1,black);
        }
        void RB_insert(int key){
        Node* newNode=create_node(key);
        Node* parent=nullNode;
        Node* currentNode=root;
        while (currentNode!=nullNode){
            parent=currentNode;
            currentNode= newNode->key>currentNode->key?
                currentNode->right:currentNode->left;
        }
        newNode->parent=parent;
        if(parent==nullNode)root=newNode;
        else{
            if(newNode->key<parent->key)parent->left=newNode;
            else parent->right=newNode;
        }
            RB_insert_fixup(newNode);
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
            Node* newNode=new Node(key,red);
            newNode->right=newNode->left=nullNode;
            return newNode;
        }
};



