/*
author: mefamex
date  : 2025-07-03
title : Pointers in Go


POINTER NEDİR?
- Pointer, başka bir değişkenin bellek adresini tutan özel bir değişken türüdür.
- Pointer'lar, verinin kopyalanması yerine orijinal veriye erişim sağlar.
- Go'da pointer'lar güvenlidir; pointer aritmetiği (C/C++'taki gibi) yoktur.
- Pointer'lar, büyük veri yapılarının kopyalanmasını engelleyerek performans artırır.
- Pointer'lar, fonksiyonlarda değişkenlerin orijinal değerlerini değiştirmek için kullanılır.
- Pointer'lar, nil değeri alabilir (hiçbir yeri göstermez).
- Pointer'lar referans semantiği sağlar; aynı veriye birden fazla yerden erişim mümkündür.



POINTER OPERATÖRLERI:
- & (Address-of): Bir değişkenin adresini verir
- * (Dereference): Pointer'ın gösterdiği değere erişir



POINTER TÜRLERI:
- *int, *string, *bool gibi temel tipler için pointer
- *struct, *slice, *map gibi karmaşık tipler için pointer
- interface{} türü için pointer



POINTER RECEIVER vs VALUE RECEIVER:
- Value Receiver : Method çağrıldığında struct'ın bir kopyası oluşturulur
- Pointer Receiver : Method çağrıldığında struct'ın orijinaline erişilir, değişiklikler kalıcıdır. (*Person)



POINTER KULLANIM ÖNERİLERİ

> Ne zaman Pointer kullanmalısınız:
1. Büyük struct'ları fonksiyona geçirirken (performans)
2. Orijinal veriyi değiştirmek istediğinizde
3. Struct method'larında kalıcı değişiklik için
4. nil kontrolü gerektiğinde
5. Interface implementation için

> Ne zaman Value kullanmalısınız:
1. Küçük veri tipleri için (int, string, bool)
2. Immutable (değişmez) davranış istediğinizde
3. Basit getter method'ları için
4. Thread-safety önemli olduğunda


*/

package main

import "fmt"

// Struct tanımlama (Pointer örnekleri için kullanılacak)
type Person struct {
	Name string
	Age  int
	City string
}

// Pointer receiver ile method tanımlama
func (p *Person) UpdateAge(newAge int) {
	p.Age = newAge // Orijinal struct'ı değiştirir
}

// Value receiver ile method tanımlama
// *Person yerine sadece Person kullanılırsa bu method value receiver olur
// Bu durumda struct'ın kopyası üzerinde çalışılır, orijinal değişmez
func (p Person) GetInfo() string {
	return fmt.Sprintf("%s, %d yaşında, %s'de yaşıyor", p.Name, p.Age, p.City)
}

// Pointer receiver ile method tanımlama (struct'ı tamamen değiştirir)
func (p *Person) UpdatePerson(name, city string, age int) {
	p.Name = name
	p.Age = age
	p.City = city
}

