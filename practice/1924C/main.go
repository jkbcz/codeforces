package main

import (
	"bufio"
	"fmt"
	"os"
)

var stdin = bufio.NewReader(os.Stdin)
var stdout = bufio.NewWriter(os.Stdout)

var mod = 999999893

func main() {
	var t int
	fmt.Fscan(stdin, &t)

	for i := 0; i < t; i++ {
		solve()
	}

	stdout.Flush()
}

func solve() {
	var n int
	fmt.Fscan(stdin, &n)

	a := fastExp(2, n/2+1) - 2
	b := fastExp(2, (n+1)/2) - 2

	p := 2 * a
	p %= mod
	q := 2*(b+2)*(b+2) - a*a

	q = ((mod + (q % mod)) % mod)

	result := p * inv(q)
	result %= mod

	fmt.Fprintln(stdout, result)
}

func fastExp(base int, exponent int) int {
	var result int = 1

	base = base % mod

	for exponent > 0 {
		if exponent%2 == 1 {
			result = (result * base) % mod
		}
		exponent = exponent >> 1
		base = (base * base) % mod
	}
	return result
}

func inv(a int) int {
	_, x, _ := extendedGcd(a, mod)

	return ((mod + (x % mod)) % mod)
}

func extendedGcd(a, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	}
	gcd, xPrime, yPrime := extendedGcd(b%a, a)
	return gcd, yPrime - (b/a)*xPrime, xPrime
}
