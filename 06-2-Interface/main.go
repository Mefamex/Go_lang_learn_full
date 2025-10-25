/*
Author : mefamex
Date   : 2025-07-05
Title  : Interface (Arayüzler) - Go



INTERFACE NEDİR? (GENEL BAKIŞ)
- Interface, bir veya daha fazla method imzası (signature) içeren soyut bir tiptir.
- Go'da interface'ler, bir tipin ilgili methodları implemente edip etmediğine bakılarak otomatik (implicit) olarak uygulanır.
- "Accept interfaces, return structs" prensibi ile esnek, sürdürülebilir ve test edilebilir kodlar yazılır.
- Polimorfizm sağlar: Farklı tipler aynı interface'i implemente edebilir ve ortak bir davranış sergileyebilir.
- Go'da çoklu kalıtım yoktur; interface'ler ile kompozisyon ve esneklik sağlanır.
- Diğer dillerden farklı olarak, Go'da bir tip, interface'teki tüm methodları implemente ettiği anda o interface'i otomatik olarak uygular (explicit anahtar kelimesine gerek yoktur).
- Interface'ler, method imzalarını tanımlar; implementasyon içermezler.


INTERFACE KULLANIM AMAÇLARI
- Verilerle değil, metotlarla etkileşim kurmayı sağlayan bir yapıdır.
- Bir türün, belirli bir davranışı (metotları) uygulayıp uygulamadığını kontrol eder.
- Kodunuzu modüler, esnek ve sürdürülebilir hale getirir.
- Farklı tiplerin aynı davranış setini paylaşmasını sağlar.
- Test edilebilirlik: Mock'lar ile bağımlılıkları izole ederek birim testleri kolaylaştırır.
- API tasarımı ve bağımlılık enjeksiyonu (dependency injection) için idealdir.
- Polimorfizm ve soyutlama sağlar.


INTERFACE'LERİN KULLANIM ALANLARI
- API ve framework tasarımı: Farklı implementasyonları desteklemek için.
- Bağımlılık enjeksiyonu: Test edilebilir ve gevşek bağlı kodlar için.
- Polimorfizm: Farklı tiplerin aynı davranışı sergilemesi için.
- Event-driven mimariler: Farklı olayları dinlemek ve işlemek için.
- Veri yapıları: Generic veri yapıları ve algoritmalar için.
- Mocking ve testler: Testlerde bağımlılıkları izole etmek için.
- Dinamik tip kontrolü: Runtime'da tip kontrolü yapmak için.
- Dinamik veri yapıları: Farklı tipleri tutabilen veri yapıları oluşturmak için.
- API'lerde dinamik tip desteği: Farklı tiplerin aynı arayüzü kullanabilmesi için.
- Event sistemleri: Olayları dinlemek ve işlemek için.
- Veri tabanı işlemleri: Farklı veri tabanı sürücüleri ile etkileşim için.
- Veri işleme: Farklı veri kaynaklarından gelen verileri işlemek için.
- Web uygulamaları: HTTP handler'lar ve middleware'ler için.
- Grafik ve oyun programlama: Farklı nesnelerin ortak davranışlarını tanımlamak için.


INTERFACE'LERİN TEMEL ÖZELLİKLERİ
- Sadece method imzalarını tanımlar, implementasyon içermez.
- Bir tip, bir interface'i uygulamak için o interface'teki tüm methodları implement etmelidir.
- Boş interface (interface{}) herhangi bir tipteki değeri tutabilir (dinamik tip).
- Interface'ler, bir tipin belirli bir interface'i uygulayıp uygulamadığını kontrol etmek için kullanılabilir.


INTERFACE'LERİN AVANTAJLARI
- Kodunuzu daha esnek, genişletilebilir ve sürdürülebilir kılar.
- Farklı tiplerin aynı interface'i uygulamasına olanak tanır.
- Test edilebilirliği artırır; mock ve stub ile kolay test imkanı sunar.
- API tasarımında ve bağımlılık enjeksiyonunda yaygın olarak kullanılır.


INTERFACE'LERİN DEZAVANTAJLARI
- Performans: Interface kullanımı, tip kontrolü ve method çağrıları nedeniyle küçük bir performans kaybına yol açabilir.
- Tip güvenliği: Boş interface (interface{}) kullanımı, tip güvenliğini azaltabilir ve runtime hatalarına sebep olabilir.
- Karmaşıklık: Çok sayıda interface ve katman, kodun okunabilirliğini ve bakımını zorlaştırabilir.



*/

package main

import (
	"fmt"
	"math"
)

// ============================================================
// 1. TEMEL INTERFACE TANIMI ve KULLANIMI
// ============================================================

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

// --- Circle için methodlar (Shape interface'i uygular) ---
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// --- Rectangle için methodlar (Shape interface'i uygular) ---
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

////
////
////

