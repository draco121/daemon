package utils

import "path"

var CronString = "*/5 * * * *"
var RootDir = GetRootDir()

const AppId = "64c3d38161c84e8b8005bdd9"
const BaseUrl = "http://localhost:4200"
const UserName = "super.user"
const Password = "Pa$$w0rd123!"

var Token = ""
var SpecDir = path.Join(RootDir, "spec.json")
var SpecSourceDir = "./sampleconfig.json"
