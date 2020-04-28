package main

type Repository struct {
	FullName     string `json:"fullName"`
	LastPushDays uint32 `json:"lastPushDays"`
	StarsCount   uint16 `json:"starsCount"`
}

type ByLastCommit []Repository

func (repos ByLastCommit) Len() int {
	return len(repos)
}

func (repos ByLastCommit) Less(i, j int) bool {
	return repos[i].LastPushDays < repos[j].LastPushDays
}

func (repos ByLastCommit) Swap(i, j int) {
	repos[i], repos[j] = repos[j], repos[i]
}

type GetProjectsRequest struct {
	ProjectsCount uint16 `json:"projectsCount"`
	Username      string `json:"username"`
}
