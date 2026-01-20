package seed

import (
	"fmt"
	"math/rand"

	"backend/models"
)

func GenerateUsers(n int) []models.User {
	var users []models.User
	for i := 1; i <= n; i++ {
		users = append(users, models.User{
			ID:       i,
			Username: fmt.Sprintf("rahul_%d", i),
			Rating:   rand.Intn(4900) + 100,
		})
	}
	return users
}
