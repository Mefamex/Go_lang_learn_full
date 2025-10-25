/*
Author : mefamex
Date   : 2025-07-06
Title  : Empty Interface - Yaygın Hatalar

YAYGIN HATALAR & EN İYİ PRATİKLER ÖZETİ:
========================================
- DO:
    - Type assertion yaparken mutlaka 'comma ok' idiomunu kullanın
    - Nil değerleri açıkça kontrol edin
    - Performans kritik kodlarda interface{} ve reflection'dan kaçının
    - Tip güvenli (type-safe) sarmalayıcı fonksiyonlar yazın
    - Beklenen tipleri mutlaka dokümante edin
    - Yapısı bilinmeyen JSON/XML parse işlemlerinde dikkatli olun
    - Go 1.18+ ile gelen generics'i alternatif olarak değerlendirin
    - Her adımda error handling yapın
    - Reflection kullanıyorsanız pointer, struct ve field kontrollerini unutmayın
    - Testlerde edge case, nil ve yanlış tipleri de kapsayın

- DON'T:
    - Tip güvenli yazılabilecek basit fonksiyonlarda interface{} kullanmayın
    - Type assertion hatalarını yok saymayın (panic oluşur)
    - Performans kritik kodlarda interface{} kullanmayın
    - Beklenen tipleri dokümante etmeyi unutmayın
    - Somut tipler daha iyi ise interface{} kullanmayın
    - API tasarımında gereksiz yere interface{} kullanmayın
    - JSON parse sonrası sayıları doğrudan int'e cast etmeyin (float64 gelir)
    - Reflection'da pointer/struct/field kontrolü yapmadan erişim yapmayın



AÇIKLAMA & KRİTİK NOTLAR:
========================
Bu dosyada, Go'da interface{} kullanımıyla ilgili en sık yapılan hatalar ve bunların güvenli, performanslı ve okunabilir şekilde nasıl önlenebileceği örneklerle gösterilmektedir. Her örnekte "BAD" (yanlış/pratikte sorunlu) ve "GOOD" (doğru/önerilen) kullanım karşılaştırmalı olarak sunulmuştur.



Hataları Önleme Prensipleri:
1. TYPE SAFETY FIRST
    - Mümkünse concrete type kullanın
    - interface{} kullanıyorsanız, type assertion'ları güvenli yapın
    - API'lerinizi type-safe tasarlayın
2. DEFENSIVE PROGRAMMING
    - Her type assertion'da 'comma ok' kullanın
    - Nil değerleri her zaman kontrol edin
    - Error handling'i atlamayın
3. PERFORMANCE CONSCIOUSNESS
    - Critical path'te interface{} kullanmayın
    - Reflection'ı az ve dikkatli kullanın
    - Hot path'ler için benchmark yapın
4. CLEAR DOCUMENTATION
    - interface{} alan fonksiyonları dokümante edin
    - Beklenen tipleri belirtin
    - Error conditions'ları açıklayın
5. THOROUGH TESTING
    - Edge case'leri test edin
    - Nil değerlerle test edin
    - Yanlış tiplerle test edin
6. MODERN ALTERNATIVES
    - Go 1.18+ için generics'i değerlendirin
    - Type-safe wrapper'lar oluşturun

Notlar:
- JSON parse işlemlerinde gerçek bir unmarshal işlemi yapılmamış, örnekler simüle edilmiştir. Gerçek uygulamada "encoding/json" paketi ile json.Unmarshal kullanılmalıdır.
- Nil interface ve nil pointer farkı, Go'da sıkça karıştırılan bir konudur. Doğru nil kontrolü için reflect.ValueOf().IsNil() kullanılmalıdır.
- Type assertion'larda her zaman 'comma ok' idiomu kullanılmalı, aksi halde panic oluşabilir.
- Go 1.18+ ile generics desteği geldiğinden, yeni projelerde mümkünse generics tercih edilmelidir.
*/

package main

import (
	"fmt"
	"reflect"
)

