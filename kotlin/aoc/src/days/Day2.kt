package days

import Day
import kotlin.math.abs

private fun incOk(lst: List<Int>): Boolean {
    var prev = 0
    lst.forEachIndexed { i, num ->
        when {
            i == 0 -> prev = num
            abs(num - prev) > 3 -> return false
            num >= prev -> return false
        }
        prev = num
    }
    return true
}
private fun decOk(lst: List<Int>): Boolean {
    var prev = 0
    lst.forEachIndexed { i, num ->
        when {
            i == 0 -> prev = num
            abs(num - prev) > 3 -> return false
            num <= prev -> return false
        }
        prev = num
    }
    return true
}

private fun buildLists(i: Int, lst: List<Int>) = buildList<List<Int>> {
    add(lst.toMutableList().apply { removeAt(i) })
    if (i > 0) {
        add(lst.toMutableList().apply { removeAt(i-1) })
    }
    if (i < lst.lastIndex) {
        add(lst.toMutableList().apply { removeAt(i+1) })
    }
}

private fun incOkRetry(lst: List<Int>): Boolean {
    var prev = 0
    lst.forEachIndexed { i, num ->
        when {
            i == 0 -> prev = num
            abs(num - prev) > 3 || num >= prev ->
                return buildLists(i, lst).any { incOk(it) }
        }
        prev = num
    }
    return true
}

private fun decOkRetry(lst: List<Int>): Boolean {
    var prev = 0
    lst.forEachIndexed { i, num ->
        when {
            i == 0 -> prev = num
            abs(num - prev) > 3 || num <= prev ->
                return buildLists(i, lst).any { decOk(it) }
        }
        prev = num
    }
    return true
}


object Day2: Day {
    override val day: Int = 2

    override fun part1(input: List<String>): Int {
        return input.mapNotNull { line ->
            line.split(' ')
                .filter { s -> s.isNotBlank() }
                .map { s -> s.toInt() }
                .takeIf { incOk(it) || decOk(it) }
        }
            .size
    }

    override fun part2(input: List<String>): Int {
        return input.mapNotNull { line ->
            line.split(' ')
                .filter { s -> s.isNotBlank() }
                .map { s -> s.toInt() }
                .takeIf { incOkRetry(it) || decOkRetry(it) }
        }
            .size
    }
}