postgrescreate :
	docker run --name simpleBankPostgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root@123 -d postgres:alpine

postgresstop: 
	docker stop simpleBankPostgres

postgresdelete:
	docker rm simpleBankPostgres

createdb :
	docker exec -it simpleBankPostgres createdb --username=root --owner=root simpleBank

dropdb:
	docker exec -it simpleBankPostgres dropdb simpleBank

migratedbup:
	migrate -path db/migration -database "postgresql://root:root@123@192.168.29.145:5432/simpleBank?sslmode=disable" -verbose up

migratedbdown: 
	migrate -path db/migration -database "postgresql://root:root@123@192.168.29.145:5432/simpleBank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

dbtest:
	go test -cover -v ./db/sqlc/...

cleantestcacahe:
	go clean -testcache

.PHONY: postgrescreate postgresdelete createdb dropdb postgresstop migratedbup migratedbdown sqlc dbtest cleantestcacahe
    
