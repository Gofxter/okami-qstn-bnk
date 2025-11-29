# Реализация сервиса Question Bank

```shell
  task gen-migrations        генерация миграция
  task migrate               migrate up
  task migrate-down:         migrate down
  task gen-mock:             генерация моков на сервис слой и БД
  task test:                 запуск тестов
  task local-run:            локальный запуск
  task docker-run            запуск из докера
```

### Сценарии для запуска сервиса
Локально:
```shell
  migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up
  go run cmd/server/main.go
```
Докер:
```shell
    echo docker
```

### Запросы curl
```shell
    echo curl
```

### Остаток:
4) Сделать выдачу теста ?!?!
5) Документация, ее генерация !!!
6) Обертка в докер !!!
7) удаление опций при удалении вопроса!!!!!
8) Обновление опций при обновлении вопроса!!!!

### Стоит улучшить, доработать
1) ```pkg/type/type.go``` - Поменять валидацию полей(усл.:backend_junior разбивать на 2 части, валидировать по отдельности)
2) Доработать структуры ответов