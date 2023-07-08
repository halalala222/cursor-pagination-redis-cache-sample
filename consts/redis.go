package consts

import "time"

const (
	DefaultSampleRedisTTL = 60 * 60 * 5 * time.Minute
	BasicPrefix           = "sample"
	SampleZSetPrefix      = "sample-ZSet"
	SamplePrefix          = "sample-string"
)
