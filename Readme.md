## Install
`go get -u github.com/sb-golang-67-bioskop`

### Running Migration
`migrate -database "postgres://postgres:postgres@localhost:5432/db-bioskop-golang?sslmode=disable" -path database/migration up`

### Available Path
- "/bioskop" **POST**