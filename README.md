### Info:
Технологический стек:
Фреймворк для API: Gin (самый популярный) или Fiber (очень быстрый, похож на Express.js).  
База данных: PostgreSQL (лучший выбор для работы с JSONB и транзакциями).  
ORM / Работа с БД: GORM (для быстрого написания кода) или sqlx / sqlc (для большего контроля над SQL).  
Миграции: golang-migrate.  
Валидация: go-playground/validator (стандарт в Go).  
Аутентификация: jwt-go или golang-jwt.  

### Cтруктура проекта
```text
warehouse-project/
├── cmd/
│   └── api/
│       └── main.go         # Точка входа
├── internal/
│   ├── models/             # Структуры БД (Warehouse, Product, User)
│   ├── handlers/           # HTTP-обработчики (контроллеры)
│   ├── repository/         # Работа с базой данных (SQL запросы)
│   ├── service/            # Бизнес-логика (перемещения, проверки)
│   └── middleware/         # Проверка прав доступа и JWT
├── pkg/                    # Вспомогательный код (конфиги, логгер)
├── migrations/             # SQL файлы миграций
├── docker-compose.yaml
└── go.mod
```
Внедряю не всё сразу:  
V1: Обычный CRUD (создал склад, удалил склад).  
V2: Добавить авторизацию и связь "Админ - Склад".  
V3: Реализовать перемещение товаров через Транзакции.  
V4: Добавить JSONB характеристики и Пагинацию.  
V5: Добавить Логирование действий.  
### QuickStart
Инициализация проекта на Go  
go mod init warehouse-app  
go get github.com/gin-gonic/gin  
go get gorm.io/gorm  
go get gorm.io/driver/postgres  

docker-compose up -d  
go run main.go  
