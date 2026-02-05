package servicesconstant

import "time"

type DateFormats = string

const (
	TIME_LAYOUT_ISO8601 DateFormats = "2006-01-02T15:04:05Z"
	TIME_LAYOUT_RFC3339 DateFormats = "2006-01-02T15:04:05Z07:00"

	TIME_LAYOUT_DATETIME_FORMAT_W_MIL DateFormats = "2006-01-02T15:04:05.000-07"
	TIME_LAYOUT_DATETIME_FORMAT       DateFormats = "2006-01-02T15:04:05"
	TIME_LAYOUT_DATETIME_FORMAT_NO_T  DateFormats = "2006-01-02 15:04:05"
	TIME_LAYOUT_DATE_FORMAT           DateFormats = "2006-01-02"
	TIME_LAYOUT_TIME_FORMAT           DateFormats = "15:04:05"
	TIME_LAYOUT_DATETIME_FORMAT_BO    DateFormats = "02/01/2006 15:04:05"
	TIME_LAYOUT_DATE_FORMAT_BO        DateFormats = "02/01/2006"
)
const (
	// Default durations
	DefaultHTTPTimeout        = 30 * time.Second
	DefaultMaxRetries         = 3
	DefaultRetryDelay         = 500 * time.Millisecond
	DefaultShutdownTimeout    = 5 * time.Second
	DefaultServerReadTimeout  = 15 * time.Second
	DefaultServerWriteTimeout = 30 * time.Second
)
