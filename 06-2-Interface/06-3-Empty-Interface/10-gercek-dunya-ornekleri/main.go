/*
Author : mefamex
Date   : 2025-07-06
Title  : Empty Interface - Gerçek Dünya Örnekleri



GERÇEK DÜNYA ÖRNEKLERİ - DETAYLI AÇIKLAMA :)
==========================================

Bu bölümde interface{}'ün production ortamlarında nasıl kullanıldığını
göreceksiniz. Gerçek dünya senaryoları, endüstri kalıpları ve büyük ölçekli çözümler ele alınmaktadır.


KULLANIM ALANLARI & SEKTÖREL ÖRNEKLER:
- Veri işleme (Data Processing):
    - JSON/XML API cevapları
    - CSV/Excel dosya okuma
    - Veritabanı sorgu sonuçları
    - Mesaj kuyruğu yükleri (message queue payloads)
- Framework geliştirme (Framework Development):
    - HTTP middleware sistemleri
    - Bağımlılık enjeksiyon konteynerleri (dependency injection)
    - Plugin mimarileri
    - Şablon motorları (template engines)
- Sistem bileşenleri (System Components):
    - Konfigürasyon yöneticileri
    - Cache sistemleri
    - Event bus'lar
    - Loglama altyapıları
- Modern alternatifler (Modern Alternatives):
    - Go 1.18+ Generics
    - Kod üretimi (code generation)
    - Tip güvenli sarmalayıcılar (type-safe wrappers)
    - Şema doğrulama (schema validation)


KURUMSAL LOG SİSTEMLERİ (Enterprise Logging Systems)
    - Yapılandırılmış loglama (JSON formatı)
    - Çoklu log seviyesi ve filtreleme
    - Bağlama duyarlı loglama (context-aware)
    - Performans optimize loglayıcılar
    - Merkezi log toplama (centralized aggregation)


KONFİGÜRASYON YÖNETİMİ (Configuration Management)
    - Ortama göre konfigürasyonlar (environment-based)
    - Canlı güncelleme (hot-reload)
    - Tip güvenli getter fonksiyonlar (type-safe getters)
    - Doğrulama ve varsayılanlar (validation & defaults)
    - Gizli veri yönetimi (secrets management)


OLAY TABANLI MİMARİLER (Event-Driven Architectures)
    - Mesaj kuyruğu entegrasyonu (message brokers)
    - Olay kaynaklı desenler (event sourcing)
    - CQRS uygulamaları
    - Mikroservisler arası iletişim (microservices communication)
    - Saga desenleri


CACHE SİSTEMLERİ (Cache Systems)
    - Çok katmanlı önbellekleme (multi-level caching)
    - Yaşam süresi desteği (TTL - Time To Live)
    - Cache geçersiz kılma stratejileri (invalidation)
    - Dağıtık cache (distributed caching)
    - Bellek verimli depolama (memory-efficient storage)


PLUGIN MİMARİLERİ (Plugin Architectures)
    - Dinamik eklenti yükleme (dynamic plugin loading)
    - Interface tabanlı sözleşmeler (interface-based contracts)
    - Güvenli sandbox çalıştırma
    - Eklenti yaşam döngüsü yönetimi (plugin lifecycle)
    - Konfigürasyon enjeksiyonu (configuration injection)


MIDDLEWARE SİSTEMLERİ
    - HTTP middleware zincirleri (chains)
    - Kimlik doğrulama ve yetkilendirme (authentication & authorization)
    - Oran sınırlama (rate limiting)
    - İstek/yanıt dönüştürme (request/response transformation)
    - Devre kesici desenleri (circuit breaker patterns)



PERFORMANS KALIPLARI (Performance Patterns):
1. Nesne havuzu (Object Pooling) - GC baskısını azaltır (reduce GC pressure)
2. Tembel yükleme (Lazy Loading) - Veriyi ihtiyaç halinde yükle
3. Önbellekleme (Caching) - Sık erişilen veriyi sakla
4. Toplu işleme (Batching) - Birden fazla öğeyi birlikte işle
5. Akış işleme (Streaming) - Büyük veri kümelerini yönet


ÖLÇEKLENEBİLİRLİK DİKKATLERİ (Scalability Considerations):
- Yatay ölçekleme stratejileri (horizontal scaling)
- Durumlu veriyle yük dengeleme (load balancing with stateful data)
- Veritabanı sharding etkileri (database sharding)
- Büyük ölçekte bellek yönetimi (memory management at scale)
- İzleme ve alarm sistemleri (monitoring & alerting)


GÜVENLİK ETKİLERİ (Security Implications):
- Girdi doğrulama kritik (input validation critical)
- Serileştirme açıkları (serialization vulnerabilities)
- Tip karışıklığı saldırıları (type confusion attacks)
- Bellek tüketimiyle DoS (memory exhaustion DoS)
- Enjeksiyon saldırı vektörleri (injection attack vectors)


DAĞITIM STRATEJİLERİ (Deployment Strategies):
- Mavi-yeşil dağıtımlar (blue-green deployments)
- Canary yayınları (canary releases)
- Özellik bayrakları entegrasyonu (feature flags integration)
- A/B test desteği (A/B testing support)
- Geri alma prosedürleri (rollback procedures)


İZLEME & GÖZLEMLENEBİLİRLİK (Monitoring & Observability):
- Metrik toplama (metrics collection)
- Dağıtık izleme (distributed tracing)
- Hata takibi (error tracking)
- Performans profilleme (performance profiling)
- İş zekası (business intelligence)


PROD HİKAYELERİ (Production War Stories):
- "JSON Bellek Sızıntısı" (The JSON Memory Leak) - Büyük dizileri unmarshal etmek
- "Tip Assertion Panic" (The Type Assertion Panic) - Nil kontrolü eksikliği
- "Performans Uçurumu" (The Performance Cliff) - Hot path'te reflection
- "Konfigürasyon Kabusu" (The Configuration Nightmare) - Tip uyumsuzlukları

*/

