import java.nio.file.Files
import java.nio.file.Paths
import java.util.stream.Stream

private fun readInput(): Stream<String> = Files.lines(Paths.get("input.txt"))

fun isVisible(x: Int, y: Int, grid: List<List<Char>>): Boolean {
    if (x == 0 || y == 0 || x == grid.size - 1 || y == grid.size - 1)
        return true

    val tree = grid[y][x]
    if (grid[y].subList(0, x).max() < tree || grid[y].subList(x + 1, grid[y].size).max() < tree)
        return true

    var isVisible = true
    for (yPos in 0 until y) {
        if (grid[yPos][x] >= tree) {
            isVisible = false
            break
        }

    }

    if (isVisible)
        return true

    isVisible = true
    for (yPos in y + 1 until grid.size) {
        if (grid[yPos][x] >= tree) {
            isVisible = false
            break
        }
    }
    return isVisible
}

fun scenicScore(x: Int, y: Int, grid: List<List<Char>>): Int {
    var totalScore = 1

    val tree = grid[y][x]
    var score = 0
    for (yPos in y - 1 downTo 0) {
        if (grid[yPos][x] <= tree)
            score++

        if (grid[yPos][x] == tree)
            break
    }
    totalScore *= score
    score = 0
    for (yPos in y + 1 until grid.size) {
        score++
        if (grid[yPos][x] >= tree)
            break
    }
    totalScore *= score
    score = 0
    for (xPos in x - 1 downTo 0) {
        score++
        if (grid[y][xPos] >= tree)
            break
    }
    totalScore *= score
    score = 0
    for (xPos in x + 1 until grid[y].size) {
        score++
        if (grid[y][xPos] >= tree)
            break
    }
    totalScore *= score
    return totalScore
}

fun part2() {
    val grid = readInput()
        .map { it.toList() }
        .toList()

    val scores = mutableListOf<Int>()
    for (y in grid.indices) {
        for (x in grid[y].indices) {
            scores.add(scenicScore(x, y, grid))
        }
    }
    println(scores.max())
}

fun part1() {
    val grid = readInput()
        .map { it.toList() }
        .toList()

    var visible = 0
    for (y in grid.indices) {
        for (x in grid[y].indices) {
            if (isVisible(x, y, grid)) {
                visible++
            }
        }
    }
    println(visible)
}

fun main() {
    part1()
    part2()
}