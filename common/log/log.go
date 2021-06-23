package log

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/zput/zxcTool/ztLog/zt_formatter"
	"io"
	"path"
	"runtime"
	"time"
)

type Options struct {
	FileName     string
	MaxAge       int    // 日志文件保留时长(天)
	RotationTime int    // 日志文件切割时间(小时)
	Release      bool   // release版
	Level        string // debug err info
}

type Option func(*Options)

func defaultOption() Options {
	return Options{
		Release: false,
		Level:   "debug",
	}
}

func InitLog(opts ...Option) (err error) {
	options := defaultOption()
	for _, o := range opts {
		o(&options)
	}

	setFormatter()

	if options.FileName != "" {
		writer := getWriter(options.FileName, options.MaxAge, options.RotationTime)
		logrus.AddHook(getLogHook(writer))
	}

	if options.Release {
		logrus.SetLevel(logrus.ErrorLevel)
		return
	}

	switch options.Level {
	case "err":
		logrus.SetLevel(logrus.ErrorLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	}

	return nil
}

func setFormatter() {
	var ztFormatter = &zt_formatter.ZtFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
		Formatter: nested.Formatter{
			//HideKeys: true,
			FieldsOrder: []string{"component", "category"},
		},
	}
	logrus.SetReportCaller(true)
	logrus.SetFormatter(ztFormatter)
}

// 将日志文件按配置的时间分割,并自动清理过期日志
func GetWriter(opts ...Option) io.Writer {
	options := defaultOption()
	for _, o := range opts {
		o(&options)
	}
	return getWriter(options.FileName, options.MaxAge, options.RotationTime)
}

func getWriter(fileName string, maxAge int, rotationTime int) io.Writer {

	// 生成的日志格式: app_SomeMachine_2000-01-01-00-00.log
	//hostname, _ := os.Hostname()
	//dir, _ := path.Split(fileName)
	//base := path.Base(fileName)
	//ext := path.Ext(fileName)
	//name := strings.TrimSuffix(base, ext)
	//fileName = fmt.Sprintf("%s%s_%s", dir, name, hostname)

	logWriter, _ := rotatelogs.New(
		//fmt.Sprintf("%s_%%Y-%%m-%%d-%%H-%%M%s", fileName, ext),
		fileName,
		rotatelogs.WithMaxAge(time.Duration(maxAge)*24*time.Hour),          // 文件最大保存天数
		rotatelogs.WithRotationTime(time.Duration(rotationTime)*time.Hour), // 日志切割时间间隔小时
	)
	return logWriter
}

func getLogHook(writer io.Writer) *lfshook.LfsHook {
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  writer,
		logrus.FatalLevel: writer,
		logrus.DebugLevel: writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.PanicLevel: writer,
	}
	return lfshook.NewHook(writeMap, &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
}
