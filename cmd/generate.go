/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2021 HereweTech Co.LTD
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

/**
 * @file generate.go
 * @package cmd
 * @author Dr.NP <np@herewe.tech>
 * @since 09/25/2023
 */

package cmd

import (
	"archive/tar"
	"bytes"
	"compress/bzip2"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

const (
	tplArchive = `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
dMbUGA==`
)

var (
	generateCmd = &cobra.Command{
		Use:   "new",
		Short: "Generate new project",
		Run:   generateRun,
	}

	projName      = ""
	projOrg       = ""
	projAuthor    = ""
	projDockerTag = ""
	projPath      = ""
	projDrone     = false
)

func generateRun(cmd *cobra.Command, args []string) {
	fmt.Println("Generate new project")
	if len(args) > 0 {
		projName = args[0]
	}

	promptProjName := promptui.Prompt{
		Label:   "Project name",
		Default: projName,
		Validate: func(input string) error {
			if input == "" {
				return errors.New("empty project name")
			}

			projName = input

			return nil
		},
	}

	promptProjOrg := promptui.Prompt{
		Label:   "Project organization",
		Default: projOrg,
		Validate: func(input string) error {
			if input == "" {
				return errors.New("empty project organization")
			}

			projOrg = input

			return nil
		},
	}

	promptProjAuth := promptui.Prompt{
		Label:   "Project author",
		Default: projAuthor,
		Validate: func(input string) error {
			if input == "" {
				return errors.New("empty project author")
			}

			projAuthor = input

			return nil
		},
	}

	promptProjDockerTag := promptui.Prompt{
		Label:   "Docker image tag",
		Default: projDockerTag,
		Validate: func(input string) error {
			if input == "" {
				return errors.New("empty project organization")
			}

			projDockerTag = input

			return nil
		},
	}

	promptProjPath := promptui.Prompt{
		Label: "Project path",
		Validate: func(input string) error {
			if input == "" {
				return errors.New("empty project path")
			}

			projPath = input

			return nil
		},
	}

	promptProjDrone := promptui.Prompt{
		Label:     "Enable DroneCI",
		IsConfirm: true,
	}

	resultProjName, err := promptProjName.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(" => ", promptProjName.Label, " : ", color.GreenString(resultProjName))

	resultProjOrg, err := promptProjOrg.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(" => ", promptProjOrg.Label, " : ", color.GreenString(resultProjOrg))

	resultProjAuthor, err := promptProjAuth.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(" => ", promptProjAuth.Label, " : ", color.GreenString(resultProjAuthor))
	promptProjDockerTag.Default = projOrg + "/" + projName

	resultProjDockerTag, err := promptProjDockerTag.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(" => ", promptProjDockerTag.Label, " : ", color.GreenString(resultProjDockerTag))

	resultProjDrone, err := promptProjDrone.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if resultProjDrone == "y" || resultProjDrone == "Y" {
		projDrone = true
	}

	fmt.Println(" => ", promptProjDrone.Label, " : ", projDrone)
	promptProjPath.Default = "./" + projName

	resultProjPath, err := promptProjPath.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(" => ", promptProjPath.Label, " : ", color.YellowString(resultProjPath))

	err = makeProjDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = unpackTpl()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func makeProjDir() error {
	fmt.Println("Making working directory")
	var err error

	err = os.MkdirAll(projPath, 0755)
	if err != nil {
		return err
	}

	err = os.Chdir(projPath)
	if err != nil {
		return err
	}

	cmdGit := exec.Command("git", "init", "--quiet", "--initial-branch=main")
	err = cmdGit.Run()
	if err != nil {
		return err
	}

	cmdMod := exec.Command("go", "mod", "init", projName)
	err = cmdMod.Run()
	if err != nil {
		return err
	}

	return nil
}

func unpackTpl() error {
	fmt.Println("Unpacking templates")
	raw, err := base64.StdEncoding.DecodeString(tplArchive)
	if err != nil {
		return err
	}

	r := bytes.NewReader(raw)
	ur := bzip2.NewReader(r)
	b, err := io.ReadAll(ur)
	if err != nil {
		return err
	}

	tr := tar.NewReader(bytes.NewBuffer(b))
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		fi := hdr.FileInfo()
		if fi.IsDir() {
			dirPath, _ := filepath.Abs(hdr.Name)
			os.MkdirAll(dirPath, 0755)
		} else {
			cb, _ := io.ReadAll(tr)
			content := string(cb)
			today := time.Now().Format("01/02/2006")
			// Write contents
			filePath, _ := filepath.Abs(strings.TrimRight(hdr.Name, ".tpl"))
			content = strings.ReplaceAll(content, "###__PROJ_NAME__###", projName)
			content = strings.ReplaceAll(content, "###__PROJ_ORG__###", projOrg)
			content = strings.ReplaceAll(content, "###__PROJ_AUTHOR__###", projAuthor)
			content = strings.ReplaceAll(content, "###__TODAY__###", today)

			err = os.WriteFile(filePath, []byte(content), 0644)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
