/*
Author : mefamex
Date   : 2025-07-06
Title  : Empty Interface - JSON ile Çalışma



JSON İLE ÇALIŞMA - DETAYLI AÇIKLAMA
====================================

Empty interface, JSON verilerle çalışırken en güçlü araçlardan biridir.
Özellikle bilinmeyen yapıdaki JSON verilerini işlemek için idealdir.

1. JSON UNMARSHAL TYPES
    var data interface{}
    json.Unmarshal([]byte(jsonString), &data)

    Possible results:
    - JSON Object → map[string]interface{}
    - JSON Array → []interface{}
    - JSON String → string
    - JSON Number → float64 (default)
    - JSON Boolean → bool
    - JSON null → nil

2. JSON VERİ TİPLERİ MAPPING
    ┌─────────────────┬─────────────────────────┐
    │ JSON Type       │ Go Type                 │
    ├─────────────────┼─────────────────────────┤
    │ object {}       │ map[string]interface{}  │
    │ array []        │ []interface{}           │
    │ string "text"   │ string                  │
    │ number 42       │ float64                 │
    │ number 42.5     │ float64                 │
    │ boolean true    │ bool                    │
    │ null            │ nil                     │
    └─────────────────┴─────────────────────────┘

3. NESTED JSON STRUCTURES
    - Deep object traversal
    - Array içinde object'ler
    - Mixed type arrays
    - Dynamic key-value structures

KULLANIM ALANLARI:
- REST API responses (unknown structure)
- Configuration files (flexible schemas)
- Webhook payloads (varying formats)
- Log aggregation (mixed data sources)
- NoSQL document parsing
- Microservices communication

PROCESSING PATTERNS:
1. Type Assertion Pattern
    if obj, ok := data.(map[string]interface{}); ok {
       // process object
    }

2. Type Switch Pattern
    switch v := data.(type) {
    case map[string]interface{}:
       // handle object
    case []interface{}:
       // handle array
    }

3. Recursive Processing
    func processJSON(data interface{}) {
        switch v := data.(type) {
        case map[string]interface{}:
            for key, value := range v {
                processJSON(value) // recursive call
            }
        }
    }

BEST PRACTICES:
- Her zaman error kontrolü yapın
- Type assertion ile güvenli tip dönüşümü
- Nested JSON için recursive işleme
- Null değerler için extra kontrol
- Large JSON için streaming parser kullanın
- Schema validation ekleyin

PERFORMANCE NOTES:
- interface{} kullanımı ~2-3x memory overhead
- Large JSON'lar için json.Decoder kullanın
- Known structure'lar için struct binding tercih edin
- Caching ile repeated parsing'i önleyin

YAYGIN HATALAR:
- Number types karıştırma (her zaman float64)
- Nil kontrolü yapmama
- Deep nesting'de memory leak
- Error handling atlamak

*/

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("\n=======================================")
	fmt.Println("    JSON İLE ÇALIŞMA")
	fmt.Println("=======================================")

	// JSON Object
	jsonObject := `{
		"name": "Akif",
		"age": 30,
		"salary": 58000.03,
		"active": true,
		"skills": ["Go", "Python", "JavaScript"],
		"address": {
			"city": "Ankara",
			"country": "Turkey",
			"coordinates": {
				"lat": 39.9334,
				"lng": 32.8597
			}
		},
		"projects": [
			{
				"name": "Project A",
				"status": "active",
				"team_size": 5
			},
			{
				"name": "Project B", 
				"status": "completed",
				"team_size": 3
			}
		]
	}`

	fmt.Println("JSON Object işleme:")
	processJSONData(jsonObject)

	// JSON Array
	jsonArray := `[
		{"name": "Ali", "age": 25, "role": "developer"},
		{"name": "Veli", "age": 30, "role": "designer"},
		{"name": "Ayşe", "age": 28, "role": "manager"}
	]`

	fmt.Println("\nJSON Array işleme:")
	processJSONData(jsonArray)

	// Simple JSON values
	jsonValues := []string{
		`"simple string"`,
		`42`,
		`3.14159`,
		`true`,
		`false`,
		`null`,
	}

	fmt.Println("\nJSON Values işleme:")
	for _, jsonValue := range jsonValues {
		processJSONData(jsonValue)
	}

	fmt.Println("\nJSON'dan alan çıkarma:")
	if name, err := safeJSONExtract(jsonObject, "name"); err == nil {
		fmt.Printf("Name: %v (%T)\n", name, name)
	}

	if age, err := safeJSONExtract(jsonObject, "age"); err == nil {
		fmt.Printf("Age: %v (%T)\n", age, age)
	}

	if skills, err := safeJSONExtract(jsonObject, "skills"); err == nil {
		fmt.Printf("Skills: %v (%T)\n", skills, skills)
	}

	fmt.Println("\nİç içe JSON işleme:")
	processNestedJSON(jsonObject)

	fmt.Println("\nJSON işleme best practices:")
	fmt.Println("1. Her zaman error kontrolü yapın")
	fmt.Println("2. Type assertion ile güvenli tip dönüşümü")
	fmt.Println("3. Nested JSON için recursive işleme")
	fmt.Println("4. Null değerler için extra kontrol")
	fmt.Println("5. Büyük JSON dosyaları için streaming parser kullanın")
}

