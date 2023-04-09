package main

import "fmt"

func hitungOperator(jumOpKali, jumOpJumlah, jumOpKurang, jumOpBagi *int, kar rune) {
	if kar == '+' {
		*jumOpJumlah += 1
	} else if kar == '-' {
		*jumOpKurang += 1
	} else if kar == '*' {
		*jumOpKali += 1
	} else if kar == '/' {
		*jumOpBagi += 1
	}
}

func cetak(jumOpKali, jumOpJumlah, jumOpKurang, jumOpBagi int) {
	if jumOpJumlah > 0{
		fmt.Println("ada ", jumOpJumlah, "operasi penjumlahan")
	}
}

func main(){
	var kar rune 
	var jumOpKali, jumOpJumlah, jumOpKurang, jumOpBagi int
	fmt.Scan(&kar)
	huruf := string(kar)
	for huruf != "." {
		hitungOperator(&jumOpKali, &jumOpJumlah, &jumOpKurang, &jumOpBagi, kar)
		fmt.Scan(&kar)
	}
}


// package main

// import (
// 	"fmt"
// 	"strings"
// 	"unicode"
// )

// func main() {
//     var input string
//     fmt.Print("Masukkan sebaris karakter: ")
//     fmt.Scanln(&input)

//     var operators = map[rune]string{
//         '+': "penjumlahan",
//         '-': "pengurangan",
//         '*': "perkalian",
//         '/': "pembagian",
//     }

//     var operatorCount = make(map[string]int)
//     var currentOperator string
//     for _, char := range input {
//         if unicode.IsDigit(char) {
//             continue
//         } else if op, ok := operators[char]; ok {
//             currentOperator = op
//             operatorCount[currentOperator]++
//         } else if char == '.' {
//             brePrintln()
//         } else {
//             fmt.Println("Karakter tidak valid")
//             return
//         }
//     }

//     var output strings.Builder
//     for op, count := range operatorCount {
//         if output.Len() > 0 {
//             output.WriteString(", ")
//         }
//         output.WriteString(fmt.Sprintf("%d %s", count, op))
//     }

//     fmt.Println(output.String())
// }