from colorama import Fore   
import random 


START = 0
STOP = 9999
max_num = 50
comparisons = 0
while_loop_execute = []


def insertion_sort(numbers):
    global comparisons
    global while_loop_execute

    if(numbers):

        print(Fore.RED+"",end="")
        print(*numbers,sep="  ")

        for i in range(1,len(numbers)):
            key=numbers[i]
            j=i-1
            
            compare_j = 1

            while(j>=0 and numbers[j]>=key):
                numbers[j+1]=numbers[j]
                j-=1
                compare_j += 1

            while_loop_execute.append(compare_j)  
            comparisons += compare_j

            numbers[j+1]=key

        print(Fore.GREEN+"",end="")
        print(*numbers,sep="  ")
        print(Fore.YELLOW+"",end="")
        print("While loop executed in each for : ", end="")
        print(*while_loop_execute,sep="  ")

def numbers_generator():
    global max_num
    return [random.randint(START, STOP) for i in range(0, max_num)]


if( __name__=='__main__'):
    numbers = numbers_generator()
    insertion_sort(numbers)
    print(f"Total number of comparisons : {comparisons}")