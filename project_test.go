package ganboard

import (
	"fmt"
	"os"
	"testing"
)

var client = Client{
	Endpoint: os.Getenv("ENDPOINT"),
	Username: os.Getenv("USERNAME"),
	Password: os.Getenv("PASSWORD"),
}

func TestGetAllProjects(t *testing.T) {
	projects, err := client.GetAllProjects()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(projects)
}
