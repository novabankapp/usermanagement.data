module github.com/novabankapp/usermanagement.data

go 1.18

require (
	github.com/gocql/gocql v1.1.0
	github.com/google/uuid v1.3.0
	github.com/scylladb/gocqlx/v2 v2.7.0
	gorm.io/gorm v1.23.3
)

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/fatih/structs v1.1.0
	github.com/golang/mock v1.4.4
	github.com/novabankapp/common.infrastructure v1.3.0
	github.com/shopspring/decimal v1.2.0
	github.com/stretchr/testify v1.7.1
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
	gorm.io/driver/postgres v1.3.1
)

replace github.com/novabankapp/common.infrastructure v1.3.0 => C:\Projects\golang\github.com\novabankapp\common.infrastructure

replace github.com/novabankapp/common.data v1.0.2 => C:\Projects\golang\github.com\novabankapp\common.data

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/snappy v0.0.3 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.10.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.2.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.9.1 // indirect
	github.com/jackc/pgx/v4 v4.14.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/novabankapp/common.data v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/scylladb/go-reflectx v1.0.1 // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)
