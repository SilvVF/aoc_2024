import days.*
import java.io.File

interface Day {
    val day: Int
    fun part1(input: List<String>): Int
    fun part2(input: List<String>): Int
}

fun main(args: Array<String>) {

    val days = listOf(
        Day1,
        Day2,
        Day3,
        Day4,
        Day5
    )

    val day = args.getOrNull(0)?.toIntOrNull() ?: 5
    val debug = args.getOrNull(1)?.toBoolean() ?: false

    val file = File("C:\\Users\\david\\dev\\aoc\\inputs\\d${day}${if (debug) "test" else ""}.txt")
    val lines = file.readLines()
    val d = days.find { it.day == day } ?: error("day not found")
    val p1 = d.part1(lines)
    val p2 = d.part2(lines)

    println("Part 1: $p1 \nPart 2: $p2")
}