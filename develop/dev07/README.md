<p><b>OR CHANNEL</b></p>
<p>Реализовать функцию, которая будет объединять один или более done-каналов в single-канал, если один из его составляющих каналов закроется.</p>
<p>Очевидным вариантом решения могло бы стать выражение при использованием select, которое бы реализовывало эту связь, однако иногда неизвестно общее число done-каналов, с которыми вы работаете в рантайме. В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or-каналов, реализовывала бы весь функционал.</p>
<p><b>ОПРЕДЕЛЕНИЕ ФУНКЦИИ:</b></p>
<p>var or func(channels ...<- chan interface{}) <- chan interface{}</p>
    
[Пример использования функции ниже](#example)


<p><b>OR CHANNEL</b></p>
<p>Implement a function that will combine one or more done channels into a single channel if one of its component channels closes.</p>
<p>The obvious solution would be to use a select expression to implement this relationship, but sometimes the total number of done channels you are working with at runtime is unknown. In this case, it is more convenient to use a call to a single function, which, having received one or more or-channels as input, would implement all the functionality.</p>
<p><b>FUNCTION DEFINITION:</b></p>
<p>var or func(channels ...<- chan interface{}) <- chan interface{}</p>
    
[An example of using the function:](#example)

<a name="example"></a>


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
