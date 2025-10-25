/*
Author : mefamex
Date   : 2025-07-06
Title  : Empty Interface - Best Practices



BEST PRACTICES ÖZETİ:
=====================
- DO :
    - Gerçekten çok biçimli (polymorphic) kullanım için tercih edin (loglama, konfigürasyon, event sistemi)
    - Type assertion yaparken mutlaka 'comma ok' idiomunu kullanın
    - Birden fazla type assertion yerine type switch tercih edin
    - Tip güvenli (type-safe) sarmalayıcı fonksiyonlar yazın
    - Beklenen tipleri mutlaka dokümante edin
    - Nil değerleri açıkça kontrol edin
    - Yapısı bilinmeyen JSON/XML parse işlemlerinde kullanın
    - Go 1.18+ ile gelen generics'i alternatif olarak değerlendirin

- DON'T :
    - Tip güvenli yazılabilecek basit fonksiyonlarda kullanmayın
    - Type assertion hatalarını yok saymayın (panic oluşur)
    - Performans kritik kodlarda kullanmayın
    - Beklenen tipleri dokümante etmeyi unutmayın
    - Somut tipler daha iyi ise interface{} kullanmayın
    - API tasarımında gereksiz yere kullanmayın

BEST PRACTICES - DETAYLI AÇIKLAMA
==================================

Empty interface güçlü bir araç olmasına rağmen, yanlış kullanıldığında ciddi sorunlara yol açabilir. Bu bölümde production-ready kod yazmanın kurallarını ve pratiklerini bulacaksınız.


NE ZAMAN KULLANIN? -
    - JSON/XML parse (yapısı bilinmeyen veri)
    - Loglama sistemleri (değişken argümanlar)
    - Konfigürasyon yönetimi
    - Event sistemleri (esnek payload)
    - Plugin mimarileri
    - Generic yardımcılar (Go 1.18 öncesi)


NE ZAMAN KULLANMAYIN? 
    - Basit tip işlemleri
    - Performans kritik kodlar
    - Tipi bilinen durumlar
    - API tasarımında tip güvenliği gerekiyorsa
    - İş mantığında gereksiz soyutlama


GÜVENLİK KALIPLARI (SAFETY PATTERNS)
    ┌─────────────────────┬─────────────────────┐
    │ - TEHLİKELİ         │ - GÜVENLİ           │
    ├─────────────────────┼─────────────────────┤
    │ value.(string)      │ value, ok := v.(T)  │
    │ Hata kontrolsüz     │ Hataları kontrol et │
    │ Derin iç içe        │ Düz mantık kur      │
    │ Nil yok saymak      │ Nil kontrolü yap    │
    │ Performansı umursama│ Profil & optimize   │
    └─────────────────────┴─────────────────────┘


TASARIM KALIPLARI
    a) Wrapper Pattern - Tip güvenli arayüzler
    b) Validation Pattern - Girdi doğrulama
    c) Factory Pattern - Nesne oluşturma
    d) Observer Pattern - Olay yönetimi
    e) Strategy Pattern - Algoritma seçimi


HATA YÖNETİMİ STRATEJİLERİ
    - Zarif hata yönetimi (graceful degradation)
    - Varsayılan değer stratejileri
    - Hata yayılımı
    - Loglama ve izleme
    - Kurtarma mekanizmaları


PROD KONTROL LİSTESİ:
	- Her zaman "comma ok" idiomunu kullan
	- Beklenen tipleri açıkça dokümante et
	- Kapsamlı hata kontrolü ekle
	- Girdi doğrulaması yap
	- Performans etkisini izle
	- Doğru loglama uygula
	- Tüm tip yolları için test yaz
	- Mümkünse tip güvenli sarmalayıcı kullan


PERFORMANS İPUÇLARI:
	- Sıcak kod yollarında interface{}'dan kaçın
	- Soğuk kod yollarında esneklik için kullanılabilir
	- Sık tekrar eden type assertion'ları cache'le
	- Optimize etmeden önce profil çıkar
	- Alternatif: Go 1.18+ generics


CODE REVIEW KONTROL LİSTESİ:
	- Type assertion'lar güvenli mi?
	- Hata yönetimi yeterli mi?
	- Performans etkisi düşünüldü mü?
	- Dokümantasyon açık mı?
	- Tüm tip yolları test edildi mi?
	- Nil durumları kontrol edildi mi?
	- Bellek sızıntısı yok mu?


GEÇİŞ STRATEJİLERİ:
	- Eski kod: Kademeli refactor
	- Yeni kod: Başta tip güvenli yaz
	- Performans sorunu: Önce profil çıkar
	- Takım: Eğitim & rehberler


İZLEME & HATA AYIKLAMA:
	- Runtime tip dağılımı
	- Performans metrikleri
	- Tip bazlı hata oranları
	- Bellek kullanımı
	- Panic sıklığı


*/

package main

import (
	"fmt"
)

