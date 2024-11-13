package day13

import java.nio.file.Files
import java.nio.file.Paths
import java.util.LinkedList

fun readInput(file: String) = Files.lines(Paths.get(file)).toList().chunked(3).map { Pair(it[0], it[1]) }

data class Sequence(var number: Int? = null, var list: MutableList<Sequence> = mutableListOf())

val sequenceComparator = Comparator { left: Sequence, right: Sequence -> compare(left, right) }

fun parseSequence(sequence: String): Sequence {
    val queue = LinkedList<Sequence>()
    var root: Sequence? = null
    var num = ""

    for (c in sequence) {
        if (c == '[') {
            val newItem = Sequence()
            if (queue.isNotEmpty()) {
                val currentItem = queue.peekLast()
                currentItem.list.add(newItem)
            }
            queue.add(newItem)
            if (root == null)
                root = newItem
        } else if (c in listOf(']', ',')) {
            if (num != "") {
                val newItem = Sequence(number = num.toInt())
                queue.peekLast().list.add(newItem)
            }

            if (c == ']') {
                if (num == "")
                    queue.peekLast().list.add(Sequence())
                queue.removeLast()
            }
            num = ""
        } else {
            num += c
        }
    }
    return root!!
}

fun compare(left: Sequence, right: Sequence): Int {
    if (left.number != null && right.number != null) {
        return left.number!!.compareTo(right.number!!)
    }

    // If one item is number convert to list
    val leftItems = if (left.list.isEmpty() && left.number != null) listOf(Sequence(number = left.number)) else left.list
    val rightItems = if (right.list.isEmpty() && right.number != null) listOf(Sequence(number = right.number)) else right.list
    for (i in rightItems.indices) {
        if (i > leftItems.size - 1)
            return -1
        val order = compare(leftItems[i], rightItems[i])
        if (order != 0)
            return order
    }

    return leftItems.size.compareTo(rightItems.size)
}
private fun sequenceItem(num: Int) = Sequence(
    number = null, list = mutableListOf(
        Sequence(number = null, list = mutableListOf(Sequence(number = num, list = mutableListOf()))), Sequence(
            number = null,
            list = mutableListOf()
        )
    )
)

fun part2(input: List<Pair<String, String>>) {
    val sorted = (input + listOf(Pair("[[2]]", "[[6]]")))
        .flatMap { it.toList() }.map { parseSequence(it) }.sortedWith(sequenceComparator)
    println((sorted.indexOf(sequenceItem(2)) + 1) * (sorted.indexOf(sequenceItem(6)) + 1))
}

fun part1(input: List<Pair<String, String>>) {
    val result = input.mapIndexed { i, pair ->
        val result = compare(parseSequence(pair.first), parseSequence(pair.second))
        if (result == -1) i + 1 else 0
    }.sum()
    println(result)
}

fun main() {
    val input = readInput("day13/input.txt")
    part1(input)
    part2(input)
}

