package day14

import java.nio.file.Files
import java.nio.file.Paths

data class Pos(val x: Int, val y: Int)
data class Line(val start: Pos, val end: Pos) {
    fun contains(pos: Pos): Boolean =
        if (pos.y == start.y) pos.x in (start.x..end.x) || pos.x in (end.x..start.x)
        else if (pos.x == start.x) pos.y in (start.y..end.y) || pos.y in (end.y..start.y)
        else false
}

fun readInput(file: String): List<Line> {
    return Files.lines(Paths.get(file))
        .map { it.split(" -> ") }
        .toList()
        .flatMap {
            it
                .map { num -> Pair(num.split(",")[0].toInt(), num.split(",")[1].toInt()) }
                .map { num -> Pos(num.first, num.second) }
                .windowed(2, 1)
                .map { (start, end) -> Line(start, end) }
        }
}

fun tryFall(dust: Pos, stopped: Set<Pos>, lines: List<Line>): Pos {
    return listOf(0, -1, 1).map { Pos(dust.x + it, dust.y + 1) }.firstOrNull { newPos ->
        lines.none { it.contains(newPos) } && newPos !in stopped
    } ?: dust
}

fun solve(lines: List<Line>) {
    val maxY = lines.maxOf { it.end.y }
    val stopped = mutableSetOf<Pos>()
    val start = Pos(500, 0)
    var dust = start
    var newPos = Pos(0, 0)

    while (newPos.y < maxY + 1) {
        newPos = tryFall(dust, stopped, lines)
        dust = if (newPos == dust) {
            stopped.add(dust)
            start
        } else {
            newPos
        }
        // Part 2
        if (newPos == start) {
            println(stopped.size)
            return
        }
    }
    // Part 1
    println(stopped.size)
}

fun part2(lines: List<Line>) {
    val linesWithFloor = lines.toMutableList()
    val maxY = lines.maxOf { it.end.y }
    linesWithFloor.add(Line(Pos(Int.MIN_VALUE, maxY + 2), Pos(Int.MAX_VALUE, maxY + 2)))
    solve(linesWithFloor)
}

fun main() {
    val lines = readInput("day14/input.txt")
    solve(lines)
    part2(lines)
}
