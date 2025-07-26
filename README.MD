### FOR DB

#### Create migrations
```
migrate create -ext sql -dir ./cmd/migrate/migrattion -seq create_users_table
```
#### Run migrations
```
migrate -database YOUR_DATABASE_URL -path PATH_TO_YOUR_MIGRATIONS up
migrate -path ./cmd/migrate/migrations -database "postgres://postgres:root@localhost:5432/social?sslmode=disable" up

drop off all the migrations
migrate -path ./cmd/migrate/migrations -database "postgres://postgres:root@localhost:5432/social?sslmode=disable" drop -f
create new migrations
migrate -path ./cmd/migrate/migrations -database "postgres://postgres:root@localhost:5432/social?sslmode=disable" up

https://chatgpt.com/c/6873e65d-0bfc-8012-9eb7-e66f34655742

```


### go packages
- go-playground validator