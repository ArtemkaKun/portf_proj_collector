package main

type Repository struct {
	FullName     string
	LastPushDays uint32
	StarsCount   uint32
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
