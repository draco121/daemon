package httpclient

import (
	"encoding/json"
	"fmt"
	"log"
	"shak-daemon/models"
	"shak-daemon/utils"

	"github.com/ddliu/go-httpclient"
)

func GetNewHttpClient() *httpclient.HttpClient {
	c := httpclient.Defaults(httpclient.Map{
		httpclient.OPT_USERAGENT: "shak-daemon",
		"Accept-Language":        "en-us",
	})
	return c
}

func Login() {
	client := GetNewHttpClient()
	res, err := client.Post(fmt.Sprintf("%s/auth/login", utils.BaseUrl), map[string]string{
		"email":    utils.UserName,
		"password": utils.Password,
	})
	if err != nil {
		fmt.Print(err)
	}
	token, err := res.ToString()
	if err != nil {
		log.Fatal(err)
	}
	utils.Token = token
}

func GetLatestSpec(spec *models.Spec) {

	client := GetNewHttpClient()
	if utils.Token == "" {
		Login()
	}
	res, err := client.WithHeader("Authorization", utils.Token).Get(fmt.Sprintf("%s/spec?appId=%s", utils.BaseUrl, utils.AppId))
	if err != nil {
		log.Fatal(err)
	}
	specId, err := res.ToString()
	specId = specId[1 : len(specId)-1]
	if err != nil {
		log.Fatal(err)
	}
	res, err = client.WithHeader("Authorization", utils.Token).Get(fmt.Sprintf("%s/spec/%s", utils.BaseUrl, specId))
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err := res.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bodyBytes, &spec)
	if err != nil {
		log.Fatal(err)
	}
}

func SendReport(report *models.Report) {
	client := GetNewHttpClient()
	if utils.Token == "" {
		Login()
	}
	res, err := client.WithHeader("Authorization", utils.Token).PutJson(fmt.Sprintf("%s/report", utils.BaseUrl), report)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode == 201 {
		fmt.Println("report submitted successfully.....")
	}
}

func UploadBundle(bundlePath string) {
	client := GetNewHttpClient()
	if utils.Token == "" {
		Login()
	}
	res, err := client.WithHeader("Authorization", utils.Token).Post(fmt.Sprintf("%s/bundle", utils.BaseUrl), map[string]string{
		"@file": bundlePath,
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode == 201 {
		fmt.Println("report submitted successfully.....")
	}
}
