/*
author: mefamex
date  : 2025-07-03
title : Pointer vs Value Receivers in Go


POINTER vs VALUE RECEIVERS NEDİR?
- Go'da method'lar iki farklı receiver türü ile tanımlanabilir:
    1.   Value Receiver: (r Type)  - Kopya üzerinde çalışır
    2. Pointer Receiver: (r *Type) - Orijinal üzerinde çalışır
- Performance: Büyük struct'lar için pointer tercih edilir
- Mutability: Struct'ı değiştirmek için pointer receiver gerekir
- Interface rules: Method sets kuralları interface implementation'ı etkiler



VALUE RECEIVER ÖZELLİKLERİ:
- Struct'ın kopyası üzerinde çalışır
- Orijinal struct'ı değiştirmez
- Küçük struct'lar için performans açısından uygun
- Immutable (değiştirilemez) davranış sağlar
- Memory safe (güvenli)



POINTER RECEIVER ÖZELLİKLERİ:
- Struct'ın orijinal değeri üzerinde çalışır
- Orijinal struct'ı değiştirebilir
- Büyük struct'lar için performans açısından uygun
- Mutable (değiştirilebilir) davranış sağlar
- Memory allocation'ı azaltır



PERFORMANS KARŞILAŞTIRMASI:
- Value receiver: Her method call'da kopya oluşturulur
- Pointer receiver: Sadece pointer adresi kopyalanır
- Büyük struct'lar için pointer receiver tercih edilir



INTERFACE IMPLEMENTATION:
- Value receiver'lar hem T hem de *T için çalışır
- Pointer receiver'lar sadece *T için çalışır
- Method sets kuralları interface implementation'ı etkiler



BEST PRACTICES DEMONSTRATION:
1. Küçük struct (<= 16 bytes) → Value receiver
2. Büyük struct (> 16 bytes) → Pointer receiver
3. Struct'ı değiştirmek → Pointer receiver
4. Read-only operations → Value receiver
5. Method chaining → Pointer receiver
6. Interface implementation → Receiver türünü dikkatli seç
7. Nil safety → Pointer receiver'da kontrol et
8. Consistency → Aynı türde receiver kullan


*/

package main

import (
	"fmt"
	"math"
	"unsafe"
)

// Küçük struct - Value receiver için uygun
type Point struct {
	X float64
	Y float64
}

// Büyük struct - Pointer receiver için uygun
type LargeStruct struct {
	Name        string
	Description string
	Data        [1000]int         // 4KB veri
	Matrix      [100][100]float64 // 80KB veri
	Buffer      []byte
}

// Orta boyut struct - Her iki receiver türü için test
type Rectangle struct {
	Width  float64
	Height float64
	Color  string
}

// Banka hesabı - Mutating operations için
type BankAccount struct {
	AccountNumber string
	Owner         string
	Balance       float64
	Transactions  []Transaction
}

type Transaction struct {
	ID     int
	Type   string
	Amount float64
}

// Counter - State değişimi için
type Counter struct {
	Value int
	Name  string
}

// Employee - Real world örneği
type Employee struct {
	ID       int
	Name     string
	Position string
	Salary   float64
	Benefits []string
}

// Interface definition
type Shape interface {
	Area() float64
	Perimeter() float64
}

