import java.io.*;

class Sp {

    static int MAX = 500;
    static int[][] lookup = new int[MAX][MAX];

    static void buildSparseTable(int arr[], int n) {
        for (int i = 0; i < n; i++)
            lookup[i][0] = arr[i];

        for (int j = 1; (1 << j) <= n; j++) {
            for (int i = 0; (i + (1 << j) - 1) < n; i++) {
                if (lookup[i][j - 1] < lookup[i + (1 << (j - 1))][j - 1])
                    lookup[i][j] = lookup[i][j - 1];
                else
                    lookup[i][j] = lookup[i + (1 << (j - 1))][j - 1];
            }
        }
    }

    static int query(int L, int R) {
        int j = (int) Math.log(R - L + 1);
        if (lookup[L][j] <= lookup[R - (1 << j) + 1][j])
            return lookup[L][j];
        else
            return lookup[R - (1 << j) + 1][j];
    }

    public static void main(String[] args) {
        int a[] = { 7, 2, 3, 0, 5, 10, 3, 12, 18 };
        int n = a.length;

        buildSparseTable(a, n);

        System.out.println(query(0, 4));
        System.out.println(query(4, 7));
        System.out.println(query(7, 8));

    }
}
