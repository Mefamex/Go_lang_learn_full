/*
author: mefamex
date  : 2025-07-03
title : Methods in Go


METHOD NEDİR?
- Method, belirli bir türle (type) ilişkilendirilmiş fonksiyondur
- Go'da method, receiver ile tanımlanır
- Method, struct'larla Object-Oriented Programming (OOP) benzeri davranış sağlar
- Method, fonksiyonun özel bir türüdür ve bir receiver'a sahiptir



METHOD SYNTAX:
func (receiver ReceiverType) methodName(parameters) returnType {
    // method body
}


RECEIVER TÜRLERİ:
1. Value Receiver: (r ReceiverType) - Kopya üzerinde çalışır
2. Pointer Receiver: (r *ReceiverType) - Orijinal değer üzerinde çalışır



METHOD vs FUNCTION:
- Function: Bağımsız çalışan kod bloğu
-   Method: Belirli bir türle ilişkilendirilmiş fonksiyon
- Function çağrısı: functionName()
-   Method çağrısı: receiver.methodName()



METHOD SETS:
- Her tür için tanımlı method'ların koleksiyonu
- Interface implementation için kritik
- Pointer receiver'lar hem T hem de *T için çalışır
- Value receiver'lar sadece T için çalışır



BEST PRACTICES:
1. Büyük struct'lar için pointer receiver kullanın
2. Method'ların tutarlılığı için aynı receiver türünü kullanın
3. Interface implementation için receiver türünü dikkatli seçin
4. Nil pointer method call'larını handle edin
5. Method chaining için pointer return edin
6. Method adları açıklayıcı olmalı
7. Exported method'lar için documentation yazın
8. Method complexity'yi düşük tutun



ADVANCED METHOD PATTERNS:
1. Method chaining: Birden fazla method'u zincirleme çağırma
2. Builder pattern: Kompleks nesne oluşturma
3. Method with multiple return values: Birden fazla değer döndüren method
4. Method with variadic parameters: Değişken sayıda parametre alan method
5. Method overriding: Embedded types ile method'ları geçersiz kılma
6. Method sets ve interface: Method set'lerin interface ile ilişkisi
7. Embedded types: Struct'ların composition ile kullanımı
8. String method implementation: fmt.Stringer interface'i ile string dönüşümü
9. Nil receiver method call'ları: Nil pointer ile method çağrısı

*/

package main

import (
	"fmt"
	"math"
)

// Temel struct tanımları
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Person struct {
	Name string
	Age  int
	City string
}

type BankAccount struct {
	AccountNumber string
	Balance       float64
	Owner         string
}

type Temperature struct {
	Celsius float64
}

// Interface tanımları
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Stringer interface {
	String() string
}

