/*
author: mefamex
date  : 2025-07-02
title : Variables and Types in Go
*/

package main

import "fmt"

func main() {

	/* GO Dilinde Degiskenler
	- Değişkenler, programda veri saklamak için kullanılır.
	- Go dilinde değişkenler, `var` anahtar kelimesi ile tanımlanır.
	- Kısa değişken tanımlama için `:=` operatörü kullanılır
	- Değişkenler, farklı veri tiplerinde olabilir: int, float64, string, bool vb.
	- Sabitler ise `const` anahtar kelimesi ile tanımlanır

	- _ : blank identifier
	*/

	var (
		val1     = 5
		val2 int = 10
		val3 float64
		val4 string = "Hello, Go!"
		val5 bool   = true
	)
	var val6, val7, val8, val9, val10 = 1, 2, 3, "Hello", false // Multiple variable declaration
	fmt.Println(val1, val2, val3, val4, val5, val6, val7, val8, val9, val10)
	fmt.Println()

	// Integer variables
	var integerVar int = 10
	var shortIntVar = 20 // Short declaration

	// Float variables
	var floatVar float64 = 3.14
	shortFloatVar := 2.718

	// String variables
	var name string = "Go"
	shortName := "Golang"

	// Boolean variables
	var isAwesome bool = true
	shortBool := false

	// Constants
	const pi = 3.14159

	// Print all variables
	fmt.Println("Integer variables:", integerVar, shortIntVar)
	fmt.Println("Float variables:", floatVar, shortFloatVar)
	fmt.Println("String variables:", name, shortName)
	fmt.Println("Boolean variables:", isAwesome, shortBool)
	fmt.Println("Constant value of pi:", pi)

	// Variable operations
	sum := integerVar + shortIntVar
	product := floatVar - shortFloatVar

	fmt.Println("Sum of integers:", sum)
	fmt.Println("Product of floats:", product)
	fmt.Println("String concatenation:", name+" "+shortName)

	////
	////
	////

	/* GO Dilinde Degisken Kavramlari

	- Declaration: Değişkenin tanımlanması
		-> var x int

	- Assign: Değişkene bir değer atanması
		-> x = 10

	- Initialization: Değişkenin tanımlanırken bir değerle başlatılması
		-> var x int = 10

	- Initial Value: Değişkenin ilk değeri
		-> Değişken tanımlanırken verilen değer
		-> Örnek: var x int = 10 (x değişkeni int tipinde olarak tanımlanır ve 10 değeri atanır)

	- Zero Value: Değişkenin varsayılan değeri
		-> Go dilinde, değişkenler tanımlandığında otomatik olarak varsayılan değer alır.
		-> Örnek:    var x int   ->   (x değişkeni int tipinde olarak tanımlanır ve varsayılan değeri 0'dır)

	- Static Typing: Değişkenlerin tiplerinin derleme zamanında belirlenmesi
		-> Go dilinde statik tipli bir dildir, yani değişkenlerin tipleri derleme zamanında belirlenir.

	- Dynamic Typing: Değişkenlerin tiplerinin çalışma zamanında belirlenmesi
		-> Go dinamik tipli bir dil değildir, ancak interface kullanarak dinamik davranışlar elde edilebilir.

	- Type Inference: Go dilinde, değişkenlerin tipleri otomatik olarak belirlenir.

	- Scope: Değişkenin erişim alanı
		-> Değişkenin tanımlandığı blok içinde geçerlidir.

	- Shadowing: Aynı isimde bir değişkenin daha dar bir kapsamda tanımlanması
		-> Daha dar kapsamda tanımlanan değişken, dışarıdaki değişkeni gölgeler.

	- [:=] Short Declaration Operator
		-> Kısa değişken tanımlama operatörü, değişkeni tanımlarken ve başlatırken kullanılır.
		-> Örnek: x := 10 (x değişkeni int tipinde olarak tanımlanır ve 10 değeri atanır)

	- [=] Assignment Operator
		-> Değişkene bir değer atamak için kullanılır.
		-> Örnek: x = 20 (x değişkenine 20 değeri atanır)

	- var (Variable Declaration)
		-> Değişken tanımlamak için kullanılır.
		-> Örnek: var x int (x değişkeni int tipinde olarak tanımlanır, ancak başlangıç değeri verilmez)

	- const (Constant Declaration)
		-> Sabit tanımlamak için kullanılır.

	*/

	////
	////
	////

	/* Go Dilinde Tipler
	- Go dilinde veri tipleri, değişkenlerin saklayabileceği veri türlerini belirler.
	- Temel veri tipleri: int, float64, string, bool
	- Go, statik tipli bir dildir, yani değişkenlerin tipleri derleme zamanında belirlenir.
	- Go, tip güvenli bir dildir, yani bir değişkenin tipi ile uyumlu olmayan bir değer atanmaya çalışılırsa hata verir.
	- Go, tip dönüşümlerini destekler, ancak bu dönüşümler açıkça belirtilmelidir.
	- Go, kullanıcı tanımlı tipleri destekler, bu sayede kendi veri tiplerini oluşturabilirsiniz.
	*/

	// Firstly complex types, then basic types
	// Complex types:
	// Array, Slice, Map, Struct, Pointer, Function, Channel, Interface

	////

	// Basic types
	fmt.Println("\n--- GO VERİ TİPLERİ ---")

	// Integer types
	var intVar int = 42                      // platform dependent size (32 or 64 bit)
	var int8Var int8 = 127                   // -128 to 127
	var int16Var int16 = 32767               // -32768 to 32767
	var int32Var int32 = 2147483647          // -2147483648 to 2147483647
	var int64Var int64 = 9223372036854775807 // -9223372036854775808 to 9223372036854775807

	fmt.Println("\nTamsayı tipleri:")
	fmt.Printf("int: %v, boyutu: %d bit\n", intVar, 32<<(^uint(0)>>63))
	fmt.Printf("int8: %v, boyutu: 8 bit\n", int8Var)
	fmt.Printf("int16: %v, boyutu: 16 bit\n", int16Var)
	fmt.Printf("int32: %v, boyutu: 32 bit\n", int32Var)
	fmt.Printf("int64: %v, boyutu: 64 bit\n", int64Var)

	// Unsigned integer types
	var uintVar uint = 42
	var uint8Var uint8 = 255
	var uint16Var uint16 = 65535
	var uint32Var uint32 = 4294967295
	var uint64Var uint64 = 18446744073709551615

	fmt.Println("\nİşaretsiz tamsayı tipleri:")
	fmt.Printf("uint: %v, boyutu: %d bit\n", uintVar, 32<<(^uint(0)>>63))
	fmt.Printf("uint8: %v, boyutu: 8 bit\n", uint8Var)
	fmt.Printf("uint16: %v, boyutu: 16 bit\n", uint16Var)
	fmt.Printf("uint32: %v, boyutu: 32 bit\n", uint32Var)
	fmt.Printf("uint64: %v, boyutu: 64 bit\n", uint64Var)

	// Float types
	var float32Var float32 = 3.14159
	var float64Var float64 = 3.14159265358979323846

	fmt.Println("\nOndalık sayı tipleri:")
	fmt.Printf("float32: %v, boyutu: 32 bit\n", float32Var)
	fmt.Printf("float64: %v, boyutu: 64 bit\n", float64Var)

	// Complex types
	var complex64Var complex64 = 1 + 2i
	var complex128Var complex128 = 1.1 + 2.2i

	fmt.Println("\nKarmaşık sayı tipleri:")
	fmt.Printf("complex64: %v, boyutu: 64 bit\n", complex64Var)
	fmt.Printf("complex128: %v, boyutu: 128 bit\n", complex128Var)

	// Other important types
	var byteVar byte = 'A' // uint8'in alias'ı
	var runeVar rune = 'Ğ' // int32'nin alias'ı, Unicode kod noktalarını temsil eder

	fmt.Println("\nDiğer tipler:")
	fmt.Printf("byte: %v (%c), boyutu: 8 bit\n", byteVar, byteVar)
	fmt.Printf("rune: %v (%c), boyutu: 32 bit\n", runeVar, runeVar)

	// Type conversions
	fmt.Println("\nTip dönüşümleri:")
	var a int = 42
	var b float64 = float64(a)
	var c string = fmt.Sprintf("%d", a)

	fmt.Printf("int -> float64: %v -> %v\n", a, b)
	fmt.Printf("int -> string: %v -> %v\n", a, c)

	// Zero values
	var zeroInt int
	var zeroFloat float64
	var zeroString string
	var zeroBool bool

	fmt.Println("\nVarsayılan değerler (Zero values):")
	fmt.Printf("int: %v, float64: %v, string: %v, bool: %v\n", zeroInt, zeroFloat, zeroString, zeroBool)

	////
	////
	////

	/* Formatting Reference for fmt.Printf


	- Common specifiers
	- %v - Default format for the value
	- %T - Type of the value
	- %% - Literal percent sign

	- Integer formats
	- %d - Decimal integer
	- %b - Binary representation
	- %o - Octal representation
	- %x - Hexadecimal (lowercase)
	- %X - Hexadecimal (uppercase)

	- Float formats
	- %f - Decimal point, no exponent
	- %.2f - Decimal point with precision (2 decimal places)
	- %e - Scientific notation

	- String formats
	- %s - Plain string
	- %q - Quoted string

	- Boolean format
	- %t - true or false

	- Character/rune format
	- %c - Character

	- Width and alignment
	- %10s - Right-align with width 10
	- %-10s - Left-align with width 10
	- %010d - Zero padding with width 10

	- Pointer
	- %p - Pointer address

	- Special width/precision
	- %*d - Width specified as an argument
	- %.*f - Precision specified as an argument
	*/

	// Examples:
	x := 15
	s := "Go"
	fmt.Printf("Integer: %d, String: %s\n", x, s)
}
