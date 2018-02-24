package postgres

import (
	"errors"
	"log"

	"github.com/sj14/web-demo/domain"
)

func (s *PostgresStore) StorePost(post domain.Post) (postID int64, err error) {
	tx, err := s.conn.Beginx()
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	// Insert User
	rows, err := tx.NamedQuery(
		"INSERT INTO posts ("+
			"user_id,"+
			"text,"+
			"created_at,"+
			"updated_at)"+

			"VALUES ("+
			":user_id,"+
			":text,"+
			":created_at,"+
			":updated_at)"+
			"RETURNING id",
		post)
	if err != nil {
		log.Println("Failed to create post in DB: ", err)
		tx.Rollback()
		return -1, err
	}

	// Get User ID
	var id int64
	if rows.Next() {
		rows.Scan(&id)
	}
	if id == 0 {
		return -1, errors.New("Failed to retrieve post id")
		tx.Rollback()
	}

	// Everything went fine, commit changes to database
	tx.Commit()
	return id, nil
}

func (s *PostgresStore) FindPostByID(postID int64) (domain.Post, error) {
	post := domain.Post{}
	err := s.conn.Get(&post, "SELECT * FROM posts WHERE id = $1", postID)
	if err != nil {
		log.Println(err)
		return domain.Post{}, err
	}

	return post, nil
}
