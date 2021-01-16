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
        Node* minNode(Node* currentNode){
            if (currentNode==nullNode)return nullNode;
            while(currentNode->left!=nullNode){
                currentNode=currentNode->left;
            }
            return currentNode;
        }
        Node* inorder_successor(Node* node){
            if (node==nullNode)return nullNode;
            if(node->right!=nullNode)return minNode(node->right);
            Node* parent=node->parent;
            while(parent!=nullNode && node==parent->right){
                node=parent;
                parent=parent->parent;
            }
            return parent;
        }
    public:
        RBTree(){
            nullNode=new Node(-1,black);
            root=nullNode;
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
        void RB_insert_fixup(Node* node){
            while(node->parent->color==red){
                Node* parentNode=node->parent;
                if(parentNode==parentNode->parent->left){
                    Node* uncle=parentNode->parent->right;
                    if(uncle->color=red){
                        node=flipColors(uncle->parent);//case 1
                    }else{
                        if(node==parentNode->right){//case 2
                            node=parentNode;
                            rotateLeft(node);
                        }
                        //case 3
                        node->parent->color=black;
                        node->parent->parent->color=red;
                        rotateRight(node->parent->parent);
                    }
                }
                else{
                    Node* uncle=parentNode->parent->left;
                    if(uncle->color=red){
                        node=flipColors(uncle->parent);//case 1
                    }else{
                        if(node==parentNode->left){//case 2
                            node=parentNode;
                            rotateRight(node);
                        }
                        //case 3
                        node->parent->color=black;
                        node->parent->parent->color=red;
                        rotateLeft(node->parent->parent);
                    }
                }
            }
        }
        Node* RB_delete(int key){
            Node* z=search(key); //z is the node that we want to delete
            if(z==nullNode){
                cout<<"There is no node with this key";
                return;
            }
            Node* y;//y is the node that gets deleted from tree
            y=(z->right==nullNode || z->left==nullNode)?z:this->inorder_successor(z);
            Node* x=y->left!=nullNode?y->left:y->right;
            x->parent=y->parent;
            if(y->parent==nullNode)root=x;
            else{
                if(y==y->parent->right)y->parent->right=x;
                else y->parent->left=x;
            }
            if(y!=z)z->key=y->key;
            if(y->color==black)this->RB_delete_fixup(x);
            return y;
        }
        void RB_delete_fixup(Node* x){
            while(x!=root && x->color==black){
                if(x==x->parent->left){
                    Node* brother=x->parent->right;
                    if(brother->color=red){
                        brother->color=black;
                        x->parent->color=red;
                        this->rotateRight(x->parent);
                        brother=x->parent->right;
                    }
                    if(brother->left->color==black && brother->right->color==black){
                        brother->color=red;
                        x=x->parent;
                    }else{
                        if(brother->right->color=black){
                            brother->left->color=black;
                            brother->color=red;
                            this->rotateRight(brother);
                            brother=x->parent->right;
                        }
                        brother->color=x->parent->color;
                        x->parent->color=black;
                        brother->right->color=black;
                        this->rotateLeft(x->parent);
                        x=root;
                    }
                }else{
                    Node* brother=x->parent->left;
                    if(brother->color=red){
                        brother->color=black;
                        x->parent->color=red;
                        this->rotateLeft(x->parent);
                        brother=x->parent->left;
                    }
                    if(brother->right->color==black && brother->left->color==black){
                        brother->color=red;
                        x=x->parent;
                    }else{
                        if(brother->left->color=black){
                            brother->right->color=black;
                            brother->color=red;
                            this->rotateLeft(brother);
                            brother=x->parent->left;
                        }
                        brother->color=x->parent->color;
                        x->parent->color=black;
                        brother->left->color=black;
                        this->rotateLeft(x->parent);
                        x=root;
                    }
                }
                x->color=black;
            }
        }
        Node* search(int key){
            Node* currentNode=root;
            if(root==nullNode)return root;
            while(currentNode!=nullNode){
                if(currentNode->key==key)return currentNode;
                if(currentNode->key<key)currentNode=currentNode->right;
                else currentNode=currentNode->left;
            }
            return nullNode;
        }
        Node* create_node(int key){
            Node* newNode=new Node(key,red);
            newNode->right=newNode->left=newNode->parent=nullNode;
            return newNode;
        }
};


int main (){
    
}