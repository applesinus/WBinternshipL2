Что выведет программа? Объяснить вывод программы.

```go
package main

// создаем структуру, подходящую под интерфейс ошибки, то есть имеющую
// метод .Error(), который возвращает string
type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

// запускаем тест, возвращающий nil
func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	// объявляем переменную и присваиваем её результату test()
	var err error
	err = test()

	// если ошибка не равна nil (true) т.к. ошибка равна нулевому указателю,
	// а не nil, печатаем error
	if err != nil {
		println("error")
		return
	}

	// если ошибка равна nil (false), значит доходим до сюда и печатаем ok
	println("ok")
}
```

Ответ:
```
error

```

Можно было бы в тесте вернуть ошибку интерфейса error, переписав test().
Тогда все было бы ок:
```go
func test() error {
	{
		// do something
	}
	return nil
}
```