func main() {

	fmt.Println("\n=======================================")
	fmt.Println("      VALUE RECEIVER ÖRNEKLERİ")
	fmt.Println("=======================================")

	// Value receiver - kopya üzerinde çalışır
	point := Point{X: 10, Y: 20}
	fmt.Printf("Orijinal point: %+v\n", point)

	// Distance calculation (read-only operation)
	distance := point.DistanceFromOrigin()
	fmt.Printf("Origin'den uzaklık: %.2f\n", distance)

	// Value receiver ile "değiştirme" denemesi
	fmt.Println("\nValue receiver ile değiştirme denemesi:")
	fmt.Printf("Değişim öncesi: %+v\n", point)
	point.MoveValue(5, 10)                      // Kopyayı değiştirir
	fmt.Printf("Değişim sonrası: %+v\n", point) // Orijinal değişmez!

	// Doğru kullanım - return değeriyle
	fmt.Println("\nValue receiver ile doğru kullanım:")
	newPoint := point.MoveValueReturn(5, 10)
	fmt.Printf("Orijinal point: %+v\n", point)
	fmt.Printf("Yeni point: %+v\n", newPoint)

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("      POINTER RECEIVER ÖRNEKLERİ")
	fmt.Println("=======================================")

	// Pointer receiver - orijinal değer üzerinde çalışır
	rect := Rectangle{Width: 10, Height: 5, Color: "Kırmızı"}
	fmt.Printf("Orijinal rectangle: %+v\n", rect)

	// Pointer receiver ile değiştirme
	fmt.Println("\nPointer receiver ile değiştirme:")
	fmt.Printf("Değişim öncesi: %+v\n", rect)
	rect.Resize(20, 15) // Orijinali değiştirir
	fmt.Printf("Değişim sonrası: %+v\n", rect)

	// Renk değiştirme
	rect.SetColor("Mavi")
	fmt.Printf("Renk değişimi sonrası: %+v\n", rect)

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("      NIL POINTER HANDLING")
	fmt.Println("=======================================")

	// Nil pointer method call
	var nilRect *Rectangle
	fmt.Println("Nil pointer method call:")
	nilRect.SafeMethod() // Nil kontrolü ile güvenli

	// Nil pointer - farklı durumları gösterelim
	fmt.Println("\nNil pointer - farklı durumlar:")

	// 1. Nil pointer - güvenli method
	fmt.Println("1. Nil pointer ile safe method:")
	nilRect.SafeMethod()

	// 2. Nil olmayan pointer
	fmt.Println("2. Nil olmayan pointer:")
	validRect := &Rectangle{Width: 5, Height: 3, Color: "Test"}
	validRect.SafeMethod()

	// 3. Nil kontrolü ile güvenli kullanım (gerçekçi senaryo)
	fmt.Println("3. Dinamik olarak oluşturulan pointer:")
	var dynamicRect *Rectangle

	// Conditional olarak pointer oluştur
	shouldCreateRect := true
	if shouldCreateRect {
		dynamicRect = &Rectangle{Width: 8, Height: 4, Color: "Dinamik"}
	}

	// Güvenli kullanım
	if dynamicRect != nil {
		dynamicRect.Resize(12, 6)
		fmt.Printf("Dynamic rectangle resize edildi: %+v\n", dynamicRect)
	} else {
		fmt.Println("Dynamic rectangle nil, resize edilemedi")
	}

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("      PERFORMANCE KARŞILAŞTIRMASI")
	fmt.Println("=======================================")

	// Küçük struct için value receiver
	smallPoint := Point{X: 1, Y: 2}
	fmt.Printf("Küçük struct boyutu: %d bytes\n", unsafe.Sizeof(smallPoint))

	// Büyük struct için pointer receiver
	largeData := LargeStruct{
		Name:        "Test",
		Description: "Large data structure",
		Buffer:      make([]byte, 1000),
	}
	fmt.Printf("Büyük struct boyutu: %d bytes\n", unsafe.Sizeof(largeData))

	// Performance demonstration
	fmt.Println("\nPerformance test:")
	fmt.Printf("Küçük struct - Value receiver kullanım: %.2f\n", smallPoint.DistanceFromOrigin())
	fmt.Printf("Büyük struct - Pointer receiver kullanım: %s\n", largeData.GetInfo())

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("      MUTABLE vs IMMUTABLE OPERATIONS")
	fmt.Println("=======================================")

	// Mutable operations - pointer receiver gerekli
	account := BankAccount{
		AccountNumber: "TR1234567890",
		Owner:         "Ali Veli",
		Balance:       1000.0,
		Transactions:  make([]Transaction, 0),
	}

	fmt.Printf("Hesap başlangıç: %+v\n", account)

	// Para yatırma (mutating operation)
	account.Deposit(500.0)
	fmt.Printf("500 TL yatırma sonrası: Balance=%.2f\n", account.Balance)

	// Para çekme (mutating operation)
	success := account.Withdraw(200.0)
	fmt.Printf("200 TL çekme sonrası: Balance=%.2f, Success=%t\n", account.Balance, success)

	// Transaction history (read-only operation)
	fmt.Printf("Transaction sayısı: %d\n", len(account.GetTransactions()))

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("      COUNTER EXAMPLE")
	fmt.Println("=======================================")

	// Counter örneği - state mutation
	counter := Counter{Value: 0, Name: "Test Counter"}
	fmt.Printf("Counter başlangıç: %+v\n", counter)

	// Increment operations
	counter.Increment()
	counter.Increment()
	counter.IncrementBy(5)
	fmt.Printf("Increment sonrası: %+v\n", counter)

	// Reset
	counter.Reset()
	fmt.Printf("Reset sonrası: %+v\n", counter)

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("      METHOD CHAINING")
	fmt.Println("=======================================")

	// Method chaining - pointer receiver gerekli
	emp := Employee{
		ID:       1,
		Name:     "Ahmet Yılmaz",
		Position: "Developer",
		Salary:   50000,
		Benefits: make([]string, 0),
	}

	fmt.Printf("Employee başlangıç: %+v\n", emp)

	// Method chaining
	emp.SetPosition("Senior Developer").
		SetSalary(60000).
		AddBenefit("Sağlık Sigortası").
		AddBenefit("Yemek Kartı")

	fmt.Printf("Method chaining sonrası: %+v\n", emp)

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("      INTERFACE IMPLEMENTATION")
	fmt.Println("=======================================")

	// Interface implementation demonstration
	var shape Shape = Rectangle{Width: 10, Height: 5, Color: "Yeşil"}
	fmt.Printf("Shape interface - Alan: %.2f\n", shape.Area())
	fmt.Printf("Shape interface - Çevre: %.2f\n", shape.Perimeter())

	// Method sets demonstration
	fmt.Println("\nMethod sets:")
	rectValue := Rectangle{Width: 3, Height: 4, Color: "Sarı"}
	rectPointer := &Rectangle{Width: 5, Height: 6, Color: "Mor"}

	// Value receiver method'ları her ikisi için de çalışır
	fmt.Printf("Value method call: %.2f\n", rectValue.Area())
	fmt.Printf("Pointer method call: %.2f\n", rectPointer.Area())

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("      COPY vs REFERENCE BEHAVIOR")
	fmt.Println("=======================================")

	// Copy behavior demonstration
	original := Rectangle{Width: 10, Height: 5, Color: "Orijinal"}
	fmt.Printf("Original: %+v\n", original)

	// Value receiver - kopya davranışı
	copy := original.GetCopy()
	copy.Width = 20 // Kopyayı değiştir
	fmt.Printf("Original (değişmez): %+v\n", original)
	fmt.Printf("Copy (değişir): %+v\n", copy)

	// Pointer receiver - reference davranışı
	reference := &original
	reference.SetColor("Değiştirilmiş")
	fmt.Printf("Original (değişir): %+v\n", original)

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("      MEMORY ALLOCATION COMPARISON")
	fmt.Println("=======================================")

	// Memory allocation demonstration
	fmt.Println("Memory allocation karşılaştırması:")

	// Value receiver - her call'da kopya
	fmt.Printf("Point struct size: %d bytes\n", unsafe.Sizeof(Point{}))
	fmt.Printf("Rectangle struct size: %d bytes\n", unsafe.Sizeof(Rectangle{}))
	fmt.Printf("LargeStruct size: %d bytes\n", unsafe.Sizeof(LargeStruct{}))

	// Pointer size (architecture dependent)
	fmt.Printf("Pointer size: %d bytes\n", unsafe.Sizeof(uintptr(0)))

	fmt.Println("\nValue receiver: Her method call'da struct kopyası oluştur")
	fmt.Println("Pointer receiver: Sadece pointer adresi kopyala")

	fmt.Println("=======================================")
}