package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// Person struct - real world örnekleri için
type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("\n=======================================")
	fmt.Println("    GERÇEK DÜNYA ÖRNEKLERİ")
	fmt.Println("=======================================")

	realWorldExamples()
}

// ============================================================
// GERÇEK DÜNYA ÖRNEKLERİ
// ============================================================

// 1. LOGGER - Production-ready logging system
type Logger struct {
	level string
	mu    sync.Mutex
}

func NewLogger(level string) *Logger {
	return &Logger{level: level}
}

func (l *Logger) Log(level string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] [%s] ", timestamp, level)

	for i, arg := range args {
		if i > 0 {
			fmt.Print(" ")
		}

		switch v := arg.(type) {
		case string:
			fmt.Print(v)
		case error:
			fmt.Printf("ERROR: %v", v)
		case map[string]interface{}:
			if jsonBytes, err := json.Marshal(v); err == nil {
				fmt.Print(string(jsonBytes))
			} else {
				fmt.Printf("%+v", v)
			}
		default:
			fmt.Printf("%v", v)
		}
	}
	fmt.Println()
}

func (l *Logger) Info(args ...interface{}) {
	l.Log("INFO", args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.Log("ERROR", args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.Log("DEBUG", args...)
}

// 2. CACHE - Thread-safe generic cache system
type Cache struct {
	data map[string]interface{}
	mu   sync.RWMutex
	ttl  map[string]time.Time
}

func NewCache() *Cache {
	cache := &Cache{
		data: make(map[string]interface{}),
		ttl:  make(map[string]time.Time),
	}

	// TTL cleanup goroutine
	go cache.cleanupExpired()

	return cache
}

func (c *Cache) Set(key string, value interface{}) {
	c.SetWithTTL(key, value, 0) // 0 = no expiration
}

func (c *Cache) SetWithTTL(key string, value interface{}, duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value
	if duration > 0 {
		c.ttl[key] = time.Now().Add(duration)
	} else {
		delete(c.ttl, key) // No expiration
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// Check TTL
	if expiry, exists := c.ttl[key]; exists && time.Now().After(expiry) {
		delete(c.data, key)
		delete(c.ttl, key)
		return nil, false
	}

	value, exists := c.data[key]
	return value, exists
}

func (c *Cache) GetString(key string) (string, error) {
	value, exists := c.Get(key)
	if !exists {
		return "", fmt.Errorf("key '%s' not found", key)
	}

	if str, ok := value.(string); ok {
		return str, nil
	}

	return "", fmt.Errorf("value is not string: %T", value)
}

func (c *Cache) GetInt(key string) (int, error) {
	value, exists := c.Get(key)
	if !exists {
		return 0, fmt.Errorf("key '%s' not found", key)
	}

	switch v := value.(type) {
	case int:
		return v, nil
	case float64:
		return int(v), nil
	default:
		return 0, fmt.Errorf("value is not number: %T", value)
	}
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, key)
	delete(c.ttl, key)
}

func (c *Cache) cleanupExpired() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for key, expiry := range c.ttl {
			if now.After(expiry) {
				delete(c.data, key)
				delete(c.ttl, key)
			}
		}
		c.mu.Unlock()
	}
}

