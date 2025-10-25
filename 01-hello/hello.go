/*
author: mefamex
date  : 2025-07-01
title : Hello World in Go
*/

// features of GO

// Programming language
// Open Source
// Cross-platform
// from Google 2009 -> 2012(1.0)
// Static
// Compiled, Fast compilation
// Simple and readable syntax
// High performance
// Built-in concurrency support (goroutines & channels)
// Powerful standard library
// Automatic memory management (Garbage Collection)
// Modern development tools (go mod, go fmt, go test, etc.)
// Interfaces
// Type Embedding
// First-class Functions
// Reflection
// Context Package
// Error Handling
// Generics (v1.18)
// Fuzzing
// Static Binary Compilation
// Defer Keyword
// Race Detector
// Advanced Package Management

package main

import "fmt"

func main() {
	fmt.Println("Hello, World!") // gömülü olarak println fonksiyonu da var ama fmt.print** kullanmak daha kapsamlı ve destekli
}

// Proje oluştururken

// - proje klasörünü oluştur ve içinde komut satırını çalıştır.
// - go mod init [folderName/projectName]
// - {code} >> [name].go içine kodlarını yaz
// - go run . ile çalıştır.

////
////
////

// Genel olarak kullanılan bazı GO CONSOLE komutları:

// - go help -------------------> user guide
// - go env --------------------> ortam değişkenleri ve yapılandırma
// - go mod init [module_name] -> modül oluşturmak ve bağımlılık yönetimini başlatmak
// - go get [package_name] -----> harici paketleri eklemek
// - go fmt [file.go] ----------> kodun okunabilirliğini ve standartlara uygunluğunu sağlamak
// - go build [file.go] --------> derlenip çalıştırılabilir dosya haline getirilmesi
// - go run [file.go] ----------> Derlemeden hızlıca kodu çalıştırmak ve test etmek
// - go test -------------------> Yazılan test dosyalarını çalıştırıp, test sonuçlarını görmek
// - go install [package] ------> Paketi derleyip Go çalışma ortamına yüklemek için, genellikle projeyi sistem genelinde kullanılabilir hale getirmek
// - go clean ------------------> Derleme sonrası önbellek ve geçici dosyaları temizlemek
// - go list -------------------> Proje içindeki modül ve paket bilgilerini listelemek için.
// - go mod tidy ---------------> Bağımlılıkları düzenlemek, kullanılmayan bağımlılıkları kaldırmak ve eksik olanları eklemek

// Genel Go İş Akışı Adımları
//     1. go mod init [module_name] -> Initialize a new module
//     2. go get [package_name] -----> Add external dependencies
//     3. go mod tidy ---------------> Clean up dependencies
//     4. go run main.go ------------> Run code without creating binary
//     5. go test ./... -------------> Run all tests in the project
//     6. go build ------------------> Create executable binary
//     7. go install ----------------> Build and install binary to GOPATH/bin
