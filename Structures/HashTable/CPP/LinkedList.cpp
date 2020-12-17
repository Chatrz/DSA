#include <iostream>
using namespace std;

class Node{
    public:
    int key;
    Node* next;
    Node* prev;
    Node(int key){
        this->key=key;
        next=NULL;
        prev=NULL;
    }
};


class LinkedList{
    private:
        Node* head;
        Node* tail;
        Node* createNode(int key){
            Node* newNode=new Node(key);
            return newNode;
        }
    public:
        LinkedList(){
            head=NULL;
            tail=NULL;
        }
        void insert_key(int key){
            Node* newNode=createNode(key);
            if(head==NULL){
                head=tail=newNode;
                return;
            }
            newNode->next=head;
            head->prev=newNode;
            head=newNode;
            return;
        };
        Node* delete_key(int key){
            Node* nodeToDel=search_key(key);
            if(nodeToDel==head){
                nodeToDel->next->prev=NULL;
                head=nodeToDel->next;
                //delete(nodeToDel);
            }else if(nodeToDel==tail){
                nodeToDel->prev->next=NULL;
                tail=nodeToDel->prev;
                //delete(nodeToDel);
            }else{
                nodeToDel->prev->next=nodeToDel->next;
                nodeToDel->next->prev=nodeToDel->prev;
                //delete(nodeToDel);
            }
            return nodeToDel;
        };
        Node* search_key(int key){
            if(head==NULL)return NULL;
            Node* currentNode=head;
            while(currentNode!=NULL){
                if(currentNode->key==key)return currentNode;
                currentNode=currentNode->next;
            }
            return NULL;

        };
        void print_list(){
            Node* current=head;
            while(current!=NULL){
                cout<<current->key<<" ";
            }
        };
};