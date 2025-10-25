/*
Author : mefamex
Date   : 2025-07-06
Title  : Empty Interface - Reflection ile Kullanım




REFLECTION İLE KULLANIM - DETAYLI AÇIKLAMA
===========================================

Reflection, Go'da runtime'da tip bilgilerini incelemek ve manipüle etmek
için kullanılan güçlü bir özelliktir. interface{} ile birlikte kullanıldığında
çok esnek sistemler oluşturulabilir.

KULLANIM ALANLARI:
- Serialization/Deserialization
- ORM libraries (database mapping)
- Configuration mapping/injection
- Generic utilities
- Validation libraries
- Testing frameworks
- Plugin systems
- Code generation tools

DİKKAT EDİLECEKLER:
- Performance overhead: Reflection, normal operasyonlardan yaklaşık 10 kat daha yavaştır.
- Runtime errors: Compile-time kontrolü yoktur, hatalar çalışma zamanında ortaya çıkar.
- Code clarity: Kodun anlaşılması ve bakımı zorlaşır.
- Debugging: Hataları bulmak ve ayıklamak zordur.

1. REFLECTION TEMEL KAVRAMLAR
	- reflect.TypeOf() - tip bilgisi
	- reflect.ValueOf() - değer bilgisi
	- reflect.Kind() - temel tip kategorisi
	- Interface{} bridge role'ü oynar

2. REFLECTION OPERATIONS
	┌─────────────────┬─────────────────────────────────┐
	│ Operation       │ Description                     │
	├─────────────────┼─────────────────────────────────┤
	│ TypeOf()        │ Get type information            │
	│ ValueOf()       │ Get value and manipulate        │
	│ Kind()          │ Get basic type category         │
	│ NumField()      │ Get struct field count          │
	│ Field(i)        │ Get struct field by index       │
	│ FieldByName()   │ Get struct field by name        │
	│ Interface()     │ Convert back to interface{}     │
	│ CanSet()        │ Check if value is settable      │
	│ Set()           │ Set new value                   │
	└─────────────────┴─────────────────────────────────┘

3. VALUE CATEGORIES (KIND)
	- Basic: Bool, Int, Float, String, Complex
	- Composite: Array, Slice, Map, Struct, Pointer
	- Functions: Func
	- Channels: Chan
	- Interfaces: Interface
	- Unsafe: UnsafePointer

4. STRUCT INTROSPECTION
	- Field discovery ve enumeration
	- Dynamic field access
	- Tag-based processing
	- Method discovery

REFLECTION PATTERNS:
1. Type Inspector Pattern
	func inspectType(v interface{}) {
		t := reflect.TypeOf(v)
		// analyze type
	}

2. Value Modifier Pattern
	func modifyValue(v interface{}) {
		rv := reflect.ValueOf(v).Elem()
		if rv.CanSet() {
			rv.Set(newValue)
		}
	}

3. Deep Copy Pattern
	func deepCopy(v interface{}) interface{} {
	   // recursive reflection-based copying
	}

PERFORMANCE CONSIDERATIONS:
- Reflection ~10-100x slower than direct access
- Memory overhead significant
- CPU intensive operations
- GC pressure increase

BEST PRACTICES:
- Cache reflection results when possible
- Use reflection sparingly in hot paths
- Validate types before reflection operations
- Handle panics with defer/recover
- Document reflection usage clearly

SECURITY CONSIDERATIONS:
- Reflection can access private fields
- Bypass type safety
- Potential for injection attacks
- Use validation before reflection ops

YAYGIN HATALAR:
- Performance impact'ini göz ardı etme
- Panic protection yapmama
- Overuse reflection (simple operations için)
- Memory leak'ler (caching issues)

*/

package main

import (
	"fmt"
	"reflect"
)

// Person struct - reflection örnekleri için
type Person struct {
	Name string
	Age  int
}

// Product struct - reflection örnekleri için
type Product struct {
	ID    int
	Name  string
	Price float64
}

func main() {
	fmt.Println("\n=======================================")
	fmt.Println("    REFLECTION İLE KULLANIM")
	fmt.Println("=======================================")

	values := []interface{}{
		Person{Name: "Reflection Test", Age: 30},
		[]int{1, 2, 3, 4, 5},
		map[string]int{"a": 1, "b": 2, "c": 3},
		"string value",
		42,
		Product{ID: 1, Name: "Laptop", Price: 1500.00},
		&Person{Name: "Pointer Person", Age: 25},
	}

	fmt.Println("Reflection ile değer inceleme:")
	for i, value := range values {
		fmt.Printf("\n--- Değer %d ---\n", i+1)
		inspectValue(value)
	}

	fmt.Println("\nStruct alanlarını dinamik olarak alma:")
	person := Person{Name: "Dinamik", Age: 25}
	fields := getStructFields(person)
	for name, value := range fields {
		fmt.Printf("%s: %v\n", name, value)
	}

	fmt.Println("\nStruct alanını dinamik olarak değiştirme:")
	personPtr := &Person{Name: "Eski İsim", Age: 20}
	fmt.Printf("Önce: %+v\n", personPtr)

	if err := setStructField(personPtr, "Name", "Yeni İsim"); err != nil {
		fmt.Printf("Hata: %v\n", err)
	} else {
		fmt.Printf("Sonra: %+v\n", personPtr)
	}

	if err := setStructField(personPtr, "Age", 35); err != nil {
		fmt.Printf("Hata: %v\n", err)
	} else {
		fmt.Printf("Final: %+v\n", personPtr)
	}

	fmt.Println("\nDeep copy örneği:")
	original := map[string]interface{}{
		"name":   "Original",
		"data":   []int{1, 2, 3},
		"nested": map[string]int{"a": 1},
	}

	copied := deepCopy(original)
	fmt.Printf("Original: %+v\n", original)
	fmt.Printf("Copied: %+v\n", copied)

	// Original'i değiştir
	original["name"] = "Modified"
	original["data"].([]int)[0] = 999

	fmt.Printf("After modification:\n")
	fmt.Printf("Original: %+v\n", original)
	fmt.Printf("Copied: %+v\n", copied)

}

