class GFG {
    static class Node {
        int data;
        Node left, right;
    };

    static void printInorder(Node node) {
        if (node == null)
            return;
        printInorder(node.left);
        System.out.print(node.data + " ");
        printInorder(node.right);
    }

    static Node buildCartesianTreeUtil(int root, int arr[], int parent[], int leftchild[], int rightchild[]) {
        if (root == -1)
            return null;

        Node temp = new Node();
        temp.data = arr[root];

        temp.left = buildCartesianTreeUtil(leftchild[root], arr, parent, leftchild, rightchild);
        temp.right = buildCartesianTreeUtil(rightchild[root], arr, parent, leftchild, rightchild);

        return temp;
    }

    static Node buildCartesianTree(int arr[], int n) {
        int[] parent = new int[n];
        int[] leftchild = new int[n];
        int[] rightchild = new int[n];

        memset(parent, -1);
        memset(leftchild, -1);
        memset(rightchild, -1);

        int root = 0, last;

        for (int i = 1; i <= n - 1; i++) {
            last = i - 1;
            rightchild[i] = -1;

            while (arr[last] <= arr[i] && last != root)
                last = parent[last];

            if (arr[last] <= arr[i]) {
                parent[root] = i;
                leftchild[i] = root;
                root = i;
            } else if (rightchild[last] == -1) {
                rightchild[last] = i;
                parent[i] = last;
                leftchild[i] = -1;
            } else {
                parent[rightchild[last]] = i;
                leftchild[i] = rightchild[last];
                rightchild[last] = i;
                parent[i] = last;
            }
        }

        parent[root] = -1;

        return (buildCartesianTreeUtil(root, arr, parent, leftchild, rightchild));
    }

    static void memset(int[] arr, int value) {
        for (int i = 0; i < arr.length; i++) {
            arr[i] = value;
        }

    }

    public static void main(String[] args) {
        int arr[] = { 5, 10, 40, 30, 28 };
        int n = arr.length;

        Node root = buildCartesianTree(arr, n);

        System.out.printf("Inorder traversal of the" + " constructed tree : \n");
        printInorder(root);
    }
}
