import java.nio.file.Files
import java.nio.file.Paths
import kotlin.math.absoluteValue

enum class Command(val cycles: Int) {
    ADDX(2),
    NOOP(1);

    companion object {
        fun from(cmd: String): Command = Command.valueOf(cmd.uppercase())
    }
}

fun part2(input: MutableList<String>) {
    val crt = IntRange(0, 5).map { " ".repeat(40).toMutableList() }.toList()
    var x = 1
    val maxSize = 6 * 40
    var beamCounter = 0

    for (line in input) {
        val splits = line.split(" ")
        val cmd = Command.from(splits[0])

        for (i in 0 until cmd.cycles) {
            val crtScreenPos = beamCounter % maxSize
            val yPos = crtScreenPos / 40
            val xPos =  crtScreenPos % 40
            if ((xPos - (x % maxSize)).absoluteValue <= 1) {
                crt[yPos][xPos] = '#'
            } else {
                crt[yPos][xPos] = '.'
            }

            if (cmd == Command.ADDX && i == cmd.cycles - 1)
                x += splits[1].toInt()

            beamCounter++
        }
    }

    crt.forEach{
        println(it.joinToString(""))
    }
}

fun part1(input: MutableList<String>) {
    val results = mutableListOf(0)
    var x = 1
    for (line in Files.lines(Paths.get("input.txt")).toList()) {
        val splits = line.split(" ")
        val cmd = Command.from(splits[0])

        for (i in 0 until cmd.cycles) {
            results.add(x * results.size)
            if (cmd == Command.ADDX && i == cmd.cycles - 1)
                x += splits[1].toInt()
        }
    }
    println(results.subList(20, results.size).filterIndexed { i, b -> i % 40 == 0 }.sum())
}

fun main() {
    val input = Files.lines(Paths.get("input.txt")).toList()
    part1(input)
    part2(input)
}