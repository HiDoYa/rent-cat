# Rent Cat

## Developer Guide

Create dev database
```
make postgres-dev
```

Create a new migration
```
migrate create -ext sql -dir migrations -seq -digits 3 create-table
```

Apply migrations
```
export POSTGRES_URL="postgres://rent-cat-user:rent-cat-pass@localhost:9945/rent-cat-db?sslmode=disable"
migrate -database ${POSTGRES_URL} -path migrations up
migrate -database ${POSTGRES_URL} -path migrations down
```

Dev database
```
psql -h localhost -p 9945 -d rent-cat-db -U rent-cat-user
```