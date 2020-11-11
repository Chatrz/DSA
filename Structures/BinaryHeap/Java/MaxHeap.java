public class MaxHeap {
    private int[] heapArr;
    private int lastEl;
    // private final int MAX_SIZE;
    public MaxHeap(){
        //creating a heap with maximum size 25
        heapArr=new int[25];
        // MAX_SIZE=25;
        lastEl=-1;
    }
    public MaxHeap(int[] heapArr){
        this.heapArr=heapArr;
        this.lastEl=heapArr.length-1;
        buildMaxHeap();
    }
    // create heap from unsorted array
    public void buildMaxHeap(){
        for (int i=(int)Math.floor((lastEl-1)/2);i>=0;i--){//watch out for the 0-index based
            maxHeapify(i);
        }
    }
    //heapify
    private void maxHeapify(int index){
        int leftEl=index*2+1;//in 0-indexed array the left element is in 2*i+
        int rightEl=index*2+2;
        int maxEl=index;
        if(leftEl<=lastEl&& heapArr[leftEl]>heapArr[index])
            maxEl=leftEl;
        if(rightEl<=lastEl && heapArr[rightEl]>heapArr[maxEl])
            maxEl=rightEl;
        if(maxEl!=index){
            int temp =heapArr[index];
            heapArr[index]=heapArr[maxEl];
            heapArr[maxEl]=temp;
            maxHeapify(maxEl);
        }
    }
    //insert
    public void insert(int key){
        //assuming we are inserting in the 25-sized heap
        if(lastEl>=24){
            System.out.println("heap overloaded");
            return;
        }
        heapArr[++lastEl]=Integer.MIN_VALUE;
        increaseKey(lastEl, key);
    } 
    //delete 
    //heap sort;
    //extract max 
    public int extractMax(){
        if (lastEl<0)
            System.out.println("heap is empty");
        int max=heapArr[0];
        heapArr[0]=heapArr[lastEl--];
        maxHeapify(0);
        return max;
    }
    //increase key;
    public void increaseKey(int index, int key){
        if(heapArr[index]>key){
            System.out.println("new key is smaller than the previous key.");
            return;
        }
        heapArr[index]=key;
        int parent=(int)Math.floor((index-1)/2);
        while (index>=0 && heapArr[parent]<heapArr[index]){
            int temp =heapArr[parent];
            heapArr[parent]=heapArr[index];
            heapArr[index]=temp;
            index=parent;
            parent=(int)Math.floor((index-1)/2);
        }
    }
    //print heap
    //print array
    public void printArr(){
        for(int i=0;i<=lastEl;i++){
            System.out.print(heapArr[i]+ " ");
        }
        System.out.println();
    }
}