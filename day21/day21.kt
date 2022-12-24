import java.nio.file.Files
import java.nio.file.Paths

data class Monkey(val id: String, val yell: Long? = null, val operator: Char? = null, val other1: String? = null, val other2: String? = null)

private const val ROOT = "root"
private const val HUMN = "humn"

private fun parseInput() = Files.lines(Paths.get("day21/input.txt"))
    .map {
        val split = it.split(" ")
        val id = split[0].replace(":", "")
        val monkey = if (split.size == 2) {
            Monkey(id = id, yell = split[1].toLong())
        } else {
            Monkey(id = id, operator = split[2][0], other1 = split[1], other2 = split[3])
        }
        monkey
    }.toList().associateBy { it.id }.toMutableMap()

fun yell(m: Monkey, monkeys: Map<String, Monkey>, part: Int): Long {
    if (m.yell != null)
        return m.yell

    val yell1 = monkeys[m.other1]!!.yell ?: yell(monkeys[m.other1]!!, monkeys, part)
    val yell2 = monkeys[m.other2]!!.yell ?: yell(monkeys[m.other2]!!, monkeys, part)

    return if (part == 2 && m.id == ROOT) {
        yell1 - yell2
    } else {
        when (m.operator) {
            '+' -> yell1 + yell2
            '-' -> yell1 - yell2
            '*' -> yell1 * yell2
            '/' -> yell1 / yell2
            else -> throw IllegalArgumentException("Unknown operator ${m.operator}")
        }
    }
}

fun part2(sourceMonkeys: Map<String, Monkey>): Long {
    val monkeys = sourceMonkeys.toMutableMap()
    var min = 0L
    var max = 100000000000000L
    var mid: Long
    while (true) {
        mid = (min + max) / 2
        monkeys[HUMN] = monkeys[HUMN]!!.copy(yell = mid)
        var yellCount = yell(monkeys[ROOT]!!, monkeys, 2)
        if (yellCount < 0) {
            max = mid
        } else if (yellCount == 0L) {
            // Almost there, closing in ...
            while (true) {
                monkeys[HUMN] = monkeys[HUMN]!!.copy(yell = --mid)
                yellCount = yell(monkeys[ROOT]!!, monkeys, 2)
                if (yellCount != 0L) {
                    return mid + 1
                }
            }
        } else {
            min = mid
        }
    }

}

private fun part1(monkeys: MutableMap<String, Monkey>) = yell(monkeys[ROOT]!!, monkeys, 1)

fun main() {
    val monkeys = parseInput()

    println(part1(monkeys))
    // Works with production data. For test data you need to switch min/max logic
    println(part2(monkeys))
}
