/*
author: mefamex
date  : 2025-07-03
title : Slices in Go

SLICE NEDİR?
- Slice, Go'da dinamik boyutlu, aynı türden elemanlar içeren bir veri yapısıdır.
- Slice'lar, dizilere (array) göre daha esnektir ve boyutları çalışma zamanında değiştirilebilir.
- Slice'lar bir diziye referans tutar, yani slice üzerinde yapılan değişiklikler diziyi de etkiler.
- Slice'lar, diziye göre daha çok tercih edilir çünkü boyutları dinamik olarak büyüyüp küçülebilir.
- Slice'lar, dizinin bir alt kümesini gösterebilir veya sıfırdan oluşturulabilir.
- Slice'lar, başlangıç adresi, uzunluk (len) ve kapasite (cap) bilgilerine sahiptir.
- Slice'lar, fonksiyonlara parametre olarak geçirildiğinde referans olarak geçilir.

- REFERANS TİPİDİR: Slice'lar, dizinin bir görünümüdür ve bellekteki veriye referans tutar.
-                -> Slice referans tiplidir. Diziler ise değer tiplidir.

- Sıfır değerli bir slice: nil'dir ve len(slice)==0'dır.
- Slice'lar, make fonksiyonu ile oluşturulabilir: make([]int, 5) // 5 elemanlı int slice -> [0 0 0 0 0]
- Slice'lar, append fonksiyonu ile dinamik olarak eleman eklenebilir: append(slice, 10)
- Slice'lar, copy fonksiyonu ile kopyalanabilir: copy(dst, src)
- Slice'lar, alt slice'lar oluşturabilir: slice[1:3] // 1. indexten başlar, 3. indexe kadar (3 dahil değil)
- Slice'lar, dizi gibi indeksleme ile erişilebilir: slice[0], slice[1], ...
- Slice'lar, sıfır değerli bir slice olarak tanımlanabilir: var nilSlice []int
- Slice'lar, boş bir slice olarak tanımlanabilir: emptySlice := []int{}


*/

package main

import "fmt"

func main() {
	fmt.Println("=========================")
	fmt.Println("SLICE OLUŞTURMA YOLLARI")
	fmt.Println("=========================")

	// 1. Bir diziden slice oluşturma
	arr := [5]int{1, 2, 3, 4, 5}
	slice1 := arr[1:4]                       // 1. indexten başlar, 4. indexe kadar (4 dahil değil)
	fmt.Println("1. Diziden slice:", slice1) // [2 3 4]

	// 2. make fonksiyonu ile slice oluşturma
	slice2 := make([]int, 3) // 3 elemanlı, sıfır değerli slice
	fmt.Println("2. make ile slice:", slice2)

	// 3. Kısa yoldan slice oluşturma
	slice3 := []string{"go", "lang", "slice"} // [] içine sayı veya "..." yazarsak dizi oluşur
	fmt.Println("3. Kısa yoldan slice:", slice3)

	////
	////
	////

	fmt.Println("\n=========================")
	fmt.Println("  SLICE TEMEL İŞLEMLERİ")
	fmt.Println("=========================")

	// len ve cap fonksiyonları
	fmt.Println("Cap (capacity): Sclice'ın kapasitesi, bellekteki toplam eleman sayısını gösterir.")
	fmt.Println("Len (length): Slice'ın uzunluğu, içindeki eleman sayısını gösterir.")
	fmt.Println("slice1 len:", len(slice1), "cap:", cap(slice1))
	fmt.Println("slice2 len:", len(slice2), "cap:", cap(slice2))
	fmt.Println("slice3 len:", len(slice3), "cap:", cap(slice3))

	// append ile eleman ekleme
	slice2 = append(slice2, 10)
	fmt.Println("append ile ekleme:", slice2)

	// Birden fazla eleman ekleme
	slice2 = append(slice2, 20, 30, 40)
	fmt.Println("çoklu ekleme:", slice2)

	// append ile başka bir slice ekleme
	slice2 = append(slice2, slice1...)
	fmt.Println("başka slice ekleme:", slice2)

	////
	////
	////

	fmt.Println("\n=========================")
	fmt.Println("  SLICE-ARRAY İLİŞKİSİ")
	fmt.Println("=========================")

	// Slice, dizinin bir görünümüdür (view)
	arr2 := [4]string{"a", "b", "c", "d"}
	slc := arr2[1:3]
	fmt.Println("Orijinal dizi:", arr2)
	fmt.Println("Slice:", slc)

	// Slice üzerinden yapılan değişiklik diziye yansır
	slc[0] = "X"
	fmt.Println("Slice değiştirildi:", slc)
	fmt.Println("Dizi de değişti:", arr2)

	////
	////
	////

	fmt.Println("\n================================")
	fmt.Println("  SLICE KOPYALAMA VE ALT SLICE")
	fmt.Println("================================")

	// copy fonksiyonu ile slice kopyalama
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Println("Kopyalanan slice:", dst)

	// Alt slice alma
	bigSlice := []int{10, 20, 30, 40, 50}
	subSlice := bigSlice[1:4]  // 1. indexten (20 değerinden) başlar, 4. indexe (50 değerine) kadar ama 4 dahil değil
	leftSlice := bigSlice[:3]  // 0. indexten başlar, 3. indexe kadar (3 dahil değil)
	rightSlice := bigSlice[2:] // 2. indexten başlar, sonuna kadar
	allSlice := bigSlice[:]    // Tüm slice'ı alır
	fmt.Println("Alt slice:", subSlice)
	fmt.Println("Sol slice [:3] :", leftSlice)
	fmt.Println("Sağ slice [2:] :", rightSlice)
	fmt.Println("Tüm slice  [:] :", allSlice)

	////
	////
	////

	fmt.Println("\n=========================")
	fmt.Println("  SLICE ÖRNEKLERİ")
	fmt.Println("=========================")

	// Nil slice
	var nilSlice []int
	fmt.Println("Nil slice:", nilSlice, "len:", len(nilSlice), "cap:", cap(nilSlice))

	// Empty slice
	emptySlice := []int{}
	fmt.Println("Empty slice:", emptySlice, "len:", len(emptySlice), "cap:", cap(emptySlice))

	// Slice ile fonksiyon kullanımı
	fmt.Println("Fonksiyon ile slice toplamı:", sumSlice([]int{1, 2, 3, 4, 5}))

	fmt.Println("=========================")
}

// Bir slice'ın elemanlarını toplayan fonksiyon
func sumSlice(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}
