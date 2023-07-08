package models

type Folder struct {
	Path        string
	Name        string
	Description string
}

type File struct {
	Path          string
	Name          string
	Description   string
	EnableWatcher bool
}

type Command struct {
	Script      string
	Name        string
	Description string
}

type Spec struct {
	Id          string
	Name        string
	Version     string
	Description string
	CronString  string
	Folders     []Folder
	Files       []File
	Commands    []Command
}
