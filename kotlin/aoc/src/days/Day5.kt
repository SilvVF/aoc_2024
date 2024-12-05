package days

import Day
import days.Target.*
import java.util.Collections.addAll

private fun parseInput(input: List<String>): Pair<Map<Int, Set<Int>>, List<List<Int>>> {
    val rules = mutableMapOf<Int, Set<Int>>()
    val inputs = mutableListOf<List<Int>>()

    input.forEach { line ->
        when {
            line.contains("|") -> {
                val (num, rule) = line.split("|").map { it.toInt() }
                rules[num] = rules.getOrDefault(num, emptySet()) + rule
            }
            line.contains(",") -> {
                val values = line.split(",").map { it.toInt() }
                inputs.add(values)
            }
        }
    }

    return rules to inputs
}

private fun fillFromLeft(src: List<Int>, dst: IntArray, rules: Map<Int, Set<Int>>): Int {
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

private fun fillFromRight(src: List<Int>, dst: IntArray, rules: Map<Int, Set<Int>>): Int {
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

        var left = 0
        var right = copy.lastIndex

        left = fillFromLeft(arr, copy, rules)
        right -= fillFromRight(arr, copy, rules)
        
        if (arr == copy.toList()) {
            total += when(target) {
                VALID -> copy[copy.size/2]
                else -> 0
            }
            continue
        }

        val currRules = buildMap {
            putAll(
                arr.slice(left..right).map { it to (rules[it] ?: emptySet()) }
            )
        }.toMutableMap()

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