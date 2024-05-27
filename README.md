## Условие

**Сайт конференции**
Приложение должно содержать следующие данные:

- Пользователь
- Доклад
- Конференция
  Реализовать API:
- Создание нового пользователя
- Поиск пользователя по логину
- Поиск пользователя по маске имя и фамилии
- Создание доклада
- Получение списка всех докладов
- Добавление доклада в конференцию
- Получение списка докладов в конференции

## Задание 01: Проектирование программной системы

docker-compose файл для запуска structurize lite вместе с worcspace.dsl в директории docs.

## Задание 02: Stateful сервис для RDBMS

1. В рамках второго задания были учтены архитектурные рекомендации, полученные в рамках проверки первого задания.
   - Согласно паттерну bounded context были выделенны модели confernece и users.
   - Добавлен прокси сервис для выделения единой точки взаимодействия с системой.
2. Были внесены соответсвующие изменения в диаграмы проекта.
3. Построено backend-приложение, взаимодействующее с postgres:
   - В рамках практической работы для простоты и наглядности, сервисы расположены в одном репозитории и находятся в реализованы в рамках одной общей структуры проекта, однако разнесены в достаточной степени друг от друга и разделены.
   - Чтобы развернуть сервис users и базу выполнить `make run`.
   - Скрипты для базы в директории `scripts`. Скрипт для создания таблицы, выполняется при поднятии окружения, заполнение тестовыми данными выполняется полсле развёртывания с помощью команды `make db-fill`.
   - `swagger.yml` файл с OpenAPI 3.0 спецификацией находится в директории `docs`,
   - После развёртывания все ручки можно просмотреть по адресу `http://localhost:8080/swagger/index.html`.
   - Пароль хранится в хэшированном виде.

## Задание 03: Stateful сервис для NoSQL

1. Внесены изменения в архитектуру в соотвествтии в условиями 3-го задания.
   - Выделена отдельная NoSQL база данных на основе MongoDB 7 для хранения данных о конференциях и докладах. Отражена в C4 диаграммах.
   - Так же в диаграммах отражено новое взаимодействие сервиса конференций с MongoDB.
2. Создано отдельное второе backend-приложение, взаимодейсвтующее с MongoDB
   - Реализованы CRUD операции для конференций и докладов на этих конференций.
   - Скрипт для создания коллекций, который запускается при деплое проекта (`make run`) и скрипт для первичного заполнения коллекции тестовыми документами, который можно выполнить после развёртывания с помощью команды `make mongo-db-fill`.
   - Создана отдельная от сервиса _users_ OpenAPI 3.0 спецификация в `docs/conferences`, специцикация _users_ перемещена в `docs/users`.
   - После развёртывания эндпоинты можно просмотреть по адресу `http://localhost:8081/swagger/index.html`.
   - Оба микросервиса - users и conferences, а так же базы - mongo и postgres разворачиваются внутри контейнров при помощи docker-compose.

## Задание 04: Аутентификация

1. В сервис пользователей добавлен эндпоинт для аутентификации - GET /api/auth
2. Для всех эндпоинтов сервиса конференций и для всех эндпоинтов (кроме создания пользователя) сервиса пользователей добавлено обязательное введение JWToken, для авторизации пользователя.
3. Для проверк работоспособности, выполнить следующие действия:
   1. Развернуть сервисы и БД - `make run`
   2. Заполнить базы тестовыми данными - `make db-fill`, `make mongo-db-fill`
   3. В swagger сервиса пользователей получить JWToken - `http://localhost:8080/swagger/index.html#/auth/get_api_auth`. В поле Authorization ввести `Basic dGVzdF91c2VybmFtZTp0ZXN0X3Bhc3N3b3Jk` (Basic auth сгененрированный из кредов тестового пользователя test_username test_password)
   4. Полученный токен применять во всех остальных эндпоинтах в поле Authorization в формате `Bearer <полученный токен>`.
   5. Чтобы завершить работу сервисов и баз данных - `make stop`

## Задание 05: Data Cache

1. Добавлен redis
2. В сервис пользователей добавлен компонент/прослойка реализующая шаблоны сквозного чтения и сковозной записи, таким образом что хэндлеры запросов получения пользователя по логину, создания пользователя и обновления пользователя не обращаются к базе postgres, а используют методы адаптера кэширования `GetUser`, `FirstSetUser` и `SetUser`, которые в свою очередь выполняют все операции по обновлению кэша записи в базу postgres. Время жизни ключей в redis установлено на 1 минуту.
3. Проведены бенчмарки при помощи утилиты wrk на эндпоинте получения пользователя по логину (`/api/v1/users/username/{username}`). Результаты замеров занесены [performance.md](docs/performance.md). По замерам виден сокращение средних значений времени отклика, максимального времени отклика, в версии без кэша оно достигает 1 секунды, тогда как в версии с кэшем нет. Так же видно что общее количество запросов прошедших за заданные 2 минуты бэнчмарка в кэшированной версии в среднем больше на 15-20%.

