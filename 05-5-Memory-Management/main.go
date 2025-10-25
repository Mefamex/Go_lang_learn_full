/*
author: mefamex
date  : 2025-07-03
title : Memory Management and Garbage Collector in Go


MEMORY MANAGEMENT NEDİR?
- Memory Management (Bellek Yönetimi), programın çalışma zamanında bellek kullanımını kontrol eden sistemdir.
- Go'da otomatik bellek yönetimi vardır; programcının manuel olarak malloc/free yapmasına gerek yoktur.
- Go'nun Garbage Collector'ı (GC) kullanılmayan bellek alanlarını otomatik olarak temizler.
- Memory management, performans ve bellek verimliliği açısından kritik öneme sahiptir.


STACK vs HEAP:
- Stack: Fonksiyon çağrıları, local değişkenler ve parametreler için kullanılır
- Heap: Dinamik olarak allocate edilen objeler için kullanılır
- Stack daha hızlıdır, Heap daha esnektir
- Go compiler otomatik olarak hangisinin kullanılacağına karar verir (escape analysis)


GARBAGE COLLECTOR (GC):
- Go'da tricolor concurrent mark-and-sweep GC kullanılır
- GC üç renk (white, gray, black) sistemi ile çalışır
- Concurrent olarak çalışır, programı durdurmaz (stop-the-world minimum)
- GC, reachable olmayan objeleri otomatik olarak siler


MEMORY ALLOCATION:
- Go runtime otomatik olarak stack vs heap kararı verir
- Küçük objeler genelde stack'te, büyük objeler heap'te tutulur
- Pointer escape ettiğinde heap'e taşınır


PERFORMANCE CONSIDERATIONS:
- Memory pool kullanımı (sync.Pool)
- Object reuse ve recycling
- GC pressure'ı azaltma teknikleri
- Memory profiling ve analiz


MEMORY MANAGEMENT TEMEL KAVRAMLAR:
- Otomatik Garbage Collection
- Stack ve Heap otomatik yönetimi
- Escape Analysis ile optimizasyon
- Memory safety (dangling pointer yok)
- Tricolor concurrent GC
- Memory profiling araçları (pprof, memstats)
- Memory optimization teknikleri (object pooling, slice capacity management)
- Unsafe package ile düşük seviyeli bellek erişimi
- Manual memory management (C/C++ tarzı)
- Memory leak detection ve prevention
- GC tuning ve konfigürasyon





MEMORY MANAGEMENT BEST PRACTICES (İyi Uygulamalar)

Bellek yönetiminde yüksek performans ve güvenlik için aşağıdaki profesyonel yaklaşımları uygulayın:

1. Gereksiz pointer kullanımından kaçının:
    - Pointer kullanımı, objelerin heap'e taşınmasına ve GC yükünün artmasına neden olabilir.
    - Mümkünse değer (value) semantiğini tercih edin.

2. Büyük slice ve map'leri önceden allocate edin:
    - Kapasiteyi baştan belirlemek, tekrar tekrar reallocation ve GC overhead'ini azaltır.
    - make([]T, 0, capacity) veya make(map[K]V, capacity) kullanın.

3. Object pooling (sync.Pool) ile sık kullanılan objeleri yeniden kullanın:
    - Özellikle kısa ömürlü ve sık oluşturulan objeler için GC baskısını azaltır.
    - sync.Pool thread-safe ve GC ile entegredir.

4. Memory leak'leri önlemek için gereksiz referansları temizleyin:
    - Kullanılmayan objelerin referanslarını nil yaparak GC'nin temizlemesini sağlayın.
    - Kapanmayan goroutine'ler ve global referanslar memory leak'e yol açabilir.

5. Object reuse ile GC pressure'ı azaltın:
    - Slice, buffer ve struct gibi objeleri yeniden kullanarak allocation sayısını düşürün.
    - slice = slice[:0] ile slice'ı sıfırlayın, kapasiteyi koruyun.

6. Profiling ve analiz araçlarını kullanın (pprof, runtime/trace):
    - Memory profili çıkararak gereksiz allocation ve leak noktalarını tespit edin.
    - go tool pprof ile heap ve allocation analizleri yapın.

7. Memory-intensive işlemler için buffering ve streaming kullanın:
    - Büyük veri işlemlerinde tüm veriyi belleğe almak yerine, parça parça işleyin.
    - bytes.Buffer, bufio.Reader/Writer gibi araçları kullanın.

8. String birleştirme için strings.Builder kullanın:
    - + operatörü ile string birleştirmek çoklu allocation'a yol açar.
    - strings.Builder ile tek seferde ve verimli şekilde birleştirme yapın.

9. Escape analysis raporlarını inceleyin:
    - go build -gcflags="-m" ile hangi değişkenlerin heap'e taşındığını analiz edin.

10. GC tuning ve memory limitlerini bilinçli yönetin:
    - GOGC, GOMEMLIMIT gibi ortam değişkenleriyle GC davranışını ayarlayın.

11. Unsafe ve düşük seviyeli memory işlemlerini dikkatli kullanın:
    - Sadece performans kritik ve güvenli olduğundan emin olduğunuz durumlarda tercih edin.

12. Goroutine ve channel kullanımında memory leak riskine dikkat edin:
    - Kapanmayan goroutine'ler ve dolu channel'lar memory leak'e yol açabilir.

13. Kütüphane ve framework'lerin memory kullanımını analiz edin:
    - Harici paketlerin allocation davranışını gözlemleyin.

14. Memory footprint'i küçük tutmak için struct ve veri tiplerini optimize edin:
    - Field sıralaması, gereksiz alanlardan kaçınma, uygun veri tipi seçimi.

15. Test ve production ortamında memory kullanımını izleyin:
    - runtime.ReadMemStats ve pprof ile düzenli monitoring yapın.




GC TUNING VE KONFIGÜRASYON

Go'nun garbage collector'ı için ayarlanabilir parametreler ve tuning yöntemleri:

- GOGC ortam değişkeni:
  * GC'nin ne kadar sıklıkla çalışacağını belirler (varsayılan: 100, yani heap %100 büyüyünce GC tetiklenir).
  * Örn: GOGC=200 ile GC daha seyrek, GOGC=50 ile daha sık çalışır.

- GOMEMLIMIT ortam değişkeni (Go 1.19+):
  * Toplam kullanılabilir heap belleği sınırını belirler (örn: GOMEMLIMIT=512MiB).
  * Limit aşıldığında GC daha agresif çalışır.

- debug.SetGCPercent():
  * Programatik olarak GC yüzdesini ayarlamanızı sağlar.
  * debug.SetGCPercent(150) ile GC threshold'u %150 yapılır.

- runtime.GC():
  * Manuel olarak GC tetiklemek için kullanılır (genellikle test ve profil amaçlı).

- pprof ve runtime/trace ile GC davranışını analiz edin.

- GC tuning, performans ve bellek kullanımı arasında denge kurmak için yapılır.
  * Yüksek GOGC değeri: Daha az GC, daha fazla memory kullanımı
  * Düşük GOGC değeri: Daha sık GC, daha az memory kullanımı, potansiyel latency artışı


*/