func main() {
	fmt.Println("=======================================")
	fmt.Println("    POINTER TEMEL KAVRAMLARI")
	fmt.Println("=======================================")

	// & operatörü: Değişkenin adresini alır
	// * operatörü: Pointer'ın gösterdiği değere erişir

	x := 42
	fmt.Println("x değişkenin değeri:", x)
	fmt.Println("x değişkenin adresi:", &x)

	// Pointer tanımlama
	var ptr *int = &x // ptr, x'in adresini tutar
	fmt.Println("ptr pointer'ının değeri (x'in adresi):", ptr)
	fmt.Println("ptr'nin gösterdiği değer (*ptr):", *ptr)

	// Pointer ile değer değiştirme
	*ptr = 100 // x'in değerini 100 yap
	fmt.Println("Pointer ile değiştirildikten sonra x:", x)

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("   FARKLI TİPLERDE POINTER KULLANIMI")
	fmt.Println("=======================================")

	// String pointer
	name := "Ahmet"
	namePtr := &name
	fmt.Println("String değeri:", name)
	fmt.Println("String pointer'ı:", namePtr)
	fmt.Println("Pointer'ın gösterdiği string:", *namePtr)

	*namePtr = "Mehmet"
	fmt.Println("Pointer ile değiştirildikten sonra name:", name)

	// Boolean pointer
	isActive := true
	boolPtr := &isActive
	fmt.Println("Boolean değeri:", isActive)
	*boolPtr = false
	fmt.Println("Pointer ile değiştirildikten sonra isActive:", isActive)

	////
	////
	////

	fmt.Println("\n================================")
	fmt.Println("    NIL POINTER VE GÜVENLIK")
	fmt.Println("================================")

	// Nil pointer
	var nilPtr *int
	fmt.Println("Nil pointer:", nilPtr)

	// Nil pointer'a değer atama
	newValue := 50
	nilPtr = &newValue
	fmt.Println("Nil pointer'a değer atandıktan sonra:", *nilPtr)

	////
	////
	////

	fmt.Println("\n==================================")
	fmt.Println("   STRUCT İLE POINTER KULLANIMI")
	fmt.Println("==================================")

	// Struct ile pointer kullanımı
	p1 := Person{Name: "Ali", Age: 25, City: "İstanbul"}
	fmt.Println("Orijinal struct:", p1)

	// Struct pointer'ı oluşturma
	p1Ptr := &p1
	fmt.Println("Struct pointer:", p1Ptr)

	// Pointer ile struct alanlarına erişim
	// Go otomatik olarak (*p1Ptr).Name şeklinde çevirir
	fmt.Println("Pointer ile Name erişimi:", p1Ptr.Name)
	fmt.Println("Pointer ile Age erişimi:", p1Ptr.Age)

	// Pointer ile struct alanını değiştirme
	p1Ptr.Age = 30
	fmt.Println("Pointer ile yaş değiştirildikten sonra:", p1)

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("   FONKSIYONLARDA POINTER KULLANIMI")
	fmt.Println("=======================================")

	// Value ile fonksiyon çağırma (kopya geçirilir)
	p2 := Person{Name: "Ayşe", Age: 20, City: "Ankara"}
	fmt.Println("Value ile fonksiyon öncesi:", p2)
	updatePersonByValue(p2)
	fmt.Println("Value ile fonksiyon sonrası:", p2) // Değişmez

	// Pointer ile fonksiyon çağırma (orijinal geçirilir)
	fmt.Println("Pointer ile fonksiyon öncesi:", p2)
	updatePersonByPointer(&p2)
	fmt.Println("Pointer ile fonksiyon sonrası:", p2) // Değişir

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("  POINTER RECEIVER vs VALUE RECEIVER")
	fmt.Println("=======================================")

	p3 := Person{Name: "Fatma", Age: 22, City: "İzmir"}
	fmt.Println("Method öncesi:", p3)

	// Value receiver ile method çağırma
	info := p3.GetInfo()
	fmt.Println("GetInfo method'u:", info)

	// Pointer receiver ile method çağırma (struct değişir)
	p3.UpdateAge(35)
	fmt.Println("UpdateAge method'u sonrası:", p3)

	// Birden fazla alanı güncelleyen pointer receiver
	p3.UpdatePerson("Zeynep", "Antalya", 28)
	fmt.Println("UpdatePerson method'u sonrası:", p3)

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("     NEW FONKSİYONU İLE POINTER")
	fmt.Println("=======================================")

	// new fonksiyonu: Sıfır değerli bir pointer döndürür
	p4 := new(Person) // *Person türünde sıfır değerli pointer
	fmt.Println("new ile oluşturulan pointer:", p4)
	fmt.Println("new ile oluşturulan struct:", *p4)

	// new ile oluşturulan pointer'a değer atama
	p4.Name = "Kemal"
	p4.Age = 40
	p4.City = "Bursa"
	fmt.Println("new pointer'ına değer atandıktan sonra:", *p4)

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("   SLICE VE MAP İLE POINTER İLİŞKİSİ")
	fmt.Println("=======================================")

	// Slice'lar zaten referans tipidir, pointer'a çok ihtiyaç duymazlar
	numbers := []int{1, 2, 3}
	fmt.Println("Orijinal slice:", numbers)
	modifySlice(numbers) // Slice referans tip olduğu için değişir
	fmt.Println("Fonksiyon sonrası slice:", numbers)

	// Ancak slice'ın kendisini değiştirmek için pointer gerekir
	modifySlicePointer(&numbers)
	fmt.Println("Slice pointer sonrası:", numbers)

	// Map'ler de referans tipidir
	scores := map[string]int{"Ali": 85, "Ayşe": 90}
	fmt.Println("Orijinal map:", scores)
	modifyMap(scores)
	fmt.Println("Fonksiyon sonrası map:", scores)

	////
	////
	////

	fmt.Println("=======================================")
}

// Value ile fonksiyon (kopya geçirilir, orijinal değişmez)
func updatePersonByValue(p Person) {
	p.Age = 999
	fmt.Println("Fonksiyon içinde (value):", p.Age)
}

// Pointer ile fonksiyon (orijinal geçirilir, değişir)
func updatePersonByPointer(p *Person) {
	p.Age = 999
	fmt.Println("Fonksiyon içinde (pointer):", p.Age)
}

// Slice fonksiyonu (slice referans tip olduğu için değişir)
func modifySlice(s []int) {
	if len(s) > 0 {
		s[0] = 100
	}
}

// Slice pointer fonksiyonu (slice'ın kendisini değiştirir)
func modifySlicePointer(s *[]int) {
	*s = append(*s, 999)
}

// Map fonksiyonu (map referans tip olduğu için değişir)
func modifyMap(m map[string]int) {
	m["Yeni"] = 95
}
