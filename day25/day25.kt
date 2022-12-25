package day25

import java.nio.file.Files
import java.nio.file.Paths
import kotlin.math.pow

fun readInput(file: String): List<String> = Files.lines(Paths.get(file)).toList()

fun decimalToFifth(number: Int): String {
    var snafu = ""
    var num = number
    while (num > 0) {
        snafu = (num % 5).toString() + snafu
        num /= 5
    }
    return snafu
}

fun snafuToDecimal(number: String): Long {
    return number.reversed().mapIndexed { pow, c ->
        val num = when (c) {
            '=' -> -2
            '-' -> -1
            else -> c.digitToInt()
        }
        num * 5.0.pow(pow.toDouble()).toLong()
    }.sum()
}

fun decimalToSnafu(number: Long): String {
    var snafu = ""
    var num = number
    while (num != 0L) {
        var rem = num % 5
        if (rem > 2)
            rem -= 5
        val remDigit = when (rem) {
            -2L -> "="; -1L -> "-"; else -> rem.toString()
        }

        snafu = remDigit + snafu
        num = num / 5 + (if (rem >= 0) 0 else 1)
    }
    return snafu
}

fun part1(input: List<String>): String {
    return decimalToSnafu(input.map(::snafuToDecimal).sum())
}

fun main() {
    val input = readInput("day25/input.txt")
    println(part1(input))
    // 1-30=-3-40-4=-0--3-3
}
