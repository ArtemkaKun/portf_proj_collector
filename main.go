package main

import "fmt"

func main() {
	projects := GetLastActiveProjects(10)

	for _, one_proj := range projects {
		fmt.Println(one_proj)
	}
}