func main() {
	fmt.Println("=======================================")
	fmt.Println("          METHODS TEMEL KAVRAMLAR")
	fmt.Println("=======================================")

	// Method'lar struct'larla OOP benzeri davranış sağlar
	fmt.Println("Go Methods Özellikleri:")
	fmt.Println("1. Receiver ile tanımlanır")
	fmt.Println("2. Value veya Pointer receiver")
	fmt.Println("3. Method sets oluşturur")
	fmt.Println("4. Interface implementation")
	fmt.Println("5. Method chaining desteği")

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("      TEMEL METHOD ÖRNEKLERİ")
	fmt.Println("=======================================")

	// Temel method kullanımı
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("Rectangle: %+v\n", rect)
	fmt.Printf("Alan: %.2f\n", rect.Area())
	fmt.Printf("Çevre: %.2f\n", rect.Perimeter())

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("      FARKLI TÜRLERİN METHOD'LARI")
	fmt.Println("=======================================")

	// Circle struct method'ları
	circle := Circle{Radius: 7}
	fmt.Printf("Circle: %+v\n", circle)
	fmt.Printf("Alan: %.2f\n", circle.Area())
	fmt.Printf("Çevre: %.2f\n", circle.Perimeter())

	// Person struct method'ları
	person := Person{Name: "Ahmet", Age: 30, City: "İstanbul"}
	fmt.Printf("Person: %+v\n", person)
	fmt.Printf("Yaş kategorisi: %s\n", person.GetAgeCategory())
	fmt.Printf("Tam bilgi: %s\n", person.GetFullInfo())

	// Method chaining
	fmt.Println("\nMethod chaining:")
	person.SetName("Mehmet").SetAge(25).SetCity("Ankara")
	fmt.Printf("Güncellenmiş person: %+v\n", person)

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("      INTERFACE IMPLEMENTATION")
	fmt.Println("=======================================")

	// Interface implementation
	var shapes []Shape
	shapes = append(shapes, Rectangle{Width: 4, Height: 3})
	shapes = append(shapes, Circle{Radius: 5})

	fmt.Println("Shape interface kullanımı:")
	for i, shape := range shapes {
		fmt.Printf("Shape %d - Alan: %.2f, Çevre: %.2f\n",
			i+1, shape.Area(), shape.Perimeter())
	}

	// Type assertion ile concrete type'a erişim
	fmt.Println("\nType assertion:")
	for i, shape := range shapes {
		switch s := shape.(type) {
		case Rectangle:
			fmt.Printf("Shape %d: Rectangle - %dx%d\n", i+1, int(s.Width), int(s.Height))
		case Circle:
			fmt.Printf("Shape %d: Circle - Radius: %d\n", i+1, int(s.Radius))
		}
	}

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("      BANKA HESABI METHOD ÖRNEKLERİ")
	fmt.Println("=======================================")

	// Banka hesabı method'ları
	account := BankAccount{
		AccountNumber: "TR1234567890",
		Balance:       1000.0,
		Owner:         "Ali Veli",
	}

	fmt.Printf("Hesap: %+v\n", account)

	// Para yatırma
	account.Deposit(500.0)
	fmt.Printf("500 TL yatırma sonrası: %.2f TL\n", account.Balance)

	// Para çekme
	if account.Withdraw(300.0) {
		fmt.Printf("300 TL çekme sonrası: %.2f TL\n", account.Balance)
	}

	// Yetersiz bakiye
	if !account.Withdraw(2000.0) {
		fmt.Println("Yetersiz bakiye!")
	}

	// Transfer
	account2 := BankAccount{
		AccountNumber: "TR0987654321",
		Balance:       500.0,
		Owner:         "Ayşe Fatma",
	}

	fmt.Printf("Transfer öncesi - Account1: %.2f, Account2: %.2f\n",
		account.Balance, account2.Balance)

	if account.Transfer(&account2, 200.0) {
		fmt.Printf("200 TL transfer sonrası - Account1: %.2f, Account2: %.2f\n",
			account.Balance, account2.Balance)
	}

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("      SICAKLIK DÖNÜŞÜM METHOD'LARI")
	fmt.Println("=======================================")

	// Temperature conversion method'ları
	temp := Temperature{Celsius: 25.0}
	fmt.Printf("Sıcaklık: %.1f°C\n", temp.Celsius)
	fmt.Printf("Fahrenheit: %.1f°F\n", temp.ToFahrenheit())
	fmt.Printf("Kelvin: %.1f K\n", temp.ToKelvin())

	// Method chaining ile sıcaklık işlemleri
	fmt.Println("\nSıcaklık işlemleri:")
	temp.SetCelsius(30.0).Add(10.0).Multiply(1.5)
	fmt.Printf("İşlemler sonrası: %.1f°C\n", temp.Celsius)

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("      STRING METHOD'LARI")
	fmt.Println("=======================================")

	// String method implementation
	fmt.Println("String method kullanımı:")
	fmt.Printf("Rectangle string: %s\n", rect.String())
	fmt.Printf("Circle string: %s\n", circle.String())
	fmt.Printf("Person string: %s\n", person.String())

	// fmt.Stringer interface automatic usage
	fmt.Printf("Automatic string formatting: %v\n", rect)
	fmt.Printf("Automatic string formatting: %v\n", circle)

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("      ADVANCED METHOD PATTERNS")
	fmt.Println("=======================================")

	// Builder pattern
	fmt.Println("Builder pattern:")
	builder := NewPersonBuilder()
	buildPerson := builder.SetName("Fatma").
		SetAge(28).
		SetCity("Bursa").
		Build()
	fmt.Printf("Built person: %+v\n", buildPerson)

	// Method with multiple return values
	fmt.Println("\nMultiple return values:")
	area, perimeter := rect.GetAreaAndPerimeter()
	fmt.Printf("Alan: %.2f, Çevre: %.2f\n", area, perimeter)

	// Method with variadic parameters
	fmt.Println("\nVariadic method:")
	person.AddHobbies("Kitap okuma", "Müzik dinleme", "Spor yapma")
	fmt.Printf("Hobbies: %v\n", person.GetHobbies())

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("      METHOD SETS VE INTERFACE")
	fmt.Println("=======================================")

	// Method sets demonstration
	fmt.Println("Method sets:")

	// Value receiver method'ları hem T hem de *T için çalışır
	rectValue := Rectangle{Width: 3, Height: 4}
	rectPointer := &Rectangle{Width: 5, Height: 6}

	fmt.Printf("Value receiver call: %.2f\n", rectValue.Area())
	fmt.Printf("Pointer receiver call: %.2f\n", rectPointer.Area())

	// Interface assignment
	var shape Shape = rectValue // Value receiver method'u interface'i implement eder
	fmt.Printf("Interface call: %.2f\n", shape.Area())

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("      EMBEDDED TYPES VE METHOD'LAR")
	fmt.Println("=======================================")

	// Embedded types (composition)
	colorRect := ColoredRectangle{
		Rectangle: Rectangle{Width: 8, Height: 6},
		Color:     "Kırmızı",
	}

	fmt.Printf("Colored Rectangle: %+v\n", colorRect)
	fmt.Printf("Alan: %.2f\n", colorRect.Area()) // Embedded method
	fmt.Printf("Renk: %s\n", colorRect.GetColor())
	fmt.Printf("Açıklama: %s\n", colorRect.GetDescription())

	////
	////
	////

	fmt.Println("\n=======================================")
	fmt.Println("      METHODS BEST PRACTICES")
	fmt.Println("=======================================")

	fmt.Println("Method Best Practices:")
	fmt.Println("1. Küçük struct'lar için value receiver kullanın")
	fmt.Println("2. Büyük struct'lar için pointer receiver kullanın")
	fmt.Println("3. Struct'ı değiştirecek method'lar için pointer receiver")
	fmt.Println("4. Tutarlılık için aynı receiver türünü kullanın")
	fmt.Println("5. Interface implementation için receiver türünü dikkatli seçin")
	fmt.Println("6. Nil pointer method call'larını handle edin")
	fmt.Println("7. Method chaining için pointer return edin")
	fmt.Println("8. Method adları açıklayıcı olmalı")
	fmt.Println("9. Exported method'lar için documentation yazın")
	fmt.Println("10. Method complexity'yi düşük tutun")

	fmt.Println("=======================================")
}

