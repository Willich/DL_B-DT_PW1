package main

import (
	crypto "crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {

	var n int64
	n = 8
	kbig := big.NewInt(16)
	key := new(big.Int)

	for n < 33 {
		key.Mul(kbig, kbig)

		r, err := crypto.Int(crypto.Reader, key)
		if err != nil {
			panic(err)
		}

		d := time_bf_1(r)
		fmt.Println("Разрядность ключа", n)
		fmt.Println("Поле ключей", key)
		fmt.Println("Случайный ключ", r)
		fmt.Println("Время перебора до совпадения с случайным ключом", d)
		fmt.Println()
		n = n * 2
		kbig = key
	}

	for n < 4097 {

		fmt.Println("Разрядность ключа", n)
		key.Mul(kbig, kbig)
		fmt.Println("Поле ключей", key)
		key_Fl := big.NewFloat(0).SetInt(key)
		fmt.Printf("Поле ключей в exp формате %.12G\n", key_Fl)

		r, err := crypto.Int(crypto.Reader, key)
		if err != nil {
			panic(err)
		}
		fmt.Println("Случайный ключ", r)
		r_Fl := big.NewFloat(0).SetInt(r) //изменение типа случайного ключа с big.Int на big.Float для вывода в экспоненциальном виде
		fmt.Printf("Случайный ключ в exp формате %.12G\n", r_Fl)

		time_bf_2(r)

		n = n * 2
		kbig = key
	}
}

func time_bf_1(x *big.Int) (duration time.Duration) {
	one := big.NewInt(1)
	count := big.NewInt(0)
	start := time.Now()
	for count.Cmp(x) == -1 {
		count.Add(count, one)
	}
	duration = time.Since(start)
	return
}

func time_bf_2(x *big.Int) {
	one := big.NewInt(1)
	count := big.NewInt(0)
	countmax := big.NewInt(1000000000)
	start := time.Now()
	for count.Cmp(countmax) == -1 { // цикл 1 млрд. переборов ключей для определения времени
		count.Add(count, one)
	}
	duration := time.Since(start) // время перебора 1 млрд. ключей

	var dur_hour float64
	dur_hour = duration.Hours() //выражение времени цикла переборов ключей в часах тип float64
	dur_hour_big := new(big.Float)
	dur_hour_big = big.NewFloat(dur_hour) // изменение типа float64 на bif.float
	duration_big := new(big.Float)

	time_r := new(big.Int)
	time_r.Div(x, countmax) // деление случайного ключа r на 1 млрд
	time_rFl := new(big.Float)
	time_rFl.SetInt(time_r) // изменение типа big. Int  на bif.float

	duration_big.Mul(time_rFl, dur_hour_big)
	fmt.Printf("Оценочное время перебора до совпадения со случайным ключом, часов, %.12G\n", duration_big)
	fmt.Println(" ")
	return
}
