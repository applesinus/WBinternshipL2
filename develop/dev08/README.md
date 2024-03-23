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
<p><i>Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).</i></p>

<hr>
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
<p><i>Shell is a regular console program that, when launched, displays a certain prompt in an interactive session to STDOUT and waits for user input via STDIN. After waiting for input, it processes the command according to its logic and, if necessary, displays the result on the screen. The interactive session is maintained until a quit command (such as \quit) is issued.</i></p>