import java.nio.file.Files
import java.nio.file.Paths
import java.util.function.Predicate

data class Range(val min: Long, val max: Long) {
    companion object {
        fun from(sections: String): Range {
            val split = sections.split("-")
            return Range(split[0].toLong(), split[1].toLong())
        }
    }

    fun contains(other: Range): Boolean = this.min <= other.min && this.max >= other.max
    fun overlaps(other: Range): Boolean = this.min >= other.min && this.min <= other.max
}

fun solve(predicate: Predicate<Pair<Range, Range>>): Long {
    return Files.lines(Paths.get("input.txt"))
        .map { it.split(",") }
        .map { Pair(Range.from(it[0]), Range.from(it[1])) }
        .filter(predicate)
        .count()
}

fun main() {
    println(solve { it.first.contains(it.second) || it.second.contains(it.first) })
    println(solve { it.first.overlaps(it.second) || it.second.overlaps(it.first) })
}
