import java.nio.file.Files
import java.nio.file.Paths

fun mod(num: Long, size: Int): Int {
    return if (num < 0) ((num % size + size) % size).toInt()
    else (num % size).toInt()
}

fun solve(numbers: List<Long>, repeat: Int = 1, key: Int = 1) {
    val numPos = numbers.mapIndexed { i, num -> Pair(i, num * key) }
    val mixed = numPos.toMutableList()

    for (i in 0 until repeat) {
        for (num in numPos) {
            val oldPos = mixed.indexOf(num)
            val newPos: Long = oldPos + num.second

            mixed.removeAt(oldPos)
            mixed.add(mod(newPos, numbers.size - 1), num)
        }
    }

    val zeroPos = mixed.indexOfFirst { it.second == 0L }
    val sum = IntRange(1, 3).sumOf {
        mixed[(zeroPos + (it * 1000)) % mixed.size].second
    }
    println(sum)
}


fun main() {
    val input = Files.lines(Paths.get("input.txt")).map { it.toLong() }.toList()
    solve(input)
    solve(input, 10, 811589153)
}
