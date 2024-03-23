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

<hr>
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