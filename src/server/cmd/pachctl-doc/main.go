package main

import (
	"regexp"

	"github.com/pachyderm/pachyderm/src/server/cmd/pachctl/cmd"
	"github.com/pachyderm/pachyderm/src/server/pkg/cmdutil"

	"github.com/spf13/cobra/doc"
)

var MD_LINK_PATTERN = regexp.MustCompile("\\.md$")

type appEnv struct{}

func main() {
	cmdutil.Main(do, &appEnv{})
}

func filePrepender(value string) string {
	// No prepending needs to be done
	return value
}

func linkHandler(value string) string {
	return MD_LINK_PATTERN.ReplaceAllLiteralString(value, ".html")
}

func do(appEnvObj interface{}) error {
	rootCmd, err := cmd.PachctlCmd()
	if err != nil {
		return err
	}
	return doc.GenMarkdownTreeCustom(rootCmd, "./doc/pachctl/", filePrepender, linkHandler)
}
