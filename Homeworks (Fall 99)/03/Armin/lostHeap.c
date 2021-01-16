#include <stdio.h>

int parent(int index) {
    return (index - 1) / 2;
}

int left(int index) {
    return 2 * index + 1;
}

int right(int index) {
    return 2 * index + 2;
}

int max(int num1, int num2) {
    return (num1 > num2) ? num1 : num2;
}

int main() {
    int n, flag;
    flag = 1;
    scanf("%d", &n);
    int array[n];
    for (int i = 0; i < n; i++) {
        scanf("%d", &array[i]);
    }
    for (int j = n - 1; j >= 0; j--) {
        if (array[j] == -1) {
            int leftIndex = left(j);
            int rightIndex = right(j);
            if (leftIndex >= n && rightIndex >= n) { // has no children
                array[j] = 1;
            } else if (leftIndex >= n) { // has 1 children at right
                array[j] = array[rightIndex] + 1;
            } else if (rightIndex >= n) { // has 1 children at left
                array[j] = array[leftIndex] + 1;
            } else { // has 2 children
                array[j] = max(array[rightIndex], array[leftIndex]) + 1;
            }
        }
        int parentIndex = parent(j);
        if (array[parentIndex] <= array[j] && array[parentIndex] != -1 && j != 0) {
            flag = 0;
            break;
        }
    }
    if (flag != 1) {
        printf("-1");
    } else {
        for (int i = 0; i < n; i++) {
            printf("%d ", array[i]);
        }
    }
    printf("\n");
    return 0;
}

