// A C++ program to find k'th smallest element in a stream 
#include<iostream> 
#include<climits> 
using namespace std; 

void swap(int *x, int *y); 

class MinHeap 
{ 
	int *harr; 
	int capacity; 
	int heap_size; 
    public: 
        MinHeap(int a[], int size);
        void buildHeap(); 
        void MinHeapify(int i);
        int parent(int i) { return (i-1)/2; } 
        int left(int i) { return (2*i + 1); } 
        int right(int i) { return (2*i + 2); } 
        int extractMin(); 
        int getMin()	 { return harr[0]; } 
	    void replaceMin(int x) { harr[0] = x; MinHeapify(0); } 
}; 

MinHeap::MinHeap (int a[], int size) 
{ 
	heap_size = size; 
	harr = a; 
} 

void MinHeap::buildHeap () 
{ 
	int i = (heap_size - 1)/2; 
	while (i >= 0) 
	{ 
		MinHeapify(i); 
		i--; 
	} 
} 

int MinHeap::extractMin () 
{ 
	if (heap_size == 0) 
		return INT_MAX; 
	int root = harr[0]; 
	if (heap_size > 1) 
	{ 
		harr[0] = harr[heap_size-1]; 
		MinHeapify(0); 
	} 
	heap_size--; 
	return root; 
} 

void MinHeap::MinHeapify (int i) 
{ 
	int l = left(i); 
	int r = right(i); 
	int smallest = i; 
	if (l < heap_size && harr[l] < harr[i]) 
		smallest = l; 
	if (r < heap_size && harr[r] < harr[smallest]) 
		smallest = r; 
	if (smallest != i) 
	{ 
		swap(&harr[i], &harr[smallest]); 
		MinHeapify(smallest); 
	} 
} 

void swap (int *x, int *y) 
{ 
	int temp = *x; 
	*x = *y; 
	*y = temp; 
} 

void kthLargest (int k) 
{ 
	int count = 0, x; 
	int *arr = new int[k]; 
	MinHeap mh(arr, k); 
	while (1) 
	{ 
		cout << "Enter next element of stream "; 
		cin >> x; 
		if (count < k-1) 
		{ 
			arr[count] = x; 
			count++; 
		} 
		else
		{ 
            if (count == k-1) 
            { 
                arr[count] = x; 
                mh.buildHeap(); 
            } 
            else
            { 
                if (x > mh.getMin()) 
                    mh.replaceMin(x);
            } 
            cout << "K'th largest element is "
                << mh.getMin() << endl; 
            count++; 
		} 
	} 
} 

int main() 
{ 
	int k = 3; 
	cout << "K is " << k << endl; 
	kthLargest(k); 
	return 0; 
} 
