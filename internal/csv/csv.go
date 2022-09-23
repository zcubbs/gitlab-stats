package csv

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/zcubbs/gitlab-stat-export/internal/utils"
	"github.com/zcubbs/gitlab-stat-export/models"
	"os"
)

func GenerateGitlabCSV(groups []models.GitlabGroup) {
	flatProjects := FlattenGitlabGroups(groups)
	utils.PrettyPrint(flatProjects)

	gitlabStatsFile, err := os.OpenFile("gitlab_stats.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer func(gitlabStatsFile *os.File) {
		err := gitlabStatsFile.Close()
		if err != nil {
			panic(err)
		}
	}(gitlabStatsFile)

	csvContent, err := gocsv.MarshalString(&flatProjects)
	err = gocsv.MarshalFile(&flatProjects, gitlabStatsFile)
	if err != nil {
		panic(err)
	}
	fmt.Println(csvContent)
}

func FlattenGitlabGroups(groups []models.GitlabGroup) []models.GitlabFlatProjectInfo {
	var flatProjects []models.GitlabFlatProjectInfo

	for _, group := range groups {
		for _, project := range group.Projects {
			for _, user := range project.Users {
				flatProjects = append(flatProjects, models.GitlabFlatProjectInfo{
					GroupID:            group.ID,
					GroupName:          group.Name,
					GroupPath:          group.Path,
					GroupDescription:   group.Description,
					GroupVisibility:    group.Visibility,
					GroupParentID:      group.ParentID,
					GroupWebURL:        group.WebURL,
					ProjectID:          project.ID,
					ProjectName:        project.Name,
					ProjectPath:        project.Path,
					ProjectDescription: project.Description,
					ProjectVisibility:  project.Visibility,
					ProjectWebURL:      project.WebURL,
					ProjectUserID:      user.ID,
					ProjectUserName:    user.Name,
					ProjectUserState:   user.State,
					ProjectUserWebURL:  user.WebURL,
				})

			}
		}
	}

	return flatProjects
}
