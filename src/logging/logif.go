package logging

func (lg Logging) IfErrError(msg string, fields F) {
	for key, val := range fields {
		if (key == "error" || key == "err") && val != nil {
			lg.Error(msg, fields)
		}
	}
}

func (lg Logging) IfErrWarn(msg string, fields F) {
	for key, val := range fields {
		if (key == "error" || key == "err") && val != nil {
			lg.Warn(msg, fields)
		}
	}
}

func (lg Logging) IfErrFatal(msg string, fields F) {
	for key, val := range fields {
		if (key == "error" || key == "err") && val != nil {
			lg.Fatal(msg, fields)
		}
	}
}
