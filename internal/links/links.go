package links

import (
	"log"

	db "github.com/khwj/hackernews/internal/pkg/db/mysql"
	"github.com/khwj/hackernews/internal/users"
)

// Link a struct that represent a link
type Link struct {
	ID          string
	Description string
	URL         string
	User        *users.User
}

// Save ...
func (link Link) Save() int64 {
	stmt, err := db.DB.Prepare("INSERT INTO links (title, address) VALUES(?,?)")
	if err != nil {
		log.Fatal("Error preparing insert statement, ", err)
	}

	result, err := stmt.Exec(link.Description, link.URL)
	if err != nil {
		log.Fatal("Error executing prepared statement, ", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal("Error, ", err)
	}

	return id
}

// GetAll ...
func GetAll() []Link {
	rows, err := db.DB.Query("SELECT id, title, address FROM links")
	if err != nil {
		log.Fatal("Error querying database, ", err)
	}
	defer rows.Close()

	var links []Link
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Description, &link.URL)
		if err != nil {
			log.Fatal("Error parsing query result, ", err)
		}
		links = append(links, link)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error during iteration, ", err)
	}

	return links
}
