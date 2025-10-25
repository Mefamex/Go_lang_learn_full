/*
author: mefamex
date  : 2025-07-02
title : loops

> GO Dili Döngüler
- Sadece "FOR" döngüsü vardır.
- Go dilinde döngüler, program akışını kontrol etmek için kullanılır.
- Go dilinde döngüler, `for` anahtar kelimesi ile tanımlanır.
- Go dilinde `for` döngüsü, C ve Java gibi dillerdeki `for` döngüsüne benzer şekilde çalışır.
- Go dilinde `for` döngüsü, üç farklı şekilde kullanılabilir:
    1. C-style klasik for döngüsü
    2. While benzeri kullanım
    3. Do-While benzeri kullanım
- `break` ve `continue` anahtar kelimeleri ile kontrol edilebilir.
- `range` anahtar kelimesi ile de kullanılabilir.
- etiketler (labels) ile de kontrol edilebilir.
- `goto` anahtar kelimesi ile de kullanılabilir, ancak bu genellikle önerilmez.

> COMPARISON OPERATORS
== Equal to
!= Not equal to
<  Less than
>  Greater than
<= Less than or equal to
>= Greater than or equal to


> LOGICAL OPERATORS
&& AND
|| OR
!  NOT


*/

package main

import "fmt"

func main() {

	// Sabit değerler
	const (
		maxCount    = 5
		maxIterator = 10
		rows        = 3
		cols        = 2
	)

	////
	////

	fmt.Println("\n--- Basic For Loop ---")
	// C-style klasik for döngüsü
	for i := 0; i < maxCount; i++ {
		fmt.Println("i:", i)
	}

	////
	////

	fmt.Println("\n--- While Style Loop ---")
	// While benzeri kullanım
	i := 0
	for i < maxCount {
		fmt.Println("i:", i)
		i++
	}

	////
	////

	fmt.Println("\n--- Do-While Style Loop ---")
	// Do-While benzeri kullanım
	for {
		fmt.Println("i:", i)
		i++
		if i >= maxIterator {
			break
		}
	}

	////
	////

	fmt.Println("\n--- Range örnekleri:")
	// Slice ile range
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Slice - Index: %d, Value: %d\n", index, value)
	}

	////
	////

	// String ile range
	text := "Merhaba"
	for i, char := range text {
		fmt.Printf("String - Index: %d, Karakter: %c\n", i, char)
	}

	////
	////

	// Map ile range
	colors := map[string]string{
		"red":   "kırmızı",
		"blue":  "mavi",
		"green": "yeşil",
	}
	for key, value := range colors {
		fmt.Printf("Map - Key: %s, Value: %s\n", key, value)
	}

	////
	////

	fmt.Println("\n--- Infinite loop ---")
	for {
		fmt.Println("This is an infinite loop") // -> This will print forever until you stop the program
		// You can use 'break' to exit the loop or 'continue' to skip to the next iteration.
		break // Break the infinite loop
	}

	////
	////

	fmt.Println("\n--- Loop with continue ---")
	for i := 0; i < 10; i++ {
		if i%2 == 0 { // Skip even numbers
			continue // Skip the rest of the loop for this iteration
		}
		fmt.Println("Odd number:", i) // -> Odd number: 1, Odd number: 3, ..., Odd number: 9
	}

	////
	////

	fmt.Println("\n--- Loop with break ---")
	for i := 0; i < 10; i++ {
		if i == 5 { // Break the loop when i is 5
			break // Exit the loop
		}
		fmt.Println("Number:", i) // -> Number: 0, Number: 1, ..., Number: 4
	}

	////
	////

	fmt.Println("\n--- Nested loops ---")
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println("i:", i, "j:", j) // -> i: 0 j: 0, i: 0 j: 1, i: 1 j: 0, ..., i: 2 j: 1
		}
	}

	////
	////

	fmt.Println("\n--- Nested Loops with Labels ---")
	// Label kullanımı - dış döngüyü kontrol etmek için
OuterLoop:
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == 1 && j == 1 {
				break OuterLoop
			}
			fmt.Println("i:", i, "j:", j)
		}
	}

	////
	////

	fmt.Println("\n--- Goto Example (Not Recommended) ---")
	// NOT: goto kullanımı genellikle önerilmez!
	// Çünkü kodun okunabilirliğini azaltır ve karmaşık kontrol akışlarına neden olabilir.
	// Go dilinde goto kullanımı, etiketler ile birlikte kullanılır.
	// Bunun yerine daha yapısal kontrol akışları tercih edilmelidir:
	// - break/continue
	// - return
	// - if/else
	// - switch
	for i := 0; i < maxCount; i++ {
		if i == 3 {
			goto End
		}
		fmt.Println("i:", i)
	}
End:
	fmt.Println("End of loop")
}
