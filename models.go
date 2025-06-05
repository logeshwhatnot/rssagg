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
	ApiKey     string    `json:"api_key"`
}

type Feed struct {
	ID         uuid.UUID `json:"id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Name       string    `json:"name"`
	Url        string    `json:"Url"`
	UserID     uuid.UUID `json:"user_id"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:         dbUser.ID,
		Created_at: dbUser.CreatedAt,
		Updated_at: dbUser.UpdatedAt,
		Name:       dbUser.Name,
		ApiKey:     dbUser.ApiKey,
	}
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:         dbFeed.ID,
		Created_at: dbFeed.CreatedAt,
		Updated_at: dbFeed.UpdatedAt,
		Name:       dbFeed.Name,
		Url:        dbFeed.Url,
		UserID:     dbFeed.UserID,
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}
	return feeds
}
