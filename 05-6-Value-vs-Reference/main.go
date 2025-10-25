/*
author: mefamex
date  : 2025-07-05
title : Value vs Reference types



VALUE vs REFERENCE TYPES - TEMEL KAVRAMLAR

Go'da veri tipleri iki ana kategoriye ayrılır >>> Value Types / Reference Types.

VALUE TYPES (Değer Tipleri):
- Veri doğrudan değişkende saklanır
- Atama yapıldığında değerin kopyası oluşturulur
- Değişiklikler orijinal veriyi etkilemez
- Stack'te saklanır (genellikle)
- Örnekler: int, float, bool, string, array, struct



REFERENCE TYPES (Referans Tipleri):
- Değişken, verinin bellekteki adresini tutar
- Atama yapıldığında adres kopyalanır, veri değil
- Değişiklikler orijinal veriyi etkiler
- Heap'te saklanır
- Örnekler: slice, map, channel, interface, func, pointer




POINTER'LAR:
- Hem value hem reference semantiği sağlayabilir
- *T türü bir pointer, T türündeki verinin adresini tutar
- & operatörü ile adres alınır
- * operatörü ile referans edilir (dereference)




MEMORY MANAGEMENT:
- Value types: Stack'te hızlı allocation/deallocation
- Reference types: Heap'te allocation, GC ile temizleme
- Escape Analysis: Compiler'ın stack vs heap kararı
- Performance: Value types genellikle daha hızlı



>>     BEST PRACTICES VE ÖNERİLER

Value ve reference semantiği seçerken aşağıdaki en iyi uygulamaları göz önünde bulundurun:


> VALUES TYPES NE ZAMAN KULLANILMALI:
- Küçük veri yapıları (primitives, küçük struct'lar)
- Immutable davranış istendiğinde (değişiklikler orijinali etkilemesin)
- Thread-safety önemli olduğunda (her goroutine kendi kopyasıyla çalışsın)
- Performans kritik basit operasyonlarda (stack allocation hızlıdır)
- API tasarımında dışarıya veri sızdırmak istenmediğinde (kopya verilir)

> REFERENCE TYPES NE ZAMAN KULLANILMALI:
- Büyük veri yapıları (slice, map, büyük struct'lar)
- Veri paylaşımı gerektiğinde (aynı veri üzerinde birden fazla yerde işlem yapılacaksa)
- Dinamik boyutlu veri yapıları (ör: slice, map, channel)
- Memory efficiency önemli olduğunda (büyük veri kopyalanmasın)
- Fonksiyonlar arası veri paylaşımı ve güncelleme gerektiğinde)

> POINTER NE ZAMAN KULLANILMALI:
- Büyük struct'ları fonksiyona geçirirken (kopya maliyetinden kaçınmak için)
- Orijinal veriyi değiştirmek istendiğinde (fonksiyon/method içinde kalıcı değişiklik)
- Nil kontrolü gerektiğinde (opsiyonel değerler için)
- Method receiver'larda kalıcı değişiklik için (ör: Update, Set gibi methodlar)
- Interface implementasyonunda pointer receiver gerekiyorsa (bazı interface'ler için)

> GENEL ÖNERİLER:
- Küçük ve immutable veri için value type tercih edin.")
- Büyük ve paylaşılan veri için reference type veya pointer kullanın.")
- API tasarımında, dışarıya reference type döndürüyorsanız dokümantasyonla belirtin.")
- Pointer ile çalışırken nil kontrollerini ihmal etmeyin.")
- Map ve slice gibi reference type'larda, shallow copy ve deep copy farkına dikkat edin.")
- Struct method receiver seçimini dikkatli yapın:")
    * Sadece okuyorsa value receiver, değiştiriyorsa pointer receiver kullanın.")
- Performans ve memory profiling için Go'nun araçlarını (pprof, escape analysis) kullanın.")



>>    YAYGIN HATALAR VE DİKKAT EDİLMESİ GEREKENLER

- Değer tipleri ile referans tipleri arasındaki farkı anlamamak
- Pointer kullanırken nil kontrollerini ihmal etmek
- Map ve slice gibi referans tiplerinde shallow copy ve deep copy farkını göz ardı etmek
- Struct method receiver seçimini yanlış yapmak (gereksiz kopyalar)
- Performans sorunlarını göz ardı etmek (escape analysis yapmamak)

*/

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// Demonstrasyon için örnek struct
type Person struct {
	Name string
	Age  int
	City string
}

