package format

import "github.com/sirupsen/logrus"

type PlainFormatter struct{}

func (PlainFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(entry.Message), nil
}
