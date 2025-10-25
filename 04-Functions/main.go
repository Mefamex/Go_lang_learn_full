/*
author: mefamex
date  : 2025-07-02
title : Functions in Go

Bu dosya Go dilindeki fonksiyonlar ve özellikleri hakkında örnekler içerir:
- Temel fonksiyonlar
- Parametreli fonksiyonlar
- Dönüş değerleri
- Çoklu dönüş değerleri
- İsimlendirilmiş dönüş değerleri
- Variadic fonksiyonlar (değişken sayıda parametre)
- Closure (Anonim fonksiyonlar)
- Defer kullanımı
- Panic ve Recover mekanizması
- Metodlar (struct ile ilişkili fonksiyonlar)

Not: Daha gelişmiş fonksiyon özellikleri için Roadmap.md dosyasındaki
ADVANCED bölümüne bakabilirsiniz.


FONKSIYON ISIMLENDIRME KURALLARI:
- camelCase kullanılır -> myFunctionIsAwesome
- İlk harf küçük olmalıdır
- Dışa aktarmak için -> İlk harf büyük olmalıdır -> ExportedFunction


*/

package main

import "fmt"

// Temel fonksiyon
func hello() {
	fmt.Println("Merhaba, Dünya!")
}

// Parametre ve argüman arasındaki fark:
// - Parametre: Fonksiyon tanımında belirtilen değişkenlerdir.
// - Argüman: Fonksiyon çağrılırken verilen değerlerdir.

// Parametreli fonksiyon
func greet(name string) {
	fmt.Printf("Merhaba, %s!\n", name)
}

// Dönüş değeri olan fonksiyon
func add(a, b int) int {
	return a + b
}

// Çoklu dönüş değeri olan fonksiyon
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("sıfıra bölme hatası")
	}
	return a / b, nil
}

// İsimlendirilmiş dönüş değerleri
func calculate(x, y int) (sum, diff int) {
	sum = x + y
	diff = x - y
	return // Çıplak return - sum ve diff otomatik döner
}

// Variadic fonksiyon (değişken sayıda parametre)
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Closure (Anonim fonksiyon)
// closure fonksiyonu: Dış fonksiyonun değişkenlerine erişebilen iç içe fonksiyondur.
// Bu örnekte, counter() fonksiyonu bir closure döndürür.
// Döndürülen closure, counter() fonksiyonunun kapsamındaki count değişkenine erişir.
// Her çağrıldığında count değişkenini artırır ve yeni değeri döndürür.
// count değişkeni closure tarafından "yakalanar" ve fonksiyon çağrıları arasında değeri korunur.
func counter() func() int {
	// count değişkeni burada tanımlanır ve closure tarafından "yakalanır"
	count := 0
	// Anonim bir fonksiyon döndürülüyor (closure)
	return func() int {
		count++ // Dış kapsamdaki count değişkenine erişim ve değiştirme
		return count
	}
}

// Defer kullanımı
func deferExample() {
	// defer, fonksiyonun sonuna eklenen ve fonksiyon bitmeden önce çalıştırılan bir komuttur.
	// Defer ile yazılan fonksiyonlar, çağrıldıkları sırada değil, fonksiyonun sonuna eklenir ve en son çalıştırılır.
	defer fmt.Println("Bu en son yazılacak")
	fmt.Println("Bu ilk yazılacak")
}

// Panic ve Recover
func panicAndRecover() {
	// Panic, programın normal akışını bozar ve geri dönüş yapmadan programı durdurur.
	// Recover, panic durumunda programın çökmesini engeller ve kontrolü geri alır.
	// Panic ve Recover, hata yönetimi için kullanılır.
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panikten kurtuldu: %v\n", r)
		}
	}()

	panic("Bu bir panik!")
}

// METHODLAR

// Methodlar, Go dilinde struct'lar üzerinde tanımlanan fonksiyonlardır.
// Methodlar, struct'ın bir örneği üzerinde çalışır ve bu örneğin verilerini kullanabilir.
// Detaylı olarak ayrı bir bölümde ele alınacaktır.

// Method (struct üzerinde)
type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r *Rectangle) scale(factor float64) {
	r.width *= factor
	r.height *= factor
}

////
////
////

func main() {
	fmt.Println("\n--- Temel Fonksiyon ---")
	hello()

	////

	fmt.Println("\n--- Parametreli Fonksiyon ---")
	greet("Ahmet")

	////

	fmt.Println("\n--- Dönüş Değeri Olan Fonksiyon ---")
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)

	////

	fmt.Println("\n--- Çoklu Dönüş Değeri ---")
	if result, err := divide(10, 2); err != nil {
		fmt.Println("Hata:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	////

	fmt.Println("\n--- İsimlendirilmiş Dönüş Değerleri ---")
	total, diff := calculate(10, 5)
	fmt.Printf("10 + 5 = %d, 10 - 5 = %d\n", total, diff)

	////

	fmt.Println("\n--- Variadic Fonksiyon ---")
	fmt.Println("Toplam:", sum(1, 2, 3, 4, 5))
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice Toplam:", sum(numbers...))

	////

	fmt.Println("\n--- Closure Örneği ---")
	count := counter()
	fmt.Println(count()) // 1
	fmt.Println(count()) // 2
	fmt.Println(count()) // 3

	////
	fmt.Println("\n--- Defer Örneği ---")
	deferExample()

	////

	fmt.Println("\n--- Panic ve Recover ---")
	panicAndRecover()

	////

	fmt.Println("\n--- Method Kullanımı ---")
	rect := Rectangle{width: 10, height: 5}
	fmt.Printf("Alan: %.2f\n", rect.area())

	rect.scale(2)
	fmt.Printf("Ölçeklendikten sonra alan: %.2f\n", rect.area())
}
