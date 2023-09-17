package main

type Segitiga struct {
	Alas float64
	Tinggi float64
}

func main() {
	func (k Segitiga) Luas() float64 {
		return (k.Alas * k.Tinggi)/2
	}
	func (k Segitiga) Keliling() float64 {
		return k.Alas * 3
	}
}