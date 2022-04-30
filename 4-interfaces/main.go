package main

func main() {
	// Estructural types
	ft := FullTimeEmployee{}
	ft.name = "Carlos"
	ft.id = 1
	ft.age = 28
	ft.endDate = "9999"
	PrintInfo(ft)
	t := TemporaryEmployee{}
	t.name = "Carlos"
	t.id = 1
	t.age = 28
	t.taxRate = 5
	PrintInfo(t)
}
