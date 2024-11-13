import java.io.File

fun solve(distinctSize: Int) {
    val answer = File("input.txt").readText()
        .windowed(distinctSize)
        .indexOfFirst { it.toSet().size == distinctSize  }
    println(answer + distinctSize)
}

fun main() {
    solve(4)
    solve(14)
}