// Person struct - best practices örnekleri için
type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("\n=======================================")
	fmt.Println("    BEST PRACTICES")
	fmt.Println("=======================================")

	fmt.Println("BEST PRACTICES DEMONSTRATİONS")
	fmt.Println("=============================")

	// 1. Logging örneği
	fmt.Println("\n1. GOOD: Logging ile empty interface:")
	logWithLevel("INFO", "Server started successfully")
	logWithLevel("ERROR", 404)
	logWithLevel("DEBUG", Person{Name: "Test User", Age: 30})
	logWithLevel("WARN", map[string]interface{}{"error": "timeout", "duration": 5000})

	// 2. Configuration parsing
	fmt.Println("\n2. GOOD: Configuration parsing:")
	config := map[string]interface{}{
		"host":     "localhost",
		"port":     8080,
		"debug":    true,
		"timeout":  30.5,
		"features": []interface{}{"auth", "cache", "monitoring"},
		"database": map[string]interface{}{
			"driver": "postgres",
			"host":   "db.example.com",
		},
	}
	parseConfig(config)

	// 3. Event system
	fmt.Println("\n3. GOOD: Event system:")
	eventBus := NewEventBus()

	eventBus.Subscribe("user.login", func(data interface{}) {
		if user, ok := data.(Person); ok {
			fmt.Printf("User logged in: %s\n", user.Name)
		}
	})

	eventBus.Subscribe("user.logout", func(data interface{}) {
		if userID, ok := data.(int); ok {
			fmt.Printf("User ID %d logged out\n", userID)
		}
	})

	eventBus.Publish(Event{Type: "user.login", Data: Person{Name: "Alice", Age: 25}})
	eventBus.Publish(Event{Type: "user.logout", Data: 12345})

	// 4. Typed cache
	fmt.Println("\n4. GOOD: Typed cache with safe getters:")
	cache := NewTypedCache()
	cache.Set("username", "john_doe")
	cache.Set("user_id", 42)
	cache.Set("is_admin", true)

	if username, err := cache.GetString("username"); err == nil {
		fmt.Printf("Username: %s\n", username)
	}

	if userID, err := cache.GetInt("user_id"); err == nil {
		fmt.Printf("User ID: %d\n", userID)
	}

	if isAdmin, err := cache.GetBool("is_admin"); err == nil {
		fmt.Printf("Is Admin: %t\n", isAdmin)
	}

	// Hatalı durumlar
	if _, err := cache.GetString("user_id"); err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}

	// 5. Validation system
	fmt.Println("\n5. GOOD: Validation system:")
	fields := map[string]interface{}{
		"name":  "John",
		"age":   25,
		"email": "john@example.com",
	}

	rules := map[string]map[string]interface{}{
		"name": {
			"required":  true,
			"minLength": 2,
			"maxLength": 50,
		},
		"age": {
			"required": true,
			"min":      18,
			"max":      120,
		},
		"email": {
			"required":  true,
			"minLength": 5,
		},
	}

	for fieldName, value := range fields {
		if fieldRules, exists := rules[fieldName]; exists {
			if err := validateField(fieldName, value, fieldRules); err != nil {
				fmt.Printf("Validation error: %v\n", err)
			} else {
				fmt.Printf("Field '%s' is valid\n", fieldName)
			}
		}
	}

	// 6. JSON-like wrapper
	fmt.Println("\n6. GOOD: JSON wrapper with type safety:")
	jsonStr := NewJSONValue("hello")
	jsonNum := NewJSONValue(42.0) // JSON numbers are float64
	jsonBool := NewJSONValue(true)

	if str, ok := jsonStr.AsString(); ok {
		fmt.Printf("String value: %s\n", str)
	}

	if num, ok := jsonNum.AsInt(); ok {
		fmt.Printf("Number value: %d\n", num)
	}

	if b, ok := jsonBool.AsBool(); ok {
		fmt.Printf("Boolean value: %t\n", b)
	}

	// BAD vs GOOD örnekleri
	fmt.Println("\n7. BAD vs GOOD examples:")

	fmt.Println("BAD: Gereksiz empty interface:")
	badResult := badAddFunction(5, 10)
	fmt.Printf("Bad result: %v (Type: %T)\n", badResult, badResult)

	fmt.Println("GOOD: Type-safe alternative:")
	goodResult := goodAddFunction(5, 10)
	fmt.Printf("Good result: %v (Type: %T)\n", goodResult, goodResult)

}

// ============================================================
// BEST PRACTICES
// ============================================================

// GOOD EXAMPLE 1: Logging ile empty interface
func logWithLevel(level string, message interface{}) {
	fmt.Printf("[%s] %v\n", level, message)
}

// GOOD EXAMPLE 2: Configuration parsing
func parseConfig(config map[string]interface{}) {
	fmt.Println("Configuration parsing:")
	for key, value := range config {
		switch v := value.(type) {
		case string:
			fmt.Printf("  %s (string): %s\n", key, v)
		case int:
			fmt.Printf("  %s (int): %d\n", key, v)
		case float64:
			fmt.Printf("  %s (float): %.2f\n", key, v)
		case bool:
			fmt.Printf("  %s (bool): %t\n", key, v)
		case []interface{}:
			fmt.Printf("  %s (array): %v\n", key, v)
		case map[string]interface{}:
			fmt.Printf("  %s (object): %v\n", key, v)
		default:
			fmt.Printf("  %s (unknown): %v (%T)\n", key, v, v)
		}
	}
}

