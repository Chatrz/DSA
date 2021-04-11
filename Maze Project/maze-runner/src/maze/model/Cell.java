package maze.model;

import java.util.Objects;

/**
 * This class stores an information about the particular cell in a grid.
 *
 * @author Philipp Malkovsky
 */
public class Cell {

    /**
     * Cell can be either a passage, a wall
     * or be a part of an escape path.
     */
    public enum Type {
        PASSAGE,
        WALL,
        ESCAPE;
    }

    /**
     * The vertical coordinate of this cell in a grid.
     */
    private final int row;

    /**
     * The horizontal coordinate of this cell in a grid.
     */
    private final int column;

    /**
     * The type of this cell: a passage, a wall or an escape.
     */
    private final Type type;

    public Cell(int row, int column, Type type) {
        this.row = row;
        this.column = column;
        this.type = type;
    }

    public int getRow() {
        return row;
    }

    public int getColumn() {
        return column;
    }

    public boolean isPassage() {
        return type == Type.PASSAGE;
    }

    public boolean isWall() {
        return type == Type.WALL;
    }

    public boolean isEscape() {
        return type == Type.ESCAPE;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        var cell = (Cell) o;
        return row == cell.row &&
            column == cell.column &&
            type == cell.type;
    }

    @Override
    public int hashCode() {
        return Objects.hash(row, column, type);
    }

    @Override
    public String toString() {
        return "Cell{" +
            "row=" + row +
            ", column=" + column +
            ", type=" + type +
            '}';
    }
}
