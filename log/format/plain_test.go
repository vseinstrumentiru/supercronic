package format

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type testWriter struct {
	c chan []byte
}

func (w testWriter) Write(p []byte) (int, error) {
	w.c <- p
	return len(p), nil
}

func TestPlainFormatter_Format(t *testing.T) {
	defaultWriter := testWriter{c: make(chan []byte, 1)}

	log := logrus.New()
	log.SetOutput(defaultWriter)
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&PlainFormatter{})

	log.Error("some plain text message")

	select {
	case log := <-defaultWriter.c:
		assert.Equal(t, "some plain text message\n", string(log))
	case <-time.After(time.Second):
		t.Fatalf("timed out waiting for out log")
	}

	log.Debug("some ", "plain ", "text ", "message")

	select {
	case log := <-defaultWriter.c:
		assert.Equal(t, "some plain text message\n", string(log))
	case <-time.After(time.Second):
		t.Fatalf("timed out waiting for out log")
	}

	b, _ := json.Marshal(struct {
		First  string   `json:"first"`
		Second int      `json:"second"`
		Third  []string `json:"third"`
	}{"val1", 2, []string{"3.1", "3.2"}})

	log.Info(string(b))
	select {
	case log := <-defaultWriter.c:
		assert.Equal(t, `{"first":"val1","second":2,"third":["3.1","3.2"]}`+"\n", string(log))
	case <-time.After(time.Second):
		t.Fatalf("timed out waiting for out log")
	}
}
