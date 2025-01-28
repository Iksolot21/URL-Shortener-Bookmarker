package services
  import (
  "database/sql"
  )

  func DeleteURL(db *sql.DB, id int) error{
      query := `DELETE FROM urls WHERE id = $1`
      _, err := db.Exec(query, id)
     return err
  }

  func GetURLById(db *sql.DB, id int) (string, error) {
  	var originalURL string
    query := `SELECT original_url FROM urls WHERE id = $1`
      err := db.QueryRow(query, id).Scan(&originalURL)
      if err != nil {
       return "", err
    }
  	return originalURL, nil
  }