package config
import (
  "database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func StartDataBase() (*sql.DB){
  
  dirDatabase := "/home/daxzbr/Documentos/programming/projects/english-app/db/videos.db"
	db, err := sql.Open("sqlite3", dirDatabase)
  if err != nil {
		panic(err)
	}
  return db
}

