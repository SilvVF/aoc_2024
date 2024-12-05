package days

import Day
import days.Target.*
import java.util.Collections.addAll

private fun parseInput(input: List<String>): Pair<Map<Int, Set<Int>>, List<IntArray>> {
    val rules = mutableMapOf<Int, MutableSet<Int>>()
    val inputs = mutableListOf<IntArray>()

    input.forEach { line ->
        when {
            line.contains("|") -> {
                val (num, rule) = line.split("|").map { it.toInt() }
                rules[num] = rules.getOrDefault(num, mutableSetOf()).apply { add(rule) }
            }
            line.contains(",") -> {
                val values = line.split(",").map { it.toInt() }.toIntArray()
                inputs.add(values)
            }
        }
    }

    return rules to inputs
}

private fun fillFromLeft(src: IntArray, dst: IntArray, rules: Map<Int, Set<Int>>): Int {
   var filled = 0
   for ((idx, x) in src.withIndex()) {
       val valid = src.slice(idx..src.lastIndex).all { y ->
           rules[y]?.contains(x) == false
       }
       if (!valid) break
       dst[idx] = x
       filled++
   }
    return filled
}

private fun fillFromRight(src: IntArray, dst: IntArray, rules: Map<Int, Set<Int>>): Int {
    var filled = 0
    for (i in src.indices.reversed()) {
        val valid = src.slice(0..i).all { y ->
            !(rules.getOrDefault(src[i], emptySet()).contains(y) && !rules.getOrDefault(y, emptySet()).contains(src[i]))
        }
        if (!valid) break
        dst[i] = src[i]
        filled++
    }
    return filled
}

enum class Target {
    VALID,
    INVALID
}

private fun countInputs(target: Target, input: List<String>): Int {
    val (rules, inputs) = parseInput(input)

    var total = 0

    for (arr in inputs) {

        val copy = IntArray(arr.size) { 0 }

        val left = fillFromLeft(arr, copy, rules)
        val right = copy.lastIndex - fillFromRight(arr, copy, rules)

        if (arr.contentEquals(copy)) {
            total += when(target) {
                VALID -> copy[copy.size/2]
                else -> 0
            }
            continue
        }

        val currRules = mutableMapOf<Int, Set<Int>>().apply {
            arr.slice(left..right).forEach {
                this[it] = rules[it] ?: emptySet()
            }
        }
        for (i in left..right) {
            for (key in currRules.keys) {
                if(currRules.values.none { it.contains(key) }) {
                    copy[i] = key
                    currRules.remove(key)
                    break
                }
            }
        }

        total += when(target) {
            INVALID ->  copy[copy.size/2]
            else -> 0
        }
    }

    return total
}

object Day5: Day {
    override val day: Int = 5

    override fun part1(input: List<String>): Int {
        return countInputs(target = VALID, input)
    }

    override fun part2(input: List<String>): Int {
        return countInputs(target = INVALID, input)
    }
}