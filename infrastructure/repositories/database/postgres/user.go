package postgres

import (
	"log"
	"github.com/sj14/web-demo/domain"
)

func (s *PostgresStore) StoreUser(user domain.User) (userID int64, err error) {

	tx, err := s.conn.Beginx()
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	// Insert User
	res, err := tx.NamedQuery(
		"INSERT INTO users ("+
			"first_name,"+
			"last_name,"+
			"email,"+
			"password,"+
			"base_type,"+
			"is_disabled,"+
			"email_token,"+
			"email_verified,"+
			"zip_code,"+
			"created_at,"+
			"failed_logins,"+
			"last_login)"+

			"VALUES ("+
			":first_name,"+
			":last_name,"+
			":email,"+
			":password,"+
			":base_type,"+
			":is_disabled,"+
			":email_token,"+
			":email_verified,"+
			":zip_code,"+
			":created_at,"+
			":failed_logins,"+
			":last_login)"+
			"RETURNING id",
		user)
	if err != nil {
		log.Println("Failed to create user in DB: ", err)
		tx.Rollback()
		return -1, err
	}

	// Get User ID
	var id int64
	err = res.Scan(id)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	user.ID = id

	// Insert all User Roles
	//for _, val := range user.AccessRoles {
	//	_, err := tx.Exec("INSERT INTO user_role(user_id, role_id) VALUES ($1, $2)", id, val)
	//	if err != nil {
	//		tx.Rollback()
	//		return domain.User{}, err
	//	}
	//
	//}

	// Everything went fine, commit changes to database
	tx.Commit()
	return id, nil
}
