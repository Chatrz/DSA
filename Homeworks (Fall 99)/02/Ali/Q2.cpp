#include <iostream>

using namespace std;
int counter = 1;

struct Node {
    int value;
    int name;
    Node *right;
    Node *left;
    Node *parent;

    Node(int value,Node *right, Node *left, Node *parent) : value(value), right(right),
                                                                       left(left), parent(parent) {}
};

void traverse(Node *rootOfSubTree) {
    if (rootOfSubTree == nullptr) {
        return;
    }
    traverse(rootOfSubTree->left);
    rootOfSubTree->value = counter;
    counter++;
    traverse(rootOfSubTree->right);
}

int main() {
    int n;
    cin >> n;
    Node *nodes[n];
    Node *root;
    for (int i = 0; i < n; ++i) {
        nodes[i] = new Node(-1,nullptr, nullptr, nullptr);
    }
    for (int i = 0; i < n; ++i) {
        int parent, leftChild, rightChild;
        cin >> parent;
        cin >> leftChild;
        cin >> rightChild;
        if (leftChild != -1) {
            nodes[parent - 1]->left = nodes[leftChild - 1];
            nodes[leftChild - 1]->parent = nodes[parent - 1];
        }
        if (rightChild != -1) {
            nodes[parent - 1]->right = nodes[rightChild - 1];
            nodes[rightChild - 1]->parent = nodes[parent - 1];
        }
    }
    for (int i = 0; i < n; ++i) {
        if (nodes[i]->parent == nullptr) {
            root = nodes[i];
            break;
        }
    }
    traverse(root);
    for (int i = 0; i < n; ++i) {
        cout << nodes[i]->value << " ";
    }

    return 0;
}