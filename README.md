# to-do-list-go
REST API для системы управления задачами (To-Do List). API позволяет создавать, просматривать, обновлять и удалять задачи.

## Используемые технологии
- Go: Язык программирования, на котором написан проект.
- PostgreSQL: База данных для хранения информации о задачах.
- Docker: Используется для контейнеризации сервиса, что облегчает его развертывание и управление зависимостями.

## Эндпоинты

### Создание задачи

- **Метод:** POST /tasks
- **Описание:** Создать новую задачу.
- **Запрос:**
   - **Заголовки:**
      - Content-Type: application/json
   - **Тело:**
     ```json
     {
       "title": "string",
       "description": "string",
       "due_date": "string (RFC3339 format)"
     }
     ```
- **Ответ:**
   - **Успех (201 Created):**
     ```json
     {
       "id": "int",
       "title": "string",
       "description": "string",
       "due_date": "string (RFC3339 format)",
       "created_at": "string (RFC3339 format)",
       "updated_at": "string (RFC3339 format)"
     }
     ```
   - **Ошибка (400 Bad Request):** Неправильный формат данных.
   - **Ошибка (500 Internal Server Error):** Проблема на сервере.

### Просмотр списка задач

- **Метод:** GET /tasks
- **Описание:** Получить список всех задач.
- **Запрос:**
   - **Заголовки:**
      - Content-Type: application/json
- **Ответ:**
   - **Успех (200 OK):**
     ```json
     [
       {
         "id": "int",
         "title": "string",
         "description": "string",
         "due_date": "string (RFC3339 format)",
         "created_at": "string (RFC3339 format)",
         "updated_at": "string (RFC3339 format)"
       }
     ]
     ```
   - **Ошибка (500 Internal Server Error):** Проблема на сервере.

### Просмотр задачи

- **Метод:** GET /tasks/{id}
- **Описание:** Получить задачу по ID.
- **Запрос:**
   - **Параметры пути:**
      - id: ID задачи (int)
   - **Заголовки:**
      - Content-Type: application/json
- **Ответ:**
   - **Успех (200 OK):**
     ```json
     {
       "id": "int",
       "title": "string",
       "description": "string",
       "due_date": "string (RFC3339 format)",
       "created_at": "string (RFC3339 format)",
       "updated_at": "string (RFC3339 format)"
     }
     ```
   - **Ошибка (404 Not Found):** Задача не найдена.
   - **Ошибка (500 Internal Server Error):** Проблема на сервере.

### Обновление задачи

- **Метод:** PUT /tasks/{id}
- **Описание:** Обновить задачу по ID.
- **Запрос:**
   - **Параметры пути:**
      - id: ID задачи (int)
   - **Заголовки:**
      - Content-Type: application/json
   - **Тело:**
     ```json
     {
       "title": "string",
       "description": "string",
       "due_date": "string (RFC3339 format)"
     }
     ```
- **Ответ:**
   - **Успех (200 OK):**
     ```json
     {
       "id": "int",
       "title": "string",
       "description": "string",
       "due_date": "string (RFC3339 format)",
       "created_at": "string (RFC3339 format)",
       "updated_at": "string (RFC3339 format)"
     }
     ```
   - **Ошибка (400 Bad Request):** Неправильный формат данных.
   - **Ошибка (404 Not Found):** Задача не найдена.
   - **Ошибка (500 Internal Server Error):** Проблема на сервере.

### Удаление задачи

- **Метод:** DELETE /tasks/{id}
- **Описание:** Удалить задачу по ID.
- **Запрос:**
   - **Параметры пути:**
      - id: ID задачи (int)
   - **Заголовки:**
      - Content-Type: application/json
- **Ответ:**
   - **Успех (204 No Content):** Задача удалена.
   - **Ошибка (404 Not Found):** Задача не найдена.
   - **Ошибка (500 Internal Server Error):** Проблема на сервере.

## Переменные окружения

Пример .env файла:

```
PORT=8888
DB_USER=postgres
DB_PASSWORD=postgres
DB_HOST=localhost
DB_PORT=5432
DB_NAME=postgres
```

## Требования

- Go 1.22+
- PostgreSQL

## Начало работы

1. Склонируйте репозиторий.
2. Создайте файл `.env` на основе примера, приведенного выше.
3. Запуск: 
   - Если хотите запустить проект без **Docker**, выполните следующие команды из корневой папки `to-do-list-go`:

    ```bash
    go mod tidy
    cd ./internal/database/migrations && goose postgres "postgresql://пользователь:пароль@хост:порт/название базы?sslmode=disable" up
    go run cmd/main.go
    ```

   - Если у вас установлена утилита **`make`**, для запуска проекта выполните команду из корневой папки `to-do-list-go`:

    ```bash
    make
    ```

   - Для запуска проекта с использованием **Docker**, выполните следующие команды из папки `to-do-list-go`:

    ```bash
    docker network create todo_network && docker-compose up -d
    ```

Теперь проект запущен и готов к использованию!

## Тестирование

Проект включает интеграционные тесты с полным покрытием логики обработки запросов и бизнес-логики.

Для запуска тестов используйте команду:

```bash
go test -v ./...
```

или 

```bash
make test
```
