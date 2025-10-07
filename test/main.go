// package for execute
package main

import "fmt"

func addition(x, y, z int) {
	fmt.Println("A + B = ", x+y)
	fmt.Println("A + C = ", x+z)
	fmt.Println("B + C = ", y+z)
}

func substraction(x, y, z int) {
	fmt.Println("A - B = ", x-y)
	fmt.Println("A - C = ", x-z)
	fmt.Println("B - C = ", y-z)
}

func divide(x, y, z float64) {
	fmt.Println("A / B = ", x/y)
	fmt.Println("A / C = ", x/z)
	fmt.Println("B / C = ", y/z)
}

func multiply(x, y, z int) {
	fmt.Println("A * B = ", x*y)
	fmt.Println("A * C = ", x*z)
	fmt.Println("B * C = ", y*z)
}

// func oddeven(x, y int) {
// 	fmt.Print("Angka yg mau di periksa? (wajib angka) >> ", x)
// 	fmt.Scanln(&x)
// 	if x%2 == 0 {
// 		fmt.Print("this is even number")
// 	} else {
// 		fmt.Println("this is odd number")
// 	}
// }

func main() {
	fmt.Println("hello world")
	fmt.Println("\n")
	fmt.Println("==input n scan==\n")
	fmt.Print("Input ur name first: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Helloo ", name, ", welcome to the program!\n")

	// fmt.Println("integers etc")
	fmt.Println("==mini calculator==")
	var a, b, c int
	fmt.Print("value a? ")
	fmt.Scanln(&a)
	fmt.Print("value b? ")
	fmt.Scanln(&b)
	fmt.Print("value c? ")
	fmt.Scanln(&c)

	addition(a, b, c)
	print("\n")
	substraction(a, b, c)
	print("\n")
	divide(float64(a), float64(b), float64(c))
	print("\n")
	multiply(a, b, c)
	print("\n")
	// oddeven(a, b)

	fmt.Println("odd even numbe checker")
	var number []int

	for {
		var x int
		fmt.Print("Angka yg mau diperiksa (wajib angka): ")
		fmt.Scanln(&x)

		// simpan ke slice
		number = append(number, x)

		// cek ganjil genap
		if x%2 == 0 {
			fmt.Println(x, "adalah bilangan genap")
		} else {
			fmt.Println(x, "adalah bilangan ganjil")
		}

		// tanya apakah lanjut atau berhenti
		var choice string
		fmt.Print("Mau cek lagi? (y/n): ")
		fmt.Scanln(&choice)

		if choice == "n" || choice == "N" {
			fmt.Println("List number:\n")
			for i, val := range number {
				fmt.Printf("%d. %d\n", i+1, val)
			}
			fmt.Println("\nok seeya\n")
			break
		}

		fmt.Println() // baris kosong biar rapi
	}

	//array
	// students := map[string]int{
	// 	"Apple":     90,
	// 	"Banana":    80,
	// 	"Chocolate": 95,
	// }
	// total := 0
	// for _, score := range students {
	// 	total += score
	// }
	// fmt.Println("Average:", total/len(students))

	//input n count

	var (
		array_name  []string
		array_value []float64
	)

	fmt.Println("==program menghitung average murid==")

	for {
		var name string
		var value float64
		fmt.Print("Name: ")
		fmt.Scanln(&name)
		fmt.Print("Grade: ")
		fmt.Scanln(&value)

		array_name = append(array_name, name)
		array_value = append(array_value, value)

		var decide string
		fmt.Println("Calculate the average nooww (y/n)>> ")
		fmt.Scanln(&decide)
		if decide == "y" || decide == "Y" {
			var total float64
			for _, val := range array_value {
				total += val
			}
			avg := total / float64(len(array_value))

			fmt.Println("\n=== Hasil ===")
			for i := range array_name {
				fmt.Printf("%s: %.2f\n", array_name[i], array_value[i])
			}
			fmt.Printf("Rata-rata nilai: %.2f\n", avg)
			break
		}
		fmt.Println("hm oke input next data.")
	}
}
