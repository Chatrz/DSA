#include <iostream>
#include <stack>
using namespace std;

struct node{
    int key;
    node *left;
    node *right;
};

void walk(node *root) {
    stack<node*> s;
    node *curr = root;
    while (!s.empty() || curr!=NULL){
        for (;curr!=nullptr; s.push(curr), curr = curr->left);
        curr = s.top();
        s.pop();
        cout << curr->key << endl;
        curr = curr->right;
    }
}   