package models
 import "time"

   type Bookmark struct {
       ID         int       `db:"id" json:"id"`
        UserID   int        `db:"user_id" json:"userId"`
      URL        string     `db:"url" json:"url"`
      Description  string    `db:"description" json:"description"`
      Tags       []string   `db:"tags" json:"tags"`
      CreatedAt time.Time `db:"created_at" json:"createdAt"`
  }