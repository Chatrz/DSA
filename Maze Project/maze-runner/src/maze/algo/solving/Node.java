package maze.algo.solving;

import java.util.Objects;

/**
 * This class stores an information about the particular
 * node in a node grid. It is used in the
 * <a href="https://en.wikipedia.org/wiki/A*_search_algorithm">
 * A* search algorithm</a> implementation.
 * <p>
 * A node stores an information about the cost of the path from
 * the start node to it and the estimated cost of the path from
 * it to the end node and calculates the final cost through it.
 *
 * @author Philipp Malkovsky
 * @see Fugitive
 */
class Node {

    /**
     * The cost of moving to neighboring nodes.
     */
    private static final int EDGE_COST = 1;

    /**
     * The vertical coordinate of this node in a grid.
     */
    private final int row;

    /**
     * The horizontal coordinate of this node in a grid.
     */
    private final int column;

    /**
     * Indicates if this node is a wall.
     */
    private final boolean isWall;

    /**
     * A parent node is saved to reconstruct a path if
     * it goes through this node. If node has no parent
     * its parent is equal to this node, i.e.
     * {@code node == node.getParent()}.
     */
    private Node parent;

    /**
     * The cost of the path from the start node to this node.
     */
    private int g;

    /**
     * The estimated cost of the path from this node to the end node.
     */
    private int h;

    /**
     * The final cost of the path from the start node to the end node
     * through this node.
     */
    private int f;

    /**
     * Creates a new node with given row and column and sets its parent
     * to itself.
     *
     * @param row    a row of a new node
     * @param column a column of a new node
     * @param isWall true if a node is a wall, false otherwise
     */
    Node(int row, int column, boolean isWall) {
        this.row = row;
        this.column = column;
        this.isWall = isWall;
        parent = this;
    }

    int getRow() {
        return row;
    }

    int getColumn() {
        return column;
    }

    boolean isWall() {
        return isWall;
    }

    Node getParent() {
        return parent;
    }

    int getFinalCost() {
        return f;
    }

    /**
     * Calculates the estimated cost of the path from this node to the
     * end node. This implementation uses a
     * <a href="https://en.wikipedia.org/wiki/Taxicab_geometry">
     * Manhattan distance</a> to calculate the heuristic.
     *
     * @param node a node to estimate a distance to
     */
    void calcHeuristicTo(Node node) {
        this.h = Math.abs(node.row - this.row)
            + Math.abs(node.column - this.column);
    }

    /**
     * Checks if the path through the given node is better
     * (i.e. cheaper) than the current path.
     *
     * @param node the node to compare
     * @return true if better, false otherwise
     */
    boolean hasBetterPath(Node node) {
        return node.g + EDGE_COST < this.g;
    }

    /**
     * Updates the path such that the given node becomes the
     * new parent and recalculates the final cost through it.
     *
     * @param node the new parent
     */
    void updatePath(Node node) {
        this.parent = node;
        this.g = node.g + EDGE_COST;
        f = g + h;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        var node = (Node) o;
        return row == node.row &&
            column == node.column &&
            isWall == node.isWall;
    }

    @Override
    public int hashCode() {
        return Objects.hash(row, column, isWall);
    }
}