Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	// создаем канал
	ch := make(chan int)

	// вызываем горутину, которая кладёт в канал значения с 0 по 9
	go func() {

		// кладём в канал значения с 0 по 9
		for i := 0; i < 10; i++ {
			// каждый раз кладя значения горутина блокируется до
			// чтения положенного значения
			ch <- i
		}
	}()

	// выводим значения из канала
	// поскольку он небуферизован, это defined behavior
	for n := range ch {
		println(n)
	}
	// цикл повиснет в дедлоке, бесконечно ожидая данные из канала, который
	// используется только в функции, которая уже завершилась
	// избежать этого можно было бы закрыв канал или любым другим образом сообщив
	// о конце работы горутины
}
```

Ответ:
```
0
1
2
3
4
5
6
7
8
9
deadlock

```
