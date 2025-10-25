/*
Author : mefamex
Date   : 2025-07-06
Title  : Empty Interface - Farklı Tipleri Saklama



FARKLI TİPLERİ SAKLAMA - DETAYLI AÇIKLAMA
==========================================

Empty interface (interface{}) kullanarak farklı tipleri aynı veri yapısında
saklama tekniklerini öğrenin. Bu bölümde şunları göreceğiz:

1. MIXED TYPE SLICE (Karışık Tip Dizileri)
    - Tek bir slice'da farklı tipler (int, string, float, bool, struct, map)
    - []interface{} kullanarak heterojen veri koleksiyonları
    - Type information'ın runtime'da korunması

2. MIXED TYPE MAP (Karışık Tip Haritaları)
    - map[string]interface{} ile anahtar-değer çiftleri
    - JSON benzeri veri yapıları oluşturma
    - Dinamik configuration ve settings yönetimi

3. DİNAMİK VERİ KOLEKSİYONLARI
    - API response'ları için esnek veri yapıları
    - Nested data structures (iç içe veri yapıları)
    - Real-world kullanım senaryoları

KULLANIM ALANLARI:
- JSON/XML parsing - bilinmeyen yapıdaki veriler
- Configuration files - farklı tip değerler
- Database query results - mixed column types
- API responses - flexible data structures
- Event data - varying payload types
- Plugin systems - arbitrary data passing

AVANTAJLAR:
- Esneklik - herhangi bir tip saklanabilir
- Runtime type information korunur
- JSON marshaling/unmarshaling için ideal
- Generic veri yapıları oluşturma

DİKKAT EDİLECEKLER:
- Type safety kaybı - compile time kontrolü yok
- Performance overhead - boxing/unboxing
- Runtime errors riski - tip kontrolü gerekli
- Memory overhead - type metadata

*/

package main

import "fmt"

// Person struct - farklı tip örnekleri için
type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("\n=======================================")
	fmt.Println("    FARKLI TİPLERİ SAKLAMA")
	fmt.Println("=======================================")

	// Mixed type slice
	mixedSlice := []interface{}{
		42,
		"string value",
		3.14159,
		true,
		Person{Name: "Ahmet", Age: 25},
		[]int{1, 2, 3},
		map[string]int{"a": 1, "b": 2},
	}

	fmt.Println("Mixed type slice:")
	for i, item := range mixedSlice {
		fmt.Printf("Index %d: %v (Type: %T)\n", i, item, item)
	}

	// Mixed type map
	mixedMap := map[string]interface{}{
		"integer": 100,
		"string":  "text",
		"float":   2.718,
		"boolean": false,
		"person":  Person{Name: "Zeynep", Age: 28},
		"slice":   []string{"a", "b", "c"},
	}

	fmt.Println("\nMixed type map:")
	for key, value := range mixedMap {
		fmt.Printf("Key: %s, Value: %v (Type: %T)\n", key, value, value)
	}

	// Örnek: Dinamik veri koleksiyonu
	fmt.Println("\nDinamik veri koleksiyonu:")
	dynamicData := []interface{}{
		"Kullanıcı verisi",
		map[string]interface{}{
			"id":    1001,
			"email": "user@example.com",
			"roles": []string{"admin", "user"},
		},
		[]interface{}{"tag1", "tag2", 123},
		Person{Name: "Dinamik Kullanıcı", Age: 35},
	}

	for i, data := range dynamicData {
		fmt.Printf("Data[%d]: %v (%T)\n", i, data, data)
	}
}
