def findMinIndex(lis):
    min=lis[0]
    minIndex=0
    for i in range(0,len(lis)):
        if(lis[i]<=min):
            minIndex=i
    return minIndex

def selectionSort(lis,size):
    for i in range(0,len(lis)-1):
        minIndex=i
        for j in range(i+1,len(lis)):
            if(lis[j]<lis[minIndex]):
                minIndex=j
        temp=lis[i]
        lis[i]=lis[minIndex]
        lis[minIndex]=temp

lis=[3,2,7,1,5,4]
selectionSort(lis,6)
print(lis)