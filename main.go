package main

import (
	"fmt"
	"game/Game2.0/logic"
	"math/rand"
	"time"
)

func WhoWin(p, d int) {
	if p > d {
		fmt.Println("Вы победили!")
	} else if p == d {
		fmt.Println("Ничья")
	} else {
		fmt.Println("Вы проиграли!")
	}
}

func main() {
	var input string
	rand.Seed(time.Now().Unix())
	fmt.Println("Здарова Чуваааак, че поиграем? Проверим твою удачу?")
	fmt.Println("Управление: Hit - взять карту, Stand - больше не брать карту, передать ход диллеру")
	fmt.Println("Статртуем? Yes/No")
	fmt.Scanln(&input)
	if input == "Yes" {
		fmt.Println("Игра началась(отсылОЧКА к пиле) =)")
		WhoWin(logic.PlayerGame(), logic.DealerGame())
	} else if input == "No" {
		fmt.Println("Golang`ера ответ, пока =(")
	} else {
		fmt.Println("Тебе сказали команды, сам разбирай свою абракадабру! Досвидули")
	}
}