// ============================================================
// 2. INTERFACE İLE POLİMORFİZM (ÇOK BİÇİMLİLİK)
// ============================================================
func printShapeInfo(s Shape) {
	fmt.Println("---------------------------")
	fmt.Printf("Type: %T\n", s)
	fmt.Printf("Alan: %.2f\n", s.Area())
	fmt.Printf("Çevre: %.2f\n", s.Perimeter())
}

////
////
////

// ============================================================
// 3. EMPTY INTERFACE (interface{}) ve DİNAMİK TİPLER
// ============================================================
func printAny(val interface{}) {
	fmt.Printf("Değer: %v, Tip: %T\n", val, val)
}

////
////
////

// ============================================================
// 4. TYPE ASSERTION (Tip Doğrulama)
// ============================================================
func typeAssertionDemo(val interface{}) {
	fmt.Println("Type Assertion Demo:")
	str, ok := val.(string)
	if ok {
		fmt.Println("String değer:", str)
	} else {
		fmt.Println("Bu bir string değil.")
	}
}

////
////
////

// ============================================================
// 5. TYPE SWITCH (Tip Kontrolü)
// ============================================================
func typeSwitchDemo(val interface{}) {
	fmt.Println("Type Switch Demo:")
	switch v := val.(type) {
	case int:
		fmt.Println("int değer:", v)
	case string:
		fmt.Println("string değer:", v)
	case Shape:
		fmt.Printf("Shape: Alan=%.2f\n", v.Area())
	default:
		fmt.Println("Bilinmeyen tip")
	}
}

////
////
////

// ============================================================
// 6. STANDART KÜTÜPHANE INTERFACE'LERİ (fmt.Stringer)
// ============================================================
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d)", p.Name, p.Age)
} // Person tipi fmt.Stringer interface'ini uygular

////
////
////

// ============================================================
// 7. NIL INTERFACE ve NIL CONCRETE DEĞER FARKI
// ============================================================
func nilInterfaceDemo() {
	var s Shape                             // nil interface
	fmt.Println("Nil interface:", s == nil) // true
	var c *Circle = nil
	s = c                                      // s artık nil concrete değer tutan bir interface
	fmt.Println("Interface nil mi?", s == nil) // false!
	fmt.Printf("Interface içeriği: %#v\n", s)
}

////
////
////

// ============================================================
// 8. POINTER RECEIVER İLE INTERFACE IMPLEMENTASYONU & TYPE ASSERTION HATASI
// ============================================================
type Writer interface {
	Write([]byte) (int, error)
}
type MyWriter struct{}

func (mw *MyWriter) Write(b []byte) (int, error) {
	fmt.Println("Yazıldı:", string(b))
	return len(b), nil
}

func pointerReceiverAndTypeAssertionDemo() {
	fmt.Println("Pointer Receiver ile Interface Implementasyonu:")
	var w Writer
	// w = MyWriter{} // HATA: *MyWriter pointer receiver ile implement etti
	w = &MyWriter{} // DOĞRU: *MyWriter pointer ile atanmalı
	w.Write([]byte("Merhaba Interface!"))

	fmt.Println("Type Assertion Hatası Örneği:")
	var val interface{} = 123
	// str := val.(string) // HATA: panic oluşur
	_, ok := val.(string)
	fmt.Println("Type assertion başarılı mı?", ok)
	if !ok {
		fmt.Println("Type assertion başarısız, panic oluşmadı.")
	}
}

////
////
////

// ============================================================
// 9. ANA FONKSİYON: KONU ÖZETİ ve DEMOLAR
// ============================================================
func main() {
	fmt.Println("============================================================")
	fmt.Println("INTERFACE TEMELLERİ ve KULLANIMI")
	fmt.Println("============================================================")

	c := Circle{Radius: 3}
	r := Rectangle{Width: 4, Height: 5}
	printShapeInfo(c)
	printShapeInfo(r)

	fmt.Println("\n--- POLİMORFİZM ÖRNEĞİ ---")
	shapes := []Shape{c, r, Circle{Radius: 1.5}}
	for _, s := range shapes {
		printShapeInfo(s)
	}

	fmt.Println("\n--- EMPTY INTERFACE (interface{}) ---")
	printAny(42)
	printAny("merhaba")
	printAny(c)

	fmt.Println("\n--- TYPE ASSERTION ---")
	typeAssertionDemo("test string")
	typeAssertionDemo(123)

	fmt.Println("\n--- TYPE SWITCH ---")
	typeSwitchDemo(99)
	typeSwitchDemo("hello")
	typeSwitchDemo(r)

	fmt.Println("\n--- STANDART KÜTÜPHANE INTERFACE ---")
	p := Person{Name: "Ali", Age: 30}
	fmt.Println(p) // fmt.Stringer interface otomatik kullanılır

	fmt.Println("\n--- NIL INTERFACE ve NIL DEĞER FARKI ---")
	nilInterfaceDemo()

	fmt.Println("\n--- BEST PRACTICES & YAYGIN HATALAR ---")
	pointerReceiverAndTypeAssertionDemo()

	fmt.Println("\n============================================================")
}
