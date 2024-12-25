module github.com/jorgeperezr/GoCommerce

go 1.21

require (
	GoCommerce/backend v0.0.0-00010101000000-000000000000
	github.com/golang-jwt/jwt/v4 v4.5.1
	github.com/gorilla/mux v1.8.1
	github.com/joho/godotenv v1.5.1
	golang.org/x/crypto v0.17.0
	gorm.io/driver/postgres v1.5.11
	gorm.io/driver/sqlite v1.5.7
	gorm.io/gorm v1.25.12
)

replace GoCommerce/backend => ./

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
