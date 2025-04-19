package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

const (
	BashpackRootVarName = "BASHPACK_ROOT"
	BashpackLibVarName  = "BASHPACK_LIB"

	BashpackRootUserDefaultRootDirectory = "/usr/local/share/bashpack"
)

var rootCmd = &cli.Command{
	Name:  "bashpack",
	Usage: "",
	Action: func(ctx context.Context, c *cli.Command) error {
		fmt.Println("test")
		return nil
	},
	Commands: []*cli.Command{{
		Name:    "cache",
		Aliases: []string{"download"},
		Action: func(ctx context.Context, c *cli.Command) error {
			return errors.New("not implemented")
		},
	}, {
		Name:    "import",
		Aliases: []string{"mainpath", "lipath"},
		Action: func(ctx context.Context, c *cli.Command) error {
			return errors.New("not implemented")
		},
	}},
}

// PackageConfig defines the structure of a bashpack package config file.
type PackageConfig struct {
	// Name is the name of the package.
	Name string `json:"name"`
	// Main is the relative path within the package repo that points to the entrypoint script (e.g.
	// "src/main.sh", "lib.sh", etc.)
	Main string `json:"main"`
}

func main() {
	if err := rootCmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

// func getBashpackRootPath() (string, error) {
// 	if override := os.Getenv(BashpackRootVarName); override != "" {
// 		return override, nil
// 	}

// }

// func getBashpackLibPath() (string, error) {
// 	if override := os.Getenv(BashpackLibVarName); override != "" {
// 		return override, nil
// 	}
// }

// func getValueForBashpackEnvVar(key string) (string, error) {
// 	currentUser, err := user.Current()
// 	if err != nil {
// 		return "", fmt.Errorf("determining user information: %w", err)
// 	}

// 	bpRootParent, err := getRootParentDirByUser(currentUser)
// 	if err != nil {

// 	}

// 	switch key {
// 	case BashpackRootVarName:
// 		if currentUser.Uid == "0" {
// 			return BashpackRootUserDefaultRootDirectory, nil
// 		} else {
// 			return filepath.Join(bpRootParent, ".bashpack"), nil
// 		}

// 	case BashpackLibVarName:
// 		if currentUser.Uid == "0" {
// 			return filepath.Join(BashpackRootUserDefaultRootDirectory, "lib"), nil
// 		}

// 	default:
// 		return "", fmt.Errorf("")
// 	}
// }

// func getRootParentDirByUser(u *user.User) (string, error) {
// 	// bpRootParent will either be the non-root user's home directory, or a system-wide path
// 	bpRootParent := u.HomeDir
// 	if bpRootParent == "" && u.Uid != "0" {
// 		return "", errors.New("could not determine directory structure to use for bashpack's root directory: non-root user has no home directory")
// 	}

// 	var root string
// 	if u.Uid == "0" {
// 		root = BashpackRootUserDefaultRootDirectory
// 	} else {
// 		root = filepath.Join(bpRootParent, ".bashpack")
// 	}

// 	return root, nil
// }