// 3. EVENT BUS - Pub/Sub system for loose coupling
type EventBus struct {
	handlers map[string][]EventHandler
	mu       sync.RWMutex
}

type EventHandler func(interface{})

type Event struct {
	Type      string
	Data      interface{}
	Timestamp time.Time
}

func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[string][]EventHandler),
	}
}

func (eb *EventBus) Subscribe(eventType string, handler EventHandler) {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	eb.handlers[eventType] = append(eb.handlers[eventType], handler)
}

func (eb *EventBus) Publish(eventType string, data interface{}) {
	eb.mu.RLock()
	handlers := make([]EventHandler, len(eb.handlers[eventType]))
	copy(handlers, eb.handlers[eventType])
	eb.mu.RUnlock()

	event := Event{
		Type:      eventType,
		Data:      data,
		Timestamp: time.Now(),
	}

	// Execute handlers asynchronously
	for _, handler := range handlers {
		go func(h EventHandler) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("Event handler panic: %v\n", r)
				}
			}()
			h(event.Data)
		}(handler)
	}
}

func (eb *EventBus) Unsubscribe(eventType string) {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	delete(eb.handlers, eventType)
}

// 4. CONFIGURATION MANAGER - Dynamic configuration system
type ConfigManager struct {
	config map[string]interface{}
	mu     sync.RWMutex
}

func NewConfigManager() *ConfigManager {
	return &ConfigManager{
		config: make(map[string]interface{}),
	}
}

func (cm *ConfigManager) LoadFromJSON(jsonData string) error {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		return fmt.Errorf("JSON parse error: %v", err)
	}

	cm.mu.Lock()
	defer cm.mu.Unlock()

	cm.config = data
	return nil
}

func (cm *ConfigManager) GetString(key string) (string, error) {
	value, err := cm.getValue(key)
	if err != nil {
		return "", err
	}

	if str, ok := value.(string); ok {
		return str, nil
	}

	return "", fmt.Errorf("key '%s' has wrong type: %T", key, value)
}

func (cm *ConfigManager) GetInt(key string) (int, error) {
	value, err := cm.getValue(key)
	if err != nil {
		return 0, err
	}

	switch val := value.(type) {
	case int:
		return val, nil
	case float64:
		return int(val), nil
	default:
		return 0, fmt.Errorf("key '%s' has wrong type: %T", key, value)
	}
}

func (cm *ConfigManager) GetBool(key string) (bool, error) {
	value, err := cm.getValue(key)
	if err != nil {
		return false, err
	}

	if b, ok := value.(bool); ok {
		return b, nil
	}

	return false, fmt.Errorf("key '%s' has wrong type: %T", key, value)
}

func (cm *ConfigManager) getValue(key string) (interface{}, error) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	value, exists := cm.config[key]
	if !exists {
		return nil, fmt.Errorf("key '%s' not found", key)
	}

	return value, nil
}

// 5. MIDDLEWARE SYSTEM - HTTP middleware with interface{}
type Context struct {
	Data map[string]interface{}
	mu   sync.RWMutex
}

