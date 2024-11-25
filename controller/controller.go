package controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Player represents a player ranking record.
type Player struct {
	Rank      int    `json:"rank"`
	Username  string `json:"username"`
	Points    int    `json:"points"`
	HeroClass string `json:"hero_class"`
}

// GetPlayersController fetches player rankings with pagination support.
func GetPlayersController(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse query parameters for pagination
		page := c.DefaultQuery("page", "1")    // Default to page 1 if not provided
		limit := c.DefaultQuery("limit", "50") // Default to 50 records per page

		// Convert page and limit to integers
		pageNum, err := strconv.Atoi(page)
		if err != nil || pageNum < 1 {
			pageNum = 1
		}
		limitNum, err := strconv.Atoi(limit)
		if err != nil || limitNum < 1 {
			limitNum = 50
		}

		offset := (pageNum - 1) * limitNum // Calculate the offset for SQL

		// Get hero_class filter
		heroClass := c.DefaultQuery("hero_class", "allclass")

		// Prepare the SQL query based on hero_class and pagination
		var rows *sql.Rows
		if heroClass == "allclass" {
			rows, err = db.Query(
				"SELECT rank, username, points, hero_class FROM player_rankings ORDER BY points DESC LIMIT $1 OFFSET $2",
				limitNum, offset,
			)
		} else {
			rows, err = db.Query(
				"SELECT rank, username, points, hero_class FROM player_rankings WHERE hero_class = $1 ORDER BY points DESC LIMIT $2 OFFSET $3",
				heroClass, limitNum, offset,
			)
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching player rankings"})
			return
		}
		defer rows.Close()

		// Parse rows into a slice of Player structs
		var players []Player
		for rows.Next() {
			var player Player
			err := rows.Scan(&player.Rank, &player.Username, &player.Points, &player.HeroClass)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning player rankings"})
				return
			}
			players = append(players, player)
		}

		// Respond with the JSON data
		c.JSON(http.StatusOK, players)
	}
}
