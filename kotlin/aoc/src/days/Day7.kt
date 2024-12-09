package days

import Day
import java.util.Stack
import kotlin.math.exp

object Day7: Day {
    override val day: Int = 7

    private fun parseInput(input: List<String>): List<Pair<Long, LongArray>> {
        return buildList {
            for (line in input) {
                val (expected, r) = line.split(":")
                val rest = r.trim().split(" ")
                val arr = LongArray(rest.size) { rest[it].toLong() }
                add(expected.toLong() to arr)
            }
        }
    }

    override fun part1(input: List<String>): Number {
        val map = parseInput(input)

        var target = 0L
        lateinit var nums: LongArray

        fun dfs(idx: Int,  value: Long): Boolean {
            return if (idx == nums.size || value > target) {
                value == target
            } else {
                 dfs(idx + 1, value + nums[idx]) ||
                 dfs(idx + 1, value.coerceAtLeast(1) * nums[idx])
            }
        }

        var total = 0L
        for ((expected, arr) in map) {
            target = expected
            nums = arr
            if (dfs(0, 0L)) {
                total += expected
            }
        }
        return total
    }

    override fun part2(input: List<String>): Number {
        val map = parseInput(input)

        var target = 0L
        lateinit var nums: LongArray

        fun dfs(idx: Int,  value: Long): Boolean {
            return if (idx == nums.size || value > target) {
                value == target
            } else {
                dfs(idx + 1, value + nums[idx]) ||
                        dfs(idx + 1, value.coerceAtLeast(1) * nums[idx]) ||
                        dfs(idx + 1, "$value${nums[idx]}".toLong())
            }
        }

        var total = 0L
        for ((expected, arr) in map) {
            target = expected
            nums = arr
            if (dfs(0, 0L)) {
                total += expected
            }
        }
        return total
    }
}