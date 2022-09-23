package main

import (
	"github.com/zcubbs/gitlab-stat-export/configs"
	"github.com/zcubbs/gitlab-stat-export/internal/csv"
	"github.com/zcubbs/gitlab-stat-export/internal/gitlab"
)

func main() {
	configs.Bootstrap()
	groups := gitlab.GetGroupsWithCredentials(configs.Config.Gitlab.Url, configs.Config.Gitlab.PrivateToken, configs.Config.Gitlab.RootGroupId)
	csv.GenerateGitlabCSV(groups)
}
