package days

import Day

private sealed interface  Token {
    data class Num(val n: Int): Token
    data object Mul: Token
    data object Sep: Token
    data object Open: Token
    data object Close: Token
    data object Do: Token
    data object Dont: Token
    data object Invalid: Token
}

private fun checkTokens(tokens: List<Token>): Boolean {
    return tokens.size == 6 &&
            tokens[0] is Token.Mul &&
            tokens[1] is Token.Open &&
            tokens[2] is Token.Num &&
            tokens[3] is Token.Sep &&
            tokens[4] is Token.Num &&
            tokens[5] is Token.Close
}

private fun parseTokens(input: String, ignoreWhenDisabled: Boolean): Int {
    var i = 0
    val tokens = mutableListOf<Token>()
    var enabled = true
    var out = 0

    while(i < input.length)  {
        val c = input[i]
        when {
            c == 'd' -> {
                tokens.clear()
                when {
                    input.slice(i..i+3) == "do()" -> {
                        i+=4
                        enabled =  true
                    }
                    input.slice(i..i+6) == "don't()" -> {
                        i+=7
                        enabled = false
                    }
                    else -> i++
                }
            }
            !enabled && ignoreWhenDisabled -> i++
            c == 'm' -> {
                tokens.clear()
                if (input.slice(i..i+2) == "mul") {
                    tokens.add(Token.Mul)
                    i+=3
                } else {
                    i++
                }
            }
            c == ',' -> tokens.add(Token.Sep).also { i++ }
            c == '(' -> tokens.add(Token.Open).also { i++ }
            c == ')' -> {
                tokens.add(Token.Close).also { i++ }
                if (checkTokens(tokens)) {
                    val (l, r)  = tokens.filterIsInstance<Token.Num>()
                    out += l.n * r.n
                }
            }
            c.isDigit() -> {
                var s = "$c"
                while(input[++i].isDigit()) {
                    s += input[i]
                }
                tokens.add(Token.Num(s.toInt()))
            }
            else -> tokens.clear().also { i++ }
        }
    }
    return out
}

object Day3: Day {

    override val day: Int = 3

    override fun part1(input: List<String>): Int {
        return parseTokens(input.joinToString(), false)
    }

    override fun part2(input: List<String>): Int {
        return parseTokens(input.joinToString(), true)
    }
}