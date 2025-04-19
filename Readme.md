
#  ToDo API
Простой таск менеджер для управления задачами с использованием Fiber, PostgreSQL.

# Функционал
- Создание задачи.
- Получение задач.
- Удаление задач.
- Обновление задач

# Быстрый старт
1. Клонировать репозиторий 
```bash
git clone https://github.com/Vladk39/ToDoApp
cd ToDoApp
```
2. Настроить .env файл
```bash
DBconfig=postgres://user:password@localhost:5432/dbname?sslmode=disable
SERVERPORT= Ваш порт на котором будет сервер
```
# Запуск приложения
1. Запуск через go run
```bash
go mod tidy
go run ./cmd
```
2. Сборка бинарного файла
```bash
go mod tidy
go build -o ToDoApp ./cmd
./ToDoApp  
```
Для Winows используется команда
```bash
ToDoApp.exe
```

# Примеры HTTP Запросов
## Создать задачу
```http
POST /tasks
Content-Type: application/json

{
  "title": "Задача 1",
  "description": "Описание задачи",
  "status": "new"
}
```
# Получить все задачи
```http
GET /tasks
```
# Обновить задачу
```http
PUT /tasks/:id
Content-Type: application/json

{
  "title": "Обновленное название",
  "description": "Новое описание",
  "status": "in_progress"
}
```
# Удалить задачу
```http
DELETE /tasks/:id
```
# Допустимые статусы задач
- new
-in_progress
-done

# Подробная документация
```http
http://localhost:port/swagger
```
# Запуск с Docker
1. Клонировать репозиторий 
```bash
git clone https://github.com/Vladk39/ToDoApp
cd ToDoApp
```
2. Создать свой .env файл
```bash
DBconfig=postgres://user:password@localhost:5432/dbname?sslmode=disable
SERVERPORT= Ваш порт на котором будет сервер
```
3. Создать Docker образ
```bash
docker build -t ToDoApp .
```
2. Запустить с контейнер с флагом, где указан собственный env файл
```bash
docker run --env-file .env -p 8080:8080 todoapp
```