// Rectangle methods - Value receiver
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// String method implementation
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle(%.1fx%.1f)", r.Width, r.Height)
}

// Multiple return values
func (r Rectangle) GetAreaAndPerimeter() (float64, float64) {
	return r.Area(), r.Perimeter()
}

// Circle methods
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle(radius:%.1f)", c.Radius)
}

// Person methods
func (p Person) GetAgeCategory() string {
	switch {
	case p.Age < 18:
		return "Çocuk"
	case p.Age < 65:
		return "Yetişkin"
	default:
		return "Yaşlı"
	}
}

func (p Person) GetFullInfo() string {
	return fmt.Sprintf("%s, %d yaşında, %s'da yaşıyor", p.Name, p.Age, p.City)
}

// Method chaining - pointer receiver
func (p *Person) SetName(name string) *Person {
	p.Name = name
	return p
}

func (p *Person) SetAge(age int) *Person {
	p.Age = age
	return p
}

func (p *Person) SetCity(city string) *Person {
	p.City = city
	return p
}

// String method
func (p Person) String() string {
	return fmt.Sprintf("Person{Name:%s, Age:%d, City:%s}", p.Name, p.Age, p.City)
}

// Hobbies field ekleyelim (slice)
var personHobbies []string

func (p *Person) AddHobbies(hobbies ...string) {
	personHobbies = append(personHobbies, hobbies...)
}

func (p Person) GetHobbies() []string {
	return personHobbies
}

// BankAccount methods
func (b *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		b.Balance += amount
	}
}

func (b *BankAccount) Withdraw(amount float64) bool {
	if amount > 0 && b.Balance >= amount {
		b.Balance -= amount
		return true
	}
	return false
}

func (b *BankAccount) Transfer(to *BankAccount, amount float64) bool {
	if b.Withdraw(amount) {
		to.Deposit(amount)
		return true
	}
	return false
}

func (b BankAccount) GetBalance() float64 {
	return b.Balance
}

// Temperature methods
func (t Temperature) ToFahrenheit() float64 {
	return (t.Celsius * 9 / 5) + 32
}

func (t Temperature) ToKelvin() float64 {
	return t.Celsius + 273.15
}

// Method chaining for Temperature
func (t *Temperature) SetCelsius(celsius float64) *Temperature {
	t.Celsius = celsius
	return t
}

func (t *Temperature) Add(celsius float64) *Temperature {
	t.Celsius += celsius
	return t
}

func (t *Temperature) Multiply(factor float64) *Temperature {
	t.Celsius *= factor
	return t
}

// Builder pattern example
type PersonBuilder struct {
	person Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{}
}

func (pb *PersonBuilder) SetName(name string) *PersonBuilder {
	pb.person.Name = name
	return pb
}

func (pb *PersonBuilder) SetAge(age int) *PersonBuilder {
	pb.person.Age = age
	return pb
}

func (pb *PersonBuilder) SetCity(city string) *PersonBuilder {
	pb.person.City = city
	return pb
}

func (pb *PersonBuilder) Build() Person {
	return pb.person
}

// Embedded types (composition)
type ColoredRectangle struct {
	Rectangle // Embedded type
	Color     string
}

func (cr ColoredRectangle) GetColor() string {
	return cr.Color
}

func (cr ColoredRectangle) GetDescription() string {
	return fmt.Sprintf("%s renkli dikdörtgen (%.1fx%.1f)",
		cr.Color, cr.Width, cr.Height)
}

// Method overriding example
func (cr ColoredRectangle) String() string {
	return fmt.Sprintf("ColoredRectangle(%.1fx%.1f, %s)",
		cr.Width, cr.Height, cr.Color)
}
