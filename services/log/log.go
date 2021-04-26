package log

import (
	"backend-services/services/constant"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

func New(appName, logDir string) (*logrus.Logger, *os.File) {
	var file *os.File

	filename := fmt.Sprintf("%s_%s.log", appName, time.Now().Format("20060102"))

	dir := logDir
	if strings.HasSuffix(dir, "/") {
		strings.TrimSuffix(dir, "/")
	}

	instance := logrus.New()
	if f, err := os.OpenFile(dir+filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755); err == nil {
		o := io.MultiWriter(os.Stdout, f)
		instance.SetOutput(o)
	} else {
		instance.Warnln(constant.ComposeMessage(constant.LogOpenFileFailed, err.Error()))
	}

	return instance, file
}
