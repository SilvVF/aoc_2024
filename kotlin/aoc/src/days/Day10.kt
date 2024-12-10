package days

import Day

object Day10: Day {
    override val day: Int = 10

    fun buildGrid(input: List<String>): List<List<Int>> {
        return input.map { line ->
            line.map { c -> c.digitToInt() }
        }
    }

    override fun part1(input: List<String>): Number {
        val grid = buildGrid(input)

        fun checkBounds(i: Int, j: Int) = i in 0..grid.lastIndex && j in 0..grid[i].lastIndex

        val seen = mutableSetOf<Pair<Int, Int>>()
        fun scoreTrailHead(i: Int, j: Int, prev: Int): Int {
            if (
                !checkBounds(i, j) ||
                grid[i][j] != prev + 1 ||
                !seen.add(i to j)
            ) {
                return 0
            }

            if (grid[i][j] == 9) {
                return 1
            }

            val dirs = listOf(
                0 to 1,
                1 to 0,
                -1 to 0,
                0 to -1
            )

            return dirs.sumOf { (x, y) ->
                scoreTrailHead(i + x, j + y, grid[i][j])
            }
        }
        var score = 0
        for (i in grid.indices) {
            for (j in grid.indices) {
                val spot = grid[i][j]
                if (spot == 0) {
                    seen.clear()
                    score += scoreTrailHead(i, j, -1)
                }
            }
        }
        return score
    }

    override fun part2(input: List<String>): Number {
        val grid = buildGrid(input)

        fun checkBounds(i: Int, j: Int) = i in 0..grid.lastIndex && j in 0..grid[i].lastIndex

        fun scoreTrailHead(i: Int, j: Int, prev: Int): Int {
            if (
                !checkBounds(i, j) ||
                grid[i][j] != prev + 1
            ) {
                return 0
            }

            if (grid[i][j] == 9) {
                return 1
            }

            val dirs = listOf(
                0 to 1,
                1 to 0,
                -1 to 0,
                0 to -1
            )

            return dirs.sumOf { (x, y) ->
                scoreTrailHead(i + x, j + y, grid[i][j])
            }
        }
        var score = 0
        for (i in grid.indices) {
            for (j in grid.indices) {
                val spot = grid[i][j]
                if (spot == 0) {
                    score += scoreTrailHead(i, j, -1)
                }
            }
        }
        return score
    }
}