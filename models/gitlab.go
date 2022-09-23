package models

type GitlabGroup struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Description string `json:"description"`
	Visibility  string `json:"visibility"`
	ParentID    int    `json:"parent_id"`
	WebURL      string `json:"web_url"`

	Users    []GitlabUser    `json:"users"`
	Projects []GitlabProject `json:"projects"`
	Groups   []GitlabGroup   `json:"groups"`
}

type GitlabProject struct {
	ID                int    `json:"id"`
	Description       string `json:"description"`
	DefaultBranch     string `json:"default_branch"`
	Visibility        string `json:"visibility"`
	SSHURLToRepo      string `json:"ssh_url_to_repo"`
	HTTPURLToRepo     string `json:"http_url_to_repo"`
	WebURL            string `json:"web_url"`
	ReadmeURL         string `json:"readme_url"`
	Name              string `json:"name"`
	NameWithNamespace string `json:"name_with_namespace"`
	Path              string `json:"path"`
	PathWithNamespace string `json:"path_with_namespace"`

	Users []GitlabUser `json:"users"`
	Tags  []GitlabTag  `json:"tags"`
}

type GitlabUser struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	State       string `json:"state"`
	AvatarURL   string `json:"avatar_url"`
	WebURL      string `json:"web_url"`
	AccessLevel string `json:"access_level"`
}

type GitlabTag struct {
	Name string `json:"name"`
}

type GitlabFlatProjectInfo struct {
	// Group
	GroupID          int    `json:"group_id" csv:"group_id"`
	GroupName        string `json:"group_name" csv:"group_name"`
	GroupPath        string `json:"group_path" csv:"group_path"`
	GroupDescription string `json:"group_description" csv:"group_description"`
	GroupVisibility  string `json:"group_visibility" csv:"group_visibility"`
	GroupParentID    int    `json:"group_parent_id" csv:"group_parent_id"`
	GroupWebURL      string `json:"group_web_url" csv:"group_web_url"`

	// PROJECT
	ProjectID                int    `json:"project_id" csv:"project_id"`
	ProjectDescription       string `json:"project_description" csv:"project_description"`
	ProjectDefaultBranch     string `json:"project_default_branch" csv:"project_default_branch"`
	ProjectVisibility        string `json:"project_visibility" csv:"project_visibility"`
	ProjectSSHURLToRepo      string `json:"project_ssh_url_to_repo" csv:"project_ssh_url_to_repo"`
	ProjectHTTPURLToRepo     string `json:"project_http_url_to_repo" csv:"project_http_url_to_repo"`
	ProjectWebURL            string `json:"project_web_url" csv:"project_web_url"`
	ProjectReadmeURL         string `json:"project_readme_url" csv:"project_readme_url"`
	ProjectName              string `json:"project_name" csv:"project_name"`
	ProjectNameWithNamespace string `json:"project_name_with_namespace" csv:"project_name_with_namespace"`
	ProjectPath              string `json:"project_path" csv:"project_path"`
	ProjectPathWithNamespace string `json:"project_path_with_namespace" csv:"project_path_with_namespace"`

	// USER
	ProjectUserID          int    `json:"project_user_id" csv:"project_user_id"`
	ProjectUserName        string `json:"project_user_name" csv:"project_user_name"`
	ProjectUserState       string `json:"project_user_state" csv:"project_user_state"`
	ProjectUserAvatarURL   string `json:"project_user_avatar_url" csv:"project_user_avatar_url"`
	ProjectUserWebURL      string `json:"project_user_web_url" csv:"project_user_web_url"`
	ProjectUserAccessLevel string `json:"project_user_access_level" csv:"project_user_access_level"`
}
