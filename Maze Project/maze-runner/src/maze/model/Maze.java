package maze.model;

import maze.algo.generation.PassageTree;
import maze.algo.solving.Fugitive;

import java.util.function.Consumer;

import static java.lang.Integer.parseInt;
import static maze.model.Cell.Type.PASSAGE;
import static maze.model.Cell.Type.WALL;

/**
 * This class encapsulates the internal representation of the maze and provides
 * methods for creating, managing and extracting information about it.
 *
 * @author Philipp Malkovsky
 */
public class Maze {

    /**
     * The height of the maze in cells.
     */
    private final int height;

    /**
     * The width of the maze in cells.
     */
    private final int width;

    /**
     * Two-dimensional array of cells representing maze.
     */
    private final Cell[][] grid;

    /**
     * Indicates if a method for solving the maze has already
     * been called. It is used to prevent recalculation.
     */
    private boolean isSolved = false;

    /**
     * Generates a new maze of given height and width.
     *
     * @param height height of a maze
     * @param width  width of a maze
     */
    public Maze(int height, int width) {
        if (height < 3 || width < 3) {
            throw new IllegalArgumentException(
                "Both the height and the width " +
                    "of the maze must be at least 3");
        }
        this.height = height;
        this.width = width;
        grid = new Cell[height][width];
        fillGrid();
    }

    /**
     * Generates a new square maze of a given size.
     *
     * @param size size of a maze
     */
    public Maze(int size) {
        this(size, size);
    }

    /**
     * Fills the cells of the new maze such that the maze becomes
     * simply connected, i.e. containing no loops and no detached walls.
     */
    private void fillGrid() {
        fillAlternately();
        fillGaps();
        makeEntranceAndExit();
        generatePassages();
    }

    /**
     * Creates a new cell with given coordinates and a type in the grid.
     *
     * @param row    a row in the grid
     * @param column a column in the grid
     * @param type   the new cell type
     */
    private void putCell(int row, int column, Cell.Type type) {
        grid[row][column] = new Cell(row, column, type);
    }

    /**
     * Fills every second cell with a passage and the others with a wall.
     * After this method, a maze looks like this:
     * <pre>
     * ██████████
     * ██  ██  ██
     * ██████████
     * ██  ██  ██
     * ██████████
     * </pre>
     */
    private void fillAlternately() {
        for (int i = 0; i < height; i++) {
            for (int j = 0; j < width; j++) {
                if ((i & 1) == 0 || (j & 1) == 0) {
                    putCell(i, j, WALL);
                } else {
                    putCell(i, j, PASSAGE);
                }
            }
        }
    }

    /**
     * If the maze has an even height or width it is needed to fill the
     * last row or column of the grid with a wall (or, otherwise, it will
     * contain passages at the outer border).
     * <pre>
     * ████████████
     * ██  ██  ██
     * ████████████
     * ██  ██  ██
     * ████████████
     * ██  ██  ██
     * </pre>
     * becomes
     * <pre>
     * ████████████
     * ██  ██  ████
     * ████████████
     * ██  ██  ████
     * ████████████
     * ████████████
     * </pre>
     */
    private void fillGaps() {
        if (height % 2 == 0) wallLastRow();
        if (width % 2 == 0) wallLastColumn();
    }

    /**
     * Fills the last column in the grid with a wall.
     */
    private void wallLastColumn() {
        for (int i = 0; i < height; i++)
            putCell(i, width - 1, WALL);
    }

    /**
     * Fills the last row in the grid with a wall.
     */
    private void wallLastRow() {
        for (int i = 0; i < width; i++)
            putCell(height - 1, i, WALL);
    }

    /**
     * Calculates the index of the passage in the last row. For a maze
     * with an odd (1) and even (2) width its indices differ:
     * <pre>
     * (1) ██████  ██
     * (2) ██████  ████
     * </pre>
     *
     * @return the index of the passage in the last row
     */
    private int getExitColumn() {
        return width - 3 + width % 2;
    }

