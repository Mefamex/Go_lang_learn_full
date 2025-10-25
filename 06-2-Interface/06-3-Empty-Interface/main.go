/*
Author : mefamex
Date   : 2025-07-06
Title  : Empty Interface (interface{}) - Go

NOT: Bu dosya sadece Empty Interface'in temel kullanımını gösterir.
Diğer alt konular (type assertion, type switch, JSON, reflection, vs.)
ayrı dosyalarda bulunmaktadır. Her dosya bağımsız çalıştırılabilir.



EMPTY INTERFACE (interface{}) NEDİR? — KAPSAMLI BAKIŞ
- Empty interface (interface{}) hiçbir method imzası içermez.
- Tüm Go tipleri (int, string, struct, vs.) otomatik olarak interface{}'ü implement eder.
- Dinamik tip desteği sağlar; compile-time'da tipi bilinmeyen değerlerle çalışmayı mümkün kılar.
- Type assertion ve type switch ile runtime'da tip kontrolü yapılabilir.
- JSON parsing, generic benzeri fonksiyonlar ve reflection işlemlerinde yaygın olarak kullanılır.
- Go 1.18+ ile "any" anahtar kelimesi interface{}'e eşdeğer olarak eklenmiştir.
- Polimorfizm ve dinamik programlama için güçlü bir araçtır.



SYNTAX:
interface{} veya any (Go 1.18+)



TEMEL KULLANIM ALANLARI:
- Farklı tipte değerleri aynı slice/map'te saklamak.
- Fonksiyonlara farklı tipte parametreler göndermek.
- JSON/XML parsing işlemleri.
- Reflection tabanlı kütüphaneler.
- Generic-like fonksiyonlar (Go 1.18 öncesi).
- Logging ve debug çıktıları.
- Configuration/settings yönetimi.
- Cache sistemleri ve dinamik veri depolama.



ÖZELLİKLERİ:
- Herhangi bir type'ı kabul eder (int, string, struct, pointer, slice, map, channel, vs.).
- Type safety compile-time'da değil, runtime'da sağlanır.
- Type assertion gerektirir (runtime'da tip kontrolü).
- Performance overhead'i vardır (boxing/unboxing, type info).
- Reflection ile birlikte güçlü dinamik programlama sağlar.



TYPE ASSERTION:
- value.(Type) ile tip kontrolü ve dönüştürme yapılır.
- "Comma ok" idiom: value, ok := value.(Type) — panic riskini önler.
- Type switch ile birden fazla tip kolayca kontrol edilebilir.



BEST PRACTICES:
1. Mümkünse concrete type kullanın (performans ve type safety için).
2. Type assertion'da "comma ok" idiomunu kullanarak panic'ten kaçının.
3. Type switch ile birden fazla tip kontrolü yapın.
4. interface{} kullanımını minimize edin; gerekliyse iyi dokümante edin.
5. Go 1.18+ projelerde generics'i tercih edin.
6. Nil değerler için ekstra kontrol yapın (nil interface != nil pointer!).
7. interface{} kabul eden fonksiyonları ve beklenen tipleri mutlaka dokümante edin.



YAYGIN HATALAR:
- Type assertion'da panic oluşması (comma ok kullanılmazsa).
- Performance overhead'ini göz ardı etmek.
- Type safety'yi kaybetmek, runtime hatalarına açık olmak.
- Gereksiz interface{} kullanımı (gereksiz soyutlama).
- Nil pointer dereference ve nil interface karışıklığı.
- Runtime type errors'ı handle etmemek.



PERFORMANCE DİKKATLERİ:
- Concrete types'a göre ~2-3x daha yavaş olabilir.
- Memory overhead: Her value boxing gerektirir (type info + value pointer).
- CPU overhead: Type assertion ve boxing/unboxing işlemleri.
- Cache locality: Interface değerleri indirect memory access gerektirir.



GERÇEK DÜNYA KULLANIMLARI:
- JSON API responses (dinamik veri).
- Configuration parsers.
- Logging ve event sistemleri.
- Plugin mimarileri.
- Data serialization/deserialization.
- Generic veri yapıları ve cache sistemleri.

*/

