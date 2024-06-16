package database

func UpsertUser(userId string, username string) error {
	_, err := db.Exec(
		`INSERT INTO "user" (id, username) VALUES ($1, $2) ON CONFLICT (id) DO UPDATE SET id = $1, username = $2;`,
		userId, username,
	)

	if err != nil {
		return err
	}

	return nil
}

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func GetUsers() (users []User, err error) {
	rows, err := db.Query(`
	SELECT * FROM "user"
	`)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Username); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, err
}

func GetUser(userId string) (user User, err error) {
	err = db.QueryRow(`SELECT * FROM "user" WHERE id = $1`, userId).Scan(&user.Id, &user.Username)
	if err != nil {
		return user, err
	}
	return user, nil
}
