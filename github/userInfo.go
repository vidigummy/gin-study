package github

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v48/github"
)

func GetUserInfo(userName string) (map[string]int, error) {
	// 내 레포에 어떤 언어가 얼마나 있는지 확인할 수 있음.
	repoList, err := getRepo(userName)
	if err != nil {
		fmt.Println(err)
	}
	// 레포 리스트 확인용(무슨 레포가 있는지)
	// fmt.Println(repoList)

	// 랭귀지맵은 사용자의 언어 map이라고 할 수 있죠
	languageMap := make(map[string]int)

	// api들 동시에 호출, 빠르게 가져오기
	// 403 뜨네? 큰일난걸지두...
	repoChan := make(chan map[string]int)
	for _, repoName := range repoList {
		go getRepoLanguages(userName, repoName, repoChan)
	}
	// 가져온 리스트 싹 다 한 레포에 집어넣기
	for i := 0; i < len(repoList); i++ {
		tmpLanguageMap := <-repoChan
		for key, value := range tmpLanguageMap {
			checkData, exists := languageMap[key]
			if !exists {
				languageMap[key] = value
			} else {
				languageMap[key] = checkData + value
			}
		}
	}
	fmt.Println(languageMap)
	return languageMap, nil
}

func getRepo(username string) ([]string, error) {
	client := github.NewClient(nil)
	ctx := context.Background()
	opt := &github.RepositoryListOptions{Type: "public"}
	repos, _, err := client.Repositories.List(ctx, username, opt)
	if err != nil {
		fmt.Println(err)
		return strings.Split("err err", " "), err
	}
	var reposList []string = make([]string, len(repos))
	repoChan := make(chan string)
	for _, repo := range repos {
		go getUrl(*repo.LanguagesURL, repoChan)
	}
	for i := 0; i < len(repos)-1; i++ {
		tmp := <-repoChan
		reposList[i] = tmp
	}
	return reposList, nil
}

func getUrl(url string, c chan<- string) {
	slicedURL := strings.Split(url, "/")
	c <- slicedURL[5]
}

func getRepoLanguages(userName string, repoName string, c chan<- map[string]int) {
	client := github.NewClient(nil)
	ctx := context.Background()
	languages, _, err := client.Repositories.ListLanguages(ctx, userName, repoName)
	if err != nil {
		fmt.Println(err)
	}
	c <- languages
}
