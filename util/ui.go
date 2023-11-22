package util

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/aethiopicuschan/penguin/static"
	"github.com/manifoldco/promptui"
)

func AskPackageName() (string, error) {
	p := promptui.Prompt{
		Label: "Package name",
		Validate: func(input string) error {
			if len(input) < 1 {
				return errors.New("package name must be at least 1 character")
			}
			r := regexp.MustCompile("^[a-z0-9]+$")
			if !r.MatchString(input) {
				return errors.New("invalid package name")
			}
			return nil
		},
	}
	return p.Run()
}

func SelectAuthorName() (author string, err error) {
	names := GetNames()
	if len(names) == 0 {
		return AskAuthorName()
	}
	s := promptui.Select{
		Label: "Author name",
		Items: append(names, "Manual input"),
	}
	i, author, err := s.Run()
	if i == len(names) {
		return AskAuthorName()
	}
	return
}

func SelectLicense() (license string, err error) {
	licenses, err := static.ListLicenses()
	if err != nil {
		return
	}
	s := promptui.Select{
		Label: "License",
		Items: licenses,
	}
	_, license, err = s.Run()
	return
}

func AskAuthorName() (string, error) {
	p := promptui.Prompt{
		Label: "Author name",
		Validate: func(input string) error {
			r := regexp.MustCompile("^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$")
			if len(input) < 40 && len(input) > 0 && r.MatchString(input) {
				return nil
			} else {
				return errors.New("invalid author name")
			}
		},
	}
	return p.Run()
}

func ConfirmModulePath(mp string) (bool, error) {
	s := promptui.Select{
		Label: fmt.Sprintf("Is the module path '%s' correct?", mp),
		Items: []string{"Yes, it is.", "No!"},
	}
	i, _, err := s.Run()
	return i == 0, err
}

func AskPublic() (bool, error) {
	s := promptui.Select{
		Label: "repository visibility",
		Items: []string{"Private", "Public"},
	}
	i, _, err := s.Run()
	return i == 1, err
}

func AskHasMain() (bool, error) {
	s := promptui.Select{
		Label: "repository type",
		Items: []string{"main", "Library"},
	}
	i, _, err := s.Run()
	return i == 0, err
}
