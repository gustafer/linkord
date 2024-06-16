package database

type Game struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func CreateGame(game *Game) (createdGameId int64, err error) {
	var id int64
	if err := db.QueryRow(
		`INSERT INTO games (title, description) VALUES ($1, $2) RETURNING id`,
		game.Title, game.Description,
	).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
