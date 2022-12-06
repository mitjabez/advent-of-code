import java.io.File

fun solve(distinctSize: Int) {
    val answer = File("input.txt").readText()
        .windowed(distinctSize)
        .mapIndexed { pos, buffer -> Pair(pos + distinctSize, buffer.toSet()) }
        .find { it.second.size == distinctSize }
        ?.first
    println(answer)
}

fun main() {
    solve(4)
    solve(14)
}