package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
	"unsafe"
)

// Örnek struct (memory allocation örnekleri için)
type Person struct {
	Name string
	Age  int
	City string
	Data []byte // Büyük veri alanı simülasyonu için
}

// Heap allocation örneği (pointer döndürür)
func createPersonOnHeap() *Person {
	p := &Person{
		Name: "Ali",
		Age:  25,
		City: "İstanbul",
		Data: make([]byte, 1024), // 1KB veri
	}
	return p // Pointer döndürülüyor, heap'e taşınır
}

// Stack allocation örneği (value döndürür)
func createPersonOnStack() Person {
	p := Person{
		Name: "Ayşe",
		Age:  30,
		City: "Ankara",
		Data: make([]byte, 100), // Küçük veri
	}
	return p // Value döndürülüyor, stack'te kalabilir
}

func main() {

	fmt.Println("\n=======================================")
	fmt.Println("      STACK vs HEAP ALLOCATION")
	fmt.Println("=======================================")

	// Stack allocation örneği
	// Local değişkenler genellikle stack'te tutulur
	localVar := 42
	localString := "Hello World"
	localArray := [5]int{1, 2, 3, 4, 5}

	fmt.Printf("Local değişken adresi: %p\n", &localVar)
	fmt.Printf("Local string adresi: %p\n", &localString)
	fmt.Printf("Local array adresi: %p\n", &localArray)

	// Heap allocation örneği
	// Pointer döndüren fonksiyonlar heap'e allocation yapar
	heapPerson := createPersonOnHeap()
	stackPerson := createPersonOnStack()

	fmt.Printf("Heap person adresi: %p\n", heapPerson)
	fmt.Printf("Stack person adresi: %p\n", &stackPerson)

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("      ESCAPE ANALYSIS ÖRNEKLERİ")
	fmt.Println("=======================================")

	// Escape Analysis: Compiler'ın objeyi stack'te mi heap'te mi
	// tutacağına karar verme sürecidir

	// Bu değişken stack'te kalır (escape etmez)
	x := 100
	fmt.Println("Stack değişkeni:", x)

	// Bu değişken heap'e escape eder (pointer döndürülüyor)
	y := escapeToHeap()
	fmt.Printf("Heap'e escape eden değişken: %d, adres: %p\n", *y, y)

	// Slice allocation
	smallSlice := make([]int, 10)  // Küçük slice, stack'te kalabilir
	bigSlice := make([]int, 10000) // Büyük slice, heap'e gider
	fmt.Printf("Küçük slice uzunluğu: %d\n", len(smallSlice))
	fmt.Printf("Büyük slice uzunluğu: %d\n", len(bigSlice))

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("     GARBAGE COLLECTOR İŞLEMLERİ")
	fmt.Println("=======================================")

	// GC çalışmadan önceki memory stats
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)
	fmt.Printf("GC öncesi Heap boyutu: %d KB\n", bToKb(m1.HeapAlloc))
	fmt.Printf("GC sayısı: %d\n", m1.NumGC)

	// Çok sayıda obje oluştur (GC tetiklemek için)
	createManyObjects()

	// Manuel GC çalıştır
	runtime.GC()

	// GC sonrası memory stats
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)
	fmt.Printf("GC sonrası Heap boyutu: %d KB\n", bToKb(m2.HeapAlloc))
	fmt.Printf("GC sayısı: %d\n", m2.NumGC)

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("      MEMORY PROFILING VE ANALİZ")
	fmt.Println("=======================================")

	// Memory istatistikleri alma
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)

	fmt.Printf("Toplam allocated memory: %d KB\n", bToKb(ms.TotalAlloc))
	fmt.Printf("Sistem memory: %d KB\n", bToKb(ms.Sys))
	fmt.Printf("Heap objects sayısı: %d\n", ms.HeapObjects)
	fmt.Printf("Goroutine sayısı: %d\n", runtime.NumGoroutine())
	fmt.Printf("GC pause süresi: %v\n", time.Duration(ms.PauseNs[(ms.NumGC+255)%256]))

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("      MEMORY OPTIMIZASYON TEKNİKLERİ")
	fmt.Println("=======================================")

	// 1. Object Pooling örneği
	fmt.Println("1. Object Pooling kullanımı:")
	demonstrateObjectPooling()

	// 2. Memory reuse
	fmt.Println("\n2. Memory reuse tekniği:")
	demonstrateMemoryReuse()

	// 3. Slice capacity optimization
	fmt.Println("\n3. Slice capacity optimizasyonu:")
	demonstrateSliceOptimization()

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("      UNSAFE PACKAGE İLE MEMORY ERİŞİMİ")
	fmt.Println("=======================================")

	// unsafe package ile memory üzerinde düşük seviyeli işlemler
	// DİKKAT: Sadece performans kritik durumlarda kullanılmalı!

	str := "Go Memory Management"
	fmt.Printf("String: %s\n", str)
	fmt.Printf("String uzunluğu: %d bytes\n", len(str))
	fmt.Printf("String memory boyutu: %d bytes\n", unsafe.Sizeof(str))

	// Struct memory layout
	p := Person{Name: "Test", Age: 25, City: "Test"}
	fmt.Println("\nStruct Memory Layout:", p)
	fmt.Printf("Person struct boyutu: %d bytes\n", unsafe.Sizeof(p))
	fmt.Printf("Name field offset: %d\n", unsafe.Offsetof(p.Name))
	fmt.Printf("Age field offset: %d\n", unsafe.Offsetof(p.Age))
	fmt.Printf("City field offset: %d\n", unsafe.Offsetof(p.City))
	fmt.Printf("Data field offset: %d\n", unsafe.Offsetof(p.Data))

	////
	////
	////

	// GC ayarlarını göster
	gcPercent := debug.SetGCPercent(-1) // Mevcut değeri al, değiştirme
	debug.SetGCPercent(gcPercent)       // Geri ata
	fmt.Printf("Mevcut GC percentage: %d\n", gcPercent)

	fmt.Println("=======================================")
}

