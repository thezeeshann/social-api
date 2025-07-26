### FOR DB

#### Create migrations
```
migrate create -ext sql -dir ./cmd/migrate/migrattion -seq create_users_table
```
#### Run migrations
```
migrate -database YOUR_DATABASE_URL -path PATH_TO_YOUR_MIGRATIONS up
```


### go packages
- go-playground validator