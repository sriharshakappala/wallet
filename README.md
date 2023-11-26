

migrate create -ext sql -dir migrations -seq create_users


migrate -path ./migrations -database "mysql://admin:password@tcp(localhost:3306)/wallet" up
