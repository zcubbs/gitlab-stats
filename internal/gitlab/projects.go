package gitlab

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/xanzy/go-gitlab"
	"github.com/zcubbs/gitlab-stat-export/models"
	"log"
	"time"
)

var s = spinner.New(spinner.CharSets[9], 100*time.Millisecond)
var client *gitlab.Client

// GetGroupsWithCredentials returns all groups
func GetGroupsWithCredentials(url string, token string, groupId int) []models.GitlabGroup {
	c, err := gitlab.NewClient(
		token,
		gitlab.WithBaseURL(fmt.Sprintf("%s/api/v4", url)),
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	client = c

	return GetGroups(groupId)
}

func GetGroups(groupId int) []models.GitlabGroup {
	var gitlabGroups []models.GitlabGroup

	groups, _, err := client.Groups.ListSubGroups(
		groupId,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to list users: %v", err)
	}

	for _, group := range groups {
		s.Start()
		fmt.Println("===============GROUP==================")
		fmt.Printf("Name=%s, Path=%s, Url=%s\n", group.Name, group.Path, group.WebURL)

		gitlabSubGroups := GetGroups(group.ID)

		gitlabGroups = append(gitlabGroups, models.GitlabGroup{
			ID:          group.ID,
			Name:        group.Name,
			Path:        group.Path,
			Description: group.Description,
			Visibility:  string(group.Visibility),
			ParentID:    group.ParentID,
			WebURL:      group.WebURL,
			Users:       GetGroupUsers(group.ID),
			Projects:    GetGroupProjects(group.ID),
			Groups:      gitlabSubGroups,
		})

		gitlabGroups = append(gitlabGroups, gitlabSubGroups...)
		s.Stop()
		s.Restart()
	}

	return gitlabGroups
}

func GetGroupProjects(groupId int) []models.GitlabProject {
	projects, _, err := client.Groups.ListGroupProjects(
		groupId,
		&gitlab.ListGroupProjectsOptions{
			ListOptions: gitlab.ListOptions{},
		},
	)
	if err != nil {
		log.Fatalf("Failed to fetch projects for group: %v - %v", groupId, err)
	}

	var gitlabProjects []models.GitlabProject

	for _, project := range projects {
		gitlabProjects = append(gitlabProjects, models.GitlabProject{
			ID:          project.ID,
			Name:        project.Name,
			Path:        project.Path,
			Description: project.Description,
			Visibility:  string(project.Visibility),
			WebURL:      project.WebURL,
			Users:       GetProjectUsers(project.ID),
		})
	}

	return gitlabProjects
}

func GetProjectUsers(projectId int) []models.GitlabUser {
	users, _, err := client.Projects.ListProjectsUsers(
		projectId,
		&gitlab.ListProjectUserOptions{
			ListOptions: gitlab.ListOptions{},
			Search:      nil,
		},
	)
	if err != nil {
		log.Fatalf("Failed to fetch users: %v", err)
	}

	var gitlabUsers []models.GitlabUser
	for _, user := range users {
		gitlabUsers = append(gitlabUsers, models.GitlabUser{
			ID:        user.ID,
			Name:      user.Name,
			Username:  user.Username,
			State:     user.State,
			AvatarURL: user.AvatarURL,
			WebURL:    user.WebURL,
		})
	}

	return gitlabUsers
}

func GetGroupUsers(groupId int) []models.GitlabUser {
	users, _, err := client.Groups.ListGroupMembers(
		groupId,
		&gitlab.ListGroupMembersOptions{
			ListOptions: gitlab.ListOptions{},
		},
	)
	if err != nil {
		log.Fatalf("Failed to fetch users: %v", err)
	}

	var gitlabUsers []models.GitlabUser
	for _, user := range users {
		gitlabUsers = append(gitlabUsers, models.GitlabUser{
			ID:          user.ID,
			Name:        user.Name,
			Username:    user.Username,
			State:       user.State,
			AvatarURL:   user.AvatarURL,
			WebURL:      user.WebURL,
			AccessLevel: GetAccessLevel(int(user.AccessLevel)),
		})
	}

	return gitlabUsers
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
