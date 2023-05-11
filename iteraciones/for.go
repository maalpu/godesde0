package iteraciones

import "fmt"

func Iterar() {
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }
	fmt.Println("-----0 a 4-----")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("---0 a 20 c/5---")
	for i := 0; i < 25; i += 5 {
		fmt.Println(i)
	}
	fmt.Println("---25 a 5 c/5---")
	for i := 25; i > 0; i -= 5 {
		fmt.Println(i)
	}
	fmt.Println("-----5 a 1------")
	for i := 5; i > 0; i-- {
		fmt.Println(i)
	}
	fmt.Println("----continue----")
	for i := 0; i < 10; i++ {
		if i == 6 {
			continue
		}
		fmt.Println(i)
	}
}
