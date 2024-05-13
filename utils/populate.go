package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/google/uuid"
	_ "github.com/david-daxzbr/english-app/handlers"
)

type Snippet struct {  
  Description  string    `json:"description"`
  Thumbnail struct {
     Default struct {
                    URL    string `json:"url"`
                } `json:"default"`
  } `json:"thumbnails"`
  VideoTitle string `json:"title"`
}

type ID struct {
  VideoID string `json:"videoId"`
}

type Item struct {
  ID      ID      `json:"id"`
  Snippet Snippet `json:"snippet"`
}

type Video struct {
  Items []Item `json:"items"`
}

func main() {
  
  dirDatabase := "/home/daxzbr/Documentos/programming/projects/english-app/db/videos.db"
	db, err := sql.Open("sqlite3", dirDatabase)
  if err != nil {
		panic(err)
	}
  defer db.Close()
  
_, err = db.Exec(`CREATE TABLE IF NOT EXISTS videos (
                    id TEXT PRIMARY KEY UNIQUE,
                    video_id TEXT NOT NULL UNIQUE,
                    video_title TEXT NOT NULL UNIQUE,
                    video_description TEXT NOT NULL UNIQUE,
                    video_thumbnail_url TEXT NOT NULL UNIQUE
                )`) 
  
  if err != nil {
    panic(err)
  }

  resp, err := http.Get("https://www.googleapis.com/youtube/v3/search?part=snippet&q=ted&maxResults=10&key=AIzaSyA0fzOrSM7nIeKveJrjolfJ7Pz0HVFV9g4") 
  if err != nil {
    panic(err)
  }

  defer resp.Body.Close()
 
  if resp.StatusCode != 200 {
    print("there was an error while makin request %d", resp.StatusCode)
  }
  
  body, err := io.ReadAll(resp.Body) 
  if err != nil {
    panic(err)
  }
  
  var video Video

  err = json.Unmarshal(body, &video)

  if err != nil {
    panic(err)
  }
  
  filteredItems := make([]Item, 0)
    for _, item := range video.Items {
        if item.ID.VideoID != "" {
            filteredItems = append(filteredItems, item)
        }
    }

  for _, item := range filteredItems {
    //fmt.Printf("id: %+v, title: %+v, description: %+v, video_thumbnail_url: %+v \n", item.ID.VideoID, item.Snippet.VideoTitle, item.Snippet.Description, item.Snippet.Thumbnail.Default.URL)

    uuid := uuid.New()
  _, err = db.Exec(`INSERT INTO videos (id, video_id, video_title, video_description, video_thumbnail_url) VALUES (?, ?, ?, ?, ?)`, uuid,  item.ID.VideoID, item.Snippet.VideoTitle, item.Snippet.Description, item.Snippet.Thumbnail.Default.URL )
  }

  rows, err := db.Query("SELECT * FROM videos")
  if err != nil {
      fmt.Println("Error querying database:", err)
      return
  }

  defer rows.Close()
  
  for rows.Next() {
    var (
      id string  
      video_title string
        video_description string
	video_id string
        video_thumbnail string  
    )
    if err := rows.Scan(&id, &video_id, &video_title, &video_description, &video_thumbnail); err != nil {
        fmt.Println("Error scanning row:", err)
        return
    }
    // Print or use the data as needed
    fmt.Println(id, video_id, video_title, video_description, video_thumbnail)
  }

  if err := rows.Err(); err != nil {
      fmt.Println("Error iterating over rows:", err)
      return
  }
} 
