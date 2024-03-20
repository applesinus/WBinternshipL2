[Jump to English](#English)

<a name="Russian"></a>
# Русский
<p id="ru"><h3>Задания Lv2 на стажировке Wildberries</h3></p>

[Задачи](#tasks-ru) / [Листинги](#listings-ru) / [Паттерны](#patterns-ru)

<a name="tasks-ru"></a>
### Задачи

<p>Каждая <b>задача</b> - одна папка с соответствующим названием в папке develop, запускается через командную строку (или если не сказанно иного, из файла main.go).</p>
<ol>
    <li>
        <p><b>БАЗОВАЯ ЗАДАЧА</b></p>
        <p>Создать программу печатающую точное время с использованием NTP -библиотеки. Инициализировать как go module. Использовать библиотеку github.com/beevik/ntp. Написать программу печатающую текущее время / точное время с использованием этой библиотеки.</p>
        <p><b>Требования:</b></p>
        <ul>
            <li>Программа должна быть оформлена как go module</li>
            <li>Программа должна корректно обрабатывать ошибки библиотеки: выводить их в STDERR и возвращать ненулевой код выхода в OS</li>
        </ul>
    </li>
    <li>
        <p><b>ЗАДАЧА НА РАСПАКОВКУ</b></p>
        <p>Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:</p>
        <ul>
            <li>"a4bc2d5e" => "aaaabccddddde"</li>
            <li>"abcd" => "abcd"</li>
            <li>"45" => "" (некорректная строка)</li>
            <li>"" => ""</li>
        </ul>
        <p><b>Дополнительно:</b></p>
        <ul>
            <li>Реализовать поддержку escape-последовательностей.</li>
            <li>Например:</li>
            <ul>
                <li>qwe\4\5 => qwe45 (*)</li>
                <li>qwe\45 => qwe44444 (*)</li>
                <li>qwe\\5 => qwe\\\\\ (*)</li>
            </ul>
            <li>В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.</li>
        </ul>
    </li>
    <li>
        <p><b>УТИЛИТА SORT</b></p>
        <p>Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры): на входе подается файл из несортированными строками, на выходе — файл с отсортированными.</p>
        <p><b>Реализовать поддержку утилитой следующий ключей:</b></p>
        <ul>
            <li>-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)</li>
            <li>-n — сортировать по числовому значению</li>
            <li>-r — сортировать в обратном порядке</li>
            <li>-u — не выводить повторяющиеся строки</li>
        </ul>
        <p><b>Дополнительно:</b> реализовать поддержку утилитой следующих ключей:</p>
        <ul>
            <li>-M — сортировать по названию месяца</li>
            <li>-b — игнорировать хвостовые пробелы</li>
            <li>-c — проверять отсортированы ли данные</li>
            <li>-h — сортировать по числовому значению с учетом суффиксов</li>
        </ul>
    </li>
    <li>
        <p><b>ПОИСК АНАГРАММ ПО СЛОВАРЮ</b></p>
        <p>Написать функцию поиска всех множеств анаграмм по словарю.</p>
        <p>Например:</p>
        <ul>
            <li>'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,</li>
            <li>'листок', 'слиток' и 'столик' - другому.</li>
        </ul>
        <p><b>Требования:</b></p>
        <ul>
            <li>Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8</li>
            <li>Выходные данные: ссылка на мапу множеств анаграмм</li>
            <li>Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого, слово из множества.</li>
            <li>Массив должен быть отсортирован по возрастанию.</li>
            <li>Множества из одного элемента не должны попасть в результат.</li>
            <li>Все слова должны быть приведены к нижнему регистру.</li>
            <li>В результате каждое слово должно встречаться только один раз.</li>
        </ul>
    </li>
    <li>
      <p><b>УТИЛИТА GREP</b></p>
      <p>Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).</p>
      <p><b>Реализовать поддержку утилитой следующий ключей:</b></p>
      <ul>
          <li>-A - "after" печатать +N строк после совпадения</li>
          <li>-B - "before" печатать +N строк до совпадения</li>
          <li>-C - "context" (A+B) печатать ±N строк вокруг совпадения</li>
          <li>-c - "count" (количество строк)</li>
          <li>-i - "ignore-case" (игнорировать регистр)</li>
          <li>-v - "invert" (вместо совпадения, исключать)</li>
          <li>-F - "fixed", точное совпадение со строкой, не паттерн</li>
          <li>-n - "line num", напечатать номер строки</li>
      </ul>
  </li>
  <li>
      <p><b>УТИЛИТА CUT</b></p>
      <p>Реализовать утилиту аналог консольной команды cut (man cut). Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.</p>
      <p><b>Реализовать поддержку утилитой следующий ключей:</b></p>
      <ul>
          <li>-f - "fields" - выбрать поля (колонки)</li>
          <li>-d - "delimiter" - использовать другой разделитель</li>
          <li>-s - "separated" - только строки с разделителем</li>
      </ul>
  </li>
  <li>
      <p><b>OR CHANNEL</b></p>
      <p>Реализовать функцию, которая будет объединять один или более done-каналов в single-канал, если один из его составляющих каналов закроется.</p>
      <p>Очевидным вариантом решения могло бы стать выражение при использованием select, которое бы реализовывало эту связь, однако иногда неизвестно общее число done-каналов, с которыми вы работаете в рантайме. В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or-каналов, реализовывала бы весь функционал.</p>
      <p><b>ОПРЕДЕЛЕНИЕ ФУНКЦИИ:</b></p>
      <p>var or func(channels ...<- chan interface{}) <- chan interface{}</p>
      <p>Пример использования функции:</p>
        
