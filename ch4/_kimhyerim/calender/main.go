package main

import "fmt"

func main() {
	months := [...]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "Jun",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}
	fmt.Println("month: ", months)
	fmt.Printf("len(months): %d, cap(months): %d\n", len(months), cap(months))

	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println("Q2: ", Q2)
	fmt.Printf("len(Q2): %d, cap(Q2): %d\n", len(Q2), cap(Q2))
	fmt.Println("summer: ", summer)
	fmt.Printf("len(summer): %d, cap(summer): %d\n", len(summer), cap(summer))

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	// fmt.Println(summer[:20]) Panic!!
	endlessSummer := summer[:5]
	fmt.Println(endlessSummer)
}
