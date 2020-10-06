package merge;

public class MergeSort {

    String parent_id;

    public MergeSort(String parent_id){
        this.parent_id = parent_id;
    }

    void merge(int[] arr, int l, int m, int r)
    {
        int n1 = m - l + 1;
        int n2 = r - m;

        int[] L = new int[n1];
        int[] R = new int[n2];

        System.arraycopy(arr, l, L, 0, n1);
        System.arraycopy(arr, m + 1, R, 0, n2);

        int i = 0, j = 0, k = l;

        while (i < n1 && j < n2) {
            if (L[i] <= R[j]) {
                arr[k] = L[i];
                i++;
            }
            else {
                arr[k] = R[j];
                j++;
            }
            k++;
        }

        while (i < n1) {
            arr[k] = L[i];
            i++;
            k++;
        }
        while (j < n2) {
            arr[k] = R[j];
            j++;
            k++;
        }
    }

    void sort(int[] arr, int l, int r)
    {
        if (l < r) {
            int m = (l + r) / 2;

            Sorter sorter1 = new Sorter(arr, l, m, parent_id + ".0");
            Sorter sorter2 = new Sorter(arr, m + 1, r, parent_id + ".1");

            sorter1.start();
            sorter2.start();

            try {
                sorter1.join();
                sorter2.join();
            } catch (InterruptedException e) {
                e.printStackTrace();
            }

            merge(arr, l, m, r);
        }
    }
}
