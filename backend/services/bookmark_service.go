package services

import (
	  "database/sql"
   "fmt"
	  "time"
	  "github.com/Iksolot21/URL-Shortener-Bookmarker/backend/models"
	  "strings"
	)

 func CreateBookmark(db *sql.DB, bookmark *models.Bookmark, userID int) error {
	 tags := strings.Join(bookmark.Tags, ",")
	   query := `INSERT INTO bookmarks (user_id, url, description, tags, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(query, userID, bookmark.URL, bookmark.Description, tags, time.Now())
	 if err != nil {
		 return fmt.Errorf("failed to create bookmark: %w", err)
	  }
	  return nil
  }

  func GetBookmarks(db *sql.DB, userID int) ([]models.Bookmark, error) {
	  var bookmarks []models.Bookmark
	 query := `SELECT id, url, description, tags FROM bookmarks WHERE user_id = $1`
	 rows, err := db.Query(query, userID)
	 if err != nil {
		return nil, fmt.Errorf("failed to get bookmarks: %w", err)
   }
	 defer rows.Close()
	  for rows.Next() {
		 var bookmark models.Bookmark
		var tags string
		 err := rows.Scan(&bookmark.ID, &bookmark.URL, &bookmark.Description, &tags)
		if err != nil {
			  return nil, fmt.Errorf("failed to scan bookmark: %w", err)
		  }
	   bookmark.Tags = strings.Split(tags,",")
		 bookmarks = append(bookmarks, bookmark)
	 }
	return bookmarks, nil
 }

   func GetBookmarkById(db *sql.DB, id int, userID int) (models.Bookmark, error) {
	   var bookmark models.Bookmark
	   var tags string
	 query := `SELECT id, url, description, tags FROM bookmarks WHERE id = $1 AND user_id = $2`
	   row := db.QueryRow(query, id, userID)
		err := row.Scan(&bookmark.ID, &bookmark.URL, &bookmark.Description, &tags)
	if err != nil {
		return bookmark, fmt.Errorf("failed to get bookmark: %w", err)
   }
		bookmark.Tags = strings.Split(tags,",")
	 return bookmark, nil

}

 func PatchBookmarkById(db *sql.DB, id int, userID int, bookmark *models.Bookmark) error {
	 tags := strings.Join(bookmark.Tags, ",")
	query := `UPDATE bookmarks SET description = $1, tags = $2 WHERE id = $3 AND user_id = $4`
	_, err := db.Exec(query, bookmark.Description, tags, id, userID)
	if err != nil {
	   return  fmt.Errorf("failed to update bookmark: %w", err)
	}
   return nil
 }