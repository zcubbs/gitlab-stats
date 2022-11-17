package main

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/xanzy/go-gitlab"
	"log"
	"os"
)

type Application struct {
	Client *gitlab.Client
}

func main() {
	Bootstrap()

	app := &Application{
		Client: GetClient(Config.Gitlab.Url, Config.Gitlab.PrivateToken),
	}

	var entries []Entry
	for _, projectId := range Config.Gitlab.ProjectIds {
		project, _, err := app.Client.Projects.GetProject(projectId, nil)
		if err != nil {
			log.Fatalf("Failed to fetch project: %v", err)
		}

		users := app.GetProjectMembers(projectId)
		for _, user := range users {
			entries = append(entries, Entry{
				ProjectID:          project.ID,
				ProjectName:        project.Name,
				UserID:             user.ID,
				UserName:           user.Name,
				UserState:          user.State,
				UserAccessLevel:    user.AccessLevel,
				UserCreatedAt:      user.CreationDate,
				UserLastActivityAt: user.LastActivityOn,
			})
		}
	}

	app.GenerateCSV(entries)
}

func (app *Application) GetProjectMembers(projectId int) []GitlabUser {
	users, _, err := app.Client.ProjectMembers.ListAllProjectMembers(
		projectId,
		&gitlab.ListProjectMembersOptions{
			ListOptions: gitlab.ListOptions{},
		},
	)
	if err != nil {
		log.Fatalf("Failed to fetch users: %v", err)
	}

	var gitlabUsers []GitlabUser
	for _, user := range users {
		gu, _, err := app.Client.Users.GetUser(user.ID, gitlab.GetUsersOptions{}, nil)
		if err != nil {
			log.Fatalf("Failed to fetch user: %v", err)
		}

		gitlabUsers = append(gitlabUsers, GitlabUser{
			ID:             user.ID,
			Name:           user.Name,
			Username:       user.Username,
			State:          user.State,
			AvatarURL:      user.AvatarURL,
			WebURL:         user.WebURL,
			AccessLevel:    GetAccessLevel(int(user.AccessLevel)),
			CreationDate:   user.CreatedAt.String(),
			LastActivityOn: gu.LastActivityOn.String(),
		})
	}

	return gitlabUsers
}

func (app *Application) GenerateCSV(entries []Entry) {
	file, err := os.OpenFile("report.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	_, err = gocsv.MarshalString(&entries)
	err = gocsv.MarshalFile(&entries, file)
	if err != nil {
		panic(err)
	}
}

func GetClient(url string, token string) *gitlab.Client {
	c, err := gitlab.NewClient(
		token,
		gitlab.WithBaseURL(fmt.Sprintf("%s/api/v4", url)),
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return c
}

func GetAccessLevel(level int) string {
	switch level {
	case 10:
		return "Guest"
	case 20:
		return "Reporter"
	case 30:
		return "Developer"
	case 40:
		return "Maintainer"
	case 50:
		return "Owner"
	default:
		return "Unknown"
	}
}

type Entry struct {
	ProjectID          int    `json:"project_id" csv:"project_id"`
	ProjectName        string `json:"project_name" csv:"project_name"`
	UserID             int    `json:"user_id" csv:"user_id"`
	UserName           string `json:"user_name" csv:"user_name"`
	UserState          string `json:"user_state" csv:"user_state"`
	UserAccessLevel    string `json:"user_access_level" csv:"user_access_level"`
	UserCreatedAt      string `json:"user_created_at" csv:"user_created_at"`
	UserLastActivityAt string `json:"user_last_activity_at" csv:"user_last_activity_at"`
}

type GitlabUser struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Username       string `json:"username"`
	State          string `json:"state"`
	AvatarURL      string `json:"avatar_url"`
	WebURL         string `json:"web_url"`
	AccessLevel    string `json:"access_level"`
	CreationDate   string `json:"creation_date"`
	LastActivityOn string `json:"last_activity_on"`
}