func NewContext() *Context {
	return &Context{
		Data: make(map[string]interface{}),
	}
}

func (c *Context) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Data[key] = value
}

func (c *Context) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists := c.Data[key]
	return value, exists
}

func (c *Context) GetString(key string) (string, bool) {
	if value, exists := c.Get(key); exists {
		if str, ok := value.(string); ok {
			return str, true
		}
	}
	return "", false
}

type Middleware func(*Context) error

type MiddlewareChain struct {
	middlewares []Middleware
}

func NewMiddlewareChain() *MiddlewareChain {
	return &MiddlewareChain{}
}

func (mc *MiddlewareChain) Use(middleware Middleware) {
	mc.middlewares = append(mc.middlewares, middleware)
}

func (mc *MiddlewareChain) Execute(ctx *Context) error {
	for _, middleware := range mc.middlewares {
		if err := middleware(ctx); err != nil {
			return err
		}
	}
	return nil
}

// 6. PLUGIN SYSTEM - Dynamic plugin loading
type Plugin interface {
	Name() string
	Execute(args ...interface{}) (interface{}, error)
}

type PluginManager struct {
	plugins map[string]Plugin
	mu      sync.RWMutex
}

func NewPluginManager() *PluginManager {
	return &PluginManager{
		plugins: make(map[string]Plugin),
	}
}

func (pm *PluginManager) Register(name string, plugin Plugin) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	pm.plugins[name] = plugin
}

func (pm *PluginManager) Execute(name string, args ...interface{}) (interface{}, error) {
	pm.mu.RLock()
	plugin, exists := pm.plugins[name]
	pm.mu.RUnlock()

	if !exists {
		return nil, fmt.Errorf("plugin '%s' not found", name)
	}

	return plugin.Execute(args...)
}

// Example plugin implementation
type MathPlugin struct{}

func (mp *MathPlugin) Name() string {
	return "math"
}

func (mp *MathPlugin) Execute(args ...interface{}) (interface{}, error) {
	if len(args) < 3 {
		return nil, fmt.Errorf("math plugin requires at least 3 args: operation, a, b")
	}

	operation, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("first argument must be operation string")
	}

	a, aOk := args[1].(float64)
	b, bOk := args[2].(float64)

	if !aOk || !bOk {
		return nil, fmt.Errorf("second and third arguments must be numbers")
	}

	switch operation {
	case "add":
		return a + b, nil
	case "subtract":
		return a - b, nil
	case "multiply":
		return a * b, nil
	case "divide":
		if b == 0 {
			return nil, fmt.Errorf("division by zero")
		}
		return a / b, nil
	default:
		return nil, fmt.Errorf("unknown operation: %s", operation)
	}
}

