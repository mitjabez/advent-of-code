package day22

import day22.Direction.*
import java.lang.IllegalStateException
import java.nio.file.Files
import java.nio.file.Paths
import java.util.LinkedList

/**
 * Code is not ideal (lots of duplicates) but gets the job done
 * Part2 rules are hardcoded for my production input
 */

fun isOnEdgeQuadrant(pos: Pos, quadrant: Quadrant, direction: Direction) =
    ((direction == LEFT && pos.x == quadrant.minX) ||
            (direction == UP && pos.y == quadrant.minY) ||
            (direction == RIGHT && pos.x == quadrant.maxX) ||
            (direction == DOWN && pos.y == quadrant.maxY))

fun findQuadrant(pos: Pos): Quadrant {
    return MOVE_RULES.keys.first { pos.x in it.minX..it.maxX && pos.y in it.minY..it.maxY }
}


fun move3D(pos: Pos, direction: Pos): Pair<Pos, Pos> {
    val oldQuadrant = findQuadrant(pos)
    val oldDirection = Direction.find(direction)

    val new2DPos = pos.move(direction)

    if (!isOnEdgeQuadrant(pos, oldQuadrant, oldDirection)) {
        return Pair(new2DPos, direction)
    }

    val newEdge = MOVE_RULES[oldQuadrant]!![oldDirection]
    val newQuadrant = newEdge!!.quadrant

    val newFace = newEdge.face
    val diffX = pos.x - oldQuadrant.minX
    val diffY = pos.y - oldQuadrant.minY

    // Only few rotations are present in my input
    val newPos: Pos = when (oldDirection) {
        RIGHT -> {
            when (newFace) {
                RIGHT -> Pos(newQuadrant.maxX, newQuadrant.maxY - diffY)
                DOWN -> Pos(newQuadrant.minX + diffY, newQuadrant.maxY)
                LEFT -> pos.move(direction)
                UP -> TODO()
            }
        }

        DOWN -> {
            when (newFace) {
                RIGHT -> Pos(newQuadrant.maxX, newQuadrant.minY + diffX)
                DOWN -> TODO()
                LEFT -> TODO()
                UP -> Pos(newQuadrant.minX + diffX, newQuadrant.minY)
            }
        }

        LEFT -> {
            when (newFace) {
                RIGHT -> Pos(newQuadrant.maxX, pos.y)
                DOWN -> TODO()
                LEFT -> Pos(newQuadrant.minX, newQuadrant.maxY - diffY)
                UP -> Pos(newQuadrant.minX + diffY, newQuadrant.minY)
            }
        }

        UP -> {
            when (newFace) {
                RIGHT -> TODO()
                DOWN -> Pos(newQuadrant.minX + diffX, newQuadrant.maxY)
                LEFT -> Pos(newQuadrant.minX, newQuadrant.minY + diffX)
                UP -> TODO()
            }
        }
    }

    return Pair(newPos, newEdge.newDirection.vector)
}

fun readInput(file: String): Pair<List<String>, MutableList<String>> {
    val lines = Files.lines(Paths.get(file)).toList()
    var num = ""
    val path = mutableListOf<String>()
    for (c in lines.last()) {
        if (c in listOf('L', 'R')) {
            if (num != "")
                path.add(num)
            path.add(c.toString())
            num = ""
        } else {
            num += c
        }
    }

    if (num != "") {
        path.add(num)
    }

    var map = lines.takeWhile { it != "" }
    val maxLineSize = map.maxOf { it.length }
    map = map.map { it.padEnd(maxLineSize) }
    return Pair(map, path)
}

fun part2(map: List<String>, path: LinkedList<String>) {
    var pos3D = Pos(x = map[0].indexOf('.'), y = 0)
    var direction3D = 0

    var j = 0
    while (path.isNotEmpty()) {
        var steps: Int
        when (val pathPart = path.remove()) {
            "R" -> {
                direction3D = mod(direction3D + 1, DIRECTIONS.size)
                continue
            }

            "L" -> {
                direction3D = mod(direction3D - 1, DIRECTIONS.size)
                continue
            }

            else -> steps = pathPart.toInt()
        }

        for (i in 0 until steps) {
//            println("#${j++} ${pos3D.y},${pos3D.x} ; (${DIRECTIONS[direction3D].y}, ${DIRECTIONS[direction3D].x})")

            val moveResult = move3D(pos3D, DIRECTIONS[direction3D])
            val newPos3D = moveResult.first

            val item = map[newPos3D.y][newPos3D.x]
            if (item == '.') {
                pos3D = newPos3D
                direction3D = Direction.find(moveResult.second).id
            } else if (item == '#') {
                break
            } else {
                throw IllegalStateException()
            }
        }
    }
    println(1000 * (pos3D.y + 1) + 4 * (pos3D.x + 1) + direction3D)
}

fun part1(map: List<String>, path: LinkedList<String>) {
    var pos = Pos(x = map[0].indexOf('.'), y = 0)
    var direction = 0

    val minMaxByLine = map.mapIndexed { i, line ->
        i to Pair(line.indexOfFirst { it != ' ' }, line.indexOfLast { it != ' ' })
    }.toMap()
    val minMaxByColumn = IntRange(0, map[0].length - 1).associateWith { x ->
        val first = map.indexOfFirst { it[x] != ' ' }
        val second = map.indexOfLast { it[x] != ' ' }

        Pair(first, second)
    }

    var j = 0

    while (path.isNotEmpty()) {
        var steps: Int
        when (val pathPart = path.remove()) {
            "R" -> {
                direction = mod(direction + 1, DIRECTIONS.size)
                continue
            }

            "L" -> {
                direction = mod(direction - 1, DIRECTIONS.size)
                continue
            }

            else -> steps = pathPart.toInt()
        }

        for (i in 0 until steps) {
//            println("#${j++} ${pos3D.y},${pos3D.x} ; (${DIRECTIONS[direction3D].y}, ${DIRECTIONS[direction3D].x})")
            val newPos = pos.move(DIRECTIONS[direction])

            if (DIRECTIONS[direction].y == 0) {
                val (minX, maxX) = minMaxByLine[newPos.y]!!
                if (newPos.x > maxX)
                    newPos.x = minX
                else if (newPos.x < minX)
                    newPos.x = maxX
            } else {
                if (newPos.x < 0) {
                    val (_, maxX) = minMaxByLine[newPos.y]!!
                    newPos.x = maxX
                } else if (newPos.x >= map[0].length) {
                    val (minX, _) = minMaxByLine[newPos.y]!!
                    newPos.x = minX
                }
                val (minY, maxY) = minMaxByColumn[newPos.x]!!
                if (newPos.y > maxY)
                    newPos.y = minY
                else if (newPos.y < minY)
                    newPos.y = maxY
            }

            val item = map[newPos.y][newPos.x]
            if (item == '.') {
                pos = newPos
            } else if (item == '#') {
                break
            } else {
                throw IllegalStateException()
            }
        }
    }
    println(1000 * (pos.y + 1) + 4 * (pos.x + 1) + direction)
}

fun mod(num: Int, size: Int): Int {
    return if (num < 0) ((num % size + size) % size)
    else (num % size)
}


fun main() {
    val (map, path) = readInput("day22/input.txt")
    part1(map, LinkedList(path.toList()))
    part2(map, LinkedList(path.toList()))
}
