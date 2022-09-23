package main

import (
	"github.com/zcubbs/gitlab-stats/configs"
	"github.com/zcubbs/gitlab-stats/internal/csv"
	"github.com/zcubbs/gitlab-stats/internal/gitlab"
)

func main() {
	configs.Bootstrap()
	groups := gitlab.GetGroupsWithCredentials(configs.Config.Gitlab.Url, configs.Config.Gitlab.PrivateToken, configs.Config.Gitlab.RootGroupId)
	csv.GenerateGitlabCSV(groups)
}
