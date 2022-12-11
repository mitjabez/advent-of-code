import java.nio.file.Files
import java.nio.file.Paths
import kotlin.math.absoluteValue
import kotlin.math.sign


data class Knot(var x: Int, var y: Int) {
    fun distance(other: Knot): Knot = Knot(x - other.x, y - other.y)
    fun absDistance(other: Knot): Int = (y - other.y).absoluteValue + (x - other.x).absoluteValue
    fun isDiagonal(other: Knot): Boolean = (x != other.x && y != other.y)
    fun moveTo(newX: Int, newY: Int) {
        this.x = newX
        this.y = newY
    }

    fun isPos(x: Int, y: Int): Boolean = this.x == x && this.y == y

    fun move(direction: String) {
        when (direction) {
            "R" -> x++
            "U" -> y++
            "L" -> x--
            "D" -> y--
            else -> throw IllegalArgumentException()
        }
    }
}

fun moveTail(head: Knot, tail: Knot) {
    if (head.isDiagonal(tail)) {
        if (head.absDistance(tail) <= 2)
            return
    } else {
        if (head.absDistance(tail) <= 1)
            return
    }

    val distance = head.distance(tail)
    if (distance.y == 0) {
        // Horizontal
        tail.x += distance.x.sign
    } else if (distance.x == 0) {
        // Vertical
        tail.y += distance.y.sign
    } else {
        // Diagonal vertical
        if (distance.y.absoluteValue > distance.x.absoluteValue) {
            tail.y = head.y - distance.y.sign
            tail.x = head.x
        } else if (distance.y.absoluteValue == distance.x.absoluteValue && distance.x.absoluteValue == distance.y.absoluteValue) {
            tail.y = head.y - distance.y.sign
            tail.x = head.x - distance.x.sign
        } else {
            // Diagonal horizontal
            tail.x = head.x - distance.x.sign
            tail.y = head.y
        }
    }
}

fun draw(head: Knot, knots: List<Knot>) {
    for (y in 4 downTo 0) {
        for (x in 0..5) {
            var didPrint = false
            if (head.isPos(x, y)) {
                print("H")
                didPrint = true
            }
            for (i in knots.indices) {
                if (knots[i].isPos(x, y)) {
                    print("${i + 1}")
                    didPrint = true
                    break
                }
            }

            if (!didPrint) {
                print(".")
            }
        }
        println()
    }
}

fun part2() {
    val input = Files.lines(Paths.get("input.txt")).toList()
    val head = Knot(0, 0)
    val tails = mutableListOf<Knot>()
    for (i in 0..9)
        tails.add(Knot(0, 0))
    val visited = mutableSetOf<Knot>()
    input.map {
        val (direction, times) = it.split(" ")
        for (i in 0 until times.toInt()) {
            head.move(direction)
            var previousTail = head
            for (tailIndex in tails.indices) {
                val tail = tails[tailIndex]
                moveTail(previousTail, tail)
                previousTail = tail
                // Last tail
                if (tailIndex == 8) {
                    visited.add(tail.copy())
                }
            }
        }
    }
    println(visited.size)
}

fun part1() {
    val input = Files.lines(Paths.get("input-test.txt")).toList()

    val head = Knot(0, 0);
    val tail = Knot(0, 0);
    val visited = mutableSetOf<Knot>()
    input.map {
        val (direction, times) = it.split(" ")
        for (i in 0 until times.toInt()) {
            head.move(direction)
            moveTail(head, tail)
            visited.add(tail.copy())
        }
    }
    println(visited.size)
}

fun main() {
    //2273 your answer is too low
    //2505 too high
    part1()
    part2()
}