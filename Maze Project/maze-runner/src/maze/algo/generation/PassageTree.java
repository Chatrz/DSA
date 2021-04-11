package maze.algo.generation;

import maze.model.Cell;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

import static java.util.stream.Collectors.toList;
import static maze.model.Cell.Type.PASSAGE;

/**
 * This class is used for creating random passages between
 * isolated passage cells such that every cell is connected
 * to the other in one way and the maze has no cycles.
 * Actually, it is a tree.
 * <p>
 * For example:
 * <pre>
 * ██  ██████
 * ██  ██  ██
 * ██████████
 * ██  ██  ██
 * ██████  ██
 * </pre>
 * can become
 * <pre>
 * ██  ██████
 * ██      ██
 * ██████  ██
 * ██      ██
 * ██████  ██
 * </pre>
 * <p>
 * Internally, it it creates a tree of edges between cells in a square
 * of doubly decreased size such that all cells are in the one
 * <a href="https://en.wikipedia.org/wiki/Component_(graph_theory)">
 * connected component</a>.
 * <p>
 * <a href="https://en.wikipedia.org/wiki/Maze_generation_algorithm#Randomized_Kruskal's_algorithm">
 * Randomized Kruskal's algorithm</a> is used.
 * <p>
 * Initially, there is a set of all possible edges in the small grid.
 * On each step of the algorithm it removes one of the edges such
 * that an edge connects two distinct connected components.
 * At the end, each cell is connected to the others in a one way.
 * <p>
 * For example, ({@code width = 9} and {@code height = 9}) were given.
 * It corresponds with the following maze where the entrance, the exit
 * and an every second cell are passages while others are walls:
 * <pre>
 * ██  ██████████████
 * ██  ██  ██  ██  ██
 * ██████████████████
 * ██  ██  ██  ██  ██
 * ██████████████████
 * ██  ██  ██  ██  ██
 * ██████████████████
 * ██  ██  ██  ██  ██
 * ██████████████  ██
 * </pre>
 * It is needed to connect these passages to make a simply connected
 * maze. After the size is decreased grid turns into the following
 * imaginary form in terms of edges:
 * <pre>
 *    _____
 * |_|_|_|_|
 * |_|_|_|_|
 * |_|_|_|_|
 * |_|_|_| |
 * </pre>
 * And, after "deleting" some random edges, a maze becomes:
 * <pre>
 *    _____
 * | |___  |
 * | |  _  |
 * |__ |   |
 * |___|_| |
 * </pre>
 * <p>
 * It then is scaled back to:
 * <pre>
 * ██  ██████████████
 * ██  ██          ██
 * ██  ██████████  ██
 * ██  ██          ██
 * ██  ██  ██████  ██
 * ██      ██      ██
 * ██████  ██  ██  ██
 * ██      ██  ██  ██
 * ██████████████  ██
 * </pre>
 * <p>
 * The algorithm output is a list of cells that connect passages
 * in an original form.
 *
 * @author Philipp Malkovsky
 */
public class PassageTree {

    /**
     * The height of the maze in an imaginary edge form.
     */
    private int height;

    /**
     * The width of the maze in an imaginary edge form.
     */
    private int width;

    /**
     * Creates a new imaginary edge form
     *
     * @param height an original height
     * @param width  an original width
     */
    public PassageTree(int height, int width) {
        this.height = (height - 1) / 2;
        this.width = (width - 1) / 2;
    }

    /**
     * Generates a random list of cells that connect passages in
     * an original form such that a maze is simply connected.
     *
     * @return a random list of cells that connect passages
     */
    public List<Cell> generate() {
        var edges = createEdges();
        Collections.shuffle(edges);
        var tree = buildRandomSpanningTree(edges);
        return createPassages(tree);
    }

    /**
     * Creates a list of all possible edges in an imaginary edge form.
     *
     * @return a list of all possible edges
     * @see Edge
     */
    private List<Edge> createEdges() {
        var edges = new ArrayList<Edge>();
        for (int column = 1; column < width; column++) {
            edges.add(new Edge(toIndex(0, column),
                               toIndex(0, column - 1)));
        }
        for (int row = 1; row < height; row++) {
            edges.add(new Edge(toIndex(row, 0),
                               toIndex(row - 1, 0)));
        }
        for (int row = 1; row < height; row++) {
            for (int column = 1; column < width; column++) {
                edges.add(new Edge(toIndex(row, column),
                                   toIndex(row, column - 1)));
                edges.add(new Edge(toIndex(row, column),
                                   toIndex(row - 1, column)));
            }
        }
        return edges;
    }

    /**
     * Transforms the coordinates in a 2-dimensional array
     * to the coordinate in a 1-dimensional array using the
     * {@code row * width + column} formula.
     *
     * @param row    the row coordinate in a 2-dimensional array
     * @param column the column coordinate in a 2-dimensional array
     * @return the coordinate in a 1-dimensional array
     */
    private int toIndex(int row, int column) {
        return row * width + column;
    }

    /**
     * Generates a list of edges that connect passages. It is a
     * <a href="https://en.wikipedia.org/wiki/Maze_generation_algorithm#Randomized_Kruskal's_algorithm">
     * Randomized Kruskal's algorithm</a> implementation.
     * <p>
     * On each step of the algorithm an edge is added to the list
     * only if it connects two disjoint subsets.
     *
     * @param edges a randomized list of all edges
     * @return a random list of edges that connect passages
     * @see DisjointSet
     */
    private List<Edge> buildRandomSpanningTree(List<Edge> edges) {
        var disjointSets = new DisjointSet(width * height);
        return edges
            .stream()
            .filter(edge -> connects(edge, disjointSets))
            .collect(toList());
    }

    /**
     * Checks if an {@code edge} connects two disjoint subsets.
     *
     * @param edge        a given edge
     * @param disjointSet a disjoint-set data structure that keeps track of subsets
     * @return true if an edge connects two disjoint subsets,
     * false otherwise
     */
    private boolean connects(Edge edge, DisjointSet disjointSet) {
        return disjointSet.union(edge.getFirstCell(), edge.getSecondCell());
    }

    /**
     * Scales and converts edges in an imaginary edge form to the cells
     * which connect passages in a original form.
     *
     * @param spanningTree a random list of edges that connect passages
     * @return a list of cells that connect passages
     */
    private List<Cell> createPassages(List<Edge> spanningTree) {
        return spanningTree
            .stream()
            .map(edge -> {
                var first = fromIndex(edge.getFirstCell());
                var second = fromIndex(edge.getSecondCell());
                return getPassage(first, second);
            }).collect(toList());
    }

    /**
     * Transforms the coordinate in a 1-dimensional array
     * back to the coordinates in a 2-dimensional array using the
     * {@code row = index / width} and {@code column = index % width}
     * formulas.
     *
     * @param index the coordinate in a 1-dimensional array
     * @return a cell with coordinates in a 2-dimensional array
     */
    private Cell fromIndex(int index) {
        var row = index / width;
        var column = index % width;
        return new Cell(row, column, PASSAGE);
    }

    /**
     * Given the coordinates of two cells that compose an edge in
     * an imaginary edge form, it scales and transforms them to
     * the coordinates of the cell that connect passages in an
     * original form. Returns a passage cell with this coordinates.
     *
     * @param first  one edge ending
     * @param second another edge ending
     * @return a passage cell with the transformed coordinates
     */
    private Cell getPassage(Cell first, Cell second) {
        var row = first.getRow() + second.getRow() + 1;
        var column = first.getColumn() + second.getColumn() + 1;
        return new Cell(row, column, PASSAGE);
    }
}
