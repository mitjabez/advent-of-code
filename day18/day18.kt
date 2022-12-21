import java.nio.file.Files
import java.nio.file.Paths
import java.util.LinkedList
import kotlin.math.absoluteValue

data class Droplet(val x: Int, val y: Int, val z: Int) {
    fun isCovered(droplet: Droplet): Boolean {
        return (x == droplet.x && y == droplet.y && (z - droplet.z).absoluteValue == 1) || (x == droplet.x && (y - droplet.y).absoluteValue == 1 && z == droplet.z) || ((x - droplet.x).absoluteValue == 1 && y == droplet.y && z == droplet.z)
    }
}

fun getAdjacent(droplet: Droplet, visited: Set<Droplet>) = listOf(
    Droplet(droplet.x, droplet.y, droplet.z - 1),
    Droplet(droplet.x, droplet.y, droplet.z + 1),
    Droplet(droplet.x, droplet.y - 1, droplet.z),
    Droplet(droplet.x, droplet.y + 1, droplet.z),
    Droplet(droplet.x - 1, droplet.y, droplet.z),
    Droplet(droplet.x + 1, droplet.y, droplet.z),
).filter { !visited.contains(it) }

fun isValidMove(droplet: Droplet, droplets: List<Droplet>): Boolean {
    val other = droplets.find { it.isCovered(droplet) }
    return (other != null && other.isCovered(droplet))
}

fun part2(droplets: List<Droplet>) {
    val minX = droplets.minOf { it.x } - 1
    val minY = droplets.minOf { it.y } - 1
    val minZ = droplets.minOf { it.z } - 1
    val maxX = droplets.maxOf { it.x } + 1
    val maxY = droplets.maxOf { it.y } + 1
    val maxZ = droplets.maxOf { it.z } + 1

    val queue = LinkedList<Droplet>()
    queue.add(Droplet(minX, minY, minZ))

    val visited = mutableSetOf<Droplet>()
    var exterior = 0

    while (queue.isNotEmpty()) {
        val droplet = queue.remove()
        if (droplet in visited)
            continue

        visited.add(droplet)

        for (dstPos in getAdjacent(droplet, visited)) {
            if (dstPos.x in minX..maxX && dstPos.y in minY..maxY && dstPos.z in minZ..maxZ) {
                if (dstPos in droplets)
                    exterior++
                else
                    queue.add(dstPos)
            }
        }
    }

    println(exterior)
}

fun part1(droplets: List<Droplet>) {
    val sidesExposed = droplets.sumOf { cube -> 6 - droplets.count { otherCube -> cube.isCovered(otherCube) } }
    println(sidesExposed)
}

fun main() {
    val input = Files.lines(Paths.get("input.txt")).map {
            val (x, y, z) = it.split(",")
            Droplet(x.toInt(), y.toInt(), z.toInt())
        }.toList()

    part1(input)
    part2(input)
}