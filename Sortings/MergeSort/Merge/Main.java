package merge;

import java.util.Random;

/**
 * Created a program that does the merge sort
 * with threads, in each sorting we create a thread to
 * sort a subarray of the main array and the return the sorted array.
 * Each thread has an IP related to the parent thread that created it,
 * so you can check the way that the array goes to be sorted.
 * To check this sorting algorithm at its best you can change the
 * size and start and stop parameters in "numberGenerator" to see a big
 * sorting.
 *
 * Created By "AM.H_NJ.Z" with Contribution of "sinamna".
 * Check our git repository "https://github.com/sinamna/DSA_Trainings.git"
 * for more Data Structures & Algorithms.
 */

class Main {

    static void printArray(int[] arr)
    {
        for (int value : arr) System.out.print(value + " ");
        System.out.println("\n");
    }

    static int[] numberGenerator() {
        Random random = new Random();
        int[] array = new int[100];
        for (int i = 0; i < 100; i++){
            array[i] = random.nextInt(1000);
        }
        return array;
    }

    public static void main(String[] args)
    {
        int[] arr = numberGenerator();

        System.out.println("Given Array");
        printArray(arr);

        MergeSort ob = new MergeSort("127");
        ob.sort(arr, 0, arr.length - 1);

        System.out.println("\nSorted array");
        printArray(arr);
    }
}
