package main

import (
	"fmt"
	"unicode"

	"golang.org/x/example/hello/reverse" //nolint:depguard

	fib "github.com/andrewizmaylov/OTUS-go26/hw01_hello_otus/Fibbonachi"

	twins "github.com/andrewizmaylov/OTUS-go26/hw01_hello_otus/Twins"

	palindrome "github.com/andrewizmaylov/OTUS-go26/hw01_hello_otus/Palindrome"
)

func main() {
	fmt.Println(reverse.String("Hello, OTUS!"))

	// Example: use Fibonacci from the fib package
  if v, ok := fib.FibN(5); ok {
    fmt.Printf("Fibonacci(5) = %d\n", v) // 5
  }
	if v, ok := fib.FibN(7); ok {
		fmt.Printf("Fibonacci(7) = %d\n", v) // 13
	}
	if v, ok := fib.FibN(10); ok {
		fmt.Printf("Fibonacci(10) = %d\n", v) // 55
	}

	ok := twins.CheckTwins(10)
	if ok {
		fmt.Printf("CheckTwins(10) = %d\n", "true")
	} else {
		fmt.Printf("CheckTwins(10) = %d\n", "false")
	}

	ok = twins.CheckTwins(223)
	if ok {
		fmt.Printf("CheckTwins(223) = %d\n", "true")
	} else {
		fmt.Printf("CheckTwins(223) = %d\n", "false")
	}

	ok = palindrome.IsPalindrome("шош")
	if !ok {
		fmt.Printf("Не палиндром")
	} else {
		fmt.Println("Палиндром")
	}

	ok = palindrome.IsPalindromeFromString("   я     ем  змея ")

	char, count, isComplex := palindrome.InspectPhrase("Жили ли ежи? Ежели ежи жили, ели ли ежи жужелиц?")

	fmt.Println(string(char), count, isComplex)

	str := "Во дворе трава, на траве дрова"
	length := calculateLength(str)
	fmt.Println(length)
}

func calculateLength(s string) (count int) {
	sl := []rune(s)
	count = 0

	for _, v := range sl {
		if v != ' ' && unicode.IsLetter(v) {
			count ++
		}
	}

	return count
}
