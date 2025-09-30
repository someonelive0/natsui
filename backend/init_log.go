package backend

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	nested_formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	DEFAULT_ROTATE_LOGFILES = 7
	DEFAULT_ROTATE_MBYTES   = 10
	MAX_ROTATE_LOGFILES     = 70
	MAX_ROTATE_MBYTES       = 100
)

// logfile is logfilename such as myserver.log
// default rotate 7 files with 10M per file. rotate_mbytes is MBytes.
func InitLogRotate(logpath, logfile, level string,
	rotate_files, rotate_mbytes uint) error {

	logrus.SetFormatter(&nested_formatter.Formatter{
		HideKeys:        true,
		TimestampFormat: "01-02 15:04:05", //time.DateTime, time.RFC3339,
		// FieldsOrder:     []string{"model", "file"},
		CallerFirst: true,
		CustomCallerFormatter: func(f *runtime.Frame) string {
			s := strings.Split(f.Function, ".")
			funcName := s[len(s)-1]
			return fmt.Sprintf(" [%s:%d %s]", path.Base(f.File), f.Line, funcName)
		},
	})
	//logrus.SetFormatter(&logrus.TextFormatter{DisableColors: true, FullTimestamp: true})
	//logrus.SetFormatter(&logrus.JSONFormatter{})

	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.Errorf("invalid log level '%s'", level)
		return err
	}
	logrus.SetLevel(lvl)
	logrus.SetReportCaller(true) // 设置在输出日志中添加文件名和方法信息

	// create logpath if not exist
	if len(logpath) > 0 {
		if _, err = os.Stat(logpath); err != nil {
			if err = os.Mkdir(logpath, 0755); err != nil {
				logrus.Errorf("create log subdir '%s' failed: %s", logpath, err)
				return err
			}
		}
	} else {
		logpath = "."
	}

	// set log file how to rotate
	if rotate_files > MAX_ROTATE_LOGFILES {
		logrus.Warnf("rotate_files %d is bigger than %d, set to %d",
			rotate_files, MAX_ROTATE_LOGFILES, MAX_ROTATE_LOGFILES)
		rotate_files = MAX_ROTATE_LOGFILES
	}
	if rotate_mbytes > MAX_ROTATE_MBYTES {
		logrus.Warnf("rotate_mbytes %dM is bigger than %dM, set to %dM",
			rotate_mbytes, MAX_ROTATE_MBYTES, MAX_ROTATE_MBYTES)
		rotate_mbytes = MAX_ROTATE_MBYTES
	}

	// don't use rotatelogs "github.com/lestrrat-go/file-rotatelogs" any more
	// because it has bugs, and not update any more
	lumberjackLogger := &lumberjack.Logger{
		// Log file abbsolute path, os agnostic
		Filename:   filepath.ToSlash(logpath + "/" + logfile),
		MaxSize:    int(rotate_mbytes), // MB
		MaxBackups: int(rotate_files),
		// MaxAge:     30,   // days
		LocalTime: true,
		Compress:  false, // disabled by default
	}
	logrus.SetOutput(io.MultiWriter(lumberjackLogger, os.Stdout))

	return nil
}
