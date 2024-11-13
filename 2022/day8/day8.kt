import java.nio.file.Files
import java.nio.file.Paths

fun getColumn(columnNo: Int, grid: List<String>): String = List(grid.size) { grid[it][columnNo] }.joinToString("")

fun firstNonZero(value: Int, defaultValue: Int): Int = if (value > 0) value else defaultValue

private fun isOnEdge(x: Int, y: Int, gridByRow: List<String>) =
    x == 0 || y == 0 || x == gridByRow.size - 1 || y == gridByRow.size - 1

fun isVisible(tree: Char, x: Int, y: Int, gridByRow: List<String>, gridByColumn: List<String>): Boolean {
    if (isOnEdge(x, y, gridByRow))
        return true

    return (gridByRow[y].substring(0, x).max() < tree || gridByRow[y].substring(x + 1, gridByRow[y].length).max() < tree) ||
            (gridByColumn[x].substring(0, y).max() < tree || gridByColumn[x].substring(y + 1, gridByColumn[x].length).max() < tree
                    )
}

fun scenicScore(tree: Char, x: Int, y: Int, gridByRow: List<String>, gridByColumn: List<String>): Int {
    if (isOnEdge(x, y, gridByRow))
        return 0

    val row = gridByRow[y]
    val column = gridByColumn[x]
    val left = firstNonZero(row.substring(0, x).reversed().indexOfFirst { it >= tree } + 1, x)
    val right = firstNonZero(row.substring(x + 1, row.length).indexOfFirst { it >= tree } + 1, row.length - x - 1)
    val top = firstNonZero(column.substring(0, y).reversed().indexOfFirst { it >= tree } + 1, y)
    val bottom = firstNonZero(column.substring(y + 1, column.length).indexOfFirst { it >= tree } + 1, column.length - y - 1)
    return left * right * top * bottom
}

fun part2(gridByRow: List<String>, gridByColumn: List<String>) {
    val score = gridByRow.flatMapIndexed { y, row ->
        row.mapIndexed { x, tree ->
            scenicScore(tree, x, y, gridByRow, gridByColumn)
        }
    }.max()
    println(score)
}

fun part1(gridByRow: List<String>, gridByColumn: List<String>) {
    val visible = gridByRow.flatMapIndexed { y, row ->
        row.mapIndexed { x, tree ->
            isVisible(tree, x, y, gridByRow, gridByColumn)
        }
    }.count { it }
    println(visible)
}

fun main() {
    val gridByRow = Files.lines(Paths.get("input.txt")).toList()
    val gridByColumn = gridByRow[0].indices.map { getColumn(it, gridByRow) }

    part1(gridByRow, gridByColumn)
    part2(gridByRow, gridByColumn)
}