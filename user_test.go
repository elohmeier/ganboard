package ganboard

import (
	"fmt"
	"testing"
)

func TestGetAllUsers(t *testing.T) {
	users, err := client.GetAllUsers()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(users)
}
