package logging

func (lg Logging) Debug(msg string, fields interface{}) {
	lg.Logrus.WithFields(lg.conv(fields)).Debug(msg)
}

func (lg Logging) Info(msg string, fields interface{}) {
	lg.Logrus.WithFields(lg.conv(fields)).Info(msg)
}

func (lg Logging) Warn(msg string, fields interface{}) {
	lg.Logrus.WithFields(lg.conv(fields)).Warn(msg)
}

func (lg Logging) Error(msg interface{}, fields interface{}) {
	lg.Logrus.WithFields(lg.conv(fields)).Error(msg)
}

func (lg Logging) Fatal(msg string, fields interface{}) {
	lg.Logrus.WithFields(lg.conv(fields)).Fatal(msg)
}
