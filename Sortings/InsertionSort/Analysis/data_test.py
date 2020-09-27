from colorama import Fore   
import random as R
import threading as TR
import datetime 
import time as T

DONE_FLAG = 0	
START = 0
STOP = 999999
max_num = 200
THREADS = 20
compare_numbers = []

def average(lst): 
    return sum(lst) / len(lst) 

def insertion_sort(numbers):
    comparisons = 0

    if(numbers):

        for i in range(1,len(numbers)):
            key=numbers[i]
            j=i-1
            
            compare_j = 1

            while(j>=0 and numbers[j]>=key):
                numbers[j+1]=numbers[j]
                j-=1
                compare_j += 1

            comparisons += compare_j

            numbers[j+1]=key

        return comparisons


def numbers_generator():
    global max_num
    return [R.randint(START, STOP) for i in range(0, max_num)]

def create_sorter(thread_id):
    global compare_numbers
    global DONE_FLAG
    start_time = datetime.datetime.now()

    numbers = numbers_generator()

    print(Fore.YELLOW,end="")
    print(f">> Sorter {thread_id + 1} start sorting {len(numbers)} numbers. Max = {max(numbers)} / Min = {min(numbers)}")
    print(Fore.RESET,end="")

    compare = insertion_sort(numbers) 
    compare_numbers.append(compare)

    finish_time = datetime.datetime.now()
    total_time = (finish_time - start_time).total_seconds()

    print(Fore.GREEN+"",end="")  
    print(f"<< Sorter {thread_id + 1} is done. Time: {total_time} seconds / Comparisons: {compare}")
    print(Fore.RESET,end="")
    DONE_FLAG += 1


if( __name__=='__main__'):
    print(Fore.RED,end="")
    print("\n>> Run\n")

    for i in range(0, THREADS):
        t = TR.Thread(target=create_sorter, args=(i,))
        t.start()

    print(Fore.RED,end="")
    print("\n>> Done\n") 

    while True:
    	if DONE_FLAG == THREADS:
    		break
    print(">> Threads Finish")		
    	
    print("\n")
    print(Fore.YELLOW,end="")
    print(f"Totally {THREADS} threads tested with average comparisons of {round(average(compare_numbers), 2)}")
