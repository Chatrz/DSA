/*
    ========================================================================
    || Data structures and Algorithems final project.                     ||
    || Title: 2D Range Tree                                               ||
    || Author: Asal Delkhosh                                              ||
    || Date: June 18th 2021                                               ||
    ========================================================================
*/
#include <iostream>
#include <string>
#include <stdlib.h>
#include <bits/stdc++.h>
#include <queue>

using namespace std;

// This is a constant used printing the three
#define COUNT 10 

// Data Structure
// ========================================================================
// ========================================================================
/**
*   This is a single point structure where
*   we store the data about a single point given.
*   @param x the x coordinate of the point 
*   @param y the y coordinate of the point
*/
class Point 
{
    public:
        float x;
        float y;
};

// Three functions to sort the array of points based on the argument and input type
bool compareByX(Point p1, Point p2) 
{
    return p1.x < p2.x;
}

bool compareByY(Point p1, Point p2) 
{
    return p1.y < p2.y;
}

bool compareByYPointer(Point* p1, Point* p2) 
{
    return p1->y < p2->y;
}

/**
* Node is our single element in the range tree.
* @param data is the highest value in the left tree of that node
* @param point is the struct that holdes the point of node
* @param left is the address of left element
* @param right is the address of right element
* @param tree is the address of the middle elements range tree
*/
class Node 
{
    public:
        float data;
        Point* point;
        Node* left;
        Node* right;
        Node* tree;
        Node(int nodeData, Point* nodePoint) // Constructor
        { 
            data = nodeData;
            left = NULL;
            right = NULL;
            point = nodePoint;
            tree = NULL;
        }
        float getMax();
};

/**
* This method gets the highest value in the left tree of an element.
*
*/
float Node::getMax() 
{
    Node* temp = left;
    while (temp->right != NULL)
    {
        temp = temp->right;
    }
    return temp->data;
}

/**
*   This method gets a list of points and then generates
*   a tree with those points, and returns the root pointer.
*
*/
Node* createRangeTreeReturnHead(int number, Point points[], int type) 
{ 
    queue<Node*> nodes;
    queue<Node*> store;
    for(int i = 0; i < number; i++)
    {
        Point* middle = new Point();
        middle->x = points[i].x;
        middle->y = points[i].y;
        if (type == 0)
            nodes.push(new Node(middle->x, middle));
        else
            nodes.push(new Node(middle->y, middle));
    }
    Node * head;
    while(!nodes.empty())
    {
        Node* temp1 = nodes.front();
        nodes.pop();
        Node* temp2 = nodes.front();
        nodes.pop();
        Node* temp = new Node(0, NULL);
        temp->left = temp1;
        temp->right = temp2;
        temp->data = temp->getMax();
        store.push(temp);
        if (nodes.size() <= 1)
        {
            Node* hold = NULL;
            if (nodes.size() == 1)
            {
                hold = nodes.front();
                nodes.pop();
            }
            while(!store.empty())
            {
                nodes.push(store.front());
                store.pop();
            }
            if (hold != NULL)
            {
                nodes.push(hold);
            } 
            if (nodes.size() == 1)
            {
                head = nodes.front();
                nodes.pop();
            }
        }
    }
    return head;
}

/**
*   This method gets the subelements of a root in a tree.
*
*/
list<Point> getChildren(Node* root, list<Point> nodeList)
{
    if (root->left == NULL && root->right == NULL) 
    {
        Point temp;
        temp.x = root->point->x;
        temp.y = root->point->y;
        nodeList.push_front(temp);
    } else 
    {
        nodeList = getChildren(root->left, nodeList);
        nodeList = getChildren(root->right, nodeList);
    }
    return nodeList;
}

/**
* This method generates the middle range trees for the 
* middle nodes.
*
*/
void createMiddleRangeTrees(Node* root) 
{
    if (root->point == NULL)
    {
        list<Point> childes;
        childes = getChildren(root, childes);
        childes.sort(compareByY);
        int number = childes.size();
        Point tempPoints[number];
        for (int i = 0; i < number; i++) 
        {
            tempPoints[i] = childes.front();
            childes.pop_front();
        }
        root->tree = createRangeTreeReturnHead(number, tempPoints, 1);
    } else {
        return;
    }
    createMiddleRangeTrees(root->left);
    createMiddleRangeTrees(root->right);
}
// ========================================================================
// ========================================================================

