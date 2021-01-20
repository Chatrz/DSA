public class SplayBST<Key extends Comparable<Key>, Value> {
    private Node root;

    private class Node {
        private Key key;
        private Value value;
        private Node left, right;

        public Node(Key key, Value value) {
            this.key = key;
            this.value = value;
        }
    }

    public boolean contains(Key key) {
        return get(key) != null;
    }

    public Value get(Key key) {
        root = splay(root, key);
        int cmp = key.compareTo(root.key);
        if (cmp == 0)
            return root.value;
        else
            return null;
    }

    public void put(Key key, Value value) {
        if (root == null) {
            root = new Node(key, value);
            return;
        }
        root = splay(root, key);
        int cmp = key.compareTo(root.key);
        if (cmp < 0) {
            Node n = new Node(key, value);
            n.left = root.left;
            n.right = root;
            root.left = null;
            root = n;
        } else if (cmp > 0) {
            Node n = new Node(key, value);
            n.right = root.right;
            n.left = root;
            root.right = null;
            root = n;
        } else {
            root.value = value;
        }
    }

    public void remove(Key key) {
        if (root == null)
            return;
        root = splay(root, key);
        int cmp = key.compareTo(root.key);
        if (cmp == 0) {
            if (root.left == null) {
                root = root.right;
            } else {
                Node x = root.right;
                root = root.left;
                splay(root, key);
                root.right = x;
            }
        }
    }

    private Node splay(Node h, Key key) {
        if (h == null)
            return null;
        int cmp1 = key.compareTo(h.key);
        if (cmp1 < 0) {
            if (h.left == null) {
                return h;
            }
            int cmp2 = key.compareTo(h.left.key);
            if (cmp2 < 0) {
                h.left.left = splay(h.left.left, key);
                h = rotateRight(h);
            } else if (cmp2 > 0) {
                h.left.right = splay(h.left.right, key);
                if (h.left.right != null)
                    h.left = rotateLeft(h.left);
            }
            if (h.left == null)
                return h;
            else
                return rotateRight(h);
        } else if (cmp1 > 0) {
            if (h.right == null) {
                return h;
            }
            int cmp2 = key.compareTo(h.right.key);
            if (cmp2 < 0) {
                h.right.left = splay(h.right.left, key);
                if (h.right.left != null)
                    h.right = rotateRight(h.right);
            } else if (cmp2 > 0) {
                h.right.right = splay(h.right.right, key);
                h = rotateLeft(h);
            }
            if (h.right == null)
                return h;
            else
                return rotateLeft(h);
        } else
            return h;
    }

    public int height() {
        return height(root);
    }

    private int height(Node x) {
        if (x == null)
            return -1;
        return 1 + Math.max(height(x.left), height(x.right));
    }

    public int size() {
        return size(root);
    }

    private int size(Node x) {
        if (x == null)
            return 0;
        else
            return 1 + size(x.left) + size(x.right);
    }

    private Node rotateRight(Node h) {
        Node x = h.left;
        h.left = x.right;
        x.right = h;
        return x;
    }

    private Node rotateLeft(Node h) {
        Node x = h.right;
        h.right = x.left;
        x.left = h;
        return x;
    }

    public static void main(String[] args) {
        SplayBST<Integer, Integer> st1 = new SplayBST<Integer, Integer>();
        st1.put(5, 5);
        st1.put(9, 9);
        st1.put(13, 13);
        st1.put(11, 11);
        st1.put(1, 1);

        SplayBST<String, String> st = new SplayBST<String, String>();
        st.put("www.cs.princeton.edu", "128.112.136.11");
        st.put("www.cs.princeton.edu", "128.112.136.12");
        st.put("www.cs.princeton.edu", "128.112.136.13");
        st.put("www.princeton.edu", "128.112.128.15");
        st.put("www.yale.edu", "130.132.143.21");
        st.put("www.simpsons.com", "209.052.165.60");
        StdOut.println("The size 0 is: " + st.size());
        st.remove("www.yale.edu");
        StdOut.println("The size 1 is: " + st.size());
        st.remove("www.princeton.edu");
        StdOut.println("The size 2 is: " + st.size());
        st.remove("non-member");
        StdOut.println("The size 3 is: " + st.size());
        StdOut.println(st.get("www.cs.princeton.edu"));
        StdOut.println("The size 4 is: " + st.size());
        StdOut.println(st.get("www.yale.com"));
        StdOut.println("The size 5 is: " + st.size());
        StdOut.println(st.get("www.simpsons.com"));
        StdOut.println("The size 6 is: " + st.size());
        StdOut.println();
    }
}
