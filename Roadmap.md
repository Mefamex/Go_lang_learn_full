# Go Programming Dili Öğrenme Yol Haritası

Go (Golang) dili için kapsamlı bir roadmap (yol haritası)

Go version 1.24.4 içerisinde (zamanında) oluşturulmuştur.

Olabildiğince kapsamlı ve eksiksiz olarak başlıklara yer verilmiştir. 

Eğer bu dilde yeniyseniz, konular birbiri ile bağlantılı olduğu için en az 2 kere tüm konuları okumanızı öneririm.

<br>

> Created by : [@mefamex](https://github.com/Mefamex) 
> Date  : 2025 july (Temmuz)

<br>


## İçindekiler

- [1. Temeller](#1-temeller)
- [2. Değişkenler ve Veri Tipleri](#2-değişkenler-ve-veri-tipleri)
- [3. Kontrol Yapıları](#3-kontrol-yapıları)
- [4. Fonksiyonlar](#4-fonksiyonlar)
- [5. Koleksiyonlar ve Veri Yapıları](#5-koleksiyonlar-ve-veri-yapıları)
- [6. Metodlar ve Interface'ler](#6-metodlar-ve-interfaceler)
- [7. Concurrency](#7-concurrency)
- [8. Error Handling](#8-error-handling)
- [9. Testing](#9-testing)
- [10. Standart Kütüphane](#10-standart-kütüphane)
- [11. Dosya İşlemleri](#11-dosya-i̇şlemleri)
- [12. Web Geliştirme](#12-web-geliştirme)
- [13. Database](#13-database)
- [14. Best Practices](#14-best-practices)
- [15. Tooling](#15-tooling)
- [16. İleri Seviye Konular](#16-i̇leri-seviye-konular)
- [17. Deployment](#17-deployment)
- [18. Ecosystem](#18-ecosystem)

<br>


## 1. Temeller

- Go dilinin tarihçesi ve felsefesi
- Basitlik, okunabilirlik, "az ama öz" felsefesini içselleştirme
- Unutulacaklar: Class'lar, kalıtım, try-catch, aşırı generic kullanımı
- Kabul edilecekler: Kompozisyon, açık hata yönetimi, goroutine ve channel
- Kurulum: Go toolchain, VS Code (Go eklentisi) veya GoLand kurulumu
- `go version`, `go env` komutlarını anlama
- "Hello World" ve temel syntax
- Paket yapısı ve modüller
- GOPATH ve çalışma alanı

<br>


## 2. Değişkenler ve Veri Tipleri

- Temel veri tipleri (int, float, string, bool, rune, byte)
- Değişken tanımlama (`var`, `:=`) ve kapsam
- Sabitler (`const`)
- Tip çıkarımı (Type inference) ve dönüşümü
- Zero values kavramı

<br>


## 3. Kontrol Yapıları

- if/else ifadeleri
- switch ve case yapıları
- for döngüleri
- range kullanımı
- break, continue, goto

<br>


## 4. Fonksiyonlar

- Fonksiyon tanımlama ve çağırma
- Parametreler ve dönüş değerleri
- Multiple return values (özellikle `(sonuc, error)` kalıbı)
- Variadic functions (değişken sayıda parametre)
- Closure functions ve anonim fonksiyonlar
- Defer, panic, recover
- Methods

<br>


## 5. Koleksiyonlar ve Veri Yapıları

- Arrays
- Slices
- Maps
- Structs
- Pointers
- Memory management, garbage collector hakkında temel bilgi
- Value vs Reference types

<br>


## 6. Metodlar ve Interface'ler

- Method tanımlama ve receiver kullanımı
- Pointer vs Value receivers
- Interface kavramı, implicit implementation
- Empty interface (`interface{}`)
- Type assertions
- Type switches
- Standart kütüphane interface’leri (`io.Reader`, `io.Writer` vb.)
- Polimorfizm ve "Accept interfaces, return structs" prensibi


<br>

## 7. Concurrency

- Goroutines: `go` anahtar kelimesi, hafif iş parçacıkları
- Channels: oluşturma, gönderme, alma, kapatma
- Buffered ve unbuffered channels
- Select statement ile çoklu kanal işlemleri
- Mutex ve senkronizasyon (`sync.Mutex`)
- WaitGroups (`sync.WaitGroup`)
- Worker pools
- Context package: iptal, timeout ve yaşam döngüsü yönetimi
- Channels ve mutexlerin birlikte kullanımı ve farkları
- Race condition tespiti ve `go run -race` kullanımı


<br>

## 8. Error Handling

- Error interface
- Hata oluşturma (`errors.New`, `fmt.Errorf`) ve yönetme
- Custom error tipleri ve sentinel errors
- Error wrapping ve unwrapping (Go 1.13+)
- Error handling best practices
- panic ve recover mekanizmalarının kontrollü kullanımı
- Üçüncü parti hata yönetim kütüphaneleri

<br>

## 9. Testing

- Unit testing (`testing` paketi)
- Table-driven tests
- Benchmarking
- Test coverage analizi
- Test suites
- Mocking
- Integration testing
- Fuzz testing
- `go test` komut satırı araçları ve parametreleri


<br>

## 10. Standart Kütüphane

- fmt, io, os, time paketleri
- encoding/json, encoding/xml
- net/http: HTTP server ve client
- database/sql ve driver kullanımı
- context ve sync paketleri
- log ve log yönetimi
- crypto paketleri (şifreleme, hash)

<br>


## 11. Dosya İşlemleri

- Dosya okuma/yazma
- JSON işlemleri
- XML işlemleri
- Klasör yönetimi
- Dosya yönetimi
- Log yönetimi ve rotasyonu

<br>


## 12. Web Geliştirme

- HTTP server oluşturma (`net/http`)
- Routing
- Middleware
- Templates
- Static files
- RESTful API geliştirme
- gRPC ve Protobuf
- WebSocket
- Web frameworkleri (Gin, Echo vb.)


<br>

## 13. Veritabanı Yönetimi

- SQL bağlantıları
- CRUD operasyonları
- Transaction yönetimi
- Migration yönetimi
- ORM kullanımı (GORM vb.)
- NoSQL veritabanları (MongoDB, Redis)
- Connection pooling


<br>

## 14. Best Practices

- Code organization
- Naming conventions
- Error handling patterns
- Logging standards and practices
- Configuration management (env, config files...)
- Dependency injection
- Code review ve static analysis tools
- Project structure
- Security best practices


<br>

## 15. Tooling

- go fmt
- go vet
- go mod
- go test
- Profiling (`pprof`)
- Debugging
- go generate ve otomatik kod üretimi
- CI/CD pipeline entegrasyonu


<br>

## 16. İleri Seviye Konular
- Generics (Go 1.18+) ve constraint tasarımı
- Reflection (`reflect` paketi)
- Unsafe package ve bellek yönetimi
- CGO ile C kütüphaneleri entegrasyonu
- Go Assembly (Plan 9 Assembly)
- Build tags ve compiler direktifleri
- Performans optimizasyonu ve profiling
- İleri concurrency pattern’leri
- Go runtime ve scheduler mimarisi
- Modül yönetimi ileri seviye (proxy, cache, versiyonlama)
- Network programming (TCP/UDP, socket)
- Internationalization (i18n) ve localization (l10n)
- Sistem programlama (process, sinyal yönetimi)
- Embedded Go kullanımı
- Bulut yerel teknolojiler ve mikroservis mimarisi
- Profiling ve tracing araçları (OpenTelemetry, Jaeger)
- İleri test teknikleri (mock generation, fuzz testing)
- Compiler optimizasyonları ve custom toolchain kullanımı
- Paralel algoritmalar ve veri yapıları
- Plugin sistemi ve dinamik modüller
- Go hafıza modeli ve data race önleme



### 16.1 İleri Seviye Konular (Gelişmiş)

- Generics (Go 1.18+)
    - Tip-güvenli, yeniden kullanılabilir fonksiyonlar ve veri yapıları
    - Constraint (kısıtlama) tasarımı: ön tanımlı ve özel kısıtlamalar
    - Type Approximation (~T) kullanımı
    - Jeneriklerin sınırları ve performans etkileri
- Reflection (`reflect` paketi)
    - Dinamik tip kontrolü, runtime tip bilgisi
    - Struct alanlarına erişim ve değiştirme
    - Dinamik metot çağırma
    - Reflection performans maliyeti ve ne zaman kaçınılmalı
- Code Generation (Kod Üretimi)
    - `go generate` kullanımı
    - Otomatik kod üretimi araçları (stringer, mockgen vb.)
    - Meta-programming teknikleri ve template tabanlı kod üretimi
- Unsafe Package
    - Bellek üzerinde düşük seviyeli işlemler
    - unsafe.Pointer, uintptr ve pointer aritmetiği
    - Struct hafıza düzeni: unsafe.Sizeof, Alignof, Offsetof
    - Güvenlik riskleri ve kullanım uyarıları
- CGO ve Harici Fonksiyonlar (FFI)
    - C kütüphaneleri ile entegrasyon
    - Veri tipi dönüşümleri ve hafıza yönetimi
    - Callback mekanizmaları
    - Performans ve build kısıtlamaları
- Go Assembly (Plan 9 Assembly)
    - Performans kritik kod yazımı
    - Plan 9 assembly sözdizimi
    - Go kodu ile assembly entegrasyonu
- Derleyici ve Bağlayıcı Direktifleri
    - `//go:noinline`, `//go:nosplit` gibi direktifler
    - `//go:linkname` ile özel erişim
    - Build tags (`//go:build`) ve platforma özel derleme
- Performans Optimizasyonu ve Profiling
    - `pprof` ile CPU, bellek, blok ve mutex profil analizleri
    - Go Execution Tracer ile detaylı izleme
    - Escape Analysis ve Bounds Check Elimination
    - Garbage Collector tuning (GOGC, debug paketleri)
- İleri Concurrency Patternleri
    - Pipeline, fan-in/fan-out, worker pool tasarımları
    - Context package derin kullanımı (iptal, timeout, değer taşıma)
- Go Runtime ve Scheduler Mimarisi
    - M (OS thread), P (processor), G (goroutine) modeli
    - Work-stealing scheduler mekanizması
    - Goroutine durumları ve sistem çağrıları yönetimi
- İleri Error Handling Teknikleri
    - Custom error tipleri
    - Error wrapping ve sentinel errors
- Modül ve Paket Yönetimi İleri Düzey
    - Versiyonlama, proxy kullanımı, modül cache yönetimi
    - Monorepo ve multi-module yapılar
- Network Programming İleri Konuları
    - TCP/UDP server ve client uygulamaları
    - Socket programlama
- Security Best Practices
    - Güvenli kod yazımı
    - Şifreleme ve veri doğrulama
- Internationalization (i18n) ve Localization (l10n)
    - Çoklu dil desteği ve yerelleştirme araçları (`golang.org/x/text` paketi)
    - Tarih, saat, para birimi formatlaması ve çeviri yönetimi
- Go ile Sistem Programlama
    - Dosya sistemleri, process yönetimi, sinyal handling
- Embedded Go
    - Gömülü sistemlerde Go kullanımı ve kısıtlamalar
- Bulut Yerel Teknolojiler ve Mikroservis Mimarisi
    - Kubernetes client, cloud SDK entegrasyonları
    - Service discovery, load balancing, circuit breaker pattern
- Profiling ve Tracing Araçları
    - OpenTelemetry, Jaeger, Zipkin entegrasyonları
- İleri Test Teknikleri
    - Mock generation, integration testing, fuzz testing
- CI/CD Otomasyonu
    - Sürekli entegrasyon ve dağıtım süreçleri
- Compiler Optimizasyonları ve Custom Toolchain Kullanımı
- Go ile Paralel Algoritmalar ve Veri Yapıları
- Plugin Paketi ve Dinamik Modüller
    - Runtime’da paylaşımlı kütüphane yükleme
    - Limitasyonlar ve uyumluluk
- Go Hafıza Modeli (Memory Model)
    - "Happens Before" ilişkisi
    - Data race önleme ve lock-free programlama

<br>

## 17. Deployment

- Binary building ve cross compilation
- Docker container oluşturma ve yönetimi
- Kubernetes ve cloud native deployment
- CI/CD pipeline kurulumu ve otomasyonu
- Monitoring (Prometheus, Grafana)
- Log yönetimi ve merkezi loglama
- Performans optimizasyonu ve tuning
- Güvenlik ve erişim yönetimi


## 18. Ecosystem

- Popüler frameworkler
- Yaygın kütüphaneler
- Tool ve utility'ler
- Community resources
- Open source projeler

