import java.nio.file.Files
import java.nio.file.Paths

private val NUMBER_REGEX = "\\d+".toRegex()

data class Operation(val operator: Char, val second: Long?) {
    fun doOp(item: Long): Long {
        val secondVal = second ?: item
        return when (operator) {
            '+' -> item + secondVal
            '*' -> item * secondVal
            else -> throw IllegalArgumentException()
        }
    }
}

data class TestInfo(val divisibleBy: Long, val trueTarget: Int, val falseTarget: Int) {
    fun isDivisible(num: Long): Boolean = num % divisibleBy == 0L
}

data class Monkey(val items: ArrayDeque<Long>, var itemsInspected: Long = 0, val operation: Operation, val testInfo: TestInfo)

fun parseOp(number: String): Long? = if (number == "old") null else number.toLong()

private fun parseInstruction(lines: List<String>): Monkey {
    val operationLine = lines[2].substringAfter("=").replace(" ", "")
    val splitter = if (operationLine.contains('*')) '*' else '+'
    val (_, secondOp) = operationLine.split(splitter)
    val testInfo = TestInfo(
        divisibleBy = NUMBER_REGEX.find(lines[3])!!.value.toLong(),
        trueTarget = NUMBER_REGEX.find(lines[4])!!.value.toInt(),
        falseTarget = NUMBER_REGEX.find(lines[5])!!.value.toInt(),
    )

    NUMBER_REGEX.findAll(operationLine).map { it.value.toLong() }.toList()
    return Monkey(
        items = ArrayDeque(NUMBER_REGEX.findAll(lines[1]).map { it.value.toLong() }.toList()),
        operation = Operation(splitter, parseOp(secondOp)),
        testInfo = testInfo
    )
}

fun solve(input: List<String>, rounds: Int, divideBy: Int) {
    val monkeys = input.windowed(6, 7).map { parseInstruction(it) }
    val divisor = monkeys.map { it.testInfo.divisibleBy }.reduce { acc, divisableBy -> acc * divisableBy }
    for (round in 0 until rounds) {
        for (monkey in monkeys) {
            while (monkey.items.isNotEmpty()) {
                val worryLevel = monkey.items.removeFirst()
                val newWorryLevel = monkey.operation.doOp(worryLevel) / divideBy
                val newMonkeyId = if (monkey.testInfo.isDivisible(newWorryLevel)) monkey.testInfo.trueTarget else monkey.testInfo.falseTarget
                monkeys[newMonkeyId].items.addLast(newWorryLevel % divisor)
                monkey.itemsInspected++
            }
        }
    }

    val monkeyBusinessLevel = monkeys
        .sortedByDescending { it.itemsInspected }
        .subList(0, 2)
        .map { it.itemsInspected }
        .reduce { a, b -> a * b }
    println(monkeyBusinessLevel)
}

fun main() {
    val input = Files.lines(Paths.get("input.txt")).toList()
    solve(input, 20, 3)
    solve(input, 10000, 1)
}
