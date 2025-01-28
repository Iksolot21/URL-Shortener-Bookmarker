package models

import "time"

type URL struct {
    ID          int       `db:"id"`
    ShortURL    string    `db:"short_url"`
    OriginalURL string    `db:"original_url"`
    CreatedAt   time.Time `db:"created_at"`
}