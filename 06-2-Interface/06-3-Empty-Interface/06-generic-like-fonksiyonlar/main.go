/*
Author : mefamex
Date   : 2025-07-06
Title  : Empty Interface - Generic-like Fonksiyonlar


/*
GENERIC-LIKE FONKSİYONLAR (interface{} ile)
------------------------------------------------------------
Bu dosyada, Go'nun generic desteği (Go 1.18 öncesi) olmadığı dönemde
tip bağımsız (generic-like) fonksiyonların nasıl yazılabileceği gösterilmektedir.
Tüm fonksiyonlar, parametre olarak interface{} kullanır ve tip kontrolü/işlemleri
reflection (reflect paketi) ile yapılır. Bu yaklaşım, farklı tiplerle çalışan
fonksiyonlar yazmayı mümkün kılar ancak bazı dezavantajları da beraberinde getirir.

Kapsanan Fonksiyonlar:
- min, max: Farklı tipler (int, float, string) için minimum/maksimum bulma
- contains, indexOf: Slice içinde eleman arama
- reverse: Slice'ı ters çevirme
- filter: Slice'ı koşula göre filtreleme
- mapFunc: Her elemana fonksiyon uygulama
- reduce: Slice'ı bir değere indirgeme
- unique: Tekil elemanları bulma
- groupBy: Slice'ı anahtara göre gruplama

Avantajlar:
- Tek kod ile birden fazla tipte çalışabilme
- Kod tekrarını azaltma (DRY)
- Yeniden kullanılabilirlik

Dezavantajlar:
- Performans kaybı (yaklaşık 10 kat yavaş)
- Derleme zamanı tip güvenliği yok, runtime hataları
- Kod karmaşıklığı ve okunabilirlikte azalma

Not: Go 1.18 ve sonrası için, bu tarz fonksiyonlar yerine
generics (type parameters) kullanılması önerilir.

*/

package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("\n=======================================")
	fmt.Println("    GENERIC-LIKE FONKSİYONLAR")
	fmt.Println("=======================================")

	// Min/Max fonksiyonu test
	fmt.Println("Generic-like min/max fonksiyonları:")
	fmt.Printf("min(10, 5): %v\n", min(10, 5))
	fmt.Printf("max(10, 5): %v\n", max(10, 5))
	fmt.Printf("min(3.14, 2.71): %v\n", min(3.14, 2.71))
	fmt.Printf("max(3.14, 2.71): %v\n", max(3.14, 2.71))
	fmt.Printf("min(\"zebra\", \"apple\"): %v\n", min("zebra", "apple"))
	fmt.Printf("max(\"zebra\", \"apple\"): %v\n", max("zebra", "apple"))

	// Contains ve indexOf fonksiyonu test
	fmt.Println("\nGeneric-like contains/indexOf fonksiyonları:")
	intSlice := []int{1, 2, 3, 4, 5}
	stringSlice := []string{"go", "python", "java", "rust"}

	fmt.Printf("Contains(intSlice, 3): %t\n", contains(intSlice, 3))
	fmt.Printf("Contains(intSlice, 10): %t\n", contains(intSlice, 10))
	fmt.Printf("IndexOf(stringSlice, \"java\"): %d\n", indexOf(stringSlice, "java"))
	fmt.Printf("IndexOf(stringSlice, \"c++\"): %d\n", indexOf(stringSlice, "c++"))

	// Reverse fonksiyonu test
	fmt.Println("\nGeneric-like reverse fonksiyonu:")
	fmt.Printf("Original intSlice: %v\n", intSlice)
	fmt.Printf("Reversed intSlice: %v\n", reverse(intSlice))
	fmt.Printf("Original stringSlice: %v\n", stringSlice)
	fmt.Printf("Reversed stringSlice: %v\n", reverse(stringSlice))

	// Filter fonksiyonu test
	fmt.Println("\nGeneric-like filter fonksiyonu:")
	evenNumbers := filter(intSlice, func(item interface{}) bool {
		if num, ok := item.(int); ok {
			return num%2 == 0
		}
		return false
	})
	fmt.Printf("Even numbers: %v\n", evenNumbers)

	longStrings := filter(stringSlice, func(item interface{}) bool {
		if str, ok := item.(string); ok {
			return len(str) > 3
		}
		return false
	})
	fmt.Printf("Long strings: %v\n", longStrings)

	// Map fonksiyonu test
	fmt.Println("\nGeneric-like map fonksiyonu:")
	squares := mapFunc(intSlice, func(item interface{}) interface{} {
		if num, ok := item.(int); ok {
			return num * num
		}
		return item
	})
	fmt.Printf("Squares: %v\n", squares)

	upperStrings := mapFunc(stringSlice, func(item interface{}) interface{} {
		if str, ok := item.(string); ok {
			return fmt.Sprintf("%s!", str)
		}
		return item
	})
	fmt.Printf("Upper strings: %v\n", upperStrings)

	// Reduce fonksiyonu test
	fmt.Println("\nGeneric-like reduce fonksiyonu:")
	sum := reduce(intSlice, 0, func(acc, item interface{}) interface{} {
		if accNum, ok := acc.(int); ok {
			if num, ok := item.(int); ok {
				return accNum + num
			}
		}
		return acc
	})
	fmt.Printf("Sum: %v\n", sum)

	concat := reduce(stringSlice, "", func(acc, item interface{}) interface{} {
		if accStr, ok := acc.(string); ok {
			if str, ok := item.(string); ok {
				if accStr == "" {
					return str
				}
				return accStr + ", " + str
			}
		}
		return acc
	})
	fmt.Printf("Concatenated: %v\n", concat)

	// Unique fonksiyonu test
	fmt.Println("\nGeneric-like unique fonksiyonu:")
	duplicateSlice := []int{1, 2, 2, 3, 3, 3, 4, 5, 5}
	fmt.Printf("Original: %v\n", duplicateSlice)
	fmt.Printf("Unique: %v\n", unique(duplicateSlice))

	// GroupBy fonksiyonu test
	fmt.Println("\nGeneric-like groupBy fonksiyonu:")
	words := []string{"apple", "ant", "banana", "bear", "cat", "car"}
	grouped := groupBy(words, func(item interface{}) interface{} {
		if str, ok := item.(string); ok && len(str) > 0 {
			return string(str[0]) // İlk karakter
		}
		return ""
	})

	for key, values := range grouped {
		fmt.Printf("'%v': %v\n", key, values)
	}

	fmt.Println("\nGeneric-like fonksiyonların avantajları:")
	fmt.Println("1. Tek kod birden fazla tip")
	fmt.Println("2. Code reusability")
	fmt.Println("3. DRY principle")

	fmt.Println("\nGeneric-like fonksiyonların dezavantajları:")
	fmt.Println("1. Performance overhead (~10x slower)")
	fmt.Println("2. Runtime errors")
	fmt.Println("3. Type safety kaybı")
	fmt.Println("4. Code complexity")

	fmt.Println("\nGo 1.18+ Generics önerisi:")
	fmt.Println("Modern Go projelerinde interface{} yerine generics kullanın!")
}