```      
sig := func(after time.Duration) <- chan interface{} {
    c := make(chan interface{})
    go func() {
        defer close(c)
        time.Sleep(after)
    }()
    return c
}
  
start := time.Now()

<-or (
    sig(2*time.Hour),
    sig(5*time.Minute),
    sig(1*time.Second),
    sig(1*time.Hour),
    sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
```

  </li>
  <li>
      <p><b>ВЗАИМОДЕЙСТВИЕ С ОС</b></p>
      <p>Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:</p>
      <ul>
          <li>cd &lt;args&gt; - смена директории (в качестве аргумента могут быть то-то и то)</li>
          <li>pwd - показать путь до текущего каталога</li>
          <li>echo &lt;args&gt; - вывод аргумента в STDOUT</li>
          <li>kill &lt;args&gt; - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)</li>
          <li>ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*</li>
      </ul>
      <p>Так же требуется поддерживать функционал fork/exec-команд</p>
      <p>Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).</p>
      <p><i>Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение 
          в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике 
          и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).</i></p>
  </li>
  <li>
      <p><b>УТИЛИТА WGET</b></p>
      <p>Реализовать утилиту wget с возможностью скачивать сайты целиком.</p>
  </li>
  <li>
      <p><b>УТИЛИТА TELNET</b></p>
      <p>Реализовать простейший telnet-клиент.</p>
      <p><b>Примеры вызовов:</b></p>
      <ul>
          <li>go-telnet --timeout=10s host port</li>
          <li>go-telnet mysite.ru 8080</li>
          <li>go-telnet --timeout=3s 1.1.1.1 123</li>
      </ul>
      <p><b>Требования:</b></p>
      <ul>
          <li>Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP. После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT</li>
          <li>Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)</li>
          <li>При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться. При подключении к несуществующему сервер, программа должна завершаться через timeout</li>
      </ul>
  </li>
  <li>
    <p><b>HTTP-СЕРВЕР</b></p>
    <p>Реализовать HTTP-сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP-библиотекой.</p>
    <p>В рамках задания необходимо:</p>
    <ul>
        <li>Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.</li>
        <li>Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.</li>
        <li>Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.</li>
        <li>Реализовать middleware для логирования запросов.</li>
    </ul>
    <p><b>Методы API:</b></p>
    <ul>
        <li>POST /create_event</li>
        <li>POST /update_event</li>
        <li>POST /delete_event</li>
        <li>GET /events_for_day</li>
        <li>GET /events_for_week</li>
        <li>GET /events_for_month</li>
    </ul>
    <p>Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&amp;date=2019-09-09). В GET методах параметры передаются через queryString, в POST через тело запроса.</p>
    <p>В результате каждого запроса должен возвращаться JSON-документ содержащий либо {"result": "..."} в случае успешного выполнения метода, либо {"error": "..."} в случае ошибки бизнес-логики.</p>
    <p>В рамках задачи необходимо:</p>
    <ul>
        <li>Реализовать все методы.</li>
        <li>Бизнес логика НЕ должна зависеть от кода HTTP сервера.</li>
        <li>В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500.</li>
        <li>Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.</li>
    </ul>
  </li>
</ol>


