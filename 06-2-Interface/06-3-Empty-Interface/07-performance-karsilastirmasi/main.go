/*
Author : mefamex
Date   : 2025-07-06
Title  : Empty Interface - Performance Karşılaştırması



PERFORMANCE KARŞILAŞTIRMASI - DETAYLI AÇIKLAMA
==============================================
Bu dosyada, Go dilinde concrete tipler (ör: int, string) ile interface{} (empty interface) kullanımı arasındaki performans ve bellek farkları gerçek kod örnekleriyle karşılaştırmalı olarak gösterilmektedir.

Kapsanan Testler ve Amaçları:
- Fonksiyon Çağrısı: Concrete tip ile yazılmış fonksiyonlar, interface{} kullanan fonksiyonlara göre çok daha hızlıdır. Çünkü interface{} ile her çağrıda tip kontrolü ve type assertion yapılır.
- Bellek Kullanımı: []int gibi concrete slice'lar, []interface{} slice'larına göre daha az bellek harcar. Çünkü interface{} her eleman için tip bilgisi de saklar.
- Type Assertion Overhead: interface{}'den gerçek tipe dönmek (type assertion) ek CPU maliyeti getirir. Özellikle döngü içinde bu maliyet çarpan etkisiyle büyür.
- Slice İşlemleri: append ve erişim işlemleri concrete tiplerde daha hızlıdır. interface{} ile yapılan işlemler hem yavaşlar hem de daha fazla bellek kullanır.

Sonuçlar ve Önemli Noktalar:
- Empty interface ile yazılan kod, concrete tiplere göre 2-10 kat arası daha yavaş çalışabilir.
- []interface{} slice'ları, []int gibi concrete slice'lara göre yaklaşık 2 kat daha fazla bellek kullanır (her eleman için tip bilgisi + değer saklanır).
- Type assertion işlemleri, özellikle döngü içinde kullanıldığında ciddi performans kaybına yol açar.
- Performans kritik (hot path) kodlarda interface{} ve reflection'dan kaçının; mümkünse doğrudan concrete tipler veya Go 1.18+ ile gelen generics kullanın.
- Reflection ve interface{} kullanımı, kodun okunabilirliğini ve bakımını da zorlaştırır.
- Generics ile hem tip güvenliği hem de performans avantajı elde edilir.

Ek Notlar:
- Benchmark sonuçları donanım, Go sürümü ve test ortamına göre değişebilir. Kendi kodunuzda mutlaka benchmark yapın.
- Esneklik gerekiyorsa, Go 1.18+ ile gelen generics'i tercih edin. interface{} sadece gerçekten çok biçimli (polymorphic) kullanım gerektiğinde kullanılmalı.

*/

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("\n=======================================")
	fmt.Println("    PERFORMANCE KARŞILAŞTIRMASI")
	fmt.Println("=======================================")

	fmt.Println("Performance karşılaştırması:")
	fmt.Println("Empty interface vs Concrete types")

	// Test verileri
	const iterations = 1000000

	// Concrete type test
	fmt.Println("\n1. Concrete Type Performance:")
	start := time.Now()
	concreteSum := 0
	for i := 0; i < iterations; i++ {
		concreteSum += addConcrete(i, i+1)
	}
	concreteTime := time.Since(start)
	fmt.Printf("Concrete type sum: %d\n", concreteSum)
	fmt.Printf("Concrete type time: %v\n", concreteTime)

	// Empty interface test
	fmt.Println("\n2. Empty Interface Performance:")
	start = time.Now()
	interfaceSum := 0
	for i := 0; i < iterations; i++ {
		result := addInterface(i, i+1)
		if sum, ok := result.(int); ok {
			interfaceSum += sum
		}
	}
	interfaceTime := time.Since(start)
	fmt.Printf("Interface sum: %d\n", interfaceSum)
	fmt.Printf("Interface time: %v\n", interfaceTime)

	// Karşılaştırma
	fmt.Printf("\nPerformance Farkı:\n")
	ratio := float64(interfaceTime) / float64(concreteTime)
	fmt.Printf("Interface/Concrete ratio: %.2fx slower\n", ratio)

	// Memory overhead testi
	fmt.Println("\n3. Memory Overhead:")
	concreteMemoryTest()
	interfaceMemoryTest()

	// Type assertion overhead
	fmt.Println("\n4. Type Assertion Overhead:")
	typeAssertionOverhead()

	// Slice operations comparison
	fmt.Println("\n5. Slice Operations Comparison:")
	sliceOperationsComparison()

	fmt.Println("\nSonuçlar:")
	fmt.Println("- Empty interface ~2-10x daha yavaş")
	fmt.Println("- Memory overhead: 2x daha fazla")
	fmt.Println("- Type assertion overhead var")
	fmt.Println("- Kritik performans gereken yerlerde kullanmayın")
}

