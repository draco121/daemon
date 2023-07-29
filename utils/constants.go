package utils

import "path"

var CronString = "{{cronstring}}"
var RootDir = GetRootDir()

const AppId = "{{appid}}"
const BaseUrl = "{{baseurl}}"
const UserName = "super.user"
const Password = "Pa$$w0rd123!"

var Token = ""
var SpecDir = path.Join(RootDir, "spec.json")
var SpecSourceDir = "./sampleconfig.json"
