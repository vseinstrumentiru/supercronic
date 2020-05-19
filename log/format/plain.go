package format

import (
	"github.com/sirupsen/logrus"
	"strings"
)

type PlainFormatter struct{}

func (PlainFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	msg := strings.TrimSuffix(entry.Message, "\n")

	if msg == "starting" || msg == "job succeeded" {
		return []byte{}, nil
	}

	return []byte(msg + "\n"), nil
}
