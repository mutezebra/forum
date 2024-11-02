package logger

func init() {
	initLogger()
	go updateLogger()
}

func Fatalf(template string, args ...interface{}) {
	loggerObj.Fatal(template, args)
}

func Errorf(template string, args ...interface{}) {
	loggerObj.Errorf(template, args...)
}

func Infof(template string, args ...interface{}) {
	loggerObj.Infof(template, args...)
}

func Debugf(template string, args ...interface{}) {
	loggerObj.Debugf(template, args...)
}

func Fatal(args ...interface{}) {
	loggerObj.Fatal(args)
}

func Info(args ...interface{}) {
	loggerObj.Info(args)
}

func Error(args ...interface{}) {
	loggerObj.Error(args)
}

func Debug(args ...interface{}) {
	loggerObj.Debug(args)
}

func KDebugf(template string, args ...interface{}) {
	kloggerObj.Debugf(template, args...)
}

func KDebug(args ...interface{}) {
	kloggerObj.Debug(args)
}

func KErrorf(template string, args ...interface{}) {
	kloggerObj.Errorf(template, args)
}

func KError(args ...interface{}) {
	kloggerObj.Error(args)
}
