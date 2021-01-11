import java.util.*;
public class recursiveSelection{
    public static void main(String[] args){
        Scanner input=new Scanner(System.in);
        int n=Integer.parseInt(input.nextLine());
        int[] arr=new int[n];
        for(int i=0;i<n;i++){
            arr[i]=input.nextInt();
        }
        selectionSort(arr, 0,arr.length-1);
        printArr(arr);
    }
    public static int findMin(int[] arr,int low,int high){
        int minIndex=low;
        for(int i=low;i<=high;i++){
            if(arr[i]<arr[minIndex]){
                minIndex=i;
            }
        }
        return minIndex;
    }
    public static void selectionSort(int[] arr,int low,int high){
        if(low==high)return;
        int minIndex=findMin(arr,low,high);
        int temp=arr[low];
        arr[low]=arr[minIndex];
        arr[minIndex]=temp;
        selectionSort(arr, low+1, high);
    }
    public static void printArr(int[] arr){
        for(int x:arr){
            System.out.print(""+x+" ");
        }
    }
}