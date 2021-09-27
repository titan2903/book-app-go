migration-up:
	migrate -path db/migration -database "mysql://titan:titan@tcp(localhost:3306)/book_db?query" up

migration-down:
	migrate -path db/migration -database "mysql://titan:titan@tcp(localhost:3306)/book_db?query" down

migration-force:
	migrate -path db/migration -database "mysql://titan:titan@tcp(localhost:3306)/book_db?query" force $(force)

create-migration:
	migrate create -ext sql -dir migration $(action)