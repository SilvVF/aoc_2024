package days

import Day
import kotlin.math.max

private data class Machine(
    val a: Pair<Long, Long>,
    val b: Pair<Long, Long>,
    val target: Pair<Long, Long>
)

object Day13: Day {
    override val day: Int = 13

    private fun parseInput(input: List<String>): List<Machine> {
        return input.chunked(4).map { lst ->
            val a = lst[0].slice(lst[0].indexOf("X+")+2..<lst[0].indexOf(',')).toLong() to lst[0].slice(lst[0].indexOf("Y+")+2..lst[0].lastIndex).toLong()
            val b = lst[1].slice(lst[1].indexOf("X+")+2..<lst[1].indexOf(',')).toLong() to lst[1].slice(lst[1].indexOf("Y+")+2..lst[1].lastIndex).toLong()
            val t = lst[2].slice(lst[2].indexOf("X=")+2..<lst[2].indexOf(',')).toLong() to lst[2].slice(lst[2].indexOf("Y=")+2..lst[2].lastIndex).toLong()
            Machine(a, b, t)
        }
    }

    override fun part1(input: List<String>): Number {

        val machines = parseInput(input)
        var result = 0L
        for (machine in machines) {

            val distA = with(machine.a) { first + second }
            val distB = with(machine.b) { 3 * (first + second) }

            val (cost, cost1) = if (distB > distA) { 3 to 1 } else { 1 to 3 }
            val (p, p1) = if (distB > distA) {
                machine.a to machine.b
            } else {
                machine.b to machine.a
            }

            val x = machine.target.first / p.first
            val y = machine.target.second / p.second

            val jump = minOf(x, y)

            var tokens = jump * cost
            var c = p.first * jump to p.second * jump

            while(tokens > 0) {
                val needX  = (machine.target.first - c.first) / p1.first
                val needy  = (machine.target.second - c.second) / p1.second

                if (needX == needy && (c.first + needX * p1.first) == machine.target.first && (c.second + needy * p1.second) == machine.target.second) {
                    tokens += (needX * cost1)
                    break
                }

                c = Pair(
                    c.first - p.first,
                    c.second - p.second
                )
                tokens -= cost
            }
            result += tokens

        }
        return result
    }

    private fun solveLinearEquations(
        a1: Long, b1: Long, c1: Long,
        a2: Long, b2: Long, c2: Long
    ): Pair<Long, Long>? {
        val determinant = a1 * b2 - a2 * b1

        // If determinant is zero, the equations are either dependent or inconsistent
        if (determinant == 0L) {
            println("No unique solution exists (either infinite solutions or no solution).")
            return null
        }
        // Calculate x and y using Cramer's rule
        val x = (c1 * b2 - c2 * b1) / determinant
        val y = (a1 * c2 - a2 * c1) / determinant

        return Pair(x, y)
    }


    override fun part2(input: List<String>): Number {
        val inc = 10000000000000L
        val machines = parseInput(input).map { it.copy(target = it.target.first + inc to it.target.second + inc) }
        var result = 0L
        for (machine in machines) {
            with(machine) {
               solveLinearEquations(
                   a.first,
                   b.first,
                   target.first,
                   a.second,
                   b.second,
                   target.second
               )?.let { s ->
                   if (a.first * s.first + b.first * s.second == target.first && a.second * s.first + b.second * s.second == target.second) {
                       result += (s.first*3) + s.second
                   }
                   println(s)
               }
            }
        }
        return result
    }
}