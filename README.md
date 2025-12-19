# Реализация сервиса Question Bank
#### Команды task:
```shell
  task gen-migrations        генерация миграций
  task migrate               migrate up
  task migrate-down          migrate down
  task gen-mock              генерация моков на сервис слой и БД
  task test-coverage         запуск тестов с отчётом покрытия
  task local-run             локальный запуск (Postgres на localhost:5432)
  task docker-up             запуск стека через docker-compose (infra/docker-compose.yml)
  task docker-down           остановка стека docker-compose
  task docker-down-volumes   остановка стека docker-compose с удалением volumes
```

### Результаты тестов и покрытие кода

Тесты запускаются командой:

```bash
  task test-coverage
```

### Сценарии для запуска сервиса
Локально:
```shell
  migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up
  CONFIG_PATH=config/config.yaml go run cmd/server/main.go
```
Докер:
```shell
  cd infra
  docker compose up --build
```

### Запросы curl
#### 1. CRUD для вопросов /questions

**Создать вопрос (POST /question-bank/questions):**
```bash
curl -X POST http://localhost:1606/question-bank/questions \
  -H "Content-Type: application/json" \
  -d '{
    "role": "backend_junior",
    "topic": "Go basics",
    "type": "single_choice",
    "options": [{"text": "Go", "is_correct": true}, {"text": "Java", "is_correct": false}],
    "difficulty": 2,
    "text": "На каком языке пишется go.mod?"
  }'
```

**Получить вопрос по ID (GET /question-bank/questions/{id}):**
```bash
curl http://localhost:1606/question-bank/questions/{id}
```

**Получить список вопросов с фильтрами (GET /question-bank/questions?role=backend_junior&topic=Go%20basics&difficulty=2):**
```bash
curl "http://localhost:1606/question-bank/questions?role=backend_junior&topic=Go%20basics&difficulty=2"
```

**Обновить вопрос (PUT /question-bank/questions/{id}):**
```bash
curl -X PUT http://localhost:1606/question-bank/questions/{id} \
  -H "Content-Type: application/json" \
  -d '{
    "role": "backend_junior",
    "topic": "Go basics",
    "difficulty": 3,
    "text": "Обновленный текст вопроса"
  }'
```

**Удалить вопрос (DELETE /question-bank/questions/{id}):**
```bash
curl -X DELETE http://localhost:1606/question-bank/questions/{id}
```

---
#### 2. CRUD для шаблонов /templates

**Создать шаблон (POST /question-bank/templates):**
```bash
curl -X POST http://localhost:1606/question-bank/templates \
  -H "Content-Type: application/json" \
  -d '{
    "role": "frontend_junior",
    "purpose": "mock_interview"
  }'
```

**Получить шаблон по ID (GET /question-bank/templates/{id}):**
```bash
curl http://localhost:1606/question-bank/templates/{id}
```

**Получить список шаблонов (GET /question-bank/templates?role=backend_junior&purpose=skills_assessment):**
```bash
curl "http://localhost:1606/question-bank/templates?role=backend_junior&purpose=skills_assessment"
```

**Обновить шаблон (PUT /question-bank/templates/{id}):**
```bash
curl -X PUT http://localhost:1606/question-bank/templates/{id} \
  -H "Content-Type: application/json" \
  -d '{
    "role": "backend_junior",
    "purpose": "skills_assessment"
  }'
```

**Удалить шаблон (DELETE /question-bank/templates/{id}):**
```bash
curl -X DELETE http://localhost:1606/question-bank/templates/{id}
```

---
#### 3. Инстанцирование теста (POST /question-bank/tests/instantiate)

**Сгенерировать тест по шаблону:**
```bash
curl -X POST http://localhost:1606/question-bank/tests/instantiate \
  -H "Content-Type: application/json" \
  -d '{
    "template_id": "YOUR_TEMPLATE_UUID"
  }'
```

### Стоит улучшить, доработать
1) ```pkg/type/type.go``` - Поменять валидацию полей(усл.:backend_junior разбивать на 2 части, валидировать по отдельности)