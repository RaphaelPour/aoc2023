//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/RaphaelPour/stellar/strings"
	"github.com/fatih/color"
)

const (
	YEAR = 2023
)

func main() {
	var day int
	if len(os.Args) == 2 {
		day = strings.ToInt(os.Args[1])
	} else {
		day = time.Now().Day()
	}
	fmt.Printf("downloading %d/%d\n", day, YEAR)

	/* download input */
	input, err := util.InputFromURL(
		YEAR,
		day,
		util.WithAoCUserAgent(
			"aoc@raphaelpour.de",
			"github.com/RaphaelPour/aoc2023",
		))
	if err != nil {
		color.Red("error getting input: %s", err)
		return
	}

	/* create day folder */
	folder := fmt.Sprintf("%02d", day)
	path := filepath.Join("days", folder)
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		color.Red("%s already exists", path)
		return
	}

	/* create folder */
	if err := os.MkdirAll(path, 0777); err != nil {
		color.Red("error creating path %s: %s", path, err)
		return
	}

	/* create part.go from template */
	templates, err := template.ParseGlob("templates/*.tpl")
	if err != nil {
		color.Red("error initializing templates: %s", err)
		return
	}

	fd, err := os.OpenFile(
		filepath.Join(path, "part.go"),
		os.O_RDWR|os.O_CREATE,
		0777,
	)
	if err != nil {
		color.Red(
			"error creating part file %s: %s",
			filepath.Join(path, "part.go"),
			err,
		)
		return
	}

	if err := templates.ExecuteTemplate(fd, "part.tpl", nil); err != nil {
		color.Red("error executing part template: %s", err)
		return
	}

	if err := fd.Close(); err != nil {
		color.Red("error writing part file: %s", err)
		return
	}

	/* Create unit test file */
	fd, err = os.OpenFile(
		filepath.Join(path, "part_test.go"),
		os.O_RDWR|os.O_CREATE,
		0777,
	)
	if err != nil {
		color.Red(
			"error creating part unittest file %s: %s",
			filepath.Join(path, "part_test.go"),
			err,
		)
		return
	}

	if err := templates.ExecuteTemplate(fd, "part_test.tpl", nil); err != nil {
		color.Red("error executing part unittest template: %s", err)
		return
	}

	if err := fd.Close(); err != nil {
		color.Red("error writing part unittest file: %s", err)
		return
	}

	/* store input */
	err = ioutil.WriteFile(filepath.Join(path, "input"), []byte(input), 0777)
	if err != nil {
		color.Red("error writing input to file: %s", err)
		return
	}

	color.Green("day %d is ready to solve :)", day)
	color.Green(" code: %s", filepath.Join(path, "part.go"))
	color.Green("tests: %s", filepath.Join(path, "part_test.go"))
}
