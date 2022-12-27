package day16

import java.nio.file.Files
import java.nio.file.Paths
import kotlin.math.max

const val START = "AA"

data class Valve(val id: String, val flow: Int, val tunnels: List<String>)
data class State(val valve: Valve, val minute: Int, val pressure: Int, val player: Int)


fun readInput(file: String): Map<String, Valve> {
    return Files.lines(Paths.get(file)).toList().map { it.split(" ") }
        .associate {
            it[1] to Valve(
                id = it[1],
                flow = it[4].substring(it[4].indexOf('=') + 1, it[4].length - 1).toInt(),
                tunnels = it.takeLast(it.size - 9).map { valve -> valve.replace(",", "") })
        }
}

fun traverseDfs(
    valve: Valve, minute: Int, pressure: Int, maxPressure: MutableList<Int>, valves: Map<String, Valve>, opened: MutableSet<String>, states:
    MutableSet<State>, maxMinutes: Int, player: Int
): Int {
    if (minute >= maxMinutes) {
        return if (player == 0)
            pressure
        else
            traverseDfs(valves[START]!!, 1, pressure, maxPressure, valves, opened.toMutableSet(), states.toMutableSet(), maxMinutes, player - 1)
    }

    val state = State(valve, minute, pressure, player)
    if (state in states) {
        return 0
    }
    states.add(state)

    var releasePressure = 0
    var hasOpened = false
    val openedBranch = opened.toMutableSet()
    if (valve.flow > 0 && valve.id !in opened) {
        hasOpened = true
        releasePressure = valve.flow * (maxMinutes - minute)
        openedBranch.add(valve.id)
    }

    for (v in valve.tunnels) {
        if (hasOpened)
            maxPressure[0] = max(
                maxPressure[0], traverseDfs(
                    valves[v]!!, minute + 2, pressure + releasePressure, maxPressure, valves, openedBranch.toMutableSet(),
                    states, maxMinutes, player
                )
            )
        maxPressure[0] =
            max(
                maxPressure[0], traverseDfs(
                    valves[v]!!, minute + 1, pressure, maxPressure, valves, opened.toMutableSet(),
                    states, maxMinutes, player
                )
            )
    }
    return maxPressure[0]
}


fun main() {
    val valves = readInput("day16/input-test.txt")
    println(traverseDfs(valves[START]!!, 1, 0, mutableListOf(0), valves, mutableSetOf(), mutableSetOf(), 30, 0))
    // part 2 is sketchy because it doesn't work for test data, but does for production!?
    println(traverseDfs(valves[START]!!, 1, 0, mutableListOf(0), valves, mutableSetOf(), mutableSetOf(), 26, 1))
}
