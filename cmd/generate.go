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
	tplArchive = `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`
	dotDroneTpl = `a2luZDogcGlwZWxpbmUKdHlwZTogZG9ja2VyCm5hbWU6ICMjI19fUFJPSl9PUkdfXyMjIzo6IyMj
X19QUk9KX05BTUUjX18jIyMKCnN0ZXBzOgogIC0gbmFtZTogYnVpbGQKICAgIGltYWdlOiBnb2xh
bmc6YWxwaW5lCiAgICBjb21tYW5kczoKICAgICAgLSAiYXBrIGFkZCBtYWtlIGdpdCIKICAgICAg
LSAibWFrZSBzd2FnIgogICAgICAtICJtYWtlIgoKICAtIG5hbWU6IGRvY2tlcgogICAgaW1hZ2U6
IHBsdWdpbnMvZG9ja2VyOmxhdGVzdAogICAgc2V0dGluZ3M6CgogIC0gbmFtZTogZGVwbG95CiAg
ICBpbWFnZTogYXBwbGVib3kvZHJvbmUtc3NoOmxhdGVzdAogICAgc2V0dGluZ3M6Cgp0cmlnZ2Vy
OgogIGV2ZW50OgogICAgLSB0YWcK`
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

	fmt.Println(" => ", promptProjPath.Label, " : ", color.HiYellowString(resultProjPath))

	err = makeProjDir()
	if err != nil {
		fmt.Println(color.HiRedString(err.Error()))
		os.Exit(1)
	}

	err = unpackTpl()
	if err != nil {
		fmt.Println(color.HiRedString(err.Error()))
		os.Exit(1)
	}

	if projDrone {
		err = dotDrone()
		if err != nil {
			fmt.Println(color.HiRedString(err.Error()))
			os.Exit(1)
		}
	}

	fmt.Println(color.HiGreenString("Project created. Run 'go mod tidy' in your project directory and enjoy it."))
}

func makeProjDir() error {
	fmt.Println(color.HiCyanString("Making working directory"))
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

	fmt.Println(" => OK")

	return nil
}

func unpackTpl() error {
	fmt.Println(color.HiCyanString("Unpacking templates"))
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

	filePath := projName + ".json"
	err = os.WriteFile(filePath, []byte("{}"), 0644)
	if err != nil {
		return err
	}

	fmt.Println(" => OK")

	return nil
}

func dotDrone() error {
	fmt.Println(color.HiCyanString("Generating .drone.yml"))
	raw, err := base64.StdEncoding.DecodeString(dotDroneTpl)
	if err != nil {
		return err
	}

	filePath := ".drone.yml"
	content := string(raw)
	content = strings.ReplaceAll(content, "###__PROJ_ORG__###", projOrg)
	content = strings.ReplaceAll(content, "###__PROJ_NAME__###", projName)

	err = os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return err
	}

	fmt.Println(" => OK")

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