// ============================================================
// REFLECTION İLE KULLANIM
// ============================================================

// inspectValue - Reflection ile değer inceleme
func inspectValue(value interface{}) {
	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)

	fmt.Printf("Değer: %v\n", value)
	fmt.Printf("Tip: %s\n", t)
	fmt.Printf("Kind: %s\n", v.Kind())

	switch v.Kind() {
	case reflect.Struct:
		fmt.Printf("Struct alanları:\n")
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fieldType := t.Field(i)
			fmt.Printf("  %s: %v (%s)\n", fieldType.Name, field.Interface(), fieldType.Type)
		}
	case reflect.Slice:
		fmt.Printf("Slice uzunluğu: %d\n", v.Len())
		for i := 0; i < v.Len() && i < 5; i++ { // İlk 5 eleman
			elem := v.Index(i)
			fmt.Printf("  [%d]: %v (%s)\n", i, elem.Interface(), elem.Type())
		}
	case reflect.Map:
		fmt.Printf("Map key sayısı: %d\n", len(v.MapKeys()))
		for i, key := range v.MapKeys() {
			if i >= 5 {
				break // İlk 5 key
			}
			value := v.MapIndex(key)
			fmt.Printf("  %v: %v\n", key.Interface(), value.Interface())
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Println("Nil pointer")
		} else {
			fmt.Println("Pointer değeri:")
			inspectValue(v.Elem().Interface())
		}
	default:
		fmt.Printf("Basit tip: %v\n", value)
	}
}

// getStructFields - Struct alanlarını dinamik olarak al
func getStructFields(value interface{}) map[string]interface{} {
	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)

	if v.Kind() != reflect.Struct {
		return nil
	}

	fields := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		fields[fieldType.Name] = field.Interface()
	}

	return fields
}

// setStructField - Struct alanını dinamik olarak değiştir
func setStructField(value interface{}, fieldName string, newValue interface{}) error {
	v := reflect.ValueOf(value)

	// Pointer kontrolü
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("value must be a pointer to struct")
	}

	// Struct kontrolü
	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("value must be a struct")
	}

	// Alan bulma
	field := v.FieldByName(fieldName)
	if !field.IsValid() {
		return fmt.Errorf("field '%s' not found", fieldName)
	}

	// Değiştirilebilir kontrolü
	if !field.CanSet() {
		return fmt.Errorf("field '%s' cannot be set", fieldName)
	}

	// Tip kontrolü ve atama
	newVal := reflect.ValueOf(newValue)
	if field.Type() != newVal.Type() {
		return fmt.Errorf("type mismatch: expected %s, got %s", field.Type(), newVal.Type())
	}

	field.Set(newVal)
	return nil
}

// callMethod - Dinamik method çağrısı
func callMethod(value interface{}, methodName string, args ...interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(value)
	method := v.MethodByName(methodName)

	if !method.IsValid() {
		return nil, fmt.Errorf("method '%s' not found", methodName)
	}

	// Argüman hazırlama
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}

	// Method çağrısı
	results := method.Call(in)

	// Sonuçları interface{} slice'ına çevir
	out := make([]interface{}, len(results))
	for i, result := range results {
		out[i] = result.Interface()
	}

	return out, nil
}

// deepCopy - Reflection ile deep copy
func deepCopy(value interface{}) interface{} {
	original := reflect.ValueOf(value)
	copy := reflect.New(original.Type()).Elem()
	copyRecursive(original, copy)
	return copy.Interface()
}

// copyRecursive - Recursive kopyalama
func copyRecursive(original, copy reflect.Value) {
	switch original.Kind() {
	case reflect.Ptr:
		if !original.IsNil() {
			copy.Set(reflect.New(original.Elem().Type()))
			copyRecursive(original.Elem(), copy.Elem())
		}
	case reflect.Struct:
		for i := 0; i < original.NumField(); i++ {
			copyRecursive(original.Field(i), copy.Field(i))
		}
	case reflect.Slice:
		copy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
		for i := 0; i < original.Len(); i++ {
			copyRecursive(original.Index(i), copy.Index(i))
		}
	case reflect.Map:
		copy.Set(reflect.MakeMap(original.Type()))
		for _, key := range original.MapKeys() {
			originalValue := original.MapIndex(key)
			copyValue := reflect.New(originalValue.Type()).Elem()
			copyRecursive(originalValue, copyValue)
			copy.SetMapIndex(key, copyValue)
		}
	default:
		copy.Set(original)
	}
}
