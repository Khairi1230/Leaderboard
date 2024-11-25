package model

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Player struct {
	Rank     int
	Username string
	Points   int
	Class    string
}

func GetPlayerRanking(db *sql.DB) ([]Player, error) {
	query := "SELECT rank, username, points, class FROM players ORDER BY points DESC"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Error fetching player ranking:", err)
		return nil, err
	}
	defer rows.Close()

	var players []Player
	for rows.Next() {
		var player Player
		err := rows.Scan(&player.Rank, &player.Username, &player.Points, &player.Class)
		if err != nil {
			log.Fatal("Error scanning row:", err)
			return nil, err
		}
		players = append(players, player)
	}

	return players, nil
}