## Задание 06: Circuit Breaker

1. В рамках этого задания был создан новый сервис проксирующий запросы в ранее созданные сервисы пользователей и конференций. Для сервиса добавлена swagger спецификация включающая в себя эндпоинты сервиса пользователей и сервиса конференций.
2. Так же в этом сервисе релизован паттерн Circuit Breaker. Он переходит в состояние *Открыт* при недоступности целевого сервиса или при получении внутренних ошибок сервера от сервиса (статус код 500).
3. Для наглядности работы и удобсвтва проверки паттерна добавлены логи по которым можно отслеживать состояние.
4. В следующем примере работы circuit breaker все запросы делаются через сваггер api-gateway - `http://localhost:8082/swagger/index.html`; пороговое значние ошибок - 2, пороговое значение успешных запросов в состоянии *полуоткрытое* - 2, таймаут нахождения в *открытом состоянии* - 10с, время достижения порогового значения ошибок в закрытом состоянии - 60с:
   1) Делаем запрос на получение списка пользователей - `16:35:04.579700 circuitbreaker.go:94: [Info] Circuit Breaker is Closed`
   2) Имитируем проблемы в работе сервиса пользователей - останаливаем его командой `docker stop user-service`
   3) Делаем запрос на получение списка пользователей -  `16:39:02.974035 circuitbreaker.go:89: [Info] Circuit Breaker is Closed`
   4) Делаем ещё один запрос на получение списка пользователей через малый промежуток вермени - `16:39:49.717647 circuitbreaker.go:83: [Info] Circuit Breaker is (Closed -> Opened)`
   5) Истекает таймаут *Открытого состояния* - `16:39:59.727621 circuitbreaker.go:105: [Info] Circuit Breaker is (Opened -> Half Opened)`
   6) Делаем ещё один запрос на получение списка пользователей - `16:40:15.979442 circuitbreaker.go:53: [Info] Circuit Breaker is (Half Opened -> Opened)`
   7) Имитируем восстановление работы сервиса пользователей - запускаем его командой `docker start user-service`
   8) Делаем ещё один запрос на получение списка пользователей - `16:43:05.582523 circuitbreaker.go:67: [Info] Circuit Breaker is Half Opened`
   9) Делаем ещё один запрос на получение списка пользователей - `16:43:49.312595 circuitbreaker.go:64: [Info] Circuit Breaker is (Half Opened -> Closed)`
   10) Делаем ещё один запрос на получение списка пользователей - `16:44:36.393038 circuitbreaker.go:94: [Info] Circuit Breaker is Closed`
5. В предыдущей проверке можно наблюдать что при достижении порогового значения ошибок (2) в заданное время (60с) circuit breaker перешёл в *состояние открытое*, затем по истечению таймаута окрытого состояния переходит в *полуоткрытое состояние*, и при повторном запросе с ошибкой обратно попадает в *открытое состояние*. После двух успешных запросов состояние переходит в *закрытое* и при последующих успшных запросах остаётся в нём.
6. В следующем примере работы circuit breaker все запросы делаются через сваггер api-gateway - `http://localhost:8082/swagger/index.html`; пороговое значние ошибок - 2, пороговое значение успешных запросов в состоянии *полуоткрытое* - 2, таймаут нахождения в *открытом состоянии* - 10с, время достижения порогового значения ошибок в закрытом состоянии - 60с::
   1) Делаем запрос на получение списка пользователей - `16:45:45.261015 circuitbreaker.go:94: [Info] Circuit Breaker is Closed`
   2) Имитируем проблемы в работе сервиса пользователей - останаливаем его командой `docker stop user-service`
   3) Делаем запрос на получение списка пользователей - `16:46:15.751019 circuitbreaker.go:89: [Info] Circuit Breaker is Closed`
   4) Делаем запрос на получение списка пользователей через значительный промежуток времени - `16:48:05.313362 circuitbreaker.go:89: [Info] Circuit Breaker is Closed`
   5) Имитируем восстановление работы сервиса пользователей - запускаем его командой `docker start user-service` 
   6) Делаем запрос на получение списка пользователей - `16:49:17.253944 circuitbreaker.go:89: [Info] Circuit Breaker is Closed`
7. В предыдущей проверке можно наблюдать что при получении повторной ошибки в промежутках времени больших заданого периода достижения порогового значения ошибок circuit breaker остаётся в *закрытом состоянии*. Таким образом *открытое состояние* достигается только при превышении порогового значения ошибок (2) в минуту.