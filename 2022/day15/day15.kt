import java.nio.file.Files
import java.nio.file.Paths
import kotlin.math.absoluteValue

fun readInput(file: String): List<Beacon> = Files.lines(Paths.get(file))
    .map {
        val split = it.replace(",", "").replace(":", "").split(" ")
        Beacon(
            sensor = Pos(x = coordinateToInt(split[2]), y = coordinateToInt(split[3])),
            beacon = Pos(x = coordinateToInt(split[8]), y = coordinateToInt(split[9])),
        )
    }.toList()

data class Pos(val x: Int, val y: Int)
data class Beacon(val sensor: Pos, val beacon: Pos) {
    val distance = (sensor.x - beacon.x).absoluteValue + (sensor.y - beacon.y).absoluteValue

    fun left(y: Int) = sensor.x - (distance - (y - sensor.y).absoluteValue)
    fun right(y: Int) = sensor.x + (distance - (y - sensor.y).absoluteValue)
}

fun coordinateToInt(coord: String): Int = coord.split("=")[1].toInt()

fun part1(beacons: List<Beacon>, y: Int) {
    val beaconPositions = beacons.map { it.beacon }.toMutableSet()

    val maxDistance = beacons.maxOf { it.distance }
    val maxX = beacons.maxOf { Math.max(it.sensor.x, it.beacon.x) } + maxDistance
    val minX = beacons.minOf { Math.min(it.sensor.x, it.beacon.x) } - maxDistance
    val pos = mutableSetOf<Int>()
    val beaconsInLine = mutableSetOf<Int>()
    for (x in minX..maxX) {
        for (b in beacons) {
            if (Pos(x, y) in beaconPositions) {
                beaconsInLine.add(x)
            }
            if ((b.sensor.x - x).absoluteValue + (b.sensor.y - y).absoluteValue <= b.distance) {
                pos.add(x)
            }
        }
    }
    println(pos.size - beaconsInLine.size)
}

private fun printFoundDistress(x: Int, y: Int) {
    println("Found: ${Pos(x, y)}, tuning frequency: ${x * 4000000L + y}")
}

fun part2(beacons: List<Beacon>, maxSize: Int) {
    for (y in 0..maxSize) {
        var x = 0
        while (x <= maxSize) {
            var isMatch = false
            // Find first beacon which matches
            for (b in beacons) {
                if (b.left(y) <= x && b.right(y) >= x) {
                    x = b.right(y) + 1
                    isMatch = true
                    break
                }
            }
            if (!isMatch) {
                printFoundDistress(x, y)
                return
            }
        }
    }
}

fun main() {
    val beacons = readInput("input.txt")
    part1(beacons, 2000000)
    part2(beacons, 4000000)
}

