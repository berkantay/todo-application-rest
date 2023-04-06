postgresinit:
	docker run --name postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres
postgresbash:
	docker exec -it postgres psql
createdb:
	docker exec -it postgres createdb --username=root --owner=root todos
dropdb:
	docker exec -it postgres dropdb todos
migrateup:
	migrate -path migrations -database "postgresql://root:password@localhost:5432/todos?sslmode=disable" -verbose up
migratedown:
	migrate -path migrations -database "postgresql://root:password@localhost:5432/todos?sslmode=disable" -verbose down
.PHONY:
	postgresinit postgresbash createdb dropdb migrateup migratedown