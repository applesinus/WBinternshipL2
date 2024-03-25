package main

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	// getting time from ntp
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		// if there's an error, print it to stderr and exit with code 1
		fmt.Fprintf(os.Stderr, "Error while getting time from ntp: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("Current time: %v\n", time)
	}
}
