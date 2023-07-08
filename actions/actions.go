package actions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path"
	"shak-daemon/models"
	"shak-daemon/utils"
	"strings"
)

func UpdateSpecAction(configPath string, spec *models.Spec) {
	configJson, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(configJson, &spec)
	if err != nil {
		panic(err)
	}
}

func InspectFolderAction(spec *models.Spec, report *models.Report) error {
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
						folderReport.LogFilePaths = append(folderReport.LogFilePaths, path.Join(folders[i].Path, files[j]))
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
	return nil
}

func InspectFileAction(spec *models.Spec, report *models.Report) error {
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
	return nil
}

func RunCommandAction(spec *models.Spec, report *models.Report) error {
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
		cmd, err := exec.Command("powershell.exe", "iex", commands[i].Script).Output()
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
	return nil
}

func CreateArchiveAction(BundleName string) error {
	bundlePath := utils.GetBundleDir(BundleName)
	bundleArchivePath := utils.GetBundleArchivePath(BundleName)
	return utils.Archive(bundlePath, bundleArchivePath)
}

func CleanUpAction(BundleName string) {
	bundlePath := utils.GetBundleDir(BundleName)
	bundleArchivePath := utils.GetBundleArchivePath(BundleName)
	cmd := exec.Command("rm", "-rf", bundlePath)
	cmd.Run()
	cmd = exec.Command("rm", "-rf", bundleArchivePath)
	cmd.Run()
}
