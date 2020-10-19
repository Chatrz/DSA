START = 100
STOP = 104
max_num = STOP - START
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
    global START
    global STOP
    return [2 * STOP - i for i in range(START, STOP)]


if( __name__=='__main__'):
    numbers = numbers_generator()
    print(" << ", end="")
    print(*numbers, sep="  ")
    selectionSort(numbers)
    print(" >> ", end="")
    print(*numbers, sep="  ")
    print(f"Total comparisons of {max_num} numbers = {swaps} times.")