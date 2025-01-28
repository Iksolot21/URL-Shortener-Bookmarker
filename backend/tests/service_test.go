package tests
import (
	"testing"
	"github.com/Iksolot21/URL-Shortener-Bookmarker/backend/services"
	"github.com/Iksolot21/URL-Shortener-Bookmarker/backend/db"
	 "github.com/Iksolot21/URL-Shortener-Bookmarker/backend/config"
	 _ "github.com/lib/pq"
	"github.com/Iksolot21/URL-Shortener-Bookmarker/backend/models"
)


func TestCreateBookmark(t *testing.T) {
  cfg, err := config.LoadConfig()
   if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}
   database, err := db.OpenDB(cfg.DatabaseURL)
	if err != nil {
		t.Fatalf("Failed to connect to db: %v", err)
   }
	defer database.Close()

	bookmark := models.Bookmark{
	   URL: "short_url",
	   Description: "description",
	   Tags: []string{"tag1", "tag2"},
	}

   err = services.CreateBookmark(database, &bookmark, 1)
	if err != nil {
	   t.Fatalf("Failed to create bookmark: %v", err)
   }

   bookmarks, err := services.GetBookmarks(database, 1)
   if err != nil {
	   t.Fatalf("Failed to get bookmarks: %v", err)
	}

   if len(bookmarks) == 0 {
	  t.Error("Bookmark was not saved")
   }
}