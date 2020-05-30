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
	PostedBy    *users.User
}

// Save ...
func (link Link) Save() int64 {
	stmt, err := db.DB.Prepare("INSERT INTO links (description, url, user_id) VALUES(?,?,?)")
	if err != nil {
		log.Fatal("Error preparing insert statement, ", err)
	}

	result, err := stmt.Exec(link.Description, link.URL, link.PostedBy.ID)
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
	rows, err := db.DB.Query(
		"SELECT l.id, l.description, l.url, u.id, u.username " +
			"FROM links l, users u " +
			"WHERE l.user_id = u.id")

	if err != nil {
		log.Fatal("Error querying database, ", err)
	}
	defer rows.Close()

	var links []Link
	for rows.Next() {
		var link Link
		var user users.User
		err := rows.Scan(&link.ID, &link.Description, &link.URL, &user.ID, &user.Username)
		if err != nil {
			log.Fatal("Error parsing query result, ", err)
		}
		link.PostedBy = &user
		links = append(links, link)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error during iteration, ", err)
	}

	return links
}
