package postgres

import (
	"errors"
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
	rows, err := tx.NamedQuery(
		"INSERT INTO users ("+
			"name,"+
			"email,"+
			"password,"+
			"is_disabled,"+
			"email_token,"+
			"email_verified,"+
			"zip_code,"+
			"created_at,"+
			"updated_at,"+
			"failed_logins)"+

			"VALUES ("+
			":name,"+
			":email,"+
			":password,"+
			":is_disabled,"+
			":email_token,"+
			":email_verified,"+
			":zip_code,"+
			":created_at,"+
			":updated_at,"+
			":failed_logins)"+
			"RETURNING id",
		user)
	if err != nil {
		log.Println("Failed to create user in DB: ", err)
		tx.Rollback()
		return -1, err
	}

	// Get User ID
	var id int64
	if rows.Next() {
		rows.Scan(&id)
	}
	if id == 0 {
		return -1, errors.New("Failed to retrieve user id")
		tx.Rollback()
	}

	// Everything went fine, commit changes to database
	tx.Commit()
	return id, nil
}

func (s *PostgresStore) UpdateUserExceptPassword(user domain.User) error {
	_, err := s.conn.Exec(
		"UPDATE users SET "+
			"name = $1,"+
			"email = $2,"+
			"is_disabled = $3,"+
			"email_token = $4,"+
			"email_verified = $5,"+
			"zip_code = $6,"+
			"created_at = $7,"+
			"failed_logins = $8"+ // no commata!
			"WHERE id = $9",
		user.Name,
		user.Email,
		user.IsDisabled,
		user.EmailToken,
		user.EmailVerified,
		user.ZipCode,
		user.CreatedAt,
		user.FailedLogins,
		user.ID)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (s *PostgresStore) UpdateUserPasswordOnly(userID int64, password string) error {
	_, err := s.conn.Exec(
		"UPDATE users SET "+
			"password = $1 "+
			"WHERE id = $2",
		password,
		userID)
	return err
}

func (s *PostgresStore) FindUserById(id int64) (domain.User, error) {
	user := domain.User{}
	err := s.conn.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		log.Println("user id:", id, "error: ", err)
		return domain.User{}, err
	}
	return user, nil
}

func (s *PostgresStore) FindUserIdByEmail(email string) (int64, error) {
	var idFromDB int64

	err := s.conn.QueryRow("SELECT id FROM users WHERE email = $1", email).Scan(&idFromDB)
	if err != nil {
		log.Println("Error in query FindUserByEmail: %v", err)
		return -1, err
	}
	return idFromDB, nil
}

//func (s *MysqlStore) DeleteUserById(id int64) error {
//	_, err := s.conn.Exec("DELETE FROM users WHERE id = $1", id)
//	if err != nil {
//		return err
//	}
//	return nil
//}
