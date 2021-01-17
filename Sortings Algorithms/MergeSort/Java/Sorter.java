package merge;

public class Sorter extends Thread{

    int[] arr;
    int l, r;
    String id;

    public Sorter(int[] arr, int l, int r, String id) {
        this.arr = arr;
        this.l = l;
        this.r = r;
        this.id = id;
    }

    @Override
    public void run() {
        System.out.println("Sorter " + id + " started a merge.");
        MergeSort mergeSort = new MergeSort(id);
        mergeSort.sort(arr, l, r);
        System.out.println("Sorter " + id + " finished the merge.");
    }
}
