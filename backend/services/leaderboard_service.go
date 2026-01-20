package services

import (
	"math/rand"
	"sort"
	"strings"
	"sync"

	"backend/models"
)

var (
	users []models.User
	mu    sync.RWMutex
)

// Seed initial users
func SeedUsers(seed []models.User) {
	mu.Lock()
	defer mu.Unlock()
	users = seed
	updateRanks()
}

// Get top users for leaderboard
func GetUsers(limit int) []models.User {
	mu.RLock()
	defer mu.RUnlock()

	if limit > len(users) {
		limit = len(users)
	}
	return users[:limit]
}

// Search users by username (substring, case-insensitive)
func SearchUsers(q string) []models.User {
	mu.RLock()
	defer mu.RUnlock()

	q = strings.ToLower(q)
	var res []models.User

	for _, u := range users {
		if len(res) >= 20 {
			break
		}
		if q != "" && strings.Contains(strings.ToLower(u.Username), q) {
			res = append(res, u)
		}
	}
	return res
}

// Simulate random score updates
func RandomScoreUpdate() {
	mu.Lock()
	defer mu.Unlock()

	for i := 0; i < 50; i++ {
		idx := rand.Intn(len(users))
		users[idx].Rating += rand.Intn(20) - 10
		if users[idx].Rating < 100 {
			users[idx].Rating = 100
		}
	}
	updateRanks()
}

// Update dense rankings (tie-aware)
func updateRanks() {
	sort.Slice(users, func(i, j int) bool {
		return users[i].Rating > users[j].Rating
	})

	rank := 1
	for i := range users {
		if i > 0 && users[i].Rating < users[i-1].Rating {
			rank = i + 1
		}
		users[i].Rank = rank
	}
}
