package tasks

import (
	"github.com/sirupsen/logrus"
	"learn-go/web/core/jobs"
	"learn-go/web/core/starter"
	"learn-go/web/domain/envelopes"
	"time"
)

func RegisterTasks() {
	scheduler := jobs.NewScheduler(10)
	err := scheduler.Start(time.Now()).JobType(jobs.RecurrenceTask).Hour(1).Do("expiredRedEnvelopeRefund", func() {
		envelopeService := envelopes.GetEnvelopeService(starter.DefaultDB())
		expiredRedEnvelopes := envelopeService.FindAllExpiredRedEnvelope()
		for _, redEnvelope := range expiredRedEnvelopes {
			go func() {
				_, err := envelopeService.Refund(redEnvelope.EnvelopeNo)
				if err != nil {
					logrus.Error(err)
				}
			}()
		}
	})
	if err != nil {
		logrus.Error(err)
	}
	scheduler.Run()
}
