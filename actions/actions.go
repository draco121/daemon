package actions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
	httpclient "shak-daemon/httpClient"
	"shak-daemon/models"
	"shak-daemon/utils"
	"strings"
)

func UpdateSpecAction() {
	fmt.Println("syncing spec from server..........")
	newSpec := models.Spec{}
	httpclient.GetLatestSpec(&newSpec)
	buf, err := json.Marshal(newSpec)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create(utils.SpecDir)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = file.Write(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("spec synced successfully...........")
}

func InspectFolderAction(spec *models.Spec, report *models.Report) error {
	fmt.Println("inspecting the folders from spec ...........")
	folders := spec.Folders
	if len(folders) == 0 {
		folderReport := models.FolderReport{
			Error: "No folders Configured to inspect",
		}
		report.FolderReports = append(report.FolderReports, folderReport)
		return nil
	}
	for i := 0; i < len(folders); i++ {
		folderReport := models.FolderReport{
			Path:        folders[i].Path,
			Name:        folders[i].Name,
			Description: folders[i].Description,
		}
		dir, err := os.Open(folders[i].Path)
		if err != nil {
			folderReport.Error = "error occured while inspecting folder: " + err.Error()
			report.FolderReports = append(report.FolderReports, folderReport)
			continue
		}
		defer dir.Close()
		stats, err := dir.Stat()
		if err != nil {
			folderReport.Error = "error occured while inspecting folder: " + err.Error()
			report.FolderReports = append(report.FolderReports, folderReport)
			continue
		}
		folderReport.Size = stats.Size()
		files, err := dir.Readdirnames(0)
		if err != nil {
			folderReport.Error = "error occured while inspecting folder: " + err.Error()
			report.FolderReports = append(report.FolderReports, folderReport)
			continue
		}
		fileCount, folderCount, logFileCount := 0, 0, 0
		for j := 0; j < len(files); j++ {
			info, err := os.Stat(path.Join(folders[i].Path, files[j]))
			if err != nil {
				continue
			} else {
				if info.IsDir() {
					folderCount++
				} else {
					fileCount++
					if strings.HasSuffix(info.Name(), ".log") {
						logFileCount++
						relPath := path.Join(report.BundleName, dir.Name(), files[j])
						folderReport.LogFilePaths = append(folderReport.LogFilePaths, relPath)
					}
				}
			}
		}
		folderReport.FileCount = fileCount
		folderReport.FolderCount = folderCount
		folderReport.LogFileCount = logFileCount
		report.FolderReports = append(report.FolderReports, folderReport)
		err = utils.Copy(folders[i].Path, fmt.Sprintf("%s/%s", utils.GetFolderDumpDir(report.BundleName), stats.Name()))
		if err != nil {
			return err
		}
	}
	fmt.Println("folder inspection completed .......")
	return nil
}

func InspectFileAction(spec *models.Spec, report *models.Report) error {
	fmt.Println("inspecting files from the spec ................")
	files := spec.Files
	if len(files) == 0 {
		fileReport := models.FileReport{
			Error: "No files Configured to inspect",
		}
		report.FileReports = append(report.FileReports, fileReport)
		return nil
	}
	for i := 0; i < len(files); i++ {
		fileReport := models.FileReport{
			Path:        files[i].Path,
			Name:        files[i].Name,
			Description: files[i].Description,
		}
		info, err := os.Stat(files[i].Path)
		if err != nil {
			fileReport.Error = "error occurred while inspecting file: " + err.Error()
		} else {
			fileReport.Size = info.Size()
			report.FileReports = append(report.FileReports, fileReport)
			err := utils.Copy(files[i].Path, fmt.Sprintf("%s/%s", utils.GetFileDumpDir(report.BundleName), info.Name()))
			if err != nil {
				return err
			}
		}
	}
	fmt.Println("file inspection completed ...........")
	return nil
}

func RunCommandAction(spec *models.Spec, report *models.Report) error {
	fmt.Println("running commands in system .........")
	commands := spec.Commands
	if len(commands) == 0 {
		commandReport := models.CommandReport{
			Error: "No files Configured to inspect",
		}
		report.CommandReports = append(report.CommandReports, commandReport)
		return nil
	}
	for i := 0; i < len(commands); i++ {
		commandReport := models.CommandReport{
			Script:      commands[i].Script,
			Name:        commands[i].Script,
			Description: commands[i].Description,
		}
		shell := ""
		args := ""
		if runtime.GOOS == "windows" {
			shell = "powershell.exe"
			args = "iex"
		} else {
			shell = "/bin/bash"
			args = "-c"
		}
		cmd, err := exec.Command(shell, args, commands[i].Script).Output()
		if err != nil {
			commandReport.Error = "error occurred while executing the command: " + err.Error()
		} else {
			output := bytes.NewBuffer(cmd).String()
			commandReport.Output = output
		}
		report.CommandReports = append(report.CommandReports, commandReport)
		file, err := os.Create(utils.GetCommandDumpDir(report.BundleName, commands[i].Name))
		if err != nil {
			return err
		}
		_, err = file.Write(cmd)
		if err != nil {
			return err
		}
	}
	fmt.Println("commands execution completed .....")
	return nil
}

func CreateArchiveAction(BundleName string) (string, error) {
	fmt.Println("creating bundle archive........")
	bundlePath := utils.GetBundleDir(BundleName)
	bundleArchivePath := utils.GetBundleArchivePath(BundleName)
	err := utils.Archive(bundlePath, bundleArchivePath)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("bundle archive created ......")
	}
	return bundleArchivePath, err
}

func CleanUpAction(BundleName string) {
	fmt.Println("cleaning up working directory...")
	bundlePath := utils.GetBundleDir(BundleName)
	bundleArchivePath := utils.GetBundleArchivePath(BundleName)
	os.RemoveAll(bundlePath)
	os.RemoveAll(bundleArchivePath)
}
