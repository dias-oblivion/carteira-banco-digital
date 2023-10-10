package config

var (
	// db
	logger *Logger
)

func Init() error {
	return nil
}

func GetLogger(p string) *Logger {
	logger = InitLogger(p)
	return logger
}
