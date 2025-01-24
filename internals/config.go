package internals

import (
	"allcaps/api/controller"
	"allcaps/api/mergehandler"
	"net/http"

	"github.com/robfig/cron"
)

var (
	APItoken      = "f3bCr4SAWwRw6bI9yHhgmgOfTlGLfPdDNnDpfMWIGFqA3ZIbcbj9Bg"
	ProductionKey = "ggVZQsaslSsuVbmfxNUM32J2YPWLAod57_GcDmKS06r7YQfCPChi5A"
	AccountToken  = "fc58ff09-85ab-401c-b46b-ed51e9bcde0b"
	Connection    = &http.Client{}
)

func InitCronJob() *cron.Cron {
	c := cron.New()
	c.AddFunc("@every 00h00m10s", mergehandler.NewMessage)
	c.AddFunc("@every 00h00m10s", controller.SyncAccount)
	c.Start()
	return c
}
