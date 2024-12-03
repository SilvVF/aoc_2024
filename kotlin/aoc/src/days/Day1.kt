package days

import Day
import java.util.PriorityQueue
import kotlin.math.abs


object Day1: Day {
    override val day: Int = 1

    override fun part1(input: List<String>): Int {
        val left = mutableListOf<Int>()
        val right =  mutableListOf<Int>()

        input.forEach { line ->
            val split = line.split(' ')
            left.add(split.first().toInt())
            right.add(split.last().toInt())
        }
        left.sort()
        right.sort()

        return left.foldIndexed(0) { index, acc, num ->
            acc + abs(num - right[index])
        }
    }

    override fun part2(input: List<String>): Int {
        val left = mutableMapOf<Int, Int>()
        val right =  mutableMapOf<Int, Int>()

        input.forEach { line ->
            val split = line.split(' ')
            val l = split.first().toInt()
            val r = split.last().toInt()
            left[l] = left.getOrDefault(l, 0) + 1
            right[r] = right.getOrDefault(r, 0) + 1
        }

        return left.entries.sumOf { (n, count) ->
            n * count * (right[n] ?: 0)
        }
    }
}