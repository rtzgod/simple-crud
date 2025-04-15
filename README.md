# Simple CRUD Notes Application
Это простой RESTful API сервис для управления заметками, разработанный на Go. Приложение следует принципам чистой архитектуры и предоставляет базовые CRUD операции для заметок.

## Обзор проекта
Приложение позволяет пользователям:
- Создавать новые заметки
- Получать список всех заметок
- Обновлять существующие заметки
- Удалять заметки

Бэкенд разработан на Go с использованием маршрутизатора Chi для обработки HTTP-запросов и PostgreSQL для хранения данных. Приложение контейнеризировано с помощью Docker для удобного развертывания.

## Требования
Для запуска приложения вам понадобится:
- Docker и Docker Compose
- Go 1.24.2 (необязательно)
- PostgreSQL (необязательно)

## Как запустить

### Использование Docker (Рекомендуется)

1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/rtzgod/simple-crud.git
   cd simple-crud
   ```
2. Скопируйте файл .env.example в файл .env:
   ```bash
   cp .env.example .env
   ```

3. Запустите приложение с помощью Docker Compose:
   ```bash
   docker-compose up -d
   ```

4. API будет доступен по адресу [`http://localhost:8080`](internal/handler/note.go)

### Локальный запуск

1. Убедитесь, что PostgreSQL запущен и доступен
2. Обновите URL подключения к базе данных в [`configs/local.yaml`](configs/local.yaml)
3. Запустите приложение:
   ```bash
   go run cmd/main.go -cfg "configs/local.yaml"
   ```
   или если есть .env файл:
   ```bash
   go run cmd/main.go
   ```

## API Endpoints

### Создание заметки
- **URL**: `/notes`
- **Метод**: `POST`
- **Тело запроса**:
  ```json
  {
    "title": "Заголовок заметки",
    "content": "Содержание заметки"
  }
  ```
- **Ответ**: Возвращает ID созданной заметки
  ```json
  {
    "id": 1
  }
  ```

### Получение всех заметок
- **URL**: `/notes`
- **Метод**: `GET`
- **Ответ**: Возвращает массив заметок
  ```json
  [
    {
      "id": 1,
      "title": "Заголовок заметки",
      "content": "Содержание заметки"
    }
  ]
  ```

### Обновление заметки
- **URL**: `/notes/{id}`
- **Метод**: `PUT`
- **Тело запроса**:
  ```json
  {
    "title": "Обновленный заголовок",
    "content": "Обновленное содержание"
  }
  ```
- **Ответ**: Статус 200 OK в случае успеха

### Удаление заметки
- **URL**: `/notes/{id}`
- **Метод**: `DELETE`
- **Ответ**: Статус 200 OK в случае успеха

## Структура проекта

```
├── cmd/
│   └── main.go         # Точка входа в приложение
├── configs/
│   └── local.yaml      # Файл конфигурации
├── db/
│   └── migrations/     # Миграции базы данных
├── internal/
│   ├── config/         # Обработка конфигурации
│   ├── handler/        # HTTP обработчики
│   ├── models/         # Модели данных
│   ├── repository/     # Операции с базой данных
│   └── service/        # Бизнес-логика
├── docker-compose.yml  # Конфигурация Docker Compose
└── Dockerfile          # Инструкции для сборки Docker
```

## Используемые технологии

- **Go**: Язык программирования
- **Chi Router**: HTTP маршрутизация
- **PostgreSQL**: База данных
- **SQLx**: SQL инструментарий
- **Golang-migrate**: Миграции базы данных
- **Docker**: Контейнеризация
- **Cleanenv**: Управление конфигурацией