import java.nio.file.Files
import java.nio.file.Paths
import java.util.stream.Collectors
import kotlin.math.max


private val NUMBER_REGEX = "\\d+".toRegex()

enum class Robot(val id: Int, val collectedMoney: Money) {
    ORE(0, Money(ores = 1)), CLAY(1, Money(clay = 1)), OBSIDIAN(2, Money(obsidian = 1)), GEODE(3, Money(geode = 1));
}

data class Money(var ores: Int = 0, var clay: Int = 0, var obsidian: Int = 0, var geode: Int = 0) {
    private fun add(money: Money) {
        ores += money.ores
        clay += money.clay
        obsidian += money.obsidian
        geode += money.geode
    }

    fun add(times: Int, money: Money) {
        for (i in 0 until times) {
            add(money)
        }
    }

    fun subtract(money: Money) {
        ores -= money.ores
        clay -= money.clay
        obsidian -= money.obsidian
        geode -= money.geode
    }

    fun hasBalance(cost: Money) = ores >= cost.ores && clay >= cost.clay && obsidian >= cost.obsidian && geode >= cost.geode
}

data class Blueprint(
    val id: Int, val oreRobotCost: Money, val clayRobotCost: Money, val obsidianRobotCost: Money, val geodeRobotCost: Money,
    var maxCostByResource: Map<Robot, Int> = emptyMap()
)

data class State(val i: Int, val balance: Money, val robots: Map<Robot, Int>)

fun parseBlueprints(input: List<String>): List<Blueprint> {
    return input.mapIndexed { id, line ->
        val splits = line.replace(Regex("Blueprint.*:"), "").split(".")
        val obsidianMatches = NUMBER_REGEX.findAll(splits[2]).toList()
        val geodeMatches = NUMBER_REGEX.findAll(splits[3]).toList()
        val bp = Blueprint(
            id = id,
            oreRobotCost = Money(ores = NUMBER_REGEX.find(splits[0])!!.value.toInt()),
            clayRobotCost = Money(ores = NUMBER_REGEX.find(splits[1])!!.value.toInt()),
            obsidianRobotCost = Money(ores = obsidianMatches.first().value.toInt(), clay = obsidianMatches.last().value.toInt()),
            geodeRobotCost = Money(ores = geodeMatches.first().value.toInt(), obsidian = geodeMatches.last().value.toInt()),
        )

        val maxCostByResource = mapOf(
            Robot.ORE to listOf(bp.oreRobotCost.ores, bp.clayRobotCost.ores, bp.obsidianRobotCost.ores, bp.geodeRobotCost.ores).max(),
            Robot.CLAY to listOf(bp.oreRobotCost.clay, bp.clayRobotCost.clay, bp.obsidianRobotCost.clay, bp.geodeRobotCost.clay).max(),
            Robot.OBSIDIAN to listOf(bp.oreRobotCost.obsidian, bp.clayRobotCost.obsidian, bp.obsidianRobotCost.obsidian, bp.geodeRobotCost.obsidian).max(),
            Robot.GEODE to 99
        )
        bp.maxCostByResource = maxCostByResource
        bp
    }
}

fun traverseDfs(
    minute: Int, balance: Money, blueprint: Blueprint, robots: MutableMap<Robot, Int>, visited: MutableSet<State>, maxGeodes: MutableList<Int>,
    maxMinutes: Int
): Int {
    if (minute >= maxMinutes)
        return balance.geode

    val state = State(minute, balance, robots)
    if (state in visited)
        return -1
    visited.add(state)

    var shouldTryNoBy = true
    for ((robotCost, robot) in listOf(
        Pair(blueprint.geodeRobotCost, Robot.GEODE),
        Pair(blueprint.obsidianRobotCost, Robot.OBSIDIAN),
        Pair(blueprint.clayRobotCost, Robot.CLAY),
        Pair(blueprint.oreRobotCost, Robot.ORE),
    )) {
        val shouldBuild = (robots[robot] ?: 0) < blueprint.maxCostByResource[robot]!!

        if (shouldBuild && balance.hasBalance(robotCost)) {
            val balanceInstance = balance.copy()
            balanceInstance.subtract(robotCost)
            val newRobots = robots.toMutableMap()
            if (newRobots.contains(robot))
                newRobots[robot] = newRobots[robot]!! + 1
            else
                newRobots[robot] = 1
            robots.forEach { (robot, count) -> balanceInstance.add(count, robot.collectedMoney) }
            maxGeodes[0] = max(traverseDfs(minute + 1, balanceInstance, blueprint, newRobots, visited, maxGeodes, maxMinutes), maxGeodes[0])
            if (robot == Robot.GEODE || robot == Robot.OBSIDIAN) {
                shouldTryNoBy = false
                break
            }
        }

    }

    if (shouldTryNoBy) {
        val balanceInstance = balance.copy()
        robots.forEach { (robot, count) -> balanceInstance.add(count, robot.collectedMoney) }
        maxGeodes[0] = max(traverseDfs(minute + 1, balanceInstance, blueprint, robots, visited, maxGeodes, maxMinutes), maxGeodes[0])
    }

    return maxGeodes[0]
}


fun part1(blueprints: List<Blueprint>) {
    val totalQuality = blueprints.parallelStream().map { blueprint ->
        val id = blueprint.id
        val quality = (id + 1) * traverseDfs(0, Money(), blueprint, mutableMapOf(Robot.ORE to 1), mutableSetOf(), mutableListOf(0), 24)
        quality
    }.collect(Collectors.toList())
        .sum()
    println(totalQuality)
}

fun part2(blueprints: List<Blueprint>) {
    val geodeProduct = blueprints.take(3).map { blueprint ->
        traverseDfs(0, Money(), blueprint, mutableMapOf(Robot.ORE to 1), mutableSetOf(), mutableListOf(0), 32)
    }.reduce { acc, geodeCount -> acc * geodeCount }
    println(geodeProduct)
}

fun main() {
    val input = parseBlueprints(Files.lines(Paths.get("input.txt")).toList())
    part1(input)
    // Lol, part 2 works for real data, fails on test data :)
    part2(input)
}
