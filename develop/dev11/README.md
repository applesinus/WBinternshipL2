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

<hr>
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