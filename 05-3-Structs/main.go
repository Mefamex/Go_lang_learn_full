/*
author: mefamex
date  : 2025-07-03
title : Structs in Go

STRUCT NEDİR?
- Struct (structure), farklı türdeki verileri bir arada tutan, kullanıcı tanımlı veri tipidir.
- Struct'lar, veri gruplama ve modelleme için kullanılır (ör: bir kişiyi temsil eden Person struct'ı).
- Struct'lar Go'da nesne yönelimli programlamanın temelini oluşturur (ancak Go'da class yoktur).
- Struct'lar değer tipidir; fonksiyona parametre olarak gönderildiğinde kopyalanır.
- Pointer ile gönderilirse orijinal veri üzerinde değişiklik yapılabilir.
- Struct'lar gömülü (embedded) olarak başka struct'larda kullanılabilir.
- Struct'lara method tanımlanabilir (receiver ile); hem value hem pointer receiver olabilir.
- Struct'lar, slice, map gibi diğer yapıları veya başka struct'ları alanlara sahip olabilir.
- Struct'lar anonim (isimsiz) olarak da tanımlanabilir ve genellikle geçici veri yapıları için kullanılır.
- JSON gibi veri formatlarına kolayca dönüştürülebilir; bunun için struct etiketleri (tags) kullanılabilir.

> TYPE / VAR ile STRUCT TANIMLAMA
- type ile tanımlanan struct'lar, bir veri tipini temsil eder.
- Struct'lar, alan isimleri ve türleri ile tanımlanır.
- var ile tanımlanan struct'lar anonim struct olarak adlandırılır ve genellikle geçici veri yapıları için kullanılır.

*/

package main

import "fmt"

// 1. Struct tanımlama
type Person struct { // underlying type: struct
	Name string
	Age  int
	City string
}

func main() {
	fmt.Println("===================================")
	fmt.Println("   STRUCT TANIMLAMA VE KULLANIMI")
	fmt.Println("===================================")

	// 2. Struct değişkeni oluşturma
	var p1 Person
	p1.Name = "Ahmet"
	p1.Age = 30
	p1.City = "İstanbul"
	fmt.Println("1. Struct değişkeni:", p1)

	// 3. Kısa yoldan struct oluşturma
	p2 := Person{"Ayşe", 25, "Ankara"}
	fmt.Println("2. Kısa yoldan struct:", p2)

	// 4. Alan isimleriyle struct oluşturma
	p3 := Person{Name: "Mehmet", City: "İzmir"}
	fmt.Println("3. Alan isimleriyle struct:", p3)

	////
	////
	////

	fmt.Println("\n=========================================")
	fmt.Println("  STRUCT İÇİNDE STRUCT VE DİĞER YAPILAR")
	fmt.Println("=========================================")
	// Struct içinde başka struct, slice veya diğer yapılar kullanma
	// Struct'lar, diğer struct'ları, slice'ları veya map'leri alan olarak içerebilir.
	// Bu, daha karmaşık veri yapıları oluşturmayı sağlar.

	// Örnek: Person struct'ı içinde Address struct'ı ve slice kullanımı
	type Address struct {
		City   string
		Street string
		Number int
	}
	type Employee struct {
		Person  // Embedded struct (gömülü struct)
		Address // Embedded struct
		Salary  float64
		Skills  []string
	}

	e1 := Employee{
		Person:  Person{Name: "Ali", Age: 28, City: "Bursa"},
		Address: Address{City: "Bursa", Street: "Atatürk", Number: 10},
		Salary:  15000,
		Skills:  []string{"Go", "Python", "SQL"},
	}
	fmt.Println("Struct içinde struct ve slice:", e1)

	////
	////
	////

	fmt.Println("\n===========================")
	fmt.Println("  ANONİM STRUCT KULLANIMI")
	fmt.Println("===========================")
	// Anonim Struct, genellikle geçici veri yapıları için kullanılır.
	// Örnek oluşturmadan kullanılır. Type ile oluşturulanlarda ise örnek oluşturulması gerekir.

	anon := struct {
		X int
		Y int
	}{X: 5, Y: 10}
	fmt.Println("Anonim struct:", anon)

	// VAR ile ANONİM STRUCT TANIMLAMA
	var anon2 struct {
		Name  string
		Age   int
		Email string
	}
	anon2.Name = "Fatma"
	anon2.Age = 20
	anon2.Email = "fatma@example.com"
	fmt.Println("Var ile anonim struct:", anon2)

	////
	////
	////

	fmt.Println("\n============================================")
	fmt.Println("  STRUCT İLE FONKSİYON VE METHOD KULLANIMI")
	fmt.Println("============================================")

	// Fonksiyona struct gönderme
	printPerson(p1)

	// Method tanımlama ve kullanma
	p1.Selamla()

	////
	////
	////

	fmt.Println("\n==============================")
	fmt.Println(" POINTER İLE STRUCT KULLANIMI")
	fmt.Println("==============================")

	// Pointer ile struct kullanımı
	// Pointer, struct'ın adresini tutar ve orijinal veriye erişim sağlar.

	p4 := &Person{Name: "Zeynep", Age: 22, City: "Antalya"}
	fmt.Println("Pointer ile struct:", p4)
	p4.Age = 23 // Pointer ile doğrudan alan değiştirilebilir
	fmt.Println("Pointer ile güncellenmiş struct:", p4)

	////
	////
	////

	fmt.Println("\n===========================")
	fmt.Println("  STRUCT ÖRNEĞİ KOPYALAMA")
	fmt.Println("===========================")
	// Struct örneği kopyalama
	// Struct'lar değer tipidir; kopyalandıklarında orijinal veriden bağımsız yeni bir kopya oluştururlar.
	// Bu nedenle, kopyalanan struct üzerinde yapılan değişiklikler orijinal struct'ı etkilemez.
	// Struct kopyalama, struct'ın tüm alanlarını kopyalar.

	p5 := p1
	p5.Name = "Fatma"
	fmt.Println("1. Orijinal struct:", p1)
	fmt.Println("2. Kopyalanmış struct:", p5)

	fmt.Println("====================\n====================\n====================")
}

// Fonksiyona struct gönderme
func printPerson(p Person) {
	fmt.Printf("Fonksiyon ile: %s, %d, %s\n", p.Name, p.Age, p.City)
}

// Struct'a method tanımlama (value receiver)
// Methodlar, struct'ın fonksiyon gibi davranmasını sağlar.
// Methodlar, struct'ın alanlarına erişebilir. Value receiver ile struct'ın kopyası üzerinde çalışılır.
// Eğer struct'ın orijinal değerini değiştirmek isterseniz pointer receiver kullanmalısınız.

func (p Person) Selamla() {
	fmt.Printf("Merhaba, ben %s!\n", p.Name)
}
