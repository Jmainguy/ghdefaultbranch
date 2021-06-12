package main

import (
	"flag"
	"fmt"
)

func main() {

	repositoryPtr := flag.String("repository", "", "Repository to rename branch on, example: Jmainguy/timesheets, not used if renameAll flag is specified")
	renameAllPtr := flag.Bool("renameAll", false, "Bool: rename default branch in all repos or not")
	askTokenPtr := flag.Bool("askToken", false, "Bool: Force user to enter token instead of trying the env variable ghdefaultbranchToken")
	defaultBranchPtr := flag.String("defaultBranch", "main", "Name of the desired defaultBranch, used with renameAll and repository flags")

	flag.Parse()

	token := getToken(*askTokenPtr)

	// Rename all repos
	if *renameAllPtr {
		repos := getRepos(token)
		for _, repo := range repos {
			if *repo.DefaultBranch != *defaultBranchPtr {
				fmt.Println("Changing default branch")
				fmt.Printf("Changing default branch for %s to %s, its current DefaultBranch is %s \n", *repo.FullName, *defaultBranchPtr, *repo.DefaultBranch)
				renameBranch(*repo.FullName, token, *repo.DefaultBranch, *defaultBranchPtr)
			}
		}
		// List all repos
	} else if *repositoryPtr == "" {
		repos := getRepos(token)
		for _, repo := range repos {
			fmt.Printf("FullName: %s, DefaultBranch: %s, HTMLURL: %s\n", *repo.FullName, *repo.DefaultBranch, *repo.HTMLURL)
		}
		// Rename single repo
	} else {
		repo := getRepo(token, *repositoryPtr)
		if *repo.DefaultBranch != *defaultBranchPtr {
			renameBranch(*repositoryPtr, token, *repo.DefaultBranch, *defaultBranchPtr)
		}
	}
}