// Point methods - Value receiver (küçük struct)
func (p Point) DistanceFromOrigin() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// Value receiver - orijinali değiştirmeye çalışır (başarısız)
// Bu method eğitim amaçlı: value receiver'ın limitation'ını gösterir
func (p Point) MoveValue(deltaX, deltaY float64) {
	// Bu assignment'lar sadece kopya üzerinde çalışır
	// Orijinal struct değişmez
	_ = deltaX // Parametre kullanıldı
	_ = deltaY // Parametre kullanıldı
	// Not: Gerçek uygulamada bu method'a gerek olmaz
	// Sadece value receiver behavior'ını göstermek için yazıldı
}

// Value receiver - return ile doğru kullanım
func (p Point) MoveValueReturn(deltaX, deltaY float64) Point {
	return Point{
		X: p.X + deltaX,
		Y: p.Y + deltaY,
	}
}

func (p Point) String() string {
	return fmt.Sprintf("Point(%.1f, %.1f)", p.X, p.Y)
}

// Rectangle methods - Pointer receiver (mutable operations)
func (r *Rectangle) Resize(width, height float64) {
	r.Width = width
	r.Height = height
}

func (r *Rectangle) SetColor(color string) {
	r.Color = color
}

// Value receiver - read-only operations
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle(%.1fx%.1f, %s)", r.Width, r.Height, r.Color)
}