// ============================================================
// GENERIC-LIKE FONKSİYONLAR
// ============================================================

// min - Generic-like minimum bulma fonksiyonu
func min(a, b interface{}) interface{} {
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)

	switch va.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if va.Int() < vb.Int() {
			return a
		}
		return b
	case reflect.Float32, reflect.Float64:
		if va.Float() < vb.Float() {
			return a
		}
		return b
	case reflect.String:
		if va.String() < vb.String() {
			return a
		}
		return b
	default:
		return nil
	}
}

// max - Generic-like maximum bulma fonksiyonu
func max(a, b interface{}) interface{} {
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)

	switch va.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if va.Int() > vb.Int() {
			return a
		}
		return b
	case reflect.Float32, reflect.Float64:
		if va.Float() > vb.Float() {
			return a
		}
		return b
	case reflect.String:
		if va.String() > vb.String() {
			return a
		}
		return b
	default:
		return nil
	}
}

// contains - Generic-like slice içerik kontrolü
func contains(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return false
	}

	for i := 0; i < s.Len(); i++ {
		if reflect.DeepEqual(s.Index(i).Interface(), item) {
			return true
		}
	}
	return false
}

// indexOf - Generic-like index bulma
func indexOf(slice interface{}, item interface{}) int {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return -1
	}

	for i := 0; i < s.Len(); i++ {
		if reflect.DeepEqual(s.Index(i).Interface(), item) {
			return i
		}
	}
	return -1
}

// reverse - Generic-like slice ters çevirme
func reverse(slice interface{}) interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return nil
	}

	length := s.Len()
	reversed := reflect.MakeSlice(s.Type(), length, length)

	for i := 0; i < length; i++ {
		reversed.Index(i).Set(s.Index(length - 1 - i))
	}

	return reversed.Interface()
}

// filter - Generic-like filtering
func filter(slice interface{}, predicate func(interface{}) bool) interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return nil
	}

	var filtered []reflect.Value
	for i := 0; i < s.Len(); i++ {
		item := s.Index(i)
		if predicate(item.Interface()) {
			filtered = append(filtered, item)
		}
	}

	result := reflect.MakeSlice(s.Type(), len(filtered), len(filtered))
	for i, item := range filtered {
		result.Index(i).Set(item)
	}

	return result.Interface()
}

// mapFunc - Generic-like mapping
func mapFunc(slice interface{}, mapper func(interface{}) interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return nil
	}

	result := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		result[i] = mapper(s.Index(i).Interface())
	}

	return result
}

// reduce - Generic-like reducing
func reduce(slice interface{}, accumulator interface{}, reducer func(interface{}, interface{}) interface{}) interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return accumulator
	}

	acc := accumulator
	for i := 0; i < s.Len(); i++ {
		acc = reducer(acc, s.Index(i).Interface())
	}

	return acc
}

// unique - Generic-like unique elements
func unique(slice interface{}) interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return nil
	}

	seen := make(map[interface{}]bool)
	var uniqueItems []reflect.Value

	for i := 0; i < s.Len(); i++ {
		item := s.Index(i)
		key := item.Interface()
		if !seen[key] {
			seen[key] = true
			uniqueItems = append(uniqueItems, item)
		}
	}

	result := reflect.MakeSlice(s.Type(), len(uniqueItems), len(uniqueItems))
	for i, item := range uniqueItems {
		result.Index(i).Set(item)
	}

	return result.Interface()
}

// groupBy - Generic-like grouping
func groupBy(slice interface{}, keyFunc func(interface{}) interface{}) map[interface{}][]interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return nil
	}

	groups := make(map[interface{}][]interface{})
	for i := 0; i < s.Len(); i++ {
		item := s.Index(i).Interface()
		key := keyFunc(item)
		groups[key] = append(groups[key], item)
	}

	return groups
}
