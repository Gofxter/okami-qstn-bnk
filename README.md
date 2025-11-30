# Реализация сервиса Question Bank
#### Документация ```http://localhost:1606/swagger```
#### Команды task:
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
1) Обертка в докер !!!
2) Спека
3) Доработать структуры ответов

### Стоит улучшить, доработать
1) ```pkg/type/type.go``` - Поменять валидацию полей(усл.:backend_junior разбивать на 2 части, валидировать по отдельности)