// Escape analysis örneği - heap'e escape eder
func escapeToHeap() *int {
	x := 42
	return &x // Pointer döndürülüyor, x heap'e escape eder
}

// Çok sayıda obje oluşturma (GC tetiklemek için)
func createManyObjects() {
	for i := 0; i < 100000; i++ {
		_ = &Person{
			Name: fmt.Sprintf("Person%d", i),
			Age:  i % 100,
			City: "TempCity",
			Data: make([]byte, 100),
		}
	}
}

// Object pooling demonstration
func demonstrateObjectPooling() {
	// sync.Pool kullanımı simülasyonu
	type Buffer struct {
		data []byte
	}

	// Buffer'ları yeniden kullanma
	buffers := make([]*Buffer, 0, 10)

	// Buffer pool simülasyonu
	getBuffer := func() *Buffer {
		if len(buffers) > 0 {
			buf := buffers[len(buffers)-1]
			buffers = buffers[:len(buffers)-1]
			return buf
		}
		return &Buffer{data: make([]byte, 0, 1024)}
	}

	putBuffer := func(buf *Buffer) {
		buf.data = buf.data[:0] // Reset
		buffers = append(buffers, buf)
	}

	// Kullanım örneği
	buf := getBuffer()
	buf.data = append(buf.data, []byte("test data")...)
	fmt.Printf("Buffer kullanıldı: %s\n", string(buf.data))
	putBuffer(buf)
	fmt.Println("Buffer pool'a geri verildi")
}

