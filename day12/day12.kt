import java.nio.file.Files
import java.nio.file.Paths
import java.util.LinkedList

data class Pos(var x: Int, var y: Int, var distance: Int = 0) {
    fun value(grid: List<String>): Char = grid[y][x]
}

fun getAdjacent(pos: Pos) = listOf(
    Pos(pos.x + 1, pos.y, distance = pos.distance + 1),
    Pos(pos.x, pos.y + 1, distance = pos.distance + 1),
    Pos(pos.x - 1, pos.y, distance = pos.distance + 1),
    Pos(pos.x, pos.y - 1, distance = pos.distance + 1)
)

fun bfs(grid: List<String>, startPos: Pos): Int {
    val visited = hashSetOf(startPos)
    val queue = LinkedList<Pos>()
    queue.add(startPos)

    while (queue.isNotEmpty()) {
        val pos = queue.peek()
        if (grid[pos.y][pos.x] == 'E') return pos.distance

        queue.remove()

        for (dstPos in getAdjacent(pos)) {
            if (isValid(pos, dstPos, grid, visited)) {
                visited.add(dstPos)
                queue.add(dstPos)
            }
        }
    }
    return -1
}

private fun isValid(srcPos: Pos, dstPos: Pos, grid: List<String>, visited: HashSet<Pos>): Boolean =
    dstPos.y >= 0 && dstPos.y < grid.size && dstPos.x >= 0 && dstPos.x < grid[0].length && !visited.contains(dstPos) && (grid[dstPos.y][dstPos.x] - grid[srcPos.y][srcPos.x] <= 1 || (grid[srcPos.y][srcPos.x] == 'S' && grid[dstPos.y][dstPos.x] == 'a'))

fun part2(grid: List<String>) {
    println(List(grid.size) { y -> bfs(grid, startPos = Pos(0, y)) }.min())
}

fun part1(grid: List<String>) {
    val startPos = grid.mapIndexed { y, row -> Pos(x = row.indexOf('S'), y = y) }.first { it.x >= 0 }
    println(bfs(grid, startPos))
}

fun main() {
    val input = Files.lines(Paths.get("input.txt")).toList()

    part1(input)
    part2(grid = input)
}