package day24

import java.nio.file.Files
import java.nio.file.Paths

val BLIZZARD = mapOf('^' to Pos(0, -1), 'v' to Pos(0, 1), '<' to Pos(-1, 0), '>' to Pos(1, 0))

data class Pos(val x: Int, val y: Int) {
    fun add(other: Pos) = Pos(x + other.x, y + other.y)
}

fun moveBlizzard(blizzard: Blizzard, mapSize: Pos): Blizzard {
    val pos = blizzard.pos
    val direction = blizzard.direction
    val newPos = when {
        direction.x == 1 && pos.x == mapSize.x - 2 -> Pos(1, pos.y)
        direction.x == -1 && pos.x == 1 -> Pos(mapSize.x - 2, pos.y)
        direction.y == 1 && pos.y == mapSize.y - 2 -> Pos(pos.x, 1)
        direction.y == -1 && pos.y == 1 -> Pos(pos.x, mapSize.y - 2)
        else -> pos.add(direction)
    }
    return blizzard.copy(pos = newPos)
}

data class Blizzard(val pos: Pos, val direction: Pos)

data class Maze(var blizzards: List<Blizzard>, val mapSize: Pos, val start: Pos, val end: Pos)


fun readInput(file: String): Maze {
    val lines = Files.lines(Paths.get(file)).toList()
    val startPos = Pos(1, 0)
    val grid = lines.indices.flatMap { y -> lines[y].indices.map { x -> Pos(x, y) } }
    return Maze(
        start = startPos,
        blizzards = grid.filter { lines[it.y][it.x] in BLIZZARD }.map { Blizzard(it, BLIZZARD[lines[it.y][it.x]]!!) },
        mapSize = Pos(lines[0].length, lines.size),
        end = Pos(lines[0].length - 2, lines.size - 1)
    )

}

fun isInGrid(pos: Pos, mapSize: Pos) = pos.x >= 1 && pos.x <= mapSize.x - 2 && pos.y >= 1 && pos.y <= mapSize.y - 2

fun movePos(pos: Pos, end: Pos, blizzards: List<Blizzard>, mapSize: Pos): Set<Pos> {
    return listOf(pos.copy(x = pos.x - 1), pos.copy(x = pos.x + 1), pos.copy(y = pos.y - 1), pos.copy(y = pos.y + 1), pos)
        .filter { p -> (p == end || isInGrid(p, mapSize)) && blizzards.find { it.pos == p } == null }
        .toSet()
}

fun draw(currentPos: Pos, start: Pos, end: Pos, mapSize: Pos, blizzards: List<Blizzard>) {
    for (y in 0 until mapSize.y) {
        for (x in 0 until mapSize.x) {
            val pos = Pos(x, y)
            val c = if (pos == currentPos) "E"
            else if (blizzards.find { it.pos == pos } != null) {
                val blizaardsInPos = blizzards.filter { it.pos == pos }
                if (blizaardsInPos.size > 1) blizaardsInPos.size.toString()
                else BLIZZARD.entries.first { it.value == blizaardsInPos[0].direction }.key
            } else if (pos in listOf(start, end)) "."
            else if (!isInGrid(pos, mapSize)) "#"
            else "."
            print(c)
        }
        println()
    }
    println()
}

fun solve(sourceBlizzards: List<Blizzard>, mapSize: Pos, startStep: Int, start: Pos, end: Pos): Int {
    var blizzards = sourceBlizzards.toList()
    var step = startStep

    // Set initial state
    for (i in 0 until startStep)
        blizzards = blizzards.map { blizzard -> moveBlizzard(blizzard, mapSize) }
    var positions = movePos(start, end, blizzards, mapSize)

    while (true) {
        val nextPositions = mutableSetOf<Pos>()
        for (pos in positions) {
//            draw(pos, start, end, mapSize, blizzards)
//            readln()
            if (pos == end)
                return step - 1
            nextPositions.addAll(movePos(pos, end, blizzards, mapSize))
        }
        blizzards = blizzards.map { blizzard -> moveBlizzard(blizzard, mapSize) }
        positions = nextPositions
        if (positions.isEmpty())
            positions.add(start)
        step++
    }
}

fun main() {
    val maze = readInput("day24/input.txt")
    val part1 = solve(maze.blizzards, maze.mapSize, 0, maze.start, maze.end)
    val part2Part1 = solve(maze.blizzards, maze.mapSize, part1, maze.end, maze.start)
    val part2Part2 = solve(maze.blizzards, maze.mapSize, part2Part1, maze.start, maze.end)

    println(part1)
    println(part2Part2)
}
