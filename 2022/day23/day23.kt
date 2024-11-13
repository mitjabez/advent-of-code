package day23

import java.nio.file.Files
import java.nio.file.Paths

data class Pos(val x: Int, val y: Int) {
    fun plus(pos: Pos) = Pos(x + pos.x, y + pos.y)
}

enum class Direction(val moveVector: Pos) {
    NORTH(Pos(0, -1)), SOUTH(Pos(0, 1)), WEST(Pos(-1, 0)), EAST(Pos(1, 0));

    val allPoleVectors: List<Pos>

    init {
        val altDir = listOf(0, -1, 1)
        allPoleVectors = if (moveVector.x == 0) altDir.map { Pos(it, moveVector.y) } else altDir.map { Pos(moveVector.x, it) }
    }
}

fun readInput(file: String): Set<Pos> {
    val input = Files.lines(Paths.get(file)).toList()
    return input.indices
        .flatMap { y -> input[0].indices.mapIndexed { x, _ -> Pos(x, y) } }
        .filter { input[it.y][it.x] == '#' }
        .toSet()
}

fun draw(elves: Set<Pos>) {
    val minX = elves.minOf { it.x }
    val minY = elves.minOf { it.y }
    val maxX = elves.maxOf { it.x }
    val maxY = elves.maxOf { it.y }

    (minY..maxY).forEach { y ->
        (minX..maxX).forEach { x ->
            if (Pos(x, y) in elves) print('#')
            else print('.')
        }
        println()
    }
    println()
}

private fun hasAdjacentElves(elfPos: Pos, allElves: Set<Pos>): Boolean {
    for (y in elfPos.y - 1..elfPos.y + 1) {
        for (x in elfPos.x - 1..elfPos.x + 1) {
            val checkPos = Pos(x, y)
            if (checkPos != elfPos && checkPos in allElves)
                return true

        }
    }
    return false
}

private fun hasNoElvesInDirection(direction: Direction, elfPos: Pos, allElves: Set<Pos>) =
    direction.allPoleVectors.all { poleVector -> elfPos.plus(poleVector) !in allElves }

fun tryMove(elfPos: Pos, directionPos: Int, allElves: Set<Pos>): Pos? {
    if (!hasAdjacentElves(elfPos, allElves))
        return null

    return (directionPos until directionPos + 4)
        .map { Direction.values()[it % 4] }
        .firstOrNull { direction -> hasNoElvesInDirection(direction, elfPos, allElves) }
        ?.moveVector?.plus(elfPos)
}

private fun score(elves: MutableSet<Pos>): Int {
    val minX = elves.minOf { it.x }
    val minY = elves.minOf { it.y }
    val maxX = elves.maxOf { it.x }
    val maxY = elves.maxOf { it.y }
    return (maxX - minX + 1) * (maxY - minY + 1) - elves.size
}

fun solve(elves: Set<Pos>) {
    val finalElves = elves.toMutableSet()
    var direction = 0
    var round = 0

    do {
        var changedPositions = 0
        finalElves
            .map { oldPos -> oldPos to tryMove(oldPos, direction, finalElves) }
            .groupBy { (_, newPos) -> newPos }
            .filter { (_, elves) -> elves.size == 1 }
            .map { it.value }
            .flatten()
            .forEach { (oldPos, newPos) ->
                changedPositions++
                finalElves.remove(oldPos)
                finalElves.add(newPos!!)
            }

        if (round == 9)
            println(score(finalElves))

        round++
        direction++
    } while (changedPositions > 0)
    println(round)
}

fun main() {
    val elves = readInput("day23/input.txt")
    solve(elves)
}
