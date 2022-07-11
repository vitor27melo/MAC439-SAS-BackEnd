package configs

import (
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DBHost     = "sas-mac439.postgres.database.azure.com"
	DBPort     = 5432
	DBUser     = "mac439"
	DBPassword = "Atila_Iamarino"
	DBName     = "postgres"
	DBType     = "postgres"
)

func GetDBType() string {
	return DBType
}

func GetPostgresConnString() string {
	conn := fmt.Sprintf("host= %s port= %d user= %s password= %s dbname= %s",
		DBHost,
		DBPort,
		DBUser,
		DBPassword,
		DBName)

	return conn
}
