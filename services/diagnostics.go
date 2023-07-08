package services

import (
	"shak-daemon/actions"
	httpclient "shak-daemon/httpClient"
	"shak-daemon/models"
	"shak-daemon/utils"
	"time"
)

type Diagnostics struct {
	Spec   models.Spec
	Report models.Report
}

func NewDiagnosticsService() Diagnostics {
	spec := models.Spec{}
	httpclient.GetLatestSpec("./sampleconfig.json", &spec)
	report := models.Report{
		SpecId:      spec.Id,
		GeneratedAt: time.Now().Format(time.RFC3339),
		BundleName:  utils.GetBundleName(),
		HostName:    utils.GetHostName(),
	}
	diagnostics := Diagnostics{
		Spec:   spec,
		Report: report,
	}
	return diagnostics
}

func (d *Diagnostics) Process() {
	actions.InspectFolderAction(&d.Spec, &d.Report)
	actions.InspectFileAction(&d.Spec, &d.Report)
	actions.RunCommandAction(&d.Spec, &d.Report)
	actions.CreateArchiveAction(d.Report.BundleName)
	//TODO:create an http client to push the bundles and reports to the server
	actions.CleanUpAction(d.Report.BundleName)
}
