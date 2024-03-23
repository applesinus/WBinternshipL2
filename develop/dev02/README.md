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

<hr>
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