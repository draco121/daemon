package models

type FolderReport struct {
	Path         string   `json:"path"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Size         int64    `json:"size"`
	FileCount    int      `json:"fileCount"`
	FolderCount  int      `json:"folderCount"`
	LogFileCount int      `json:"logFileCount"`
	LogFilePaths []string `json:"logFilePaths"`
	Error        string   `json:"error"`
}

type FileReport struct {
	Path        string `json:"path"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Size        int64  `json:"size"`
	Error       string `json:"error"`
}

type CommandReport struct {
	Script      string `json:"script"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Output      string `json:"output"`
	Error       string `json:"error"`
}

type Report struct {
	AppId          string          `json:"appId"`
	SpecId         string          `json:"specId"`
	GeneratedAt    string          `json:"generatedAt"`
	BundleName     string          `json:"bundleName"`
	BundleStatus   string          `json:"bundleStatus"`
	HostName       string          `json:"hostName"`
	FolderReports  []FolderReport  `json:"folderReports"`
	FileReports    []FileReport    `json:"fileReports"`
	CommandReports []CommandReport `json:"commandReports"`
}
