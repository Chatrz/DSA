import java.util.*;

public class Main {

    public static void main(String[] args) {
        /////////////////////////////////////////////////getting nodes
        Scanner scanner = new Scanner(System.in);
        int n = scanner.nextInt();
        // Map<Integer, Node> nodes = new HashMap<Integer, Node>();
        Node[] nodes = new Node[n + 1];
        for (int i = 0; i < n + 1; i++)
            nodes[i] = new Node(0, i);
        nodes[0].isParent = false;
        for (int i = 0; i < n; i++) {
            int index = scanner.nextInt();
            int left = scanner.nextInt();
            int right = scanner.nextInt();
            Node headNode = nodes[index];
            if (left != -1) {
                nodes[left].isParent = false;
                headNode.left = nodes[left];
            }
            if (right != -1) {
                nodes[right].isParent = false;
                headNode.right = nodes[right];
            }
        }
        //////////////////////////////////////////////////////////////////// setting root of tree
        BST tree = new BST();
        for (Node node : nodes) {
            if (node.isParent)
                tree.setRoot(node);
        }
        ////////////////////////////////////////////////////////////////////// setting keys
        tree.inorder(1);
        //////////////////////////////////////////////////////////////////////////////////
        for (Node node : nodes) {
            if (node.index != 0) {
                System.out.print(node.value);
                System.out.print(" ");
            }
        }
        System.out.println();
    }
}

class Node {
    int value;
    Node right;
    Node left;
    int index;
    boolean isParent;

    public Node(int value, int index) {
        this.value = value;
        this.index = index;
        isParent = true;
        right = null;
        left = null;
    }
}

class BST {
    private Node root;

    public BST() {
        root = null;
    }

    public void setRoot(Node root) {
        this.root = root;
    }

    public Node getRoot() {
        return root;
    }

    void inorder(int key) {
        if (root == null)
            return;
        Stack<Node> s = new Stack<Node>();
        Node curr = root;
        while (curr != null || s.size() > 0) {
            while (curr != null) {
                s.push(curr);
                curr = curr.left;
            }
            curr = s.pop();
            curr.value = key;
            key++;
            curr = curr.right;
        }
    }
}
