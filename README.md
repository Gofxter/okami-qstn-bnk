# Реализация сервиса Question Bank

7 часов + 19:00 - N

### Запуск локально
Два сценария для запуска сервиса
```
$ task local-run
```
или
```
$ migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up
$ !FIXME ДОБАВИТЬ ИНТ ТЕСТЫ
$ go run cmd/server/main.go
```