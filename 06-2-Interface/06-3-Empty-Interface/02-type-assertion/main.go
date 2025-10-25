/*
Author : mefamex
Date   : 2025-07-06
Title  : Empty Interface - Type Assertion

TYPE ASSERTION - DETAYLI AÇIKLAMA
==================================

Type assertion, interface{} içindeki değerin gerçek tipini kontrol etme ve
o tipe dönüştürme işlemidir. Bu bölümde şunları öğreneceksiniz:

 1. TEMEL TYPE ASSERTION
    Syntax: value.(Type)
    - Direkt dönüştürme (panic riski var)
    - Runtime'da tip kontrolü yapma
    - Concrete type'a erişim

 2. GÜVENLİ TYPE ASSERTION (COMMA OK İDIOM)
    Syntax: value, ok := variable.(Type)
    - Panic'ten korunma
    - Boolean dönüş değeri ile kontrol
    - Production code için önerilen yöntem

3. ÇOK TİPLİ KONTROL
  - Birden fazla tip için assertion
  - Error handling best practices
  - Default değer stratejileri

KULLANIM ALANLARI:
- JSON unmarshaling sonrası tip kontrolü
- API response'larında değer çıkarma
- Configuration parsing
- Plugin interface'lerinde tip belirleme
- Generic fonksiyonlarda tip spesializasyonu

SYNTAX ÖRNEKLERİ:
// Tehlikeli (panic riski)
str := value.(string)

// Güvenli (önerilen)
str, ok := value.(string)

	if ok {
	    // string olarak kullan
	}

// Type switch alternatifi
switch v := value.(type) {
case string:

	// string işlemleri

case int:

	    // int işlemleri
	}

BEST PRACTICES:
- Her zaman "comma ok" idiomunu kullanın
- Panic recovery mekanizması ekleyin
- Type assertion yerine type switch tercih edin
- Error handling stratejisi belirleyin

YAYGIN HATALAR:
- Panic riski göz ardı etme
- Type assertion zincirleme
- Performance impact'ini unutma
- Nil değer kontrolü yapmama
*/

package main

import "fmt"

// Person struct - type assertion örnekleri için
type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("\n=======================================")
	fmt.Println("    TYPE ASSERTION")
	fmt.Println("=======================================")

	values := []interface{}{
		"Merhaba Dünya",
		42,
		Person{Name: "Kemal", Age: 35},
		3.14,
		true,
		[]interface{}{1, "test", 3.14},
	}

	fmt.Println("Güvenli type assertion:")
	for _, value := range values {
		safeTypeAssertion(value)
	}

	fmt.Println("\nTehlikeli type assertion:")
	for _, value := range values {
		dangerousTypeAssertion(value)
	}

	fmt.Println("\nÇoklu tip kontrolü:")
	for _, value := range values {
		multipleTypeAssertion(value)
	}

	fmt.Println("\nType assertion best practices:")
	fmt.Println("1. Her zaman 'comma ok' idiomunu kullanın")
	fmt.Println("2. Panic riskini göz önünde bulundurun")
	fmt.Println("3. Type switch'i tercih edin (daha temiz)")
	fmt.Println("4. Sık kullanılan tip kontrollerini fonksiyon haline getirin")
}

// ============================================================
// TYPE ASSERTION
// ============================================================

// safeTypeAssertion - Güvenli tip dönüşümü
func safeTypeAssertion(value interface{}) {
	// String assertion
	if str, ok := value.(string); ok {
		fmt.Printf("String değer: %s (uzunluk: %d)\n", str, len(str))
		return
	}

	// Int assertion
	if num, ok := value.(int); ok {
		fmt.Printf("Integer değer: %d (karesi: %d)\n", num, num*num)
		return
	}

	// Person assertion
	if person, ok := value.(Person); ok {
		fmt.Printf("Person: %s, %d yaşında\n", person.Name, person.Age)
		return
	}

	fmt.Printf("Bilinmeyen tip: %T\n", value)
}

// dangerousTypeAssertion - Tehlikeli tip dönüşümü (panic riski)
func dangerousTypeAssertion(value interface{}) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic yakalandı: %v\n", r)
		}
	}()

	// Bu panic oluşturabilir!
	str := value.(string)
	fmt.Printf("String değer: %s\n", str)
}

// multipleTypeAssertion - Birden fazla tip kontrolü
func multipleTypeAssertion(value interface{}) {
	// Float64 kontrolü
	if f64, ok := value.(float64); ok {
		fmt.Printf("Float64: %.2f\n", f64)
		return
	}

	// Bool kontrolü
	if b, ok := value.(bool); ok {
		fmt.Printf("Boolean: %t\n", b)
		return
	}

	// Slice kontrolü
	if slice, ok := value.([]interface{}); ok {
		fmt.Printf("Slice: %v (uzunluk: %d)\n", slice, len(slice))
		return
	}

	fmt.Printf("Desteklenmeyen tip: %T\n", value)
}
