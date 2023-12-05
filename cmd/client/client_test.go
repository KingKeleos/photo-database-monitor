package client

import (
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient("localhost:8081")
	if err != nil {
		t.Errorf("Error occured: %v", err)
	}

	err = client.UpdatePeople(1)
	if err != nil {
		t.Errorf("Error occured: %v", err)
	}

	time.Sleep(1 * time.Second)

	err = client.UpdateProjects(5)
	if err != nil {
		t.Errorf("Error occured: %v", err)
	}

	time.Sleep(1 * time.Second)

	err = client.UpdateParticipantsToProject("test-project", 5)
	if err != nil {
		t.Errorf("Error occured: %v", err)
	}

	time.Sleep(1 * time.Second)

	err = client.UpdateParticipantsToProject("test-project2", 1)
	if err != nil {
		t.Errorf("Error occured: %v", err)
	}

	time.Sleep(1 * time.Second)

	err = client.UpdatePostsToProject("test-project2", 3)
	if err != nil {
		t.Errorf("Error occured: %v", err)
	}

	time.Sleep(1 * time.Second)

	err = client.NewProjects(3)
	if err != nil {
		t.Errorf("Error occured: %v", err)
	}

	time.Sleep(1 * time.Second)

	err = client.ActiveProjects(1)
	if err != nil {
		t.Errorf("Error occured: %v", err)
	}

	time.Sleep(1 * time.Second)

	err = client.FinishedProjects(1)
	if err != nil {
		t.Errorf("Error occured: %v", err)
	}
}
