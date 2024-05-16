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
3. Создано отдельное второе backend-приложение, взаимодейсвтующее с MongoDB
    - Реализованы CRUD операции для конференций и докладов на этих конференций.
    - Скрипт для создания коллекций, который запускается при деплое проекта (`make run`) и скрипт для первичного заполнения коллекции тестовыми документами, который можно выполнить после развёртывания с помощью команды `make mongo-db-fill`.
    - Создана отдельная от сервиса *users* OpenAPI 3.0 спецификация в `docs/conferences`, специцикация *users* перемещена в `docs/users`.
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