// Value receiver (struct'ın kopyası üzerinde çalışır)
func (p Person) GetInfoValue() string {
	p.Age = 999 // Bu değişiklik orijinal struct'ı etkilemez
	return fmt.Sprintf("Name: %s, Age: %d, City: %s", p.Name, p.Age, p.City)
}

// Pointer receiver (orijinal struct üzerinde çalışır)
func (p *Person) UpdateAgePointer(newAge int) {
	p.Age = newAge // Bu değişiklik orijinal struct'ı etkiler
}

// Reference semantiği için örnek struct
type PersonRef struct {
	Name    string
	Age     int
	Friends []string       // Slice (reference type)
	Scores  map[string]int // Map (reference type)
}

func main() {
	fmt.Println("================================================")
	fmt.Println("    VALUE TYPES - TEMEL DAVRANIŞLAR")
	fmt.Println("================================================")

	// VALUE TYPES: int, float, bool, string, array, struct

	// 1. Primitive types (int, float, bool, string)
	fmt.Println("1. Primitive Types:")
	a := 42
	b := a // b, a'nın kopyasını alır
	b = 100
	fmt.Printf("   a = %d, b = %d (a değişmedi)\n", a, b)

	str1 := "Hello"
	str2 := str1 // str2, str1'in kopyasını alır
	str2 = "World"
	fmt.Printf("   str1 = %s, str2 = %s (str1 değişmedi)\n", str1, str2)

	// 2. Array (value type)
	fmt.Println("\n2. Array (Value Type):")
	arr1 := [3]int{1, 2, 3}
	arr2 := arr1 // arr2, arr1'in tam kopyasını alır
	arr2[0] = 999
	fmt.Printf("   arr1 = %v, arr2 = %v (arr1 değişmedi)\n", arr1, arr2)

	// 3. Struct (value type)
	fmt.Println("\n3. Struct (Value Type):")
	p1 := Person{Name: "Ali", Age: 25, City: "İstanbul"}
	p2 := p1 // p2, p1'in tam kopyasını alır
	p2.Age = 30
	fmt.Printf("   p1.Age = %d, p2.Age = %d (p1 değişmedi)\n", p1.Age, p2.Age)

	////
	////
	////

	fmt.Println("\n================================================")
	fmt.Println("    REFERENCE TYPES - TEMEL DAVRANIŞLAR")
	fmt.Println("================================================")

	// REFERENCE TYPES: slice, map, channel, interface, func, pointer

	// 1. Slice (reference type)
	fmt.Println("1. Slice (Reference Type):")
	slice1 := []int{1, 2, 3}
	slice2 := slice1 // slice2, slice1 ile aynı diziyi gösterir
	slice2[0] = 999
	fmt.Printf("   slice1 = %v, slice2 = %v (slice1 değişti!)\n", slice1, slice2)

	// 2. Map (reference type)
	fmt.Println("\n2. Map (Reference Type):")
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map1 // map2, map1 ile aynı veriyi gösterir
	map2["a"] = 999
	fmt.Printf("   map1 = %v, map2 = %v (map1 değişti!)\n", map1, map2)

	// 3. Pointer (reference semantiği)
	fmt.Println("\n3. Pointer (Reference Semantiği):")
	x := 42
	ptr1 := &x
	ptr2 := ptr1 // ptr2, ptr1 ile aynı adresi gösterir
	*ptr2 = 999
	fmt.Printf("   x = %d, *ptr1 = %d, *ptr2 = %d (x değişti!)\n", x, *ptr1, *ptr2)

	////
	////
	////

	fmt.Println("\n================================================")
	fmt.Println("    FUNCTION PARAMETER PASSING")
	fmt.Println("================================================")

	// Fonksiyonlara parametre geçirme davranışları

	fmt.Println("1. Value Types - Fonksiyona Geçirme:")
	num := 42
	fmt.Printf("    Fonksiyon öncesi: %d\n ", num)
	modifyInt(num)
	fmt.Printf("   Fonksiyon sonrası: %d (değişmedi)\n", num)

	fmt.Println("\n2. Reference Types - Fonksiyona Geçirme:")
	numbers := []int{1, 2, 3}
	fmt.Printf("    Fonksiyon öncesi: %v\n ", numbers)
	modifySlice(numbers)
	fmt.Printf("   Fonksiyon sonrası: %v (değişti!)\n", numbers)

	fmt.Println("\n3. Pointer - Fonksiyona Geçirme:")
	value := 42
	fmt.Printf("    Fonksiyon öncesi: %d\n ", value)
	modifyByPointer(&value)
	fmt.Printf("   Fonksiyon sonrası: %d (değişti!)\n", value)

	////
	////
	////

	fmt.Println("\n================================================")
	fmt.Println("    STRUCT METHODS: VALUE vs POINTER RECEIVER")
	fmt.Println("================================================")

	person := Person{Name: "Ayşe", Age: 28, City: "Ankara"}
	fmt.Printf("        Orijinal person : %+v\n", person)

	// Value receiver - orijinal struct değişmez
	info := person.GetInfoValue()
	fmt.Printf(" Value receiver sonrası : %+v\n", person)
	fmt.Printf("            Dönen bilgi : %s\n", info)

	// Pointer receiver - orijinal struct değişir
	person.UpdateAgePointer(35)
	fmt.Printf("Pointer receiver sonrası : %+v\n", person)

	////
	////
	////

	fmt.Println("\n================================================")
	fmt.Println("    COMPLEX SCENARIOS - KARIŞIK DURUMLAR")
	fmt.Println("================================================")

	// Struct içinde reference type fields
	personRef := PersonRef{
		Name:    "Mehmet",
		Age:     30,
		Friends: []string{"Ali", "Veli"},
		Scores:  map[string]int{"math": 85, "physics": 90},
	}

	// Struct kopyalandığında reference fields paylaşılır
	personRefCopy := personRef
	personRefCopy.Age = 35             // Bu değişmez (value field)
	personRefCopy.Friends[0] = "Ahmet" // Bu değişir (reference field)
	personRefCopy.Scores["math"] = 95  // Bu değişir (reference field)

	fmt.Printf("Orijinal: Age=%d, Friends=%v, Scores=%v\n",
		personRef.Age, personRef.Friends, personRef.Scores)
	fmt.Printf("Kopya   : Age=%d, Friends=%v, Scores=%v\n",
		personRefCopy.Age, personRefCopy.Friends, personRefCopy.Scores)

	////
	////
	////

	fmt.Println("\n================================================")
	fmt.Println("    MEMORY LAYOUT VE PERFORMANCE")
	fmt.Println("================================================")

	// Value types - Stack allocation (genellikle)
	valueType := 42
	fmt.Printf("Value type adresi: %p\n", &valueType)
	fmt.Printf("Value type boyutu: %d bytes\n", unsafe.Sizeof(valueType))

	// Reference types - Heap allocation
	refType := []int{1, 2, 3}
	fmt.Printf("Slice header adresi: %p\n", &refType)
	fmt.Printf("Slice header boyutu: %d bytes\n", unsafe.Sizeof(refType))
	fmt.Printf("Slice data adresi: %p\n", &refType[0])

	// Struct vs Pointer performance
	largeStruct := Person{Name: "Test", Age: 25, City: "Test"}
	structPtr := &largeStruct

	fmt.Printf("Struct adresi: %p\n", &largeStruct)
	fmt.Printf("Struct boyutu: %d bytes\n", unsafe.Sizeof(largeStruct))
	fmt.Printf("Pointer boyutu: %d bytes\n", unsafe.Sizeof(structPtr))

	////
	////
	////

	fmt.Println("\n================================================")
	fmt.Println("    INTERFACE VE POLYMORPHISM")
	fmt.Println("================================================")

	// Interface ile value vs pointer behavior
	var emptyInterface interface{}

	// Value type interface'e atanırsa
	emptyInterface = 42
	fmt.Printf("Interface value type: %v, type: %T\n", emptyInterface, emptyInterface)

	// Pointer type interface'e atanırsa
	emptyInterface = &valueType
	fmt.Printf("Interface pointer type: %v, type: %T\n", emptyInterface, emptyInterface)

	////
	////
	////

	fmt.Println("\n================================================")
	fmt.Println("    DEEP COPY vs SHALLOW COPY")
	fmt.Println("================================================")

	// Shallow copy (varsayılan davranış)
	original := PersonRef{
		Name:    "Original",
		Age:     25,
		Friends: []string{"Friend1", "Friend2"},
		Scores:  map[string]int{"test": 100},
	}

	shallowCopy := original
	shallowCopy.Name = "Shallow Copy"
	shallowCopy.Friends[0] = "Modified Friend"

	fmt.Printf("Original: %+v\n", original)
	fmt.Printf("Shallow Copy: %+v\n", shallowCopy)

	// Deep copy (manuel implementasyon)
	deepCopy := deepCopyPersonRef(original)
	deepCopy.Name = "Deep Copy"
	deepCopy.Friends[0] = "Deep Modified Friend"

	fmt.Printf("After Deep Copy Original: %+v\n", original)
	fmt.Printf("Deep Copy: %+v\n", deepCopy)

	////
	////
	////

	fmt.Println("\n================================================")
	fmt.Println("    COMMON PITFALLS VE HATALAR")
	fmt.Println("================================================")

	// 1. Slice append ile kapasite aşımı
	demonstrateSlicePitfall()

	// 2. Map nil kontrolü
	demonstrateMapPitfall()

	// 3. Pointer nil kontrolü
	demonstratePointerPitfall()

	// 4. Interface nil kontrolü
	demonstrateInterfacePitfall()

	fmt.Println("================================================")
}

