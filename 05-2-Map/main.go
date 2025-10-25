/*
author: mefamex
date  : 2025-07-03
title : Maps in Go

MAP NEDİR?
- Map, Go'da anahtar-değer (key-value) ilişkisiyle veri saklayan dinamik bir veri yapısıdır.
- Map'lerde anahtarlar (key) benzersizdir ve her anahtar bir değere (value) karşılık gelir.
- Map'ler, Python'daki dict, JavaScript'teki object, Java'daki HashMap'e benzer.
- Map'ler referans tiptir; fonksiyona gönderildiğinde orijinal veriyi değiştirebilir.
- Map'ler dinamik olarak büyüyüp küçülebilir.
- Map'lerde anahtar olarak karşılaştırılabilir (comparable) tipler kullanılabilir (int, string, bool, array, struct).
- Map'ler sıfır değerli olarak nil'dir ve kullanılmadan önce oluşturulmalıdır.
- Nil map'e eleman eklenemez, önce make ile oluşturulmalı.
- Map'te olmayan bir anahtara erişirseniz, ilgili value tipinin sıfır değeri (zero value) döner.
- Map'ler thread-safe değildir; eşzamanlı erişim için sync.Map veya mutex kullanılmalıdır.
- Map'ler sıralı değildir; range ile dolaşırken anahtar sırası rastgele gelir.
    ->(alfabetik veya sayıya göre belki olabilir, güvenmemek lazım)

> map[keyType]valueType
- Anahtarlar benzersizdir, aynı anahtar birden fazla kez kullanılamaz.
- Değerler aynı türde olmalıdır.


*/

package main

import "fmt"

func main() {
	fmt.Println("=========================")
	fmt.Println("MAP OLUŞTURMA YOLLARI")
	fmt.Println("=========================")

	// 1. make fonksiyonu ile map oluşturma
	m1 := make(map[string]int) // string anahtar, int değer
	fmt.Println("1. make ile map:", m1)

	// 2. Kısa yoldan map oluşturma
	m2 := map[string]string{"ad": "Ahmet", "şehir": "İstanbul"}
	fmt.Println("2. Kısa yoldan map:", m2)

	// 3. Boş map tanımlama (nil map)
	var m3 map[int]bool // nil map, kullanılmadan önce make ile oluşturulmalı
	fmt.Println("3. Nil map:", m3)
	// m3["test"] = true // Bu satır hata verir, çünkü m3 nil'dir ve eleman eklenemez
	m3 = make(map[int]bool) // Önce make ile oluşturulmalı
	fmt.Println("3. Nil map oluşturulduktan sonra:", m3)
	m3[0] = true // Artık bu satır çalışır, çünkü m3 artık boş bir map

	////
	////
	////

	fmt.Println("\n=========================")
	fmt.Println("  MAP TEMEL İŞLEMLERİ")
	fmt.Println("=========================")

	// Eleman ekleme
	m1["bir"] = 1
	m1["iki"] = 2
	fmt.Println("Eleman ekleme (1,2):", m1)

	// Eleman okuma
	fmt.Println("m1[\"bir\"]:", m1["bir"])

	// Eleman güncelleme
	m1["bir"] = 10
	fmt.Println("Güncellenmiş m1[\"bir\"]:", m1["bir"])

	// Eleman silme
	delete(m1, "iki")
	fmt.Println("Eleman silindikten sonra:", m1)

	// Map uzunluğu (eleman sayısı)
	fmt.Println("m1 uzunluğu:", len(m1))

	// Kopyalama
	m4 := m1 // Referans kopyalama, m1 ve m4 aynı map'i gösterir
	fmt.Println("m4 (m1'in referansı):", m4)
	// m1'i değiştirmek m4'ü de etkiler
	m1["bir"] = 100
	fmt.Println("Güncellenmiş m1 ve m4:", m1, m4)
	// Go'da map'ler arasında doğrudan tüm key-value'ları tek satırda kopyalamanın yerleşik bir yolu yoktur.
	// Yani, for döngüsü dışında map'in tüm içeriğini başka bir map'e kopyalamak mümkün değildir.
	// Aşağıdaki gibi bir fonksiyon veya for döngüsü kullanmak zorunludur.
	m5 := make(map[string]int)
	for k, v := range m1 {
		m5[k] = v // m1'in değerlerini m5'e kopyala
	}

	////
	////
	////

	fmt.Println("\n=========================")
	fmt.Println("MAP'TE ANAHTAR VAR MI? (ok idiom)")
	fmt.Println("=========================")

	// Bir anahtarın map'te olup olmadığını kontrol etme
	value, exists := m1["bir"]
	if exists {
		fmt.Println("'bir' anahtarı var, değeri:", value)
	} else {
		fmt.Println("'bir' anahtarı yok")
	}

	// Kısa kullanım
	if v, ok := m1["iki"]; ok {
		fmt.Println("'iki' anahtarı var, değeri:", v)
	} else {
		fmt.Println("'iki' anahtarı yok")
	}

	////
	////
	////

	fmt.Println("\n=============================")
	fmt.Println("  MAP ÜZERİNDE DÖNGÜ (range)")
	fmt.Println("=============================")

	// Map üzerinde range ile gezinme
	for k, v := range m2 {
		fmt.Printf("Anahtar: %s, Değer: %s\n", k, v)
	}

	////
	////
	////

	fmt.Println("\n=================================")
	fmt.Println("  MAP OF SLICE & MAP OF STRUCT")
	fmt.Println("=================================")

	// Map of slice
	students := map[string][]int{
		"Ali":  {90, 85, 80},
		"Ayşe": {95, 88, 92},
	}
	fmt.Println("Map of slice:", students)

	// Map of struct
	type Person struct {
		Name string
		Age  int
	}
	people := map[string]Person{
		"ahmet": {Name: "Ahmet", Age: 30},
		"ayse":  {Name: "Ayşe", Age: 25},
	}
	fmt.Println("Map of struct:", people)

	////
	////
	////

	fmt.Println("\n================================")
	fmt.Println("   MAP İLE FONKSİYON KULLANIMI")
	fmt.Println("================================")

	// Map'i fonksiyona gönderme
	demoMap := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println("Toplam:", sumMapValues(demoMap))

	fmt.Println("=========================")
}

// Map'in tüm değerlerini toplayan fonksiyon
func sumMapValues(myMap map[string]int) int {
	total := 0
	for _, v := range myMap {
		total += v
	}
	return total
}
