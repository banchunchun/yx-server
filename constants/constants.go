package constants

var (
	TRANSCODER_STARTED       = "__STARTED__%__"
	TRANSCODER_RUNNING       = "__TRANSCODER_RUNNING__"
	TRANSCODER_COMPLETE      = "transcoding(pass0) finished"
	TRANSCODER_ERROR_WARNING = "__ERROR__$"
	TRANSCODER_HANG_ERROR    = "hang about 120 seconds"
	TRAMSCPDER_PROCESS       = "finished, speed"
	STATUS_RUNNING           = "RUNNING"
	STATUS_ERROR             = "ERROR"
	STATUS_OTHER             = "OTHER"
	STATUS_SUCCESS           = "SUCCESS"
	// yx redis key
	YX_KEY_PREFIX = "yx:task:%d"
)