[Задачи](#tasks-ru) / [Листинги](#listings-ru) / [Паттерны](#patterns-ru)

<a name="listings-ru"></a>
### Листинги

<p>Каждый <b>листинг</b> - один файл с соответствующим названием в папке listing.</p>
<ol>
  <li>
    <p>Что выведет программа? Объяснить вывод программы.</p>

```
package main
import (
  "fmt"
)
  
func main() {
  a := [5]int{76, 77, 78, 79, 80}
  var b []int = a[1:4]
  fmt.Println(b)
}
```

  </li>
  <li>
    <p>Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и порядок их вызовов.</p>

```
package main
 
import (
    "fmt"
)
 
func test() (x int) {
    defer func() {
        x++
    }()
    x = 1
    return
}
 
 
func anotherTest() int {
    var x int
    defer func() {
        x++
    }()
    x = 1
    return x
}
 
 
func main() {
    fmt.Println(test())
    fmt.Println(anotherTest())
}
```

  </li>
  <li>
    <p>Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.</p>

```
package main
 
import (
    "fmt"
    "os"
)
 
func Foo() error {
    var err *os.PathError = nil
    return err
}
 
func main() {
    err := Foo()
    fmt.Println(err)
    fmt.Println(err == nil)
}
```

  </li>
  <li>
    <p>Что выведет программа? Объяснить вывод программы.</p>

```
package main
 
func main() {
    ch := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            ch <- i
        }
    }()
 
    for n := range ch {
        println(n)
    }
}
```

  </li>
  <li>
    <p>Что выведет программа? Объяснить вывод программы.</p>

```
package main
 
type customError struct {
     msg string
}
 
func (e *customError) Error() string {
    return e.msg
}
 
func test() *customError {
     {
         // do something
     }
     return nil
}
 
func main() {
    var err error
    err = test()
    if err != nil {
        println("error")
        return
    }
    println("ok")
}
```

  </li>
  <li>
    <p>Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передаче их в качестве аргументов функции.</p>

```
package main
 
import (
  "fmt"
)
 
func main() {
  var s = []string{"1", "2", "3"}
  modifySlice(s)
  fmt.Println(s)
}
 
func modifySlice(i []string) {
  i[0] = "3"
  i = append(i, "4")
  i[1] = "5"
  i = append(i, "6")
}
```

  </li>
  <li>
    <p>Что выведет программа? Объяснить вывод программы.</p>

```
package main
 
import (
    "fmt"
    "math/rand"
    "time"
)
 
func asChan(vs ...int) <-chan int {
   c := make(chan int)
 
   go func() {
       for _, v := range vs {
           c <- v
           time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
      }
 
      close(c)
  }()
  return c
}
 
func merge(a, b <-chan int) <-chan int {
   c := make(chan int)
   go func() {
       for {
           select {
               case v := <-a:
                   c <- v
              case v := <-b:
                   c <- v
           }
      }
   }()
 return c
}
 
func main() {
 
   a := asChan(1, 3, 5, 7)
   b := asChan(2, 4 ,6, 8)
   c := merge(a, b )
   for v := range c {
       fmt.Println(v)
   }
}
```

  </li>
</ol>


[Задачи](#tasks-ru) / [Листинги](#listings-ru) / [Паттерны](#patterns-ru)

<a name="patterns-ru"></a>
### Паттерны

<p>Каждый <b>паттерн</b> - один файл с соответствующим названием в папке patterns. Тестовый запуск из файла main в папке patterns через интерактивный ввод.</p>
<ol>
  <li>Паттерн «фасад».</li>
  <li>Паттерн «строитель»</li>
  <li>Паттерн «посетитель»</li>
  <li>Паттерн «комманда»</li>
  <li>Паттерн «цепочка вызовов»</li>
  <li>Паттерн «фабричный метод»</li>
  <li>Паттерн «стратегия»</li>
  <li>Паттерн «состояние»</li>
</ol>


<a name="English"></a>
# English
<p id="en"><h3>Tasks on Lv2 at the Wildberries internship</h3></p>

[Tasks](#tasks-en) / [Listings](#listings-en) / [Patterns](#patterns-en)

<a name="tasks-en"></a>
### Tasks

<p>Each <b>task</b> is one folder with the appropriate name in the develop folder, launched via the command line (or unless otherwise stated, from the main.go file).</p>
<ol>
     <li>
         <p><b>BASIC TASK</b></p>
         <p>Create a program that prints the exact time using the NTP library. Initialize as go module. Use the library github.com/beevik/ntp. Write a program that prints the current time / exact time using this library.</p>
         <p><b>Requirements:</b></p>
         <ul>
             <li>The program must be formatted as a go module</li>
             <li>The program must correctly handle library errors: output them to STDERR and return a non-zero exit code to the OS</li>
         </ul>
     </li>
     <li>
         <p><b>UNPACKING TASK</b></p>
         <p>Create a Go function that performs primitive unpacking of a string containing repeated characters/runes, for example:</p>
         <ul>
             <li>"a4bc2d5e" => "aaaabccddddde"</li>
             <li>"abcd" => "abcd"</li>
             <li>"45" => "" (invalid string)</li>
             <li>"" => ""</li>
         </ul>
         <p><b>Additional:</b></p>
         <ul>
             <li>Implement support for escape sequences.</li>
             <li>For example:</li>
             <ul>
                 <li>qwe\4\5 => qwe45 (*)</li>
                 <li>qwe\45 => qwe44444 (*)</li>
                 <li>qwe\\5 => qwe\\\\\ (*)</li>
             </ul>
             <li>If an incorrect string was passed, the function should return an error. Write unit tests.</li>
         </ul>
     </li>
     <li>
         <p><b>SORT UTILITY</b></p>
         <p>Sort the lines in the file by analogy with the console utility sort (man sort - see the description and main parameters): the input is a file of unsorted lines, the output is a file with sorted lines.</p>
         <p><b>Implement the utility's support for the following keys:</b></p>
         <ul>
             <li>-k - specifying the column to be sorted (words in a line can act as columns, the default separator is space)</li>
             <li>-n - sort by numeric value</li>
             <li>-r - sort in reverse order</li>
             <li>-u - do not print duplicate lines</li>
         </ul>
         <p><b>Additional:</b> implement the utility support for the following keys:</p>
         <ul>
             <li>-M - sort by month name</li>
             <li>-b - ignore trailing spaces</li>
             <li>-c - check if the data is sorted</li>
             <li>-h - sort by numeric value, taking into account suffixes</li>
         </ul>
     </li>
     <li>
         <p><b>SEARCH FOR ANAGRAMS IN THE DICTIONARY</b></p>
         <p>Write a function to search for all sets of anagrams in a dictionary.</p>
         <p>For example:</p>
         <ul>
             <li>'penny', 'heel' and 'chopper' - belong to the same set</li>
             <li>'leaf', 'ingot' and 'table' - to another.</li>
         </ul>
         <p><b>Requirements:</b></p>
         <ul>
             <li>Input data for the function: a link to an array, each element of which is a word in Russian in utf8 encoding</li>
             <li>Output: link to the map of sets of anagrams</li>
             <li>Key is the first word of a set found in the dictionary. The value is a reference to an array, each element of which is a word from the set.</li>
             <li>The array must be sorted in ascending order.</li>
             <li>Sets of one element should not be included in the result.</li>
             <li>All words must be converted to lowercase.</li>
             <li>As a result, each word should appear only once.</li>
         </ul>
     </li>
     <li>
       <p><b>GREP UTILITY</b></p>
       <p>Implement the filtering utility by analogy with the console utility (man grep - see the description and main parameters).</p>
       <p><b>Implement the utility's support for the following keys:</b></p>
       <ul>
           <li>-A - "after" print +N lines after match</li>
           <li>-B - "before" print +N lines until match</li>
           <li>-C - "context" (A+B) print ±N lines around match</li>
           <li>-c - "count" (number of lines)</li>
           <li>-i - "ignore-case" (ignore case)</li>
           <li>-v - "invert" (instead of matching, exclude)</li>
           <li>-F - "fixed", exact match with string, not a pattern</li>
           <li>-n - "line num", print line number</li>
       </ul>
   </li>
   <li>
<li>
       <p><b>CUT UTILITY</b></p>
       <p>Implement a utility similar to the console command cut (man cut). The utility should receive lines via STDIN, split them by delimiter (TAB) into columns and output the requested ones.</p>
       <p><b>Implement the utility's support for the following keys:</b></p>
       <ul>
           <li>-f - "fields" - select fields (columns)</li>
           <li>-d - "delimiter" - use a different delimiter</li>
           <li>-s - "separated" - only lines with a separator</li>
       </ul>
   </li>
   <li>
       <p><b>OR CHANNEL</b></p>
       <p>Implement a function that will combine one or more done channels into a single channel if one of its component channels closes.</p>
       <p>The obvious solution would be to use a select expression to implement this relationship, but sometimes the total number of done channels you are working with at runtime is unknown. In this case, it is more convenient to use a call to a single function, which, having received one or more or-channels as input, would implement all the functionality.</p>
       <p><b>FUNCTION DEFINITION:</b></p>
       <p>var or func(channels ...<- chan interface{}) <- chan interface{}</p>
       <p>An example of using the function:</p>
        
```
sig := func(after time.Duration) <- chan interface{} {
     c := make(chan interface{})
     go func() {
         defer close(c)
         time.Sleep(after)
     }()
     return c
}
  
start := time.Now()

<-or (
     sig(2*time.Hour),
     sig(5*time.Minute),
     sig(1*time.Second),
     sig(1*time.Hour),
     sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
```

   </li>
   <li>
       <p><b>INTERACTION WITH THE OS</b></p>
       <p>You need to implement your own UNIX shell utility that supports a number of simple commands:</p>
       <ul>
           <li>cd &lt;args&gt; - change directory (the argument can be this and that)</li>
           <li>pwd - show the path to the current directory</li>
           <li>echo &lt;args&gt; - output the argument to STDOUT</li>
           <li>kill &lt;args&gt; - “kill” the process passed as an argument (example: such and such an example)</li>
           <li>ps - displays general information on running processes in the format *such and such format*</li>
       </ul>
       <p>It is also required to support the functionality of fork/exec commands</p>
       <p>Additionally, it is necessary to support the pipeline on pipes (linux pipes, example cmd1 | cmd2 | .... | cmdN).</p>
       <p><i>Shell is a regular console program that, when launched, displays a certain prompt in an interactive session
           to STDOUT and waits for user input via STDIN. After waiting for input, it processes the command according to its logic
           and, if necessary, displays the result on the screen. The interactive session is maintained until a quit command (such as \quit) is issued.</i></p>
   </li>
   <li>
       <p><b>WGET UTILITY</b></p>
       <p>Implement the wget utility with the ability to download entire sites.</p>
   </li>
   <li>
       <p><b>TELNET UTILITY</b></p>
       <p>Implement a simple telnet client.</p>
       <p><b>Call examples:</b></p>
       <ul>
           <li>go-telnet --timeout=10s host port</li>
           <li>go-telnet mysite.ru 8080</li>
           <li>go-telnet --timeout=3s 1.1.1.1 123</li>
       </ul>
       <p><b>Requirements:</b></p>
       <ul>
           <li>The program must connect to the specified host (ip or domain name + port) via TCP. After connecting, the STDIN of the program must be written to the socket, and the data received from the socket must be output to STDOUT</li>
           <li>Optionally, you can pass a timeout for connecting to the server to the program (via the --timeout argument, default 10s)</li>
           <li>When you press Ctrl+D, the program should close the socket and exit. If the socket is closed on the server side, the program must also exit. When connecting to a non-existent server, the program must terminate after timeout</li>
       </ul>
   </li>
   <li>
     <p><b>HTTP SERVER</b></p>
     <p>Implement an HTTP server for working with a calendar. As part of the assignment, you must work strictly with the standard HTTP library.</p>
     <p>As part of the task you must:</p>
     <ul>
         <li>Implement helper functions for serializing domain objects to JSON.</li>
         <li>Implement auxiliary functions for parsing and validating parameters of the /create_event and /update_event methods.</li>
         <li>Implement HTTP handlers for each API method using helper functions and domain objects.</li>
         <li>Implement middleware for request logging.</li>
     </ul>
     <p><b>API methods:</b></p>
     <ul>
         <li>POST /create_event</li>
         <li>POST /update_event</li>
         <li>POST /delete_event</li>
         <li>GET /events_for_day</li>
         <li>GET /events_for_week</li>
         <li>GET /events_for_month</li>
     </ul>
     <p>Parameters are sent in the form www-url-form-encoded (i.e. regular user_id=3&amp;date=2019-09-09). In GET methods, parameters are passed through queryString, in POST through the request body.</p>
     <p>As a result of each request, a JSON document must be returned containing either {"result": "..."} in case of successful execution of the method, or {"error": "..."} in case of a business logic error. </p>
     <p>As part of the task you must:</p>
     <ul>
         <li>Implement all methods.</li>
         <li>Business logic should NOT depend on the HTTP server code.</li>
         <li>In case of a business logic error, the server must return HTTP 503. In case of an input data error (invalid int, for example), the server must return HTTP 400. In case of other errors, the server must return HTTP 500.</li>
         <li>The web server must run on the port specified in the config and log every processed request.</li>
     </ul>
   </li>
</ol>



[Tasks](#tasks-en) / [Listings](#listings-en) / [Patterns](#patterns-en)

<a name="listings-en"></a>
### Listings

<p>Each <b>listing</b> is one file with the corresponding name in the listing folder.</p>
<ol>
   <li>
     <p>What will the program output? Explain the output of the program.</p>

```
package main
import(
   "fmt"
)
  
func main() {
   a := [5]int{76, 77, 78, 79, 80}
   var b []int = a[1:4]
   fmt.Println(b)
}
```

   </li>
   <li>
     <p>What will the program output? Explain the output of the program. Explain how defers work and the order in which they are called.</p>

```
package main
 
import (
     "fmt"
)
 
func test() (x int) {
     defer func() {
         x++
     }()
     x = 1
     return
}
 
 
func anotherTest() int {
     var x int
     defer func() {
         x++
     }()
     x = 1
     return x
}
 
 
func main() {
     fmt.Println(test())
     fmt.Println(anotherTest())
}
```

   </li>
   <li>
     <p>What will the program output? Explain the output of the program. Explain the internal structure of interfaces and their difference from empty interfaces.</p>

```
package main
 
import (
     "fmt"
     "os"
)
 
func Foo() error {
     var err *os.PathError = nil
     return err
}
 
func main() {
     err := Foo()
     fmt.Println(err)
     fmt.Println(err == nil)
}
```

   </li>
   <li>
     <p>What will the program output? Explain the output of the program.</p>

```
package main
 
func main() {
     ch := make(chan int)
     go func() {
         for i := 0; i < 10; i++ {
             ch <- i
         }
     }()
 
     for n := range ch {
         println(n)
     }
}
```

   </li>
   <li>
     <p>What will the program output? Explain the output of the program.</p>

```
package main
 
type customError struct {
      msg string
}
 
func (e *customError) Error() string {
     return e.msg
}
 
func test() *customError {
      {
          //do something
      }
      return nil
}
 
func main() {
     var err error
     err = test()
     if err != nil {
         println("error")
         return
     }
     println("ok")
}
```

   </li>
   <li>
     <p>What will the program output? Explain the output of the program. Tell us about the internal structure of slices and what happens when passing them as function arguments.</p>

```
package main
 
import(
   "fmt"
)
 
func main() {
   var s = []string{"1", "2", "3"}
   modifySlice(s)
   fmt.Println(s)
}
 
func modifySlice(i []string) {
   i[0] = "3"
   i = append(i, "4")
   i[1] = "5"
   i = append(i, "6")
}
```

   </li>
   <li>
     <p>What will the program output? Explain the output of the program.</p>

```
package main
 
import (
     "fmt"
     "math/rand"
     "time"
)
 
func asChan(vs ...int) <-chan int {
    c := make(chan int)
 
    go func() {
        for _, v := range vs {
            c <- v
            time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
       }
 
       close(c)
   }()
   return c
}
 
func merge(a, b <-chan int) <-chan int {
    c := make(chan int)
    go func() {
        for {
            select {
                case v := <-a:
                    c <- v
               case v := <-b:
                    c <- v
            }
       }
    }()
  return c
}
 
func main() {
 
    a := asChan(1, 3, 5, 7)
    b := asChan(2, 4 ,6, 8)
    c := merge(a, b)
    for v := range c {
        fmt.Println(v)
    }
}
```

   </li>
</ol>


[Tasks](#tasks-en) / [Listings](#listings-en) / [Patterns](#patterns-en)

<a name="patterns-en"></a>
### Patterns

<p>Each <b>pattern</b> is one file with the corresponding name in the patterns folder. Test run from the main file in the patterns folder via interactive input.</p>
<ol>
  <li>Facade pattern</li>
  <li>Builder pattern</li>
  <li>Visitor pattern</li>
  <li>Command pattern</li>
  <li>Chain of responsibility pattern</li>
  <li>Factory method pattern</li>
  <li>Strategy pattern</li>
  <li>State pattern</li>
</ol>