func realWorldExamples() {
	fmt.Println("GERÇEK DÜNYA KULLANIM ÖRNEKLERİ")
	fmt.Println("===============================")

	// 1. Logger örneği
	fmt.Println("\n1. PRODUCTION LOGGER SİSTEMİ")
	fmt.Println("-----------------------------")
	logger := NewLogger("INFO")

	logger.Info("Application started")
	logger.Error("Database connection failed", fmt.Errorf("timeout"))
	logger.Debug("User data:", map[string]interface{}{
		"user_id": 12345,
		"action":  "login",
		"ip":      "192.168.1.100",
	})

	// 2. Cache sistemi
	fmt.Println("\n2. GENERIC CACHE SİSTEMİ")
	fmt.Println("-------------------------")
	cache := NewCache()

	// Different types in cache
	cache.Set("user_name", "john_doe")
	cache.Set("user_id", 42)
	cache.SetWithTTL("session_token", "abc123xyz", 5*time.Second)
	cache.Set("user_data", map[string]interface{}{
		"email":      "john@example.com",
		"role":       "admin",
		"last_login": time.Now(),
	})

	// Type-safe retrieval
	if username, err := cache.GetString("user_name"); err == nil {
		fmt.Printf("Username: %s\n", username)
	}

	if userID, err := cache.GetInt("user_id"); err == nil {
		fmt.Printf("User ID: %d\n", userID)
	}

	// Test TTL
	fmt.Printf("Session token exists: %t\n", func() bool {
		_, exists := cache.Get("session_token")
		return exists
	}())

	// 3. Event Bus sistemi
	fmt.Println("\n3. EVENT BUS SİSTEMİ")
	fmt.Println("--------------------")
	eventBus := NewEventBus()

	// Subscribe to events
	eventBus.Subscribe("user.created", func(data interface{}) {
		if user, ok := data.(map[string]interface{}); ok {
			fmt.Printf("New user created: %s\n", user["name"])
		}
	})

	eventBus.Subscribe("order.placed", func(data interface{}) {
		if order, ok := data.(map[string]interface{}); ok {
			fmt.Printf("Order placed: #%v for $%.2f\n", order["id"], order["amount"])
		}
	})

	// Publish events
	eventBus.Publish("user.created", map[string]interface{}{
		"id":    1001,
		"name":  "Alice Johnson",
		"email": "alice@example.com",
	})

	eventBus.Publish("order.placed", map[string]interface{}{
		"id":       "ORD-001",
		"amount":   99.99,
		"user_id":  1001,
		"products": []string{"laptop", "mouse"},
	})

	// Give time for async handlers
	time.Sleep(100 * time.Millisecond)

	// 4. Configuration Manager
	fmt.Println("\n4. CONFIGURATION MANAGER")
	fmt.Println("------------------------")
	configManager := NewConfigManager()

	configJSON := `{
		"app_name": "MyApp",
		"port": 8080,
		"debug": true,
		"database": {
			"host": "localhost",
			"port": 5432,
			"name": "myapp_db"
		},
		"features": ["auth", "logging", "metrics"]
	}`

	if err := configManager.LoadFromJSON(configJSON); err != nil {
		fmt.Printf("Config load error: %v\n", err)
	} else {
		if appName, err := configManager.GetString("app_name"); err == nil {
			fmt.Printf("App Name: %s\n", appName)
		}

		if port, err := configManager.GetInt("port"); err == nil {
			fmt.Printf("Port: %d\n", port)
		}

		if debug, err := configManager.GetBool("debug"); err == nil {
			fmt.Printf("Debug Mode: %t\n", debug)
		}
	}

	// 5. Middleware sistemi
	fmt.Println("\n5. MIDDLEWARE SİSTEMİ")
	fmt.Println("---------------------")
	chain := NewMiddlewareChain()

	// Authentication middleware
	chain.Use(func(ctx *Context) error {
		ctx.Set("user_id", 12345)
		ctx.Set("authenticated", true)
		fmt.Println("✓ Authentication middleware")
		return nil
	})

	// Logging middleware
	chain.Use(func(ctx *Context) error {
		if userID, exists := ctx.Get("user_id"); exists {
			fmt.Printf("✓ Logging middleware: User %v\n", userID)
		}
		return nil
	})

	// Authorization middleware
	chain.Use(func(ctx *Context) error {
		if auth, ok := ctx.GetString("role"); !ok || auth != "admin" {
			ctx.Set("role", "user") // Default role
		}
		fmt.Println("✓ Authorization middleware")
		return nil
	})

	// Execute middleware chain
	ctx := NewContext()
	if err := chain.Execute(ctx); err != nil {
		fmt.Printf("Middleware error: %v\n", err)
	}

	// 6. Plugin sistemi
	fmt.Println("\n6. PLUGIN SİSTEMİ")
	fmt.Println("-----------------")
	pluginManager := NewPluginManager()

	// Register plugins
	pluginManager.Register("math", &MathPlugin{})

	// Execute plugin operations
	operations := []struct {
		op string
		a  float64
		b  float64
	}{
		{"add", 10, 5},
		{"subtract", 10, 5},
		{"multiply", 10, 5},
		{"divide", 10, 5},
	}

	for _, op := range operations {
		if result, err := pluginManager.Execute("math", op.op, op.a, op.b); err == nil {
			fmt.Printf("%.1f %s %.1f = %.2f\n", op.a, op.op, op.b, result)
		} else {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
