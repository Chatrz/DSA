package maze.algo.generation;

import static java.util.stream.IntStream.range;

/**
 * @author Philipp Malkovsky
 * @see <a href="https://en.wikipedia.org/wiki/Disjoint-set_data_structure">
 * Disjoint-set data structure</a> implementation.
 */
public class DisjointSet {

    /**
     * Representatives for disjoint subsets. If the set consists
     * only of the one element its parent equals to its id.
     * Otherwise, its parent is the next element up the tree.
     *
     * @see <a href="https://en.wikipedia.org/wiki/Disjoint-set_data_structure#Representation">
     * Disjoint-set representation</a>
     */
    private int[] parent;

    /**
     * Heights of the trees corresponding to the subsets.
     * A subset with one element has a rank of zero.
     *
     * @see <a href="https://en.wikipedia.org/wiki/Disjoint-set_data_structure#by_rank">
     * Union by rank</a>
     */
    private int[] rank;

    /**
     * The number of disjoint subsets.
     */
    private int size;

    /**
     * Constructs a disjoint set of {@code size} disjoint subsets.
     *
     * @param size a number of disjoint subsets
     */
    public DisjointSet(int size) {
        this.size = size;
        parent = new int[size];
        rank = new int[size];
        range(0, size).forEach(this::makeSet);
    }

    /**
     * Initializes a particular set.
     *
     * @param i an id of the set to initialize
     */
    private void makeSet(int i) {
        parent[i] = i;
        rank[i] = 0;
    }

    /**
     * Returns the number of disjoint subsets.
     *
     * @return the number of disjoint subsets
     */
    public int getSize() {
        return size;
    }

    /**
     * Finds a representative for the set. If the set consists
     * only of the one element its parent equals to its id.
     *
     * @param i id of the set
     * @return a representative parent pointer for this set
     * @see <a href="https://en.wikipedia.org/wiki/Disjoint-set_data_structure#Path_compression">
     * Path compression</a> is used.
     */
    public int find(int i) {
        if (i != parent[i])
            parent[i] = find(parent[i]);
        return parent[i];
    }

    /**
     * Merges two disjoint sets into one by the ids
     * if their ids are not in the same set already.
     *
     * @param i an id of the first set
     * @param j an id of the second set
     * @return true if sets are merge, false if ids are in the same set
     * @see <a href="https://en.wikipedia.org/wiki/Disjoint-set_data_structure#by_rank">
     * Union by rank</a> is used.
     */
    public boolean union(int i, int j) {
        var iRoot = find(i);
        var jRoot = find(j);
        if (iRoot == jRoot)
            return false;
        if (rank[iRoot] < rank[jRoot]) {
            parent[iRoot] = jRoot;
        } else {
            parent[jRoot] = iRoot;
            if (rank[iRoot] == rank[jRoot])
                rank[iRoot]++;
        }
        size--;
        return true;
    }
}