// Value type fonksiyonu - orijinal değişmez
func modifyInt(x int) {
	x = 999
	fmt.Printf("   Fonksiyon içinde: %d\n", x)
}

// Reference type fonksiyonu - orijinal değişir
func modifySlice(s []int) {
	if len(s) > 0 {
		s[0] = 999
	}
	fmt.Printf("   Fonksiyon içinde: %v\n", s)
}

// Pointer fonksiyonu - orijinal değişir
func modifyByPointer(x *int) {
	*x = 999
	fmt.Printf("   Fonksiyon içinde: %d\n", *x)
}

// Deep copy implementation
func deepCopyPersonRef(original PersonRef) PersonRef {
	copy := PersonRef{
		Name:    original.Name,
		Age:     original.Age,
		Friends: make([]string, len(original.Friends)),
		Scores:  make(map[string]int),
	}

	// Friends slice'ını deep copy et
	for i, friend := range original.Friends {
		copy.Friends[i] = friend
	}

	// Scores map'ini deep copy et
	for key, value := range original.Scores {
		copy.Scores[key] = value
	}

	return copy
}

// Common pitfalls demonstration
func demonstrateSlicePitfall() {
	fmt.Println("\n1. Slice Append Kapasitesi Aşımı:")
	slice := make([]int, 2)
	slice[0] = 1
	slice[1] = 2
	fmt.Printf("   Orijinal slice: %v, cap: %d\n", slice, cap(slice))

	// Capacity aşıldığında yeni backing array oluşur
	newSlice := append(slice, 3)
	newSlice[0] = 999
	fmt.Printf("   Append sonrası orijinal: %v\n", slice)
	fmt.Printf("   Append sonrası yeni: %v, cap: %d\n", newSlice, cap(newSlice))
}

