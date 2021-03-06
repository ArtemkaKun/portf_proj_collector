package main

type Repository struct {
	FullName      string `json:"fullName"`
	LastPushDays  uint32 `json:"lastPushDays"`
	StarsCount    uint16 `json:"starsCount"`
	WatchersCount uint16 `json:"watchersCount"`
	ForksCount    uint16 `json:"forksCount"`
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
