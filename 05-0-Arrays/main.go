/*
author: mefamex
date  : 2025-07-02
title : Arrays in Go

GO Dilinde Dizi (Array) Nedir?

- Dizi, aynı türdeki birden fazla veriyi tek bir değişkende saklamamızı sağlayan temel bir veri yapısıdır.
- Dizi boyutu sabittir; tanımlandıktan sonra değiştirilemez.
- Dizi elemanlarına indeks numarası ile erişilir ve indeksler 0'dan başlar.
- Dizi elemanları bellekte ardışık olarak saklanır, bu da hızlı erişim sağlar.
- Dizi tanımlanırken boyut belirtilebilir: örneğin [5]int gibi.
- Boyut belirtilmeden de değer atayarak dizi oluşturulabilir: [...]int{1,2,3}.
- Dizi oluşturulduğunda tüm elemanlar varsayılan (sıfır) değer ile başlar (int için 0, string için "", bool için false vb.).
- Dizi elemanlarına erişim ve atama işlemleri indeks numarası ile yapılır: dizi[0] = 10 gibi.
- Dizi yalnızca aynı türdeki verileri saklayabilir; farklı türler için slice, struct veya map gibi yapılar tercih edilir.
- Dizi, sabit boyutlu ve homojen veri saklama ihtiyacı olduğunda kullanılır.
- Go dilinde diziler, fonksiyonlara parametre olarak geçirildiğinde kopyalanır (referans değil, değer olarak).
- Daha esnek ve dinamik veri yapıları için slice kullanılır. (sonraki konu)
*/

package main

import "fmt"



func main() {
	fmt.Println("=========================\n=========================\n=========================")

	// 1. Sabit boyutlu boş dizi tanımlama
	var numbers [5]int
	fmt.Println("1. Boş dizi:", numbers)

	// 2. Sabit boyutlu ve değer atanmış dizi
	var moreNumbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("2. Dolu dizi:", moreNumbers)

	// 3. Boyut belirtmeden değer atama (boyut otomatik belirlenir)
	anotherNumbers := [...]int{6, 7, 8, 9, 10}
	//                 []  int{6, 7, 8, 9, 10}  -> slice
	fmt.Println("3. Boyut belirtilmeden dizi:", anotherNumbers)

	// 4. Kısmî değer atama (atanmayanlar sıfır değeri alır)
	var partialNumbers = [5]int{1, 2}
	fmt.Println("4. Kısmî değer atanmış dizi:", partialNumbers)

	// 5. Belirli indekslere değer atama
	var indexedNumbers = [5]int{0: 100, 3: 400}
	fmt.Println("5. Belirli indekslere değer atanmış dizi:", indexedNumbers)

	// 6. String dizisi tanımlama
	var stringArray = [3]string{"go", "lang", "array"}
	fmt.Println("6. String dizi:", stringArray)

	// 7. Boolean dizi tanımlama
	boolArray := [2]bool{true, false}
	fmt.Println("7. Boolean dizi:", boolArray)

	// 8. Çok boyutlu dizi tanımlama
	// Çok boyutlu dizilerdeki elemanların türleri aynı olmalıdır
	// alt boyutlar da aynı türde olmalıdır
	var matrix [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("8. Çok boyutlu dizi:", matrix)
	fmt.Println("8. Çok boyutlu dizinin 2. satır 3. sütun elemanı:", matrix[1][2])


	// 9. Kısmî çok boyutlu dizi tanımlama
	var partialMatrix = [2][3]int{{1}, {4, 5}}
	fmt.Println("9. Kısmî çok boyutlu dizi:", partialMatrix)
	fmt.Println("")

	////
	////
	////

	// 1. Tek tek indeks ile değer atama
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50
	fmt.Println("Diziye atanan değerler:", numbers)

	// 2. Döngü ile değer atama
	for i := 0; i < len(numbers); i++ {
		numbers[i] = (i + 1) * 100
	}
	fmt.Println("Döngü ile değer atanmış dizi:", numbers)

	// 3. Range ile değer atama
	// index, value := range numbers
	for i := range numbers {
		numbers[i] = i * 10
	}
	fmt.Println("Range ile değer atanmış dizi:", numbers)

	// 4. Çok boyutlu diziye değer atama
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = (i + 1) * (j + 1)
		}
	}

	// 5. Kopyalama ile değer atama
	// Kopylama işlemi, dizilerde değer olarak yapılır. Slice'larda ise referans olarak yapılır.
	// yeni kopya dizideki eleman değişiklikleri, orijinal diziyi etkilemez.
	var copyNumbers = numbers
	// copyNumbers := numbers
	fmt.Println("Başka bir diziye kopyalama ile atama:", copyNumbers)

	////
	////
	////

	// Dizi elemanlarına erişim
	fmt.Println("\nİlk eleman:", numbers[0])
	fmt.Println("İkinci eleman:", numbers[1])
	fmt.Println("Beşinci eleman:", numbers[4])

	// Dizi uzunluğu
	fmt.Println("\nDizi uzunluğu:", len(numbers))

	// Slice konusuna bakmayı unutmayın!

	fmt.Println("=========================\n=========================\n=========================")
}
