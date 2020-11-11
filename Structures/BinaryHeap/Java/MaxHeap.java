public class MaxHeap{
    private int[] heapArr;
    private int size;
    public MaxHeap(){
        //creating a heap with maximum size 25
        heapArr=new int[25];
        size=0;
    }

    // create heap from unsorted array
    //heapify
    private void maxHeapify(int index){
        int leftEl=index*2+1;//in 0-indexed array the left element is in 2*i+
        int rightEl=index*2+2;
        int maxEl=index;
        if(leftEl<=size && heapArr[leftEl]>heapArr[index])maxEl=leftEl;
        if(rightEl<=size && heapArr[rightEl]>heapArr[maxEl])maxEl=rightEl;
        if(maxEl!=index){
            int temp =heapArr[index];
            heapArr[index]=heapArr[maxEl];
            heapArr[maxEl]=temp;
            maxHeapify(maxEl);
        }
    }
    //insert 
    //delete 
    //heap sort;
    //extract max 
    //increase key;
    //print heap

}