package main

import (
	"fmt"
)

// Person - Örnek struct tipi
type Person struct {
	Name string
	Age  int
}

// Container - Empty interface kullanarak farklı tipleri depolayan basit container
type Container struct {
	Items []interface{}
}

func main() {
	fmt.Println("=======================================")
	fmt.Println("    EMPTY INTERFACE TEMEL KULLANIM")
	fmt.Println("=======================================")

	fmt.Println("1. Farklı tipte değerleri yazdırma:")
	printAny(42)
	printAny("Merhaba")
	printAny(3.14)
	printAny(true)
	printAny(Person{Name: "Ali", Age: 30})

	fmt.Println("\n2. Fonksiyon dönüşü:")
	result := processValue("Test değeri")
	fmt.Println(result)

	fmt.Println("\n3. Slice kullanımı:")
	items := []interface{}{1, "string", 3.14, true}
	printSlice(items)

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("    CONTAINER ÖRNEĞİ")
	fmt.Println("=======================================")

	container := &Container{}

	fmt.Println("Container'a farklı tipte değerler ekleme:")

	// Farklı tipte değerler ekle
	container.Add(42)
	container.Add("Merhaba")
	container.Add(3.14)
	container.Add(Person{Name: "Ahmet", Age: 25})
	container.Add([]int{1, 2, 3})

	fmt.Printf("\nContainer boyutu: %d\n", container.Size())

	fmt.Println("\nContainer içeriği:")
	// Tüm elemanları yazdır
	for i := 0; i < container.Size(); i++ {
		item := container.Get(i)
		fmt.Printf("  [%d]: %v (%T)\n", i, item, item)
	}

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("        DİĞER KONULAR")
	fmt.Println("=======================================")
	fmt.Println("Diğer Empty Interface konuları ayrı klasörlerde:")
	fmt.Println("- cd 01-farkli-tipleri-saklama && go run main.go")
	fmt.Println("- cd 02-type-assertion && go run main.go")
	fmt.Println("- cd 03-type-switch && go run main.go")
	fmt.Println("- cd 04-json-ile-calisma && go run main.go")
	fmt.Println("- cd 05-reflection-ile-kullanim && go run main.go")
	fmt.Println("- cd 06-generic-like-fonksiyonlar && go run main.go")
	fmt.Println("- cd 07-performance-karsilastirmasi && go run main.go")
	fmt.Println("- cd 08-best-practices && go run main.go")
	fmt.Println("- cd 09-yaygin-hatalar && go run main.go")
	fmt.Println("- cd 10-gercek-dunya-ornekleri && go run main.go")
	fmt.Println("\nHer klasör bağımsız çalıştırılabilir!")
	fmt.Println("=======================================")
}

////
////
////

// ============================================================
// EMPTY INTERFACE TEMEL KULLANIM
// ============================================================

// printAny - Empty interface ile herhangi bir değeri yazdır
func printAny(value interface{}) {
	fmt.Printf("Değer: %v, Tip: %T\n", value, value)
}

// processValue - Empty interface'i parametre olarak al ve işle
func processValue(value interface{}) string {
	return fmt.Sprintf("İşlenen değer: %v (%T)", value, value)
}

// printSlice - Empty interface slice'ını yazdır
func printSlice(items []interface{}) {
	fmt.Println("Slice içeriği:")
	for i, item := range items {
		fmt.Printf("  [%d]: %v (%T)\n", i, item, item)
	}
}

////
////
////

// ============================================================
// CONTAINER ÖRNEĞİ
// ============================================================

func (c *Container) Add(item interface{}) {
	c.Items = append(c.Items, item)
}

func (c *Container) Get(index int) interface{} {
	if index >= 0 && index < len(c.Items) {
		return c.Items[index]
	}
	return nil
}

func (c *Container) Size() int {
	return len(c.Items)
}