// Concrete type fonksiyonu
func addConcrete(a, b int) int {
	return a + b
}

// Empty interface fonksiyonu
func addInterface(a, b interface{}) interface{} {
	if numA, ok := a.(int); ok {
		if numB, ok := b.(int); ok {
			return numA + numB
		}
	}
	return 0
}

// Memory test fonksiyonları
func concreteMemoryTest() {
	const size = 100000
	data := make([]int, size)

	start := time.Now()
	for i := 0; i < size; i++ {
		data[i] = i
	}
	duration := time.Since(start)

	fmt.Printf("Concrete slice creation: %v\n", duration)
	fmt.Printf("Memory per int: 8 bytes\n")
	fmt.Printf("Total memory: %d bytes\n", size*8)
}

func interfaceMemoryTest() {
	const size = 100000
	data := make([]interface{}, size)

	start := time.Now()
	for i := 0; i < size; i++ {
		data[i] = i
	}
	duration := time.Since(start)

	fmt.Printf("Interface slice creation: %v\n", duration)
	fmt.Printf("Memory per interface{}: 16 bytes (type info + value)\n")
	fmt.Printf("Total memory: %d bytes\n", size*16)
}

// Type assertion overhead testi
func typeAssertionOverhead() {
	const iterations = 1000000
	var val interface{} = 42

	// Direkt erişim (benchmark için - gerçekte mümkün değil)
	start := time.Now()
	sum1 := 0
	for i := 0; i < iterations; i++ {
		// Bu sadece karşılaştırma için - interface{}'den direkt erişim mümkün değil
		sum1 += 42
	}
	directTime := time.Since(start)

	// Type assertion ile erişim
	start = time.Now()
	sum2 := 0
	for i := 0; i < iterations; i++ {
		if num, ok := val.(int); ok {
			sum2 += num
		}
	}
	assertionTime := time.Since(start)

	// Unsafe type assertion (panic riski)
	start = time.Now()
	sum3 := 0
	for i := 0; i < iterations; i++ {
		sum3 += val.(int) // Panic riski var!
	}
	unsafeTime := time.Since(start)

	fmt.Printf("Direct access time: %v\n", directTime)
	fmt.Printf("Safe type assertion time: %v\n", assertionTime)
	fmt.Printf("Unsafe type assertion time: %v\n", unsafeTime)

	fmt.Printf("Safe assertion overhead: %.2fx\n", float64(assertionTime)/float64(directTime))
	fmt.Printf("Unsafe assertion overhead: %.2fx\n", float64(unsafeTime)/float64(directTime))
}

// Slice operations karşılaştırması
func sliceOperationsComparison() {
	const size = 100000

	// Concrete slice append
	start := time.Now()
	concreteSlice := make([]int, 0, size)
	for i := 0; i < size; i++ {
		concreteSlice = append(concreteSlice, i)
	}
	concreteAppendTime := time.Since(start)

	// Interface slice append
	start = time.Now()
	interfaceSlice := make([]interface{}, 0, size)
	for i := 0; i < size; i++ {
		interfaceSlice = append(interfaceSlice, i)
	}
	interfaceAppendTime := time.Since(start)

	// Concrete slice access
	start = time.Now()
	concreteSum := 0
	for i := 0; i < len(concreteSlice); i++ {
		concreteSum += concreteSlice[i]
	}
	concreteAccessTime := time.Since(start)

	// Interface slice access (with type assertion)
	start = time.Now()
	interfaceSum := 0
	for i := 0; i < len(interfaceSlice); i++ {
		if num, ok := interfaceSlice[i].(int); ok {
			interfaceSum += num
		}
	}
	interfaceAccessTime := time.Since(start)

	fmt.Printf("Concrete slice append: %v\n", concreteAppendTime)
	fmt.Printf("Interface slice append: %v\n", interfaceAppendTime)
	fmt.Printf("Append ratio: %.2fx slower\n", float64(interfaceAppendTime)/float64(concreteAppendTime))

	fmt.Printf("Concrete slice access: %v\n", concreteAccessTime)
	fmt.Printf("Interface slice access: %v\n", interfaceAccessTime)
	fmt.Printf("Access ratio: %.2fx slower\n", float64(interfaceAccessTime)/float64(concreteAccessTime))

	// Memory usage comparison
	fmt.Printf("\nMemory usage:\n")
	fmt.Printf("Concrete slice: %d bytes\n", len(concreteSlice)*8)
	fmt.Printf("Interface slice: %d bytes\n", len(interfaceSlice)*16)
	fmt.Printf("Memory ratio: %.2fx more\n", float64(len(interfaceSlice)*16)/float64(len(concreteSlice)*8))
}
