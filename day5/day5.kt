import java.nio.file.Files
import java.nio.file.Paths
import java.util.stream.Stream

private fun readInput(): Stream<String> = Files.lines(Paths.get("input.txt"))

data class Rule(val count: Int, val from: Int, val to: Int) {
    companion object {
        private val NUMBER_REGEX = "\\d+".toRegex()
        fun from(line: String): Rule {
            val rules = NUMBER_REGEX.findAll(line).toList()
            return Rule(rules[0].value.toInt(), rules[1].value.toInt(), rules[2].value.toInt())
        }
    }
}

private fun move(rule: Rule, crane: (List<ArrayDeque<Char>>, Rule, Int) -> Char, stacks: List<ArrayDeque<Char>>): List<ArrayDeque<Char>> {
    for (i in 1..rule.count) {
        val crate = crane.invoke(stacks, rule, i)
        stacks[rule.to - 1].addFirst(crate)
    }
    return stacks
}

fun solve(crane: (List<ArrayDeque<Char>>, Rule, Int) -> Char) {
    val initialStacks = readInput()
        .takeWhile { !it.startsWith(" 1") }
        .map { it.chunked(4).map { crate -> crate[1] } }
        .toList()
        .fold(List(9) { ArrayDeque<Char>() }) { stacks, input ->
            stacks.mapIndexed { i, stack ->
                val crate = if (i < input.count()) input[i] else ' '
                if (crate != ' ') {
                    stack.addLast(crate)
                }
                stack
            }
        }

    val answer = readInput()
        .filter { it.startsWith("move") }
        .map { Rule.from(it) }
        .toList()
        .fold(initialStacks) { stacks, rule -> move(rule, crane, stacks) }
        .filter { it.isNotEmpty() }
        .map { it.first() }
        .joinToString("")

    println(answer)
}

fun main() {
    solve { stacks, rule, _ -> stacks[rule.from - 1].removeFirst() }
    solve { stacks, rule, index -> stacks[rule.from - 1].removeAt(rule.count - index) }
}