    /**
     * Puts entrance and exit passages to upper left and lower right
     * corners. For example:
     * <pre>
     * ████████████
     * ██  ██  ████
     * ████████████
     * ██  ██  ████
     * ████████████
     * ████████████
     * </pre>
     * becomes
     * <pre>
     * ██  ████████
     * ██  ██  ████
     * ████████████
     * ██  ██  ████
     * ██████  ████
     * ██████  ████
     * </pre>
     */
    private void makeEntranceAndExit() {
        putCell(0, 1, PASSAGE);
        putCell(height - 1, getExitColumn(), PASSAGE);
        if (height % 2 == 0)
            putCell(height - 2, getExitColumn(), PASSAGE);
    }

    /**
     * Creates random passages between isolated passage cells such
     * that every cell is connected to the other in one way and
     * has no cycles. For example:
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
     *
     * @see PassageTree
     */
    private void generatePassages() {
        new PassageTree(height, width)
            .generate()
            .forEach(putCell());
    }

    /**
     * Puts a cell in the corresponding place in grid.
     *
     * @return lambda to put a cell
     */
    private Consumer<Cell> putCell() {
        return cell -> grid[cell.getRow()][cell.getColumn()] = cell;
    }

    /**
     * Finds a path in the maze from its entrance to its exit.
     * For example:
     * <pre>
     * ██░░██████████
     * ██░░░░░░██  ██
     * ██████░░██  ██
     * ██    ░░    ██
     * ██████░░██████
     * ██    ░░░░░░██
     * ██████████░░██
     * </pre>
     *
     * <p>If this method is called several times, the path is not
     * recalculated. It is stored in the grid so it is returned
     * immediately.</p>
     *
     * @return string representation of the maze containing a path
     * @see Fugitive
     */
    public String findEscape() {
        if (!isSolved) {
            new Fugitive(grid, getEntrance(), getExit())
                .findEscape()
                .forEach(putCell());
            isSolved = true;
        }
        return toString(true);
    }

    /**
     * Return the entrance cell.
     *
     * @return the entrance cell
     */
    private Cell getEntrance() {
        return grid[0][1];
    }

    /**
     * Return the exit cell.
     *
     * @return the exit cell
     */
    private Cell getExit() {
        return grid[height - 1][getExitColumn()];
    }

    /**
     * Return the string representation of the grid. The path
     * from the entrance to the exit can be displayed if it
     * is already found and {@code showEscape} is {@code true}.
     * The path is found on demand.
     *
     * <p>
     * For example:<br>
     * if path is already found and {@code showEscape} is
     * {@code true}
     * <pre>
     * ██░░██████████
     * ██░░░░░░██  ██
     * ██████░░██  ██
     * ██    ░░    ██
     * ██████░░██████
     * ██    ░░░░░░██
     * ██████████░░██
     * </pre>
     * if {@code showEscape} is {@code false}
     * <pre>
     * ██  ██████████
     * ██      ██  ██
     * ██████  ██  ██
     * ██          ██
     * ██████  ██████
     * ██          ██
     * ██████████  ██
     * </pre>
     *
     * @param showEscape show path or not
     * @return string representation of the maze
     */
    private String toString(boolean showEscape) {
        var sb = new StringBuilder();
        for (var row : grid) {
            for (var cell : row) {
                if (cell.isWall()) {
                    sb.append("██");
                } else if (showEscape && cell.isEscape()) {
                    sb.append("▓▓");
                } else {
                    sb.append("  ");
                }
            }
            sb.append('\n');
        }
        return sb.toString();
    }

    /**
     * Return the string representation of the grid.
     * The path is never displayed. For example:
     * <pre>
     * ██  ██████████
     * ██      ██  ██
     * ██████  ██  ██
     * ██          ██
     * ██████  ██████
     * ██          ██
     * ██████████  ██
     * </pre>
     *
     * @return string representation of the maze
     */
    @Override
    public String toString() {
        return toString(false);
    }

