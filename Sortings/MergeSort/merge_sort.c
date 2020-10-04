#include <stdio.h> 
#include <stdlib.h> 

void merge(int arr[], int left, int middle, int right) 
{ 
	int i = 0, j = 0; 
	int n1 = middle - left + 1; 
	int n2 = right - middle; 

	int left_arr[n1], right_arr[n2]; 

	for (i = 0; i < n1; i++) 
		left_arr[i] = arr[left + i]; 
	for (j = 0; j < n2; j++) 
		right_arr[j] = arr[middle + 1 + j]; 

	i = 0; 
	j = 0; 
	int k = left; 
	while (i < n1 && j < n2) { 
		if (left_arr[i] <= right_arr[j]) { 
			arr[k] = left_arr[i]; 
			i++; 
		} 
		else { 
			arr[k] = right_arr[j]; 
			j++; 
		} 
		k++; 
	} 

	while (i < n1) { 
		arr[k] = left_arr[i]; 
		i++; 
		k++; 
	} 

	while (j < n2) { 
		arr[k] = right_arr[j]; 
		j++; 
		k++; 
	} 
} 

void mergeSort(int arr[], int left, int right) 
{ 
	if (left < right) { 
		int middle = left + (right - left) / 2; 

		mergeSort(arr, left, middle); 
		mergeSort(arr, middle + 1, right); 

		merge(arr, left, middle, right); 
	} 
} 

void printArray(int A[], int size) 
{ 
	int i; 
	for (i = 0; i < size; i++) 
		printf("%d ", A[i]); 
	printf("\n"); 
} 

int main() 
{ 
	int arr[] = { 12, 11, 13, 5, 6, 7 }; 
	int arr_size = sizeof(arr) / sizeof(arr[0]); 

	printf("Given array is \n"); 
	printArray(arr, arr_size); 

	mergeSort(arr, 0, arr_size - 1); 

	printf("\nSorted array is \n"); 
	printArray(arr, arr_size); 
	return 0; 
} 