// Memory reuse demonstration
func demonstrateMemoryReuse() {
	// Slice'ı yeniden kullanma
	slice := make([]int, 0, 100)

	for round := 1; round <= 3; round++ {
		slice = slice[:0] // Reset length, capacity'yi koru

		for i := 0; i < 10; i++ {
			slice = append(slice, i)
		}

		fmt.Printf("Round %d: len=%d, cap=%d\n", round, len(slice), cap(slice))
	}
}

// Slice capacity optimization
func demonstrateSliceOptimization() {
	// Kötü: Her seferinde yeniden allocation
	var badSlice []int
	for i := 0; i < 1000; i++ {
		badSlice = append(badSlice, i) // Çok sayıda reallocation
	}

	// İyi: Önceden allocation
	goodSlice := make([]int, 0, 1000) // Capacity önceden belirlendi
	for i := 0; i < 1000; i++ {
		goodSlice = append(goodSlice, i) // Reallocation yok
	}

	fmt.Printf("Bad slice len=%d, cap=%d\n", len(badSlice), cap(badSlice))
	fmt.Printf("Good slice len=%d, cap=%d\n", len(goodSlice), cap(goodSlice))
}

// Utility function: bytes to kilobytes
func bToKb(b uint64) uint64 {
	return b / 1024
}
