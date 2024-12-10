package days

import Day

object Day9: Day {
    override val day: Int = 9

    override fun part1(input: List<String>): Number {
        val block = mutableListOf<String>()

        for ((id, chunk) in input.flatMap { it.asIterable() }.chunked(2).withIndex()) {
            repeat(chunk[0].digitToInt()) {
                block.add("$id")
            }
            if (chunk.size == 2) {
                repeat(chunk[1].digitToInt()) {
                    block.add(".")
                }
            }
        }

        fun getNext(): Pair<Int, Int> {
            val free = block.indexOfFirst { it == "." }
            val file = block.indexOfLast { it != "." }
            return free to file
        }

        var (free, last) = getNext()
        while(last != -1 && free != -1 && free < last) {
            block[free] = block[last]
            block[last] = "."
            val (l, r) = getNext()

            free = l
            last = r
        }
        return block.foldIndexed(0L) { idx, acc, v ->
            v.toLongOrNull()?.let{ acc + (idx * it) } ?: acc
        }
    }

    override fun part2(input: List<String>): Number {
        val block = mutableListOf<String>()
        val str = StringBuilder()

        for ((id, chunk) in input.flatMap { it.asIterable() }.chunked(2).withIndex()) {
            repeat(chunk[0].digitToInt()) {
                block.add("$id")
                str.append('1')
            }
            if (chunk.size == 2) {
                repeat(chunk[1].digitToInt()) {
                    block.add(".")
                    str.append('0')
                }
            }
        }
        var offset = block.lastIndex
        fun getNext(): Pair<IntRange, IntRange> {

            val fidx = block.slice(0..offset).indexOfLast { it != "." }
            val file = block.slice(0..fidx).takeLastWhile { it == block[fidx] }

            val range = (fidx-file.lastIndex..fidx)
            offset = fidx-file.size

            val search = buildString {
                repeat(file.size) { append('0') }
            }

            val eidx = str.indexOf(search)
            val emptyRange = (eidx..eidx+file.lastIndex)

            return emptyRange to range
        }

        do {
            val (emptyRange, fileRange) = getNext()
            if (emptyRange.first <= -1 || fileRange.first > fileRange.last || emptyRange.first > fileRange.first) {
                continue
            }
            for ((empty, file) in emptyRange.zip(fileRange)) {
                str[empty] = '1'
                str[file] = '0'

                block[empty] = block[file]
                block[file] = "."
            }
        } while (offset > block.indexOf("."))

        return block.foldIndexed(0L) { idx, acc, v ->
            v.toLongOrNull()?.let{ acc + (idx * it) } ?: acc
        }
    }
}