func demonstrateMapPitfall() {
	fmt.Println("\n2. Map Nil Kontrolü:")
	var m map[string]int
	fmt.Printf("   Nil map: %v\n", m)

	// Nil map'e yazma panic yaratır
	// m["key"] = 1 // Bu panic yaratır!

	// Doğru kullanım
	m = make(map[string]int)
	m["key"] = 1
	fmt.Printf("   İnitialized map: %v\n", m)
}

func demonstratePointerPitfall() {
	fmt.Println("\n3. Pointer Nil Kontrolü:")
	var ptr *int
	fmt.Printf("   Nil pointer: %v\n", ptr)

	// Nil pointer'ı dereference etmek panic yaratır
	// fmt.Println(*ptr) // Bu panic yaratır!

	// Doğru kullanım
	if ptr != nil {
		fmt.Printf("   Pointer value: %d\n", *ptr)
	} else {
		fmt.Println("   Pointer is nil")
	}
}

func demonstrateInterfacePitfall() {
	fmt.Println("\n4. Interface Nil Kontrolü:")
	var i interface{}
	fmt.Printf("   Nil interface: %v, type: %T\n", i, i)

	// Interface nil kontrolü
	if i == nil {
		fmt.Println("   Interface is nil")
	}

	// Interface type assertion
	if val, ok := i.(int); ok {
		fmt.Printf("   Interface contains int: %d\n", val)
	} else {
		fmt.Println("   Interface doesn't contain int")
	}

	// Reflection ile type kontrolü
	fmt.Printf("   Interface type via reflection: %v\n", reflect.TypeOf(i))
}
