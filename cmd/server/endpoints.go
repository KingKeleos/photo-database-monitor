package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/KingKeleos/photo-database-monitor/graphite"
)

type Handler struct {
}

func createMap(name string, content float64) map[string]float64 {
	return map[string]float64{
		name: content,
	}
}

func (h Handler) UpdateProjects() {
	response, err := http.Get("http://localhost:90/projects")
	if err != nil {
		log.Fatal("Could not reach Database-Exporter for Metric, err: " + err.Error())
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var projects []Project
	json.Unmarshal(body, &projects)

	count := len(projects)

	metric := createMap("overview.project_count", float64(count))

	err = graphite.Client.SendData(metric)
	if err != nil {
		fmt.Printf("Error occured sending data to Graphite: %v", err)
	}
	time.Sleep(500 * time.Millisecond)
}

func (h Handler) UpdatePeople() {
	response, err := http.Get("http://localhost:90/people")
	if err != nil {
		log.Fatal("Could not reach Database-Exporter for Metric, err: " + err.Error())
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var people PeopleList
	json.Unmarshal(body, &people)

	count := len(people.People)
	metric := createMap("overview.people_count", float64(count))

	err = graphite.Client.SendData(metric)
	if err != nil {
		fmt.Printf("Error occured sending data to Graphite: %v", err)
	}
	time.Sleep(500 * time.Millisecond)
}

func (h Handler) UpdateParticipantsToProjects() {
	response, err := http.Get("http://localhost:90/projects")
	if err != nil {
		log.Fatal("Could not reach Database-Exporter for Metric, err: " + err.Error())
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var projects []Project
	json.Unmarshal(body, &projects)

	for id, project := range projects {
		id += 1
		url := fmt.Sprintf("http://localhost:90/projects/%d/participants", id)
		response, err := http.Get(url)
		if err != nil {
			log.Fatal("Could not reach Database-Exporter for Metric, err: " + err.Error())
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var socials []Socials
		json.Unmarshal(body, &socials)

		count := len(socials)

		label := fmt.Sprintf("projects.%d_%s.participants", id, project.Name)
		metric := createMap(label, float64(count))

		err = graphite.Client.SendData(metric)
		if err != nil {
			fmt.Printf("Error occured sending data to Graphite: %v", err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func (h Handler) UpdatePostsToProject() {
	response, err := http.Get("http://localhost:90/projects")
	if err != nil {
		log.Fatal("Could not reach Database-Exporter for Metric, err: " + err.Error())
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var projects []Project
	json.Unmarshal(body, &projects)

	for id, project := range projects {
		id += 1
		url := fmt.Sprintf("http://localhost:90/projects/%d/posts", id)
		response, err := http.Get(url)
		if err != nil {
			log.Fatal("Could not reach Database-Exporter for Metric, err: " + err.Error())
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var posts []Posts
		json.Unmarshal(body, &posts)

		count := len(posts)

		label := fmt.Sprintf("projects.%d_%s.posts", id, project.Name)
		metric := createMap(label, float64(count))

		err = graphite.Client.SendData(metric)
		if err != nil {
			fmt.Printf("Error occured sending data to Graphite: %v", err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func (h Handler) UpdateSocials() {
	response, err := http.Get("http://localhost:90/people")
	if err != nil {
		log.Fatal("Could not reach Database-Exporter for Metric, err: " + err.Error())
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var people PeopleList
	json.Unmarshal(body, &people)

	for id, project := range people.People {
		id += 1
		url := fmt.Sprintf("http://localhost:90/people/%d/socials", id)
		response, err := http.Get(url)
		if err != nil {
			log.Fatal("Could not reach Database-Exporter for Metric, err: " + err.Error())
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var socials []Socials
		json.Unmarshal(body, &socials)

		count := len(socials)

		label := fmt.Sprintf("people.%d_%s.socials", id, project.Name)
		metric := createMap(label, float64(count))

		err = graphite.Client.SendData(metric)
		if err != nil {
			fmt.Printf("Error occured sending data to Graphite: %v", err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func (h Handler) NewProjects() {
	response, err := http.Get("http://localhost:90/projects")
	if err != nil {
		log.Fatal("Could not reach Database-Exporter for Metric, err: " + err.Error())
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var projects []Project
	json.Unmarshal(body, &projects)

	var newProjects []Project
	for _, project := range projects {
		if project.State == "Created" {
			newProjects = append(newProjects, project)
		}
	}

	count := len(newProjects)

	metric := createMap("projects.states.new_projects", float64(count))

	err = graphite.Client.SendData(metric)
	if err != nil {
		fmt.Printf("Error occured sending data to Graphite: %v", err)
	}
	time.Sleep(500 * time.Millisecond)
}

func (h Handler) ActiveProjects() {
	response, err := http.Get("http://localhost:90/projects")
	if err != nil {
		log.Fatal("Could not reach Database-Exporter for Metric, err: " + err.Error())
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var projects []Project
	json.Unmarshal(body, &projects)

	var activeProjects []Project
	for _, project := range projects {
		if project.State == "Active" {
			activeProjects = append(activeProjects, project)
		}
	}

	count := len(activeProjects)

	metric := createMap("projects.states.active_projects", float64(count))

	err = graphite.Client.SendData(metric)
	if err != nil {
		fmt.Printf("Error occured sending data to Graphite: %v", err)
	}
	time.Sleep(500 * time.Millisecond)
}

func (h Handler) FinishedProjects() {
	response, err := http.Get("http://localhost:90/projects")
	if err != nil {
		log.Fatal("Could not reach Database-Exporter for Metric, err: " + err.Error())
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var projects []Project
	json.Unmarshal(body, &projects)

	var finishedProjects []Project
	for _, project := range projects {
		if project.State == "Finished" {
			finishedProjects = append(finishedProjects, project)
		}
	}

	count := len(finishedProjects)

	metric := createMap("projects.states.finished_projects", float64(count))

	err = graphite.Client.SendData(metric)
	if err != nil {
		fmt.Printf("Error occured sending data to Graphite: %v", err)
	}
	time.Sleep(500 * time.Millisecond)
}
