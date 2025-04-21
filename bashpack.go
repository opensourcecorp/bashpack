package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/urfave/cli/v3"
)

const (
	BashpackRootVarName = "BASHPACK_ROOT"

	BashpackRootUserDefaultRootDirectory = "/usr/local/share/bashpack"

	BashpackRootFlag = "bashpack-root"
)

func newRootCmd() *cli.Command {
	cmd := &cli.Command{
		Name:  "bashpack",
		Usage: "a package manager for bash",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        BashpackRootFlag,
				Usage:       "The root directory for bashpack's cache, etc.",
				DefaultText: fmt.Sprintf("${HOME}/.local/share/bashpack for non-root user, and %s for root user", BashpackRootUserDefaultRootDirectory),
				Sources:     cli.EnvVars(BashpackRootVarName),
			},
		},
		Action: rootAction,
		Commands: []*cli.Command{
			{
				Name:   "run",
				Usage:  "Runs a script/program defined as a bashpack package",
				Action: runAction,
			}, {
				Name: "download",
				Usage: fmt.Sprint(
					"Fetches bashpack packages to the local cache. ",
					"Will be run behind the scenes for other operations that need it, ",
					"but can also be run directly to separate download vs. usage stages if desired.",
				),
				Aliases: []string{"cache"},
				Action:  downloadAction,
			}, {
				Name: "import",
				Usage: fmt.Sprint(
					"Returns the path on disk to the entrypoint of a bashpack package. ",
					"Aliased to 'mainpath/libpath' as well for clarity, but 'import' is the primary name because of how library usage within bash is intended ",
					"(e.g. 'source $(bashpack import github.com/person/library)')",
				),
				Aliases: []string{"mainpath", "libpath"},
				Action:  importAction,
			},
		},
	}

	return cmd
}

func rootAction(ctx context.Context, cmd *cli.Command) error {
	fmt.Println("bashpack root command")
	bpRoot := initBashpackRootPath(cmd.String(BashpackRootFlag))
	fmt.Printf("--bashpack-root: %s\n", bpRoot)
	return errors.New("not implemented")
}

func runAction(ctx context.Context, c *cli.Command) error {
	fmt.Println("bashpack run subcommand")
	return errors.New("not implemented")
}

func downloadAction(ctx context.Context, c *cli.Command) error {
	fmt.Println("bashpack download subcommand")
	return errors.New("not implemented")
}

func importAction(ctx context.Context, c *cli.Command) error {
	fmt.Println("bashpack import subcommand")
	return errors.New("not implemented")
}

// PackageConfig defines the structure of a bashpack package config file.
type PackageConfig struct {
	// Name is the name of the package.
	Name string `json:"name"`
	// Main is the relative path within the package repo that points to the entrypoint script (e.g.
	// "src/main.sh", "lib.sh", etc.)
	Main string `json:"main"`
}

// intentionally does not return an error and instead exits with status code 1 if any error is
// encountered, because this is only to be used as an early-stage value determination and any error
// should be fatal.
func initBashpackRootPath(override string) string {
	if override != "" {
		if err := os.MkdirAll(override, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "initializing bashpack root directory: %s\n", err.Error())
			os.Exit(1)
		}
		return override
	}

	currentUser, err := user.Current()
	if err != nil {
		fmt.Fprintf(os.Stderr, "determining user information: %s\n", err.Error())
		os.Exit(1)
	}

	bpRoot, err := getRootDirByUser(currentUser)
	if err != nil {
		fmt.Fprintf(os.Stderr, "determining bashpack root directory: %s\n", err.Error())
		os.Exit(1)
	}

	if err := os.MkdirAll(bpRoot, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "initializing bashpack root directory: %s\n", err.Error())
		os.Exit(1)
	}

	return bpRoot
}

func getRootDirByUser(u *user.User) (string, error) {
	// bpRootParent will either be the non-root user's home directory, or a system-wide path
	bpRootParent := u.HomeDir
	if bpRootParent == "" && u.Uid != "0" {
		return "", errors.New("could not determine directory structure to use for bashpack's root directory: non-root user has no home directory")
	}

	var root string
	if u.Uid == "0" {
		root = BashpackRootUserDefaultRootDirectory
	} else {
		root = filepath.Join(bpRootParent, ".local", "share", "bashpack")
	}

	return root, nil
}

func getBashpackLibPath(rootOverride string) string {
	return filepath.Join(initBashpackRootPath(rootOverride), "lib")
}
