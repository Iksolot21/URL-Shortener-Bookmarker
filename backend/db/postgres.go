package db

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/lib/pq"

)

func OpenDB(dbUrl string) (*sql.DB, error) {
    db, err := sql.Open("postgres", dbUrl)
    if err != nil {
        return nil, fmt.Errorf("failed to open database: %w", err)
    }
     err = db.Ping()
   if err != nil {
     return nil, fmt.Errorf("unable to reach database: %w", err)
  }
log.Println("Database opened")
    return db, nil
}