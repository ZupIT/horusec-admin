package logger

import log "github.com/sirupsen/logrus"

type Jaeger struct {
	*log.Entry
}

func NewJaeger() *Jaeger {
	return &Jaeger{Entry: WithPrefix("jaeger")}
}

func (j *Jaeger) Error(msg string) {
	j.Entry.Error(msg)
}

func (j *Jaeger) Infof(msg string, args ...interface{}) {
	j.Entry.Debugf(msg, args...)
}
