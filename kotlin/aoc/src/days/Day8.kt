package days

import Day
import kotlin.math.abs

data class Node (
    val x: Int,
    val y: Int,
    val antenna: Char,
    val anti: MutableList<Char> = mutableListOf()
)


object Day8: Day {
    override val day: Int = 8

    override fun part1(input: List<String>): Number {
        val antennas = mutableMapOf<Char, MutableList<Node>>()
        val grid = input.mapIndexed { i, line ->
            line.mapIndexed { j, c ->
                Node(i, j, c).also {
                    if (c != '.') {
                        if (antennas.contains(c)) {
                            antennas[c]!!.add(it)
                        } else {
                            antennas[c] = mutableListOf(it)
                        }
                    }
                }
            }
        }
        for (i in grid.indices) {
            for (j in grid.indices) {
                val node = grid[i][j]
                if(node.antenna == '.') {
                    continue
                }
                for (other in antennas[node.antenna].orEmpty()) {
                    if (other.x == i && other.y == j) {
                        continue
                    }

                    val distX = i - other.x
                    val distY = j - other.y
                    val exp = 1

                    if ((i + (distX * exp)) in 0..grid.lastIndex && (j + (distY * exp) in 0..grid[i + (distX * exp)].lastIndex)) {
                        grid[i + (distX * exp)][j + (distY * exp)].anti.add(node.antenna)
                    }
                }
            }
        }
        return grid.sumOf { row ->
            row.count { it.anti.isNotEmpty() }
        }
    }

    override fun part2(input: List<String>): Number {
        val antennas = mutableMapOf<Char, MutableList<Node>>()
        val grid = input.mapIndexed { i, line ->
            line.mapIndexed { j, c ->
                Node(i, j, c).also {
                    if (c != '.') {
                        if (antennas.contains(c)) {
                            antennas[c]!!.add(it)
                        } else {
                            antennas[c] = mutableListOf(it)
                        }
                    }
                }
            }
        }
        for (i in grid.indices) {
            for (j in grid.indices) {
                val node = grid[i][j]
                if(node.antenna == '.') {
                    continue
                }
                for (other in antennas[node.antenna].orEmpty()) {
                    if (other.x == i && other.y == j) {
                        continue
                    }

                    grid[other.x][other.y].anti.add(node.antenna)

                    val distX = i - other.x
                    val distY = j - other.y
                    var exp = 1

                    while ((i + (distX * exp)) in 0..grid.lastIndex && (j + (distY * exp) in 0..grid[i + (distX * exp)].lastIndex)) {
                        grid[i + (distX * exp)][j + (distY * exp)].anti.add(node.antenna)
                        exp += 1
                    }
                }
            }
        }
        return grid.sumOf { row ->
            row.count { it.anti.isNotEmpty() }
        }
    }
}