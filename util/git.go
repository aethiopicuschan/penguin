package util

import (
	"net/http"
	"slices"

	"github.com/go-git/go-git/v5/config"
)

// FIXME このためだけにgo-gitを使うのはどうなのか？
func GetNames() (names []string) {
	for _, scope := range []config.Scope{config.LocalScope, config.GlobalScope, config.SystemScope} {
		c, err := config.LoadConfig(scope)
		if err == nil && c != nil {
			if c.User.Name != "" {
				names = append(names, c.User.Name)
			}
		}
	}
	slices.Sort(names)
	slices.Compact(names)
	return
}

func IsRepositoryExist(r string) (exist bool, err error) {
	resp, err := http.Get(r)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	exist = resp.StatusCode == 200
	return
}

func Create(name string, public bool) (err error) {
	visibility := "--private"
	if public {
		visibility = "--public"
	}
	return Execute("gh", "repo", "create", name, visibility)
}

func Clone(name string) (err error) {
	return Execute("gh", "repo", "clone", name)
}

func AddAll() (err error) {
	return Execute("git", "add", ".")
}

func Commit(m string) (err error) {
	return Execute("git", "commit", "-m", m)
}

func Push(branch string) (err error) {
	return Execute("git", "push", "origin", branch)
}
