import random

START = 0
STOP = 9999
max_num = 100
swaps = 0

def selectionSort(numbers):
    global swaps
    for i in range(0, len(numbers)-1):
        minIndex = i
        for j in range(i+1, len(numbers)):
            if(numbers[j] < numbers[minIndex]):
                minIndex = j
                swaps += 1
        temp = numbers[i]
        numbers[i] = numbers[minIndex]
        numbers[minIndex] = temp


def numbers_generator():
    global max_num
    return [random.randint(START, STOP) for i in range(0, max_num)]


if( __name__=='__main__'):
    numbers = numbers_generator()
    print(" << ", end="")
    print(*numbers, sep="  ")
    selectionSort(numbers)
    print(" >> ", end="")
    print(*numbers, sep="  ")
    print(f"Total comparisons of {max_num} numbers = {swaps} times.")