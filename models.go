package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/logeshwhatnot/rssagg/internal/database"
)

type User struct {
	ID         uuid.UUID `json:"id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Name       string    `json:"name"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:         dbUser.ID.UUID,
		Created_at: dbUser.CreatedAt,
		Updated_at: dbUser.UpdatedAt,
		Name:       dbUser.Name,
	}
}
