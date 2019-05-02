package api2go

import (
	"time"
)

// The Etagger interface can be optionally implemented to directly return a
// hash used for cache control.
//
type Etagger interface {
	GetEtag() (string, error)
}

// the optional Timestamper interface returns the timestamp of the resource
// used for cache control.
//
type Timestamper interface {
	GetTimestamp() (time.Time, error)
}