// ============================================================
// JSON İLE ÇALIŞMA
// ============================================================

// processJSONData - JSON verilerini empty interface ile işle
func processJSONData(jsonData string) {
	var data interface{}
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		fmt.Printf("JSON parse hatası: %v\n", err)
		return
	}

	fmt.Printf("JSON Tip: %T\n", data)

	switch v := data.(type) {
	case map[string]interface{}:
		fmt.Println("JSON Object:")
		for key, value := range v {
			fmt.Printf("  %s: %v (%T)\n", key, value, value)
		}
	case []interface{}:
		fmt.Println("JSON Array:")
		for i, value := range v {
			fmt.Printf("  [%d]: %v (%T)\n", i, value, value)
		}
	default:
		fmt.Printf("JSON Value: %v (%T)\n", v, v)
	}
}

// extractFromJSON - JSON'dan belirli alanları çıkar
func extractFromJSON(jsonData string, key string) interface{} {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		return nil
	}
	return data[key]
}

// safeJSONExtract - Güvenli JSON alan çıkarma
func safeJSONExtract(jsonData string, key string) (interface{}, error) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		return nil, fmt.Errorf("JSON parse error: %v", err)
	}

	value, exists := data[key]
	if !exists {
		return nil, fmt.Errorf("key '%s' not found", key)
	}

	return value, nil
}

// processNestedJSON - İç içe JSON işleme
func processNestedJSON(jsonData string) {
	var data interface{}
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		fmt.Printf("JSON parse hatası: %v\n", err)
		return
	}

	processJSONValue(data, 0)
}

// processJSONValue - JSON değerini recursively işle
func processJSONValue(value interface{}, depth int) {
	indent := ""
	for i := 0; i < depth; i++ {
		indent += "  "
	}

	switch v := value.(type) {
	case map[string]interface{}:
		fmt.Printf("%sObject {\n", indent)
		for key, val := range v {
			fmt.Printf("%s  %s: ", indent, key)
			if isSimpleType(val) {
				fmt.Printf("%v (%T)\n", val, val)
			} else {
				fmt.Println()
				processJSONValue(val, depth+2)
			}
		}
		fmt.Printf("%s}\n", indent)
	case []interface{}:
		fmt.Printf("%sArray [\n", indent)
		for i, val := range v {
			fmt.Printf("%s  [%d]: ", indent, i)
			if isSimpleType(val) {
				fmt.Printf("%v (%T)\n", val, val)
			} else {
				fmt.Println()
				processJSONValue(val, depth+2)
			}
		}
		fmt.Printf("%s]\n", indent)
	default:
		fmt.Printf("%s%v (%T)\n", indent, v, v)
	}
}

// isSimpleType - Basit tip kontrolü
func isSimpleType(value interface{}) bool {
	switch value.(type) {
	case string, int, float64, bool, nil:
		return true
	default:
		return false
	}
}
