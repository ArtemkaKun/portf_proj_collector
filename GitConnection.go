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

func GetLastActiveProjects(projectCount uint16, user string) []Repository {
	allProjects := GetMyProjects(user)
	SortProjects(allProjects)

	var lastRepos []Repository
	for i := uint16(0); i < projectCount; i++ {
		lastRepos = append(lastRepos, allProjects[i])
	}

	return lastRepos
}

func SortProjects(projects []Repository) {
	sort.Sort(ByLastCommit(projects))
}

func GetMyProjects(user string) []Repository {
	reposList, _, err := Client.Repositories.List(context.Background(), user, nil)
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
	decodedRepo.StarsCount = uint16(data.GetStargazersCount())
	decodedRepo.WatchersCount = uint16(data.GetWatchersCount())
	decodedRepo.ForksCount = uint16(data.GetForksCount())

	return decodedRepo
}

func CalcDaysToLastPush(lastPushTimestamp github.Timestamp) uint32 {
	return uint32(math.Round(time.Now().Sub(lastPushTimestamp.Time).Hours() / 24))
}

func CalcAllStarts(user string) (starsCount uint16) {
	for _, project := range GetMyProjects(user) {
		starsCount += project.StarsCount
	}

	return
}

func CalcAllWatchers(user string) (watchersCount uint16) {
	for _, project := range GetMyProjects(user) {
		watchersCount += project.WatchersCount
	}

	return
}

func CalcAllForks(user string) (forksCount uint16) {
	for _, project := range GetMyProjects(user) {
		forksCount += project.ForksCount
	}

	return
}
