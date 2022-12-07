import java.nio.file.Files
import java.nio.file.Paths
import java.util.stream.Stream

private fun readInput(): Stream<String> = Files.lines(Paths.get("input.txt"))

data class Node(val name: String, var size: Int = 0, val isDir: Boolean, val parent: Node?) {
    fun addSize(newSize: Int) {
        this.size += newSize
        var parent = this.parent
        while (parent != null) {
            parent.size += newSize
            parent = parent.parent
        }
    }
}

fun readDirs(): List<Node> {
    var currentDir: Node? = null
    val dirs = mutableListOf<Node>()

    val input = readInput()
        .filter{ !it.startsWith("dir") && !it.startsWith("$ ls") && !it.startsWith("dir")}
        .toList()

    for (line in input) {
        if (line.startsWith("$ cd ..")) {
            currentDir = currentDir?.parent
        } else if (line.startsWith("$ cd")) {
            currentDir = Node(name = line.substring(5), isDir = true, parent = currentDir)
            dirs.add(currentDir)
        } else {
            val (size, fileName) = line.split(" ")
            val file = Node(name = fileName, isDir = false, parent = currentDir)
            file.addSize(size.toInt())
        }
    }

    return dirs
}

fun part1(dirs: List<Node>) {
    println(dirs
        .filter { it.size <= 100000 }
        .sumOf { it.size })
}

fun part2(dirs: List<Node>) {
    val usedSize = dirs.first { it.name == "/" }.size
    val freeSpace = 70000000 - usedSize
    println(dirs
        .filter { freeSpace + it.size >= 30000000 }
        .minBy { it.size }
        .size
    )
}

fun main() {
    val dirs = readDirs()

    part1(dirs)
    part2(dirs)
}