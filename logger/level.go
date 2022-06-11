package logger

const (
	LevelDebug Level = "DEBUG"
	LevelInfo        = "INFO"
	LevelError       = "ERROR"
	LevelNoOp        = "NOOP"

	DefaultLevel = LevelInfo
)

type Level string

func (l defaultLogger) shouldLog(msgLevel Level) bool {
	switch l.level {
	case LevelNoOp:
		return false
	case LevelError:
		return msgLevel == LevelError
	case LevelInfo:
		return msgLevel == LevelInfo || msgLevel == LevelError
	case LevelDebug:
		fallthrough
	default:
		return true
	}
}

func (l *defaultLogger) SetLevel(level Level) {
	l.level = level
}
