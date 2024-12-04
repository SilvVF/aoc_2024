package days

import Day
import kotlin.math.abs

object Day4: Day {
    override val day: Int = 4

    override fun part1(input: List<String>): Int {
        val arr =  input.map { line -> line.toList() }
        return arr.foldIndexed(0) { i, total, line ->
            total + line.foldIndexed(0) { j, acc, _ ->
                acc + countPossibleXMAS(i, j, arr)
            }
        }
    }

    override fun part2(input: List<String>): Int {
        val arr =  input.map { line -> line.toList() }
        return arr.foldIndexed(0) { i, total, line ->
            total + line.foldIndexed(0) { j, acc, _ ->
                acc + if(checkXMas(i, j, arr)) 1 else 0
            }
        }
    }
}

private fun checkXMas(i: Int, j: Int, arr: List<List<Char>>): Boolean {
    if (arr[i][j] != 'A') {
        return false
    }
    if (i == 0 || j == 0 || i >= arr.lastIndex || j >= arr.lastIndex) {
        return false
    }

    val tl = arr[i-1][j-1]
    val tr = arr[i-1][j+1]
    val bl = arr[i+1][j-1]
    val br = arr[i+1][j+1]

    return when {
        tl == 'M' && tr == 'S' -> bl == 'M' && br == 'S'
        tl == 'M' && tr == 'M' -> bl == 'S' && br == 'S'
        tl == 'S' && tr == 'S' -> bl == 'M' && br == 'M'
        tl == 'S' && tr == 'M' -> bl == 'S' && br == 'M'
        else -> false
    }
}

private fun countPossibleXMAS(i: Int, j: Int, arr: List<List<Char>>): Int {
    if (arr[i][j] != 'X') return 0

    val upOk = { i - 3 >= 0 }
    val downOk = { i + 3 <= arr.lastIndex }
    val rightOK = { j + 3 <= arr[i].lastIndex }
    val leftOK = { j - 3 >= 0 }

    val dirs = listOf(
        -3 to 0,
        3 to 0,
        0 to -3,
        0 to 3,
        3 to 3,
        -3 to -3,
        3 to -3,
        -3 to 3
    )

    return dirs.count { (id, jd) ->
        if (id < 0 && !upOk() || id > 0 && !downOk()) {
            return@count false
        }
        if (jd > 0 && !rightOK() || jd < 0 && !leftOK()) {
            return@count false
        }

        val iList =  when {
            id > 0 -> (0..id).toList()
            id < 0 -> (0 downTo id).toList()
            else -> listOf(0, 0, 0, 0)
        }
        val jList = when {
            jd > 0 -> (0..jd).toList()
            jd < 0 -> (0 downTo jd).toList()
            else -> listOf(0, 0, 0, 0)
        }
        buildString {
            repeat(4) { idx ->
                append(arr[i+iList[idx]][j+jList[idx]])
            }
        } == "XMAS"
    }
}