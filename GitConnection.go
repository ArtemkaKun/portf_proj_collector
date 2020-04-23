package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v31/github"
	"math"
	"sort"
	"time"
)

var Client = *github.NewClient(nil)

const MY_NICKNAME = "artemkakun"

func GetLastActiveProjects(project_count uint16) []Repository {
	allProjects := GetMyProjects()
	SortProjects(allProjects)

	var lastRepos []Repository
	for i := uint16(0); i < project_count; i++ {
		lastRepos = append(lastRepos, allProjects[i])
	}

	return lastRepos
}

func SortProjects(projects []Repository) {
	sort.Sort(ByLastCommit(projects))
}

func GetMyProjects() []Repository {
	reposList, _, err := Client.Repositories.List(context.Background(), MY_NICKNAME, nil)
	if err != nil {
		fmt.Println(fmt.Errorf("Cannot read from client: %v\n", err))
	}

	var allRepos []Repository
	for _, oneRepo := range reposList {
		allRepos = append(allRepos, DecodeReposData(oneRepo))
	}

	return allRepos
}

func DecodeReposData(data *github.Repository) Repository {
	decodedRepo := Repository{}

	decodedRepo.FullName = data.GetFullName()
	decodedRepo.LastPushDays = CalcDaysToLastPush(data.GetPushedAt())
	decodedRepo.StarsCount = uint32(data.GetStargazersCount())

	return decodedRepo
}

func CalcDaysToLastPush(lastPushTimestamp github.Timestamp) uint32 {
	return uint32(math.Round(time.Now().Sub(lastPushTimestamp.Time).Hours() / 24))
}