// Searching methods
// ========================================================================
// ========================================================================
list<Point*> searchY(float y1, float y2, Node* root, list<Point*>nodes, float oy1, float oy2)
 {
    if (root->point != NULL)
    {
        if (root->point->y >= oy1 && root->point->y <= oy2)
            nodes.push_front(root->point);
        return nodes;
    } 
    if (y1 > oy1 && y2 < oy2)
    {
        list<Point> temp;
        temp = getChildren(root, temp);
        int number = temp.size();
        for(int i = 0; i < number; i++) 
        {
            Point* pointTemp = new Point();
            pointTemp->x = temp.front().x;
            pointTemp->y = temp.front().y;
            nodes.push_front(pointTemp);
            temp.pop_front();
        }
    } else {
        if (root->point != NULL)
        {
            if (root->point->y >= oy1 && root->point->y <= oy2)
                nodes.push_front(root->point);
        } 
        if (root->data >= y1 && root->data <= y2) 
        {
            nodes = searchY(y1, root->data, root->left, nodes, oy1, oy2);
            nodes = searchY(root->data, y2, root->right, nodes, oy1, oy2);
        } else if (root->data > y2) 
        {
            nodes = searchY(y1, y2, root->left, nodes, oy1, oy2);
        }
        else if (root->data < y1) 
        {
            nodes = searchY(y1, y2, root->right, nodes, oy1, oy2);
        }
    }
    return nodes;
}

list<Point*> searchX(float x1, float x2, Node* root, list<Point*>nodes, float ox1, float oy1, float ox2, float oy2)
{
    if (root->point != NULL)
    {
        if (root->point->x >= ox1 && root->point->x <= ox2)
            if (root->point->y >= oy1 && root->point->y <= oy2)
                nodes.push_front(root->point);
        return nodes;
    }
    if (x1 > ox1 && x2 < ox2)
    {
        if (root->tree != NULL)
            nodes = searchY(oy1, oy2, root->tree, nodes, oy1, oy2);
        else {
            if (root->point->x >= x1 && root->point->x <= x2)
            {
                if (root->point->y >= oy1 && root->point->y <= oy2)
                    nodes.push_front(root->point);
            }
        }
    } else {
        if (root->point != NULL)
        {
            if (root->point->x >= x1 && root->point->x <= x2)
            {
                if (root->point->y >= oy1 && root->point->y <= oy2)
                    nodes.push_front(root->point);
            }
        } 
        if (root->data >= x1 && root->data <= x2) 
        {
            nodes = searchX(x1, root->data, root->left, nodes, ox1, oy1, ox2, oy2);
            nodes = searchX(root->data, x2, root->right, nodes, ox1, oy1, ox2, oy2);
        } else if (root->data > x2) 
        {
            nodes = searchX(x1, x2, root->left, nodes, ox1, oy1, ox2, oy2);
        }
        else if (root->data < x1) 
        {
            nodes = searchX(x1, x2, root->right, nodes, ox1, oy1, ox2, oy2);
        }
    }
    return nodes;
}
// ========================================================================
// ========================================================================

// For showing the output
void print2DUtil(Node *root, int space) 
{ 
    if (root == NULL) 
        return; 
    space += COUNT; 
    print2DUtil(root->right, space); 
    cout<<endl; 
    for (int i = COUNT; i < space; i++) 
        cout<<" ";
    if (root->point != NULL) 
        cout<<root->point->x<<","<<root->point->y<<"\n"; 
    else
        cout<<root->data<<"\n";
    print2DUtil(root->left, space); 
} 
  
void print2D(Node *root) 
{ 
    print2DUtil(root, 0); 
} 

int main() 
{
    // Inputs
    int number;
    cout << "Enter the number of points: ";
    cin >> number;

    Point points[number]; // Array of points

    cout << "Enter the $x$ values: " << endl;
    for(int i = 0; i < number; i++)
    {
        cin >> points[i].x;
    }
    cout << "Enter the $y$ values: " << endl;
    for(int i = 0; i < number; i++)
    {
        cin >> points[i].y;
    }

    int n = sizeof(points) / sizeof(points[0]);
    sort(points, points + n, compareByX); // Sorting points by X value

    // X range tree
    Node* treeHead = createRangeTreeReturnHead(number, points, 0);

    // 2D range tree generates here
    createMiddleRangeTrees(treeHead);

    // Testing the three
    int orders;
    cout << "How many inputs ? ";
    cin >> orders;
    for (int i = 0; i < orders; i++)
    {
        cout << "======================================================" << endl;
        cout << "======================================================" << endl;

        float x1, y1, x2, y2;

        cout << "Enter : ";
        cin >> x1;
        cin >> y1;
        cin >> x2;
        cin >> y2;

        list<Point*> temp;
        list<Point*> second;
        temp = searchX(x1, x2, treeHead, temp, x1, y1, x2, y2);
        temp.sort(compareByYPointer);

        int number = temp.size();
        if (number == 0)
        {
            cout << "None" << endl;
        }
        for(int j = 0; j < number; j++) 
        {
            cout << temp.front()->x << " ";
            second.push_back(temp.front());
            temp.pop_front();
        }
        cout << "\n";
        for(int j = 0; j < number; j++) 
        {
            cout << second.front()->y << " ";
            second.pop_front();
        }
        cout << "\n";

        cout << "======================================================" << endl;
        cout << "======================================================" << endl;
    }
    return 0;
}