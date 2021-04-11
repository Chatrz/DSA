package maze.algo.solving;

import maze.model.Cell;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.PriorityQueue;
import java.util.Set;

import static java.util.Comparator.comparingInt;
import static maze.model.Cell.Type.ESCAPE;

/**
 * This class is used for finding an escape path from the maze
 * entrance to the maze exit. It is the implementation of the
 * <a href="https://en.wikipedia.org/wiki/A*_search_algorithm">
 * A* search algorithm</a>.
 * <p>
 * The implementation is pretty close to the Wikipedia article.
 *
 * @author Philipp Malkovsky
 */
public class Fugitive {

    /**
     * Moves to up, left, right and down from the current cell.
     */
    private static final int[][] DELTAS = {{-1, 0}, {0, -1}, {0, 1}, {1, 0}};

    /**
     * The height of the maze in nodes.
     */
    private int height;

    /**
     * The width of the maze in node.
     */
    private int width;

    /**
     * Two-dimensional array of nodes representing maze.
     */
    private Node[][] grid;

    /**
     * The start point to find a path from.
     */
    private Node start;

    /**
     * The end point to find a path to.
     */
    private Node end;

    /**
     * A priority queue to perform the selection of minimum
     * estimated cost node on every step of the algorithm.
     */
    private PriorityQueue<Node> open = new PriorityQueue<>(comparingInt(Node::getFinalCost));

    /**
     * Already processed nodes.
     */
    private Set<Node> closed = new HashSet<>();

    /**
     * Constructs a new object with given grid of cells
     * and start and end cells. Creates a grid of nodes
     * based on that.
     *
     * @param grid  a grid of cells of a maze
     * @param start a start point to find a path from
     * @param end   an end point to find a path to
     */
    public Fugitive(Cell[][] grid, Cell start, Cell end) {
        this.height = grid.length;
        this.width = grid[0].length;
        this.grid = new Node[height][width];
        this.start = new Node(start.getRow(), start.getColumn(), false);
        this.end = new Node(end.getRow(), end.getColumn(), false);
        createNodes(grid);
    }

    /**
     * For each cell in a given grid it creates the corresponding
     * node in a grid of nodes. Calculates an estimated cost
     * to the end for each node.
     *
     * @param grid a given grid of cells
     */
    private void createNodes(Cell[][] grid) {
        for (int i = 0; i < height; i++) {
            for (int j = 0; j < width; j++) {
                var node = new Node(i, j, grid[i][j].isWall());
                node.calcHeuristicTo(end);
                this.grid[i][j] = node;
            }
        }
    }

    /**
     * Find a path from the start to the end using
     * <a href="https://en.wikipedia.org/wiki/A*_search_algorithm">
     * * A* search algorithm</a>.
     *
     * @return a cell list leading from the start to the end
     */
    public List<Cell> findEscape() {
        open.add(start);
        while (!open.isEmpty()) {
            var cur = open.poll();
            if (isEnd(cur))
                return reconstructPath(cur);
            closed.add(cur);
            updateNeighbors(cur);
        }
        return new ArrayList<>();
    }

    /**
     * Check if a node is the end point to find a path to.
     *
     * @param currentNode a given node
     * @return true if node is the end point, false otherwise
     */
    private boolean isEnd(Node currentNode) {
        return currentNode.equals(end);
    }

    /**
     * Reconstructs the path from the given node to the
     * start node, i.e. node having no parent. Returns a
     * list of cells in the format: start -> ... -> cur.
     *
     * @param cur a node to reconstruct path to
     * @return a list of cells containing a path from
     * the start node to the given cell
     */
    private List<Cell> reconstructPath(Node cur) {
        var path = new LinkedList<Cell>();
        path.add(toCell(cur));
        while (cur.getParent() != cur) {
            var parent = cur.getParent();
            path.addFirst(toCell(parent));
            cur = parent;
        }
        return path;
    }

    /**
     * Converts a node back to the cell format.
     * Cell type is escape path.
     *
     * @param node a given node
     * @return a converted cell
     */
    private Cell toCell(Node node) {
        return new Cell(node.getRow(), node.getColumn(), ESCAPE);
    }

    /**
     * Updates an estimated and a final costs of neighboring
     * cells according to the
     * <a href="https://en.wikipedia.org/wiki/A*_search_algorithm">
     * A* search algorithm</a>.
     *
     * @param cur a node which neighbors are updated
     */
    private void updateNeighbors(Node cur) {
        for (var delta : DELTAS) {
            var row = cur.getRow() + delta[0];
            var column = cur.getColumn() + delta[1];
            if (inBounds(row, column)) {
                var node = grid[row][column];
                if (!node.isWall() && !closed.contains(node)) {
                    if (open.contains(node)) {
                        if (node.hasBetterPath(cur)) {
                            open.remove(node);
                        } else {
                            continue;
                        }
                    }
                    node.updatePath(cur);
                    open.add(node);
                }
            }
        }
    }

    /**
     * Checks if given cell indices are in bounds
     * of the 2-dimensional array.
     *
     * @param row    a row index
     * @param column a column index
     * @return true if cell indices are in bounds, false otherwise
     */
    private boolean inBounds(int row, int column) {
        return row >= 0 && row < height
            && column >= 0 && column < width;
    }
}

