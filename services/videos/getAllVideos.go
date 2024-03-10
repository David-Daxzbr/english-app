package videos

import (
	"log"
	"github.com/david-daxzbr/english-app/config"
	"github.com/david-daxzbr/english-app/dto"
)

func GetAllVideos() []*dto.VideoDto {

  db := config.StartDataBase()
  db.Query("SELECT * FROM videos")
  
  defer db.Close()
  rows, err := db.Query("SELECT * FROM videos")
  if err != nil {
    panic(err)
  }

  defer rows.Close()
  
  videos := []*dto.VideoDto{}

  for rows.Next() {
    
    var (
      id string  
      video_title string
      video_description string
	    video_id string
      video_thumbnail string  
    )

    if err := rows.Scan(&id, &video_id, &video_title, &video_description, &video_thumbnail); err != nil {
        log.Fatal(err)
    }
   
    videos = append(videos, &dto.VideoDto{
      Id: id,
      VideoId: video_id,
      VideoTitle: video_title,
      VideoDescription: video_description,
      VideoThumbnail: video_thumbnail,
    }) 
  }

  return videos

}

