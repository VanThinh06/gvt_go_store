postgres:
	docker run --name my-postgres -p 8080:5432 -e POSTGRES_USER=pgVanThinh -e POSTGRES_PASSWORD=VanThinh2512. -e POSRGRES_DB=pt01 -d postgres

createdb:
	docker exec -it my-postgres createdb --username=pgVanThinh --owner=pgVanThinh pt01_StoreMobile

dropdb:
	docker exec -it my-postgres dropdb pt01_StoreMobile

runsqlc: 
    docker run --rm -v ${pwd}:/src -w /src kjconroy/sqlc generate

migrateup:
	migrate -path db/migration -database "postgresql://pgVanThinh:VanThinh2512.@localhost:8080/pt01_StoreMobile?sslmode=disable" -verbose up 

migratedown:
	migrate -path db/migration -database "postgresql://pgVanThinh:VanThinh2512.@localhost:8080/pt01_StoreMobile?sslmode=disable" -verbose down 

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown runsqlc
