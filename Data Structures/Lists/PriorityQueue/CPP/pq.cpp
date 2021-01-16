#include <iostream>
#include <queue>

using namespace std;

void showpq(priority_queue<int> gq)
{
	priority_queue<int> g = gq;
	while (!g.empty()) {
		cout << '\t' << g.top();
		g.pop();
	}
	cout << '\n';
}

int main()
{
	priority_queue<int> pq;
	pq.push(10);
	pq.push(30);
	pq.push(20);
	pq.push(5);
	pq.push(1);

	cout << "The priority queue is : ";
	showpq(pq);

	cout << "\npq.size() : " << pq.size();
	cout << "\npq.top() : " << pq.top();

	cout << "\npq.pop() : ";
	pq.pop();
	showpq(pq);

	return 0;
}
