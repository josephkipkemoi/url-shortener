package models

import (
	"log"
	"url-shortener/database"
)

type Url struct {
	Id            int    `json:"id"`
	Original_Url  string `json:"original_url"`
	Shortened_Url string `json:"shortened_url"`
}

func (u *Url) InsertUrl() {
	const insertUrlSql = `INSERT INTO urlshortener(original_url, shortened_url) VALUES (?, ?)`
	shortStr := database.GenerateShortKey()

	statement, err := database.DB.Prepare(insertUrlSql)
	checkErr(err)

	_, err = statement.Exec(u.Original_Url, shortStr)
	checkErr(err)

	u.Shortened_Url = shortStr

	log.Println("record inserted successuly...")
}

func GetUrl(url string) string {
	const getUrlSql = `SELECT original_url FROM urlshortener WHERE shortened_url = ?`
	row, err := database.DB.Query(getUrlSql, url)
	checkErr(err)

	var urlString string
	for row.Next() {
		row.Scan(&urlString)
	}

	return urlString
}