// GOOD EXAMPLE 3: Event system
type Event struct {
	Type string
	Data interface{}
}

type EventHandler func(interface{})

type EventBus struct {
	handlers map[string][]EventHandler
}

func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[string][]EventHandler),
	}
}

func (eb *EventBus) Subscribe(eventType string, handler EventHandler) {
	eb.handlers[eventType] = append(eb.handlers[eventType], handler)
}

func (eb *EventBus) Publish(event Event) {
	if handlers, exists := eb.handlers[event.Type]; exists {
		for _, handler := range handlers {
			handler(event.Data)
		}
	}
}

// GOOD EXAMPLE 4: Generic cache with type-safe getters
type TypedCache struct {
	data map[string]interface{}
}

func NewTypedCache() *TypedCache {
	return &TypedCache{
		data: make(map[string]interface{}),
	}
}

func (c *TypedCache) Set(key string, value interface{}) {
	c.data[key] = value
}

func (c *TypedCache) Get(key string) (interface{}, bool) {
	value, exists := c.data[key]
	return value, exists
}

// Type-safe getters
func (c *TypedCache) GetString(key string) (string, error) {
	value, exists := c.data[key]
	if !exists {
		return "", fmt.Errorf("key '%s' not found", key)
	}
	if str, ok := value.(string); ok {
		return str, nil
	}
	return "", fmt.Errorf("value is not string: %T", value)
}

func (c *TypedCache) GetInt(key string) (int, error) {
	value, exists := c.data[key]
	if !exists {
		return 0, fmt.Errorf("key '%s' not found", key)
	}
	if num, ok := value.(int); ok {
		return num, nil
	}
	return 0, fmt.Errorf("value is not int: %T", value)
}

func (c *TypedCache) GetBool(key string) (bool, error) {
	value, exists := c.data[key]
	if !exists {
		return false, fmt.Errorf("key '%s' not found", key)
	}
	if b, ok := value.(bool); ok {
		return b, nil
	}
	return false, fmt.Errorf("value is not bool: %T", value)
}

// BAD EXAMPLE 1: Gereksiz empty interface kullanımı
func badAddFunction(a, b interface{}) interface{} {
	// Bu fonksiyon tip güvenliği kaybeder ve mantıksız
	return fmt.Sprintf("%v + %v", a, b) // String döndürür!
}

// GOOD EXAMPLE: Type-safe alternative
func goodAddFunction(a, b int) int {
	return a + b
}

// BAD EXAMPLE 2: Type assertion without checking
func badProcessValue(value interface{}) {
	// Bu panic oluşturabilir!
	// str := value.(string) // Tehlikeli!
	// fmt.Println(str)
}

// GOOD EXAMPLE: Safe type assertion
func goodProcessValue(value interface{}) {
	if str, ok := value.(string); ok {
		fmt.Printf("String value: %s\n", str)
	} else {
		fmt.Printf("Not a string: %T\n", value)
	}
}

// GOOD EXAMPLE 5: Validation with empty interface
func validateField(fieldName string, value interface{}, rules map[string]interface{}) error {
	if value == nil {
		if required, ok := rules["required"].(bool); ok && required {
			return fmt.Errorf("field '%s' is required", fieldName)
		}
		return nil
	}

	// Type-specific validations
	switch v := value.(type) {
	case string:
		if minLen, ok := rules["minLength"].(int); ok && len(v) < minLen {
			return fmt.Errorf("field '%s' must be at least %d characters", fieldName, minLen)
		}
		if maxLen, ok := rules["maxLength"].(int); ok && len(v) > maxLen {
			return fmt.Errorf("field '%s' must be at most %d characters", fieldName, maxLen)
		}
	case int:
		if min, ok := rules["min"].(int); ok && v < min {
			return fmt.Errorf("field '%s' must be at least %d", fieldName, min)
		}
		if max, ok := rules["max"].(int); ok && v > max {
			return fmt.Errorf("field '%s' must be at most %d", fieldName, max)
		}
	}

	return nil
}

// GOOD EXAMPLE 6: JSON-like data structure with type safety
type JSONValue struct {
	value interface{}
}

func NewJSONValue(value interface{}) *JSONValue {
	return &JSONValue{value: value}
}

func (j *JSONValue) AsString() (string, bool) {
	str, ok := j.value.(string)
	return str, ok
}

func (j *JSONValue) AsInt() (int, bool) {
	// Handle both int and float64 (JSON numbers)
	switch v := j.value.(type) {
	case int:
		return v, true
	case float64:
		return int(v), true
	default:
		return 0, false
	}
}

func (j *JSONValue) AsBool() (bool, bool) {
	b, ok := j.value.(bool)
	return b, ok
}

func (j *JSONValue) AsArray() ([]interface{}, bool) {
	arr, ok := j.value.([]interface{})
	return arr, ok
}

func (j *JSONValue) AsObject() (map[string]interface{}, bool) {
	obj, ok := j.value.(map[string]interface{})
	return obj, ok
}
