package annotations

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"github.com/mitchellh/go-homedir"
)

var home, _ = homedir.Dir()
var annotationsDirectory = filepath.Join(home, ".annotorious")
var annotationsDB = filepath.Join(annotationsDirectory, "annotorious.db")

type Api struct {
	DB *sqlx.DB
}

func (api *Api) Init() error {
	if _, err := os.Stat(annotationsDirectory); os.IsNotExist(err) {
		os.Mkdir(annotationsDirectory, 0755)
	}
	var err error
	api.DB, err = sqlx.Open("sqlite3", annotationsDB)
	if err != nil {
		return err
	}
	statement, err := api.DB.Prepare(`CREATE TABLE IF NOT EXISTS annotations (
		id VARCHAR PRIMARY KEY, 
		canvas VARCHAR, 
		manifest VARCHAR,
		annotation TEXT, 
		created_at DATETIME, 
		updated_at DATETIME)`)
	statement.Exec()
	if err != nil {
		return err
	}

	api.DB.Exec("CREATE UNIQUE INDEX IF NOT EXISTS annotation_id ON annotations (id);")
	api.DB.Exec("CREATE INDEX IF NOT EXISTS canvas ON annotations (canvas);")
	return nil
}

func (api *Api) Save(id string, manifest string, canvas string, annotation string) error {
	statement, _ := api.DB.Prepare(`INSERT INTO annotations (id, created_at, manifest, canvas, annotation) 
						VALUES (?, strftime('%Y-%m-%d %H-%M-%S','now'), ?, ?, ?)`)
	statement.Exec(id, manifest, canvas, annotation)
	return nil
}

func (api *Api) Update(id string, annotation string) error {
	statement, _ := api.DB.Prepare(`UPDATE annotations SET annotation=?, 
									updated_at=strftime('%Y-%m-%d %H-%M-%S','now')
									WHERE id=?`)
	statement.Exec(annotation, id)
	return nil
}

func (api *Api) Delete(id string) error {
	statement, _ := api.DB.Prepare(`DELETE FROM annotations WHERE id=?`)
	statement.Exec(id)
	return nil
}

func (api *Api) Get(canvas string) string {
	rows, err := api.DB.Query("SELECT annotation FROM annotations WHERE canvas=?", canvas)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var annotations []json.RawMessage
	for rows.Next() {
		var annotation []byte
		err = rows.Scan(&annotation)
		if err != nil {
			log.Println(err)
		}
		annotations = append(annotations, annotation)
	}
	j, err := json.Marshal(annotations)
	if err != nil {
		log.Println(err)
	}
	return string(j)
}
