from colorama import Fore    
import time
def insertion_sort(numbers: list):
    if(numbers):
        for i in range(1,len(numbers)):
            key=numbers[i]
            j=i-1
            while(j>=0 and numbers[j]>=key):
                numbers[j+1]=numbers[j]
                j-=1
            numbers[j+1]=key
            time.sleep(0.8)
            printColoredList(numbers,i)
        print(Fore.GREEN+"",end="")
        print(*numbers,sep="  ")

    
def printColoredList(numbers: list,index: int):

    for i in range(0,len(numbers)):
        if(i>=index):
            print(Fore.RED+f"{numbers[i]} ",end=" ")
        else:
            print(Fore.GREEN+f"{numbers[i]} ",end=" ")
    print()

def strToIntList(strList: str):
    try:
        strNums=strList.split(" ")
        intList=[]
        for i in strNums:
            intList.append(int(i))
        return intList
    except ValueError:
        print(Fore.RED+"You should enter space sperated numbers in one line "+Fore.RESET,end="")


if( __name__=='__main__'):
    numbers=input("please enter the numbers u wish to be sorted\n")
    numbers=strToIntList(numbers)
    insertion_sort(numbers)
    print(Fore.RESET,end="")
    # printColoredList(numbers,)

    
