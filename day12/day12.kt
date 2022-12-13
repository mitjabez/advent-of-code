import java.nio.file.Files
import java.nio.file.Paths

data class Pos(var x: Int, var y: Int) {
    fun value(grid: List<String>): Char = grid[y][x]
}

class MazeRunner(private val grid: List<String>, private val startPos: Pos) {

    private var moves = 0
    private var minMoves = Int.MAX_VALUE
    private val visited = mutableSetOf<Pos>()
    private var counter = 0L


    fun findShortest(): Int {
        move(startPos)
        return minMoves
    }

    private fun move(pos: Pos) {
        if (grid[pos.y][pos.x] == 'E') {
            minMoves = Math.min(minMoves, moves)
            // Try to find more
            return
        }

        if (counter++ % 100000L == 0L)
            draw(grid, visited)

        visited.add(pos)
        moves++
//        println("$pos=${grid[pos.y][pos.x]}")

        for (dstPos in listOf(
            Pos(pos.x + 1, pos.y),
            Pos(pos.x, pos.y + 1),
            Pos(pos.x - 1, pos.y),
            Pos(pos.x, pos.y - 1)
        )) {
            if (isValid(pos, dstPos))
                move(dstPos)
        }

        visited.remove(pos)
        moves--
    }

    private fun isValid(srcPos: Pos, dstPos: Pos): Boolean = dstPos.y >= 0 && dstPos.y < grid.size &&
            dstPos.x >= 0 && dstPos.x < grid[0].length &&
            !visited.contains(dstPos) &&
            (grid[dstPos.y][dstPos.x] - grid[srcPos.y][srcPos.x] in 0..1 ||
                    (grid[srcPos.y][srcPos.x] == 'z' && grid[dstPos.y][dstPos.x] == 'E') ||
                    (grid[srcPos.y][srcPos.x] == 'S' && grid[dstPos.y][dstPos.x] == 'a'))
}

fun draw(grid: List<String>, visited: Set<Pos>) {
    Thread.sleep(500)
    for (i in 0 .. 10)
        println()

    grid.forEachIndexed { y, row ->
        row.forEachIndexed { x, c ->
            print(if (Pos(x, y) in visited) "." else c)
        }
        println()
    }
}

fun part1(grid: List<String>) {
    val startPos = grid.mapIndexed { y, row -> Pos(x = row.indexOf('S'), y = y) }.first { it.x >= 0 }
    println(MazeRunner(grid, startPos).findShortest())
}

fun main() {
    val input = Files.lines(Paths.get("input.txt")).toList()

    part1(input)
}