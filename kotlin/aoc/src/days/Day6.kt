package days

import Day
import days.Dir.*

private sealed interface Space {
    data object Empty: Space
    data object Passed: Space
    data class Occ(val hits: Lazy<MutableList<Dir>> = lazy { mutableListOf() }): Space
    data object Block: Space
}

private enum class Dir {
   UP, RIGHT, DOWN, LEFT;

    fun getNext(): Dir {
        return if (ordinal == entries.lastIndex) {
            Dir.entries[0]
        } else {
            Dir.entries[this.ordinal+1]
        }
    }
}

data class Pos(
    var x: Int,
    var y: Int
)

object Day6: Day {
    override val day: Int = 6

    private fun buildGrid(input: List<String>): Triple<Dir, Pos, Array<Array<Space>>> {
        lateinit var start: Pos
        lateinit var dir: Dir
        val grid = Array(input.size) { i ->
            Array(input[i].length) { j ->
                when(val c = input[i][j]) {
                    '.' -> Space.Empty
                    '#' -> Space.Occ()
                    '^', '<', '>', 'v' -> Space.Passed.also {
                        start = Pos(i, j)
                        dir = when(c) {
                            '^' -> UP
                            '<' -> LEFT
                            '>' -> RIGHT
                            'v' -> DOWN
                            else -> error("invalid dir")
                        }
                    }
                    else -> error("invalid input")
                }
            }
        }
        return Triple(dir, start, grid)
    }

    override fun part1(input: List<String>): Int {
        var (dir, pos, grid) = buildGrid(input)

        while (grid.checkBounds(pos.x, pos.y)) {
           val (nr, nc) = when(dir) {
               UP -> pos.x-1 to pos.y
               RIGHT -> pos.x to pos.y+1
               DOWN -> pos.x+1 to pos.y
               LEFT -> pos.x to pos.y-1
           }
           if (!grid.checkBounds(nr, nc)) {
               break
           }

           val next = grid[nr][nc]
           if (next is Space.Occ) {
               dir = dir.getNext()
               continue
           }
            grid[nr][nc] = Space.Passed
            pos.x = nr
            pos.y = nc
        }
        return grid.sumOf { line -> line.filterIsInstance<Space.Passed>().size }
    }

    override fun part2(input: List<String>): Int {
        var (dir, pos, grid) = buildGrid(input)
        var out = 0
        grid[pos.x][pos.y] = Space.Passed

        while (grid.checkBounds(pos.x, pos.y)) {
            val (nr, nc) = when(dir) {
                UP -> pos.x-1 to pos.y
                RIGHT -> pos.x to pos.y+1
                DOWN -> pos.x+1 to pos.y
                LEFT -> pos.x to pos.y-1
            }
            if (!grid.checkBounds(nr, nc)) {
                break
            }
            when(val next = grid[nr][nc]) {
                Space.Block,
                Space.Passed -> {
                    pos.x = nr
                    pos.y = nc
                    continue
                }
                is Space.Occ -> {
                    next.hits.value.add(dir)
                    dir = dir.getNext()
                    continue
                }
                Space.Empty -> {
                    grid[nr][nc] = Space.Occ()
                    grid[nr][nc] = if(grid.checkLoop(pos.x, pos.y, Pos(nr, nc), dir)) {
                        out++
                        Space.Block
                    } else {
                        Space.Passed
                    }
                    pos.x = nr
                    pos.y = nc
                }
            }
        }
        return out
    }
}

private fun Array<Array<Space>>.raycastNextDirForHit(row: Int, col: Int, dir: Dir): Pos? {
    when(dir) {
        UP -> {
            val idx = this[row].slice(col+1..this[row].lastIndex).indexOfFirst { it is Space.Occ }
            if (idx != -1) {
                return Pos(row, idx + 1 + col)
            }
        }
        RIGHT -> {
            for (i in row..lastIndex) {
                if (this[i][col] is Space.Occ) {
                    return Pos(i, col)
                }
            }
        }
        DOWN -> {
            val idx = this[row].slice(0..<col).indexOfLast { it is Space.Occ }
            if (idx != -1) {
                return Pos(row, idx)
            }
        }
        LEFT -> {
            for (i in row downTo 0) {
                if (this[i][col] is Space.Occ) {
                    return Pos(i, col)
                }
            }
        }
    }
    return null
}

private fun Array<Array<Space>>.checkLoop(row: Int, col: Int, block: Pos, dir: Dir): Boolean {
    val seen = mutableMapOf<Dir, MutableSet<Pos>>()
    seen[dir] = mutableSetOf(block)
    var cd: Dir = dir

    var pos = raycastNextDirForHit(row, col, dir)
    while(pos != null) {
        var (i, j) = pos
        cd = cd.getNext()

        val hitAlready = (this[i][j] as? Space.Occ)?.hits?.takeIf { it.isInitialized() }?.value?.contains(cd) == true
        if (seen.getOrPut(cd) { mutableSetOf() }.contains(pos) || hitAlready) {
            return true
        }
        seen[cd]!!.add(pos)

        when(cd) {
            UP -> i += 1
            DOWN -> i -= 1
            RIGHT -> j -= 1
            LEFT -> j += 1
        }
        pos = raycastNextDirForHit(i, j, cd)
    }
    return false
}

private fun Array<Array<Space>>.checkBounds(row: Int, col: Int): Boolean {
    return row in 0..lastIndex && col in 0..this[row].lastIndex
}