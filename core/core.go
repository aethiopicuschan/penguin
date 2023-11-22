package core

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/aethiopicuschan/penguin/static"
	"github.com/aethiopicuschan/penguin/util"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

func Main(cmd *cobra.Command, args []string) (err error) {
	for _, c := range []string{"go", "gh", "git"} {
		if !util.IsExistCommand(c) {
			return fmt.Errorf("'%s' is not installed", c)
		}
	}

	// 必要な情報を聞いてまわる
	packageName, err := util.AskPackageName()
	if err != nil {
		return
	}
	authorName, err := util.SelectAuthorName()
	if err != nil {
		return
	}
	license, err := util.SelectLicense()
	if err != nil {
		return
	}
	hasMain, err := util.AskHasMain()
	if err != nil {
		return
	}
	public, err := util.AskPublic()
	if err != nil {
		return
	}

	// 最終確認
	mp := path.Join("github.com", authorName, packageName)
	confirm, err := util.ConfirmModulePath(mp)
	if err != nil {
		return
	}
	if !confirm {
		return errors.New("canceled")
	}

	// ペンギンっぽいので
	s := spinner.New(spinner.CharSets[12], 100*time.Millisecond)
	s.Start()
	defer s.Stop()

	repoUrl, err := url.JoinPath("https://github.com", authorName, packageName)
	if err != nil {
		return
	}

	// 既にないか確認
	if util.IsExist(packageName) {
		return fmt.Errorf("'%s' is already exists", packageName)
	}
	repExist, err := util.IsRepositoryExist(repoUrl)
	if err != nil {
		return err
	}
	if repExist {
		return fmt.Errorf("'github.com/%s/%s' is already exists", authorName, packageName)
	}

	// GitHubリポジトリを作成してClone
	util.Create(packageName, public)
	util.Clone(packageName)

	// ファイルを展開したりする
	if err = os.Chdir(packageName); err != nil {
		return
	}
	if err = util.Execute("go", "mod", "init", mp); err != nil {
		return
	}
	if err = static.CopyStatic(); err != nil {
		return
	}
	if err = static.CopyTemplates(static.Copy{
		Name:    packageName,
		License: license,
		Author:  authorName,
		Year:    time.Now().Format("2006"),
		HasMain: hasMain,
	}); err != nil {
		return
	}

	// Git add & commit & push
	if err = util.AddAll(); err != nil {
		return
	}
	if err = util.Commit("Initial commit"); err != nil {
		return
	}
	if err = util.Push("main"); err != nil {
		return
	}

	// ブラウザを開く
	if err = util.Open(repoUrl); err != nil {
		return
	}

	return nil
}
