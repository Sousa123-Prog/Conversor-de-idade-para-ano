package main

import (
	"fmt"
	"time"
)

func main() {
	var idade int
	for {
		fmt.Println("Por favor diga a sua idade")
		_, err := fmt.Scanln(&idade)
		if err != nil {
			fmt.Println("Idade inválida, por favor repita")
			continue
		}
		Ano := time.new(time.Now().Year())

		sub := Ano - idade

		fmt.Println("O ano atual é %d", Ano)
		fmt.Println("O ano que você nasceu foi %d", sub)
	}
}