// Copy behavior demonstration
func (r Rectangle) GetCopy() Rectangle {
	return Rectangle{
		Width:  r.Width,
		Height: r.Height,
		Color:  r.Color,
	}
}

// Nil pointer safe method
func (r *Rectangle) SafeMethod() {
	if r == nil {
		fmt.Println("Rectangle is nil - safe method called")
		return
	}
	fmt.Printf("Rectangle: %+v\n", r)
}

// LargeStruct methods - Pointer receiver (performans için)
func (ls *LargeStruct) GetInfo() string {
	return fmt.Sprintf("LargeStruct: %s - %s", ls.Name, ls.Description)
}

func (ls *LargeStruct) UpdateData(index int, value int) {
	if index >= 0 && index < len(ls.Data) {
		ls.Data[index] = value
	}
}

// BankAccount methods - Pointer receiver (mutable operations)
func (ba *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		ba.Balance += amount
		ba.Transactions = append(ba.Transactions, Transaction{
			ID:     len(ba.Transactions) + 1,
			Type:   "Deposit",
			Amount: amount,
		})
	}
}

func (ba *BankAccount) Withdraw(amount float64) bool {
	if amount > 0 && ba.Balance >= amount {
		ba.Balance -= amount
		ba.Transactions = append(ba.Transactions, Transaction{
			ID:     len(ba.Transactions) + 1,
			Type:   "Withdraw",
			Amount: amount,
		})
		return true
	}
	return false
}

// Read-only operation - Value receiver
func (ba BankAccount) GetTransactions() []Transaction {
	return ba.Transactions
}

func (ba BankAccount) GetBalance() float64 {
	return ba.Balance
}

// Counter methods - Pointer receiver (state mutation)
func (c *Counter) Increment() {
	c.Value++
}

func (c *Counter) IncrementBy(amount int) {
	c.Value += amount
}

func (c *Counter) Reset() {
	c.Value = 0
}

// Read-only operation - Value receiver
func (c Counter) GetValue() int {
	return c.Value
}

func (c Counter) String() string {
	return fmt.Sprintf("Counter{%s: %d}", c.Name, c.Value)
}

// Employee methods - Method chaining with pointer receiver
func (e *Employee) SetPosition(position string) *Employee {
	e.Position = position
	return e
}

func (e *Employee) SetSalary(salary float64) *Employee {
	e.Salary = salary
	return e
}

func (e *Employee) AddBenefit(benefit string) *Employee {
	e.Benefits = append(e.Benefits, benefit)
	return e
}

// Read-only operations - Value receiver
func (e Employee) GetInfo() string {
	return fmt.Sprintf("%s (%s) - $%.2f", e.Name, e.Position, e.Salary)
}

func (e Employee) String() string {
	return fmt.Sprintf("Employee{ID:%d, Name:%s, Position:%s, Salary:%.2f}",
		e.ID, e.Name, e.Position, e.Salary)
}
