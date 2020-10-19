import random
import time

START = 0
STOP = 99
max_num = 10

def bubble_sort(numbers): 
    n = len(numbers) 
    for i in range(n): 
        for j in range(0, n-i-1): 
            if numbers[j] > numbers[j+1] : 
                numbers[j], numbers[j+1] = numbers[j+1], numbers[j] 
                print(" > ", end="")
                print(*numbers, sep=" ")
                time.sleep(1)
        print(" < ", end="")            
        print(*numbers, sep=" ")
        time.sleep(1)

def numbers_generator():
    global max_num
    return [random.randint(START, STOP) for i in range(0, max_num)]


if( __name__=='__main__'):
    numbers = numbers_generator()
    print(" << ", end="")
    print(*numbers, sep="  ")
    bubble_sort(numbers)
    print(" >> ", end="")
    print(*numbers, sep="  ")