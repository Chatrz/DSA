package maze.algo.generation;

/**
 * This class stores an information about the particular edge in a passage tree.
 * It stores two coordinates corresponding to the cells locations in a grid.
 * Cells locations are consider as they would be in 1-dimensional array and
 * are calculated using the {@code row * width + column} formula.
 *
 * @author Philipp Malkovsky
 * @see PassageTree
 */
class Edge {

    /**
     * The coordinate of the first cell.
     */
    private final int firstCell;
    /**
     * The coordinate of the second cell.
     */
    private final int secondCell;

    /**
     * Creates a new edge with given cells coordinates.
     *
     * @param firstCell  the coordinate of the first cell
     * @param secondCell the coordinate of the second cell
     */
    Edge(int firstCell, int secondCell) {
        this.firstCell = firstCell;
        this.secondCell = secondCell;
    }

    /**
     * @return the first cell coordinate
     */
    int getFirstCell() {
        return firstCell;
    }

    /**
     * @return the second cell coordinate
     */
    int getSecondCell() {
        return secondCell;
    }
}
