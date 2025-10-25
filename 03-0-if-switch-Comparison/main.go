/*
author: mefamex
date  : 2025-07-02
title : if, switch, comparison operators
*/

package main

import "fmt"

func main() {
	// Sabit değerler
	const (
		minValue = 5
		maxValue = 10
		midValue = 15
	)

	////
	////

	fmt.Println("\n--- Karşılaştırma Operatörleri ---")
	// COMPARISON OPERATORS
	// == Equal to      : Eşit mi?
	// != Not equal to  : Eşit değil mi?
	// <  Less than     : Küçük mü?
	// >  Greater than  : Büyük mü?
	// <= Less than or equal to    : Küçük veya eşit mi?
	// >= Greater than or equal to : Büyük veya eşit mi?

	// LOGICAL OPERATORS
	// && AND : Ve (her iki koşul da doğru olmalı)
	// || OR  : Veya (koşullardan biri doğru olmalı)
	// !  NOT : Değil (koşulun tersini alır)

	a, b := minValue, maxValue

	////
	////

	// Tüm karşılaştırma operatörlerinin kullanımı
	fmt.Printf("a = %v, b = %v\n", a, b) // Değerleri yazdır
	fmt.Printf("a == b: %v\n", a == b)   // Eşit mi?
	fmt.Printf("a != b: %v\n", a != b)   // Eşit değil mi?
	fmt.Printf("a < b : %v\n", a < b)    // Küçük mü?
	fmt.Printf("a > b : %v\n", a > b)    // Büyük mü?
	fmt.Printf("a <= b: %v\n", a <= b)   // Küçük veya eşit mi?
	fmt.Printf("a >= b: %v\n", a >= b)   // Büyük veya eşit mi?

	////
	////

	fmt.Println("\n--- If-Else Örnekleri ---")
	// Basit if-else kullanımı
	if a < b {
		fmt.Printf("%v küçüktür %v\n", a, b)
	} else if a == b {  // sınırsız else if bloğu yapılabilir
		fmt.Printf("%v eşittir %v\n", a, b)
	} else {
		fmt.Printf("%v büyüktür %v\n", a, b)
	}

	////
	////

	// If ile değişken tanımlama (scope örneği)
	if x := "fine"; x == "fine" {
		fmt.Println("x değeri 'fine'")
	} else {
		fmt.Println("x değeri 'fine' değil")
	}

	////
	////

	fmt.Println("\n--- Switch Örnekleri ---")
	// 1. Koşullu switch örneği
	switch c := midValue; {
	case c < minValue:
		fmt.Printf("%v, %v'den küçük\n", c, minValue)
	case c >= minValue && c < maxValue:
		fmt.Printf("%v, %v ile %v arasında\n", c, minValue, maxValue)
	case c >= maxValue:
		fmt.Printf("%v, %v veya daha büyük\n", c, maxValue)
		fallthrough // fallthrough sonraki case'i de çalıştırır
	default:
		fmt.Println("Varsayılan durum")
	}

	////
	////

	// 2. Değer karşılaştırmalı switch
	switch day := "Pazartesi"; day {
	case "Pazartesi", "Salı", "Çarşamba", "Perşembe", "Cuma":
		fmt.Println("Hafta içi")
	case "Cumartesi", "Pazar":
		fmt.Println("Hafta sonu")
	default:
		fmt.Println("Geçersiz gün")
	}  // pythonda da in [pazartesi, salı....] vardı yani.. çok bir şey değil

	////
	////

	// 3. Type switch örneği
	var i interface{} = 42
	switch v := i.(type) {
	case int:
		fmt.Printf("Integer: %v\n", v)
	case string:
		fmt.Printf("String: %v\n", v)
	case bool:
		fmt.Printf("Boolean: %v\n", v)
	default:
		fmt.Printf("Bilinmeyen tip: %T\n", v)
	}
}
