/*
Author : mefamex
Date   : 2025-07-06
Title  : Empty Interface - Type Switch



TYPE SWITCH - DETAYLI AÇIKLAMA
===============================

Type switch, birden fazla tipi aynı anda kontrol etmek için kullanılan
Go'nun güçlü özelliklerinden biridir. Type assertion'ın gelişmiş versiyonudur.

1. TEMEL TYPE SWITCH SYNTAX
    switch v := value.(type) {
    case string:
        - v string tipinde
    case int:
        - v int tipinde
    default:
        bilinmeyen tip
    }

2. GELİŞMİŞ TYPE SWITCH
    - Multiple case'ler: case int, int64:
    - Nil kontrolü: case nil:
    - Conditional logic içinde tip kontrolü
    - Nested type switch'ler

3. İÇ İÇE VERİ YAPILARI
    - Slice içindeki farklı tipler
    - Map değerlerinde tip kontrolü
    - JSON-like veri yapıları işleme

AVANTAJLAR:
- Type assertion'dan daha okunabilir
- Birden fazla tipi tek blokta kontrol
- Automatic type conversion
- Default case ile beklenmeyen tiplerı handle etme
- Performance optimizasyonu (compiler seviyesinde)

KULLANIM ALANLARI:
- JSON/XML parsing - mixed data types
- API response processing
- Configuration file parsing
- Event handling systems
- Protocol buffer processing
- Plugin architecture implementations

TYPE SWITCH vs TYPE ASSERTION:
┌─────────────────┬─────────────────┬─────────────────┐
│ Özellik         │ Type Switch     │ Type Assertion  │
├─────────────────┼─────────────────┼─────────────────┤
│ Çoklu tip       │    Mükemmel     │    Zor          │
│ Okunabilirlik   │    Yüksek       │    Orta         │
│ Performance     │    İyi          │    İyi          │
│ Error handling  │    Kolay        │    Manuel       │
│ Tek tip kontrolü│    Fazla        │    İdeal        │
└─────────────────┴─────────────────┴─────────────────┘

PATTERN'LER:
1. Basic Pattern: Tek tip kontrolü
2. Multi-type Pattern: Benzer tiplerı grupla
3. Recursive Pattern: İç içe veri yapıları
4. Validator Pattern: Veri doğrulama
5. Processor Pattern: Tip bazlı işleme

BEST PRACTICES:
- Default case her zaman ekleyin
- Case'leri mantıksal sırada düzenleyin
- Complex logic'i ayrı fonksiyonlara taşıyın
- Type-specific error handling yapın

*/

package main

import "fmt"

// Person struct - type switch örnekleri için
type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("\n=======================================")
	fmt.Println("    TYPE SWITCH")
	fmt.Println("=======================================")

	values := []interface{}{
		"Kısa",
		"Bu çok uzun bir string metindir",
		42,
		-15,
		0,
		Person{Name: "Ali", Age: 25},
		Person{Name: "Ayşe", Age: 16},
		[]interface{}{1, 2, 3},
		[]int{10, 20, 30},
		[]string{"go", "is", "awesome"},
		map[string]interface{}{"key": "value"},
		nil,
		3.14159,
	}

	fmt.Println("Temel type switch:")
	for _, value := range values {
		processWithTypeSwitch(value)
	}

	fmt.Println("\nGelişmiş type switch:")
	for _, value := range values {
		advancedTypeSwitch(value)
	}

	fmt.Println("\nİç içe type switch:")
	complexData := []interface{}{
		map[string]interface{}{
			"name":   "John",
			"age":    30,
			"active": true,
		},
		[]interface{}{"hello", 42, true},
		"simple string",
	}

	for _, data := range complexData {
		nestedTypeSwitch(data)
		fmt.Println("---")
	}

	fmt.Println("\nType switch avantajları:")
	fmt.Println("1. Birden fazla tipi tek blokta kontrol edebilir")
	fmt.Println("2. Type assertion'dan daha okunabilir")
	fmt.Println("3. Performans açısından daha iyi")
	fmt.Println("4. Default case ile beklenmeyen tiplerı handle edebilir")
}

// ============================================================
// TYPE SWITCH
// ============================================================

// processWithTypeSwitch - Type switch ile tip kontrolü
func processWithTypeSwitch(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Printf("String: '%s' (uzunluk: %d)\n", v, len(v))
	case int:
		fmt.Printf("Integer: %d (çift mi: %t)\n", v, v%2 == 0)
	case float64:
		fmt.Printf("Float: %.2f (yuvarlama: %.0f)\n", v, v)
	case bool:
		fmt.Printf("Boolean: %t (tersi: %t)\n", v, !v)
	case Person:
		fmt.Printf("Person: %s (%d yaşında)\n", v.Name, v.Age)
	case []interface{}:
		fmt.Printf("Slice: %d eleman\n", len(v))
	case map[string]interface{}:
		fmt.Printf("Map: %d key\n", len(v))
	case nil:
		fmt.Println("Nil değer")
	default:
		fmt.Printf("Bilinmeyen tip: %T, değer: %v\n", v, v)
	}
}

// advancedTypeSwitch - Gelişmiş type switch
func advancedTypeSwitch(value interface{}) {
	switch v := value.(type) {
	case string:
		if len(v) > 10 {
			fmt.Printf("Uzun string: %s...\n", v[:10])
		} else {
			fmt.Printf("Kısa string: %s\n", v)
		}
	case int:
		if v > 0 {
			fmt.Printf("Pozitif integer: %d\n", v)
		} else if v < 0 {
			fmt.Printf("Negatif integer: %d\n", v)
		} else {
			fmt.Println("Sıfır")
		}
	case Person:
		if v.Age >= 18 {
			fmt.Printf("Yetişkin: %s (%d)\n", v.Name, v.Age)
		} else {
			fmt.Printf("Çocuk: %s (%d)\n", v.Name, v.Age)
		}
	case []int:
		fmt.Printf("Int slice: %v (toplam: %d)\n", v, sum(v))
	case []string:
		fmt.Printf("String slice: %v (birleşik: %s)\n", v, joinStrings(v))
	default:
		fmt.Printf("İşlenmeyen tip: %T\n", v)
	}
}

// nestedTypeSwitch - İç içe type switch
func nestedTypeSwitch(value interface{}) {
	switch outer := value.(type) {
	case map[string]interface{}:
		fmt.Println("Map içeriği:")
		for key, val := range outer {
			fmt.Printf("  Key: %s, ", key)
			switch inner := val.(type) {
			case string:
				fmt.Printf("String Value: %s\n", inner)
			case int:
				fmt.Printf("Int Value: %d\n", inner)
			case bool:
				fmt.Printf("Bool Value: %t\n", inner)
			default:
				fmt.Printf("Other Value: %v (%T)\n", inner, inner)
			}
		}
	case []interface{}:
		fmt.Println("Slice içeriği:")
		for i, item := range outer {
			fmt.Printf("  [%d]: ", i)
			switch v := item.(type) {
			case string:
				fmt.Printf("String: %s\n", v)
			case int:
				fmt.Printf("Int: %d\n", v)
			default:
				fmt.Printf("Other: %v\n", v)
			}
		}
	default:
		fmt.Printf("Basit tip: %T = %v\n", outer, outer)
	}
}

// ============================================================
// HELPER FUNCTIONS
// ============================================================

func sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func joinStrings(strings []string) string {
	result := ""
	for i, str := range strings {
		if i > 0 {
			result += " "
		}
		result += str
	}
	return result
}
