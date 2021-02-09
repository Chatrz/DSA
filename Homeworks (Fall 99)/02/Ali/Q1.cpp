#include <iostream>

using namespace std;

class SuperQueue {

    struct QueueNode {
        int data;
        int count;
        QueueNode *next;

        QueueNode(int d, int c) {
            data = d;
            count = c;
            next = nullptr;
        }
    };

private:
    QueueNode *front, *rear;

public:
    SuperQueue() {
        front = nullptr;
        rear = nullptr;
    }

    void push(int number, int times) {
        QueueNode *newNode = new QueueNode(number, times);
        if (rear == nullptr) {
            rear = newNode;
            front = newNode;
        } else {
            rear->next = newNode;
            rear = newNode;
        }
    }

    void pop(int times) {

        while (times > 0) {
            if (times - front->count >= 0) {
                times -= front->count;
                QueueNode *temp = front;
                front = front->next;
                if (front == nullptr)
                    rear = nullptr;
                delete (temp);
            } else {
                front->count -= times;
                break;
            }
        }
    }

    void getFront() {
        if (front == nullptr)
            cout<<("empty")<<endl;
        else
            cout<<(front->data)<<endl;
    }
};

int main() {
    SuperQueue superQueue;
    int x;
    cin >> x;
    for (int i = 0; i < x; ++i) {
        char operand;
        string command;
        cin >> operand;
        if (operand == '+') {
            int number, times;
            cin >> number;
            cin >> times;
            superQueue.push(number, times);
        } else if (operand == '-') {
            int times;
            cin >> times;
            superQueue.pop(times);
        } else if (operand == '?') {
            superQueue.getFront();
        }
    }
    return 0;
}