// Person struct - hata örnekleri için
type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("\n=======================================")
	fmt.Println("    YAYGIN HATALAR")
	fmt.Println("=======================================")

	////
	////
	////

	// HATA 1: Panic oluşturan type assertion
	fmt.Println("\n1. PANIC OLUŞTURAN TYPE ASSERTION")
	fmt.Println("----------------------------------")
	fmt.Println("BAD: Panic oluşturan type assertion")

	var value interface{} = 42

	// BAD - Bu panic oluşturacak
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("- Panic yakalandı: %v\n", r)
			}
		}()

		// Bu satır panic oluşturacak çünkü 42 string değil
		str := value.(string) // PANIC!
		fmt.Printf("Bu satır çalışmayacak: %s\n", str)
	}()

	// GOOD - Güvenli type assertion
	if str, ok := value.(string); ok {
		fmt.Printf("- String değer: %s\n", str)
	} else {
		fmt.Printf("- Değer string değil: %T (değer: %v)\n", value, value)
	}

	fmt.Println("ÇÖZÜM: Her zaman 'comma ok' idiomunu kullanın!")

	////
	////
	////

	// HATA 2: Nil interface karmaşası
	fmt.Println("\n2. NIL INTERFACE KARMAŞASI")
	fmt.Println("---------------------------")
	fmt.Println("BAD: Nil interface karışıklığı")

	// Case 1: Gerçek nil interface
	var nilInterface interface{} = nil
	fmt.Printf("nilInterface == nil: %t\n", nilInterface == nil)

	// Case 2: Nil pointer ama non-nil interface
	var nilPointer *Person = nil
	var nonNilInterface interface{} = nilPointer
	fmt.Printf("nonNilInterface == nil: %t (YANLIŞ!)\n", nonNilInterface == nil)

	// Case 3: Nil check'in doğru yolu
	fmt.Printf("nonNilInterface value is nil: %t\n",
		reflect.ValueOf(nonNilInterface).IsNil())

	// GOOD - Doğru nil kontrolü
	fmt.Println("\n- Doğru nil kontrolü:")
	checkNilCorrectly := func(value interface{}) bool {
		if value == nil {
			return true
		}
		v := reflect.ValueOf(value)
		switch v.Kind() {
		case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func:
			return v.IsNil()
		default:
			return false
		}
	}

	fmt.Printf("nilInterface is nil: %t\n", checkNilCorrectly(nilInterface))
	fmt.Printf("nonNilInterface is nil: %t\n", checkNilCorrectly(nonNilInterface))

	fmt.Println("ÇÖZÜM: Reflect.ValueOf().IsNil() kullanın!")

	////
	////
	////

	// HATA 3: Performance overhead'i göz ardı etme
	fmt.Println("\n3. PERFORMANCE OVERHEAD'İ GÖZ ARDI ETME")
	fmt.Println("---------------------------------------")
	fmt.Println("BAD: Performance overhead'i göz ardı etme")

	// BAD - Critical path'te interface{} kullanımı
	badCalculateSum := func(numbers []interface{}) int {
		sum := 0
		for _, num := range numbers {
			if n, ok := num.(int); ok {
				sum += n
			}
		}
		return sum
	}

	// GOOD - Concrete type kullanımı
	goodCalculateSum := func(numbers []int) int {
		sum := 0
		for _, num := range numbers {
			sum += num
		}
		return sum
	}

	// Test verileri
	concreteNumbers := []int{1, 2, 3, 4, 5}
	interfaceNumbers := []interface{}{1, 2, 3, 4, 5}

	fmt.Printf("Bad approach result: %d\n", badCalculateSum(interfaceNumbers))
	fmt.Printf("Good approach result: %d\n", goodCalculateSum(concreteNumbers))

	fmt.Println("ÇÖZÜM: Performance-critical kod'da concrete type kullanın!")

	////
	////
	////

	// HATA 4: Type safety kaybı
	fmt.Println("\n4. TYPE SAFETY KAYBI")
	fmt.Println("--------------------")
	fmt.Println("BAD: Type safety kaybı")

	// BAD - Type safety kaybeden fonksiyon
	badMath := func(operation string, a, b interface{}) interface{} {
		switch operation {
		case "add":
			// Bu çok tehlikeli - tip kontrolü yok!
			return fmt.Sprintf("%v + %v", a, b) // String döndürür!
		case "multiply":
			// Bu da runtime'da patlar
			if numA, ok := a.(int); ok {
				if numB, ok := b.(int); ok {
					return numA * numB
				}
			}
			return "error" // Tutarsız dönüş tipi!
		default:
			return nil
		}
	}

	// GOOD - Type-safe alternatif
	goodMath := func(operation string, a, b int) (int, error) {
		switch operation {
		case "add":
			return a + b, nil
		case "multiply":
			return a * b, nil
		default:
			return 0, fmt.Errorf("unknown operation: %s", operation)
		}
	}

	// Test
	badResult := badMath("add", 5, 10)
	fmt.Printf("- Bad result: %v (Type: %T)\n", badResult, badResult)

	goodResult, err := goodMath("add", 5, 10)
	if err == nil {
		fmt.Printf("- Good result: %v (Type: %T)\n", goodResult, goodResult)
	}

	fmt.Println("ÇÖZÜM: Type-safe API tasarlayın!")

	////
	////
	////

	// HATA 5: Gereksiz interface{} kullanımı
	fmt.Println("\n5. GEREKSIZ INTERFACE{} KULLANIMI")
	fmt.Println("---------------------------------")
	fmt.Println("BAD: Gereksiz interface{} kullanımı")

	// BAD - Basit durum için interface{} kullanımı
	badPrint := func(values ...interface{}) {
		for _, v := range values {
			fmt.Printf("%v ", v)
		}
		fmt.Println()
	}

	// GOOD - Specific case için specific function
	printInts := func(values ...int) {
		for _, v := range values {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}

	printStrings := func(values ...string) {
		for _, v := range values {
			fmt.Printf("%s ", v)
		}
		fmt.Println()
	}

	// BAD kullanım
	fmt.Print("- Bad: ")
	badPrint(1, "hello", 3.14, true)

	// GOOD kullanım
	fmt.Print("- Good ints: ")
	printInts(1, 2, 3, 4, 5)
	fmt.Print("- Good strings: ")
	printStrings("hello", "world", "go")

	fmt.Println("ÇÖZÜM: Gerçekten gerekli olmadıkça interface{} kullanmayın!")

	////
	////
	////

	// HATA 6: Runtime type errors
	fmt.Println("\n6. RUNTIME TYPE ERRORS")
	fmt.Println("----------------------")
	fmt.Println("BAD: Runtime type errors'ı handle etmemek")

	// BAD - Error handling yok
	badProcess := func(data interface{}) {
		// Bu runtime'da patlar
		str := data.(string)
		fmt.Printf("Processed: %s\n", str)
	}

	// GOOD - Proper error handling
	goodProcess := func(data interface{}) error {
		str, ok := data.(string)
		if !ok {
			return fmt.Errorf("expected string, got %T", data)
		}
		fmt.Printf("Processed: %s\n", str)
		return nil
	}

	// Test
	testData := []interface{}{"valid string", 42, true, nil}

	fmt.Println("- Bad approach:")
	for _, data := range testData {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("Panic for %v: %v\n", data, r)
				}
			}()
			badProcess(data)
		}()
	}

	fmt.Println("\n- Good approach:")
	for _, data := range testData {
		if err := goodProcess(data); err != nil {
			fmt.Printf("Error for %v: %v\n", data, err)
		}
	}

	fmt.Println("ÇÖZÜM: Her zaman error handling yapın!")

	////
	////
	////

	// HATA 7: Reflection ile ilgili hatalar
	fmt.Println("\n7. REFLECTION İLE İLGİLİ HATALAR")
	fmt.Println("--------------------------------")
	fmt.Println("BAD: Reflection ile ilgili hatalar")

	// BAD - Reflection abuse
	badGetField := func(obj interface{}, fieldName string) interface{} {
		v := reflect.ValueOf(obj)
		// Pointer kontrolü yok!
		// Struct kontrolü yok!
		// Field var mı kontrolü yok!
		field := v.FieldByName(fieldName)
		return field.Interface() // PANIC!
	}

	// GOOD - Safe reflection
	goodGetField := func(obj interface{}, fieldName string) (interface{}, error) {
		v := reflect.ValueOf(obj)

		// Pointer kontrolü
		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				return nil, fmt.Errorf("object is nil")
			}
			v = v.Elem()
		}

		// Struct kontrolü
		if v.Kind() != reflect.Struct {
			return nil, fmt.Errorf("object is not a struct")
		}

		// Field kontrolü
		field := v.FieldByName(fieldName)
		if !field.IsValid() {
			return nil, fmt.Errorf("field '%s' not found", fieldName)
		}

		return field.Interface(), nil
	}

	// Test
	person := Person{Name: "Test", Age: 30}

	fmt.Println("- Bad reflection (will panic):")
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Panic: %v\n", r)
			}
		}()
		result := badGetField(&person, "InvalidField")
		fmt.Printf("Result: %v\n", result)
	}()

	fmt.Println("\n- Good reflection:")
	if name, err := goodGetField(person, "Name"); err == nil {
		fmt.Printf("Name: %v\n", name)
	}

	if _, err := goodGetField(person, "InvalidField"); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("ÇÖZÜM: Reflection'da her adımı kontrol edin!")

	////
	////
	////

	// HATA 8: JSON unmarshaling hataları
	fmt.Println("\n8. JSON UNMARSHALING HATALARI")
	fmt.Println("-----------------------------")
	fmt.Println("BAD: JSON unmarshaling hataları")

	jsonData := `{"name": "John", "age": 30, "score": 95.5}`

	// BAD - Type assumption yok
	badParseJSON := func(data string) {
		var result interface{}
		fmt.Printf("JSON: %s\n", data)

		// Parse başarılı olduğunu varsayıyor
		// Type assertion kontrolü yok
		obj := result.(map[string]interface{})
		name := obj["name"].(string)
		age := obj["age"].(int) // Bu hata! JSON'da number float64'tür

		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}

	// GOOD - Proper JSON handling
	goodParseJSON := func(data string) error {
		var result interface{}
		if err := fmt.Errorf("simulated unmarshal"); err != nil {
			return fmt.Errorf("JSON parse error: %v", err)
		}

		obj, ok := result.(map[string]interface{})
		if !ok {
			return fmt.Errorf("expected JSON object")
		}

		name, ok := obj["name"].(string)
		if !ok {
			return fmt.Errorf("name field is not string")
		}

		// JSON'da numbers float64'tür
		ageFloat, ok := obj["age"].(float64)
		if !ok {
			return fmt.Errorf("age field is not number")
		}
		age := int(ageFloat)

		fmt.Printf("Name: %s, Age: %d\n", name, age)
		return nil
	}

	fmt.Println("- Bad JSON parsing:")
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Panic: %v\n", r)
			}
		}()
		badParseJSON(jsonData)
	}()

	fmt.Println("\n- Good JSON parsing:")
	if err := goodParseJSON(jsonData); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("ÇÖZÜM: JSON'da number'lar float64'tür, kontrol edin!")
}
