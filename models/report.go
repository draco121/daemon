package models

type FolderReport struct {
	Path         string
	Name         string
	Description  string
	Size         int64
	FileCount    int
	FolderCount  int
	LogFileCount int
	LogFilePaths []string
	Error        string
}

type FileReport struct {
	Path        string
	Name        string
	Description string
	Size        int64
	Error       string
}

type CommandReport struct {
	Script      string
	Name        string
	Description string
	Output      string
	Error       string
}

type Report struct {
	SpecId         string
	GeneratedAt    string
	BundleName     string
	HostName       string
	FolderReports  []FolderReport
	FileReports    []FileReport
	CommandReports []CommandReport
}
