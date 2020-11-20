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
        Node* rotateRight(Node* node){     
            Node* leftNode=node->left;
            node->left=leftNode->right;
            if(leftNode->right!=nullNode)leftNode->right->parent=node;
            leftNode->parent=node->parent;
            if(node->parent==nullNode)
                root=leftNode;
            else{
                if(node==node->parent->right)node->parent->right=leftNode;
                else node->parent->left=leftNode;
            }
            leftNode->right=node;
            node->parent=leftNode;
            return leftNode;
        }
        Node* rotateLeft(Node* node){
            Node* rightNode=node->right;
            node->right=rightNode->left;
            if(rightNode->left!=nullNode)rightNode->left->parent=node;
            rightNode->parent=node->parent;
            if(node->parent==nullNode)
                root=rightNode;
            else{
                if(node==node->parent->right)node->parent->right=rightNode;
                else node->parent->left=rightNode;
            }
            rightNode->left=node;
            node->parent=rightNode;
            return rightNode;
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
            newNode->right=newNode->left=newNode->parent=nullNode;
            return newNode;
        }
};



