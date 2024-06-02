package commands

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	_ "embed"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
)

//go:embed templates/README.tmpl.md
var readmeTemplate string

//go:embed templates/postinstall.tmpl
var postinstallTemplate string

type InitCommand struct {
	RepoName string `arg:"" optional:"" help:"The name of the port repository"`
	AppName  string `optional:"" help:"The name of the app"`
	AppLink  string `optional:"" help:"The primary link for the app"`
	GitName  string `optional:"" help:"The name to use for the Thanks To section"`
	GitURL   string `optional:"" help:"The profile URL to use for the Thanks To section"`
}

func getGitName() string {
	cmd := exec.Command("git", "config", "user.name")
	stdout, err := cmd.Output()
	// We don't care about errors here, just return empty
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(stdout))
}

func ensureSet(str string) error {
	if len(str) == 0 {
		return errors.New("required input")
	}
	return nil
}

func configureInteractiveFromArgs(ctx *Context, args ...string) {
	for _, arg := range args {
		if arg == "" {
			ctx.Interactive = true
		}
	}
}

func (i *InitCommand) Run(ctx *Context) error {
	var err error

	if len(i.GitName) == 0 {
		i.GitName = getGitName()
	}

	if len(i.GitURL) == 0 {
		i.GitURL = fmt.Sprintf("https://github.com/%s", i.GitName)
	}

	configureInteractiveFromArgs(ctx, i.RepoName, i.AppName, i.AppLink, i.GitName, i.GitURL)

	if ctx.Interactive {
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().Value(&i.RepoName).Title("Repo name").Validate(ensureSet),
				huh.NewInput().Value(&i.AppName).Title("App name").Validate(ensureSet),
				huh.NewInput().Value(&i.AppLink).Title("App link").Validate(ensureSet),
				huh.NewInput().Value(&i.GitName).Title("Git name").Validate(ensureSet),
				huh.NewInput().Value(&i.GitURL).Title("GitHub profile URL").Validate(ensureSet),
			),
		)
		form.WithTheme(huh.ThemeCatppuccin())

		err := form.Run()
		if err != nil {
			log.Fatal("Problem when running interactive form", "err", err)
		}
	}

	i.RepoName, err = filepath.Abs(i.RepoName)

	if err != nil {
		log.Fatal("Failed to resolve path", "err", err)
	}

	if stat, err := os.Stat(i.RepoName); err == nil && stat.IsDir() {
		log.Fatal("Destination already exists, refusing to overwite!", "path", i.RepoName)
	}

	r, err := git.PlainClone(i.RepoName, false, &git.CloneOptions{
		URL:           "https://github.com/catppuccin/template/",
		ReferenceName: "main",
	})

	if err != nil {
		log.Fatal("Problem when cloning template repository", "err", err)
	}

	log.Info("Cloning the template...")
	dotGitPath := filepath.Join(i.RepoName, ".git")
	err = os.RemoveAll(dotGitPath)
	if err != nil {
		log.Errorf("Failed to remove %s directory (%s), continuing...", dotGitPath, err.Error())
	}

	_, err = git.InitWithOptions(r.Storer, nil, git.InitOptions{
		DefaultBranch: plumbing.Main,
	})
	if err != nil {
		log.Error("Failed to init repository, continuing...", "err", err)
	}

	tmpl := template.New("README")
	vars := map[string]any{
		"appLink": func() string {
			return i.AppLink
		},
		"appName": func() string {
			return i.AppName
		},
		"githubName": func() string {
			return i.GitName
		},
		"githubURL": func() string {
			return i.GitURL
		},
		"repoName": func() string {
			return filepath.Base(i.RepoName)
		},
		"repositoryURL": func() string {
			return fmt.Sprintf("https://github.com/%s/%s", i.GitName, strings.ToLower(i.AppName))
		},
	}

	tmpl = tmpl.Funcs(vars)
	tmpl, err = tmpl.Parse(readmeTemplate)
	if err != nil {
		log.Fatal("Problem when parsing README template", "err", err)
	}

	fi, err := os.Create(filepath.Join(i.RepoName, "README.md"))
	if err != nil {
		log.Fatal("Problem when opening README.md for writing", "err", err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			log.Fatal("Problem when closing README.md", "err", err)
		}
	}()

	log.Info("Creating README.md...")
	err = tmpl.Execute(fi, nil)
	if err != nil {
		log.Fatal("Problem when executing README template", "err", err)
	}

	log.Info("Initialisation complete, enjoy!")

	tmpl = template.New("postinstall")
	tmpl = tmpl.Funcs(vars)
	tmpl, err = tmpl.Parse(postinstallTemplate)
	if err != nil {
		log.Fatal("Problem when parsing postinstall template", "err", err)
	}

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, nil)
	if err != nil {
		log.Fatal("Problem when executing postinstall template", "err", err)
	}

	fmt.Println()
	str := buf.String()
	for _, line := range strings.Split(strings.TrimSuffix(str, "\n"), "\n") {
		log.Info(line)
	}

	return nil
}
