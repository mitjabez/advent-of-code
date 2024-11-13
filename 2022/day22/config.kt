package day22

enum class Direction(val id: Int, val vector: Pos) {
    RIGHT(0, Pos(1, 0)), DOWN(1, Pos(0, 1)), LEFT(2, Pos(-1, 0)), UP(3, Pos(0, -1));

    companion object {
        fun find(vector: Pos) = Direction.values().first { it.vector == vector }
    }
}

val DIRECTIONS = listOf(Pos(1, 0), Pos(0, 1), Pos(-1, 0), Pos(0, -1))

data class Edge(val quadrant: Quadrant, val face: Direction, val newDirection: Direction)

const val QUADRANT_SIZE = 50

data class Pos(var x: Int, var y: Int) {
    fun move(steps: Int, direction: Pos): Pos {
        return Pos(x + steps * direction.x, y + steps * direction.y)
    }

    fun move(direction: Pos): Pos {
        return Pos(x + direction.x, y + direction.y)
    }
}

val quadrantA = Quadrant("A", 50, 0)
val quadrantB = Quadrant("B", 100, 0)
val quadrantC = Quadrant("C", 50, 50)
val quadrantD = Quadrant("D", 50, 100)
val quadrantE = Quadrant("E", 0, 100)

val quadrantF = Quadrant("F", 0, 150)

data class MoveRule(val rightEdge: Edge, val bottomEdge: Edge, val leftEdge: Edge, val topEdge: Edge)

val MOVE_RULES = mapOf(
    quadrantA to mapOf(
        Direction.RIGHT to Edge(quadrantB, Direction.LEFT, Direction.RIGHT),
        Direction.DOWN to Edge(quadrantC, Direction.UP, Direction.DOWN),
        Direction.LEFT to Edge(quadrantE, Direction.LEFT, Direction.RIGHT),
        Direction.UP to Edge(quadrantF, Direction.LEFT, Direction.RIGHT)
    ),
    quadrantB to mapOf(
        Direction.RIGHT to Edge(quadrantD, Direction.RIGHT, Direction.LEFT),
        Direction.DOWN to Edge(quadrantC, Direction.RIGHT, Direction.LEFT),
        Direction.LEFT to Edge(quadrantA, Direction.RIGHT, Direction.LEFT),
        Direction.UP to Edge(quadrantF, Direction.DOWN, Direction.UP)
    ),
    quadrantC to mapOf(
        Direction.RIGHT to Edge(quadrantB, Direction.DOWN, Direction.UP),
        Direction.DOWN to Edge(quadrantD, Direction.UP, Direction.DOWN),
        Direction.LEFT to Edge(quadrantE, Direction.UP, Direction.DOWN),
        Direction.UP to Edge(quadrantA, Direction.DOWN, Direction.UP)
    ),
    quadrantD to mapOf(
        Direction.RIGHT to Edge(quadrantB, Direction.RIGHT, Direction.LEFT),
        Direction.DOWN to Edge(quadrantF, Direction.RIGHT, Direction.LEFT),
        Direction.LEFT to Edge(quadrantE, Direction.RIGHT, Direction.LEFT),
        Direction.UP to Edge(quadrantC, Direction.DOWN, Direction.UP)
    ),
    quadrantE to mapOf(
        Direction.RIGHT to Edge(quadrantD, Direction.LEFT, Direction.RIGHT),
        Direction.DOWN to Edge(quadrantF, Direction.UP, Direction.DOWN),
        Direction.LEFT to Edge(quadrantA, Direction.LEFT, Direction.RIGHT),
        Direction.UP to Edge(quadrantC, Direction.LEFT, Direction.RIGHT)
    ),
    quadrantF to mapOf(
        Direction.RIGHT to Edge(quadrantD, Direction.DOWN, Direction.UP),
        Direction.DOWN to Edge(quadrantB, Direction.UP, Direction.DOWN),
        Direction.LEFT to Edge(quadrantA, Direction.UP, Direction.DOWN),
        Direction.UP to Edge(quadrantE, Direction.DOWN, Direction.UP)
    ),
)
data class Quadrant(val id: String, val minX: Int, val minY: Int) {
    val maxX = minX + QUADRANT_SIZE - 1
    val maxY = minY + QUADRANT_SIZE - 1
}