    /**
     * Parses a serialized maze representation and
     * constructs a new maze from it.
     * <p>
     * Maze is serialized in the following form:
     * <pre>
     * height width
     * cell[0][0] cell[0][1] ... cell[0][width - 1]
     * cell[1][0] cell[1][1] ... cell[1][width - 1]
     * ...
     * cell[height - 1][0] cell[height - 1][1] ... cell[height - 1][width - 1],
     * </pre>
     * where cell[i][j] is 1 if the cell type is a wall and 0 if
     * the cell type is a passage. The escape path is not serialized.
     * <br>
     * For example:
     * <pre>
     * 5 7
     * 1 0 1 1 1 1 1
     * 1 0 0 0 1 0 1
     * 1 1 1 0 1 0 1
     * 1 0 0 0 0 0 1
     * 1 1 1 1 1 0 1
     * </pre>
     * (a serialized form)<br>
     * corresponds to
     * <pre>
     * ██  ██████████
     * ██      ██  ██
     * ██████  ██  ██
     * ██          ██
     * ██████████  ██
     * </pre>
     * (a regular form)<br>
     *
     * @param str a serialized maze representation
     * @return a new maze from a given string
     */
    public static Maze load(String str) {
        try {
            var whole = str.split("\n");
            var size = whole[0].split(" ");
            var height = parseInt(size[0]);
            var width = parseInt(size[1]);
            var grid = new Cell[height][width];
            for (int i = 0; i < height; i++) {
                var row = whole[i + 1].split(" ");
                for (int j = 0; j < width; j++)
                    grid[i][j] = new Cell(
                        i, j, intToType(parseInt(row[j]))
                    );
            }
            return new Maze(height, width, grid);
        } catch (Exception e) {
            throw new IllegalArgumentException(
                "Cannot load the maze. " +
                    "It has an invalid format"
            );
        }
    }

    /**
     * Creates a maze instance with given height, width and grid.
     *
     * @param height height of a maze
     * @param width  width of a maze
     * @param grid   grid of cells of a maze
     */
    private Maze(int height, int width, Cell[][] grid) {
        this.height = height;
        this.width = width;
        this.grid = grid;
    }

    /**
     * Converts 1 to the WALL and 0 to the PASSAGE.
     * The path is not serialized so there are only 2 choices.
     *
     * @param val value to convert
     * @return converted WALL or PASSAGE
     */
    private static Cell.Type intToType(int val) {
        return val == 1 ? WALL : PASSAGE;
    }

    /**
     * Converts the maze to the serialized form.
     * <p>
     * Maze is serialized in the following form:
     * <pre>
     * height width
     * cell[0][0] cell[0][1] ... cell[0][width - 1]
     * cell[1][0] cell[1][1] ... cell[1][width - 1]
     * ...
     * cell[height - 1][0] cell[height - 1][1] ... cell[height - 1][width - 1],
     * </pre>
     * where cell[i][j] is 1 if the cell type is a wall and 0 if
     * the cell type is a passage. The escape path is not serialized.
     * <br>
     * For example:
     * <pre>
     * 5 7
     * 1 0 1 1 1 1 1
     * 1 0 0 0 1 0 1
     * 1 1 1 0 1 0 1
     * 1 0 0 0 0 0 1
     * 1 1 1 1 1 0 1
     * </pre>
     * (a serialized form)<br>
     * corresponds to
     * <pre>
     * ██  ██████████
     * ██      ██  ██
     * ██████  ██  ██
     * ██          ██
     * ██████████  ██
     * </pre>
     * (a regular form)<br>
     *
     * @return string in a serialized form
     */
    public String export() {
        var sb = new StringBuilder();
        sb.append(height).append(' ')
          .append(width).append('\n');
        for (var row : grid) {
            for (var cell : row)
                sb.append(typeToInt(cell))
                  .append(' ');
            sb.append('\n');
        }
        return sb.toString();
    }

    /**
     * Converts WALL to the 1 and PASSAGE to the 0.
     * The path is not serialized so there are only 2 choices.
     *
     * @param cell value to convert
     * @return converted 1 or 0
     */
    private int typeToInt(Cell cell) {
        return cell.isWall() ? 1 : 0;
    }
}
