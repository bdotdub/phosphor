package domain

import (
	"errors"
	"sync"
	"time"
)

// NewTrace initialises and returns a new Trace
func NewTrace() *Trace {
	return &Trace{
		frames: make([]*Frame, 0),
	}
}

// Trace represents a full trace of a request
// comprised of a number of frames
type Trace struct {
	sync.Mutex

	frames []*Frame
}

// AppendFrame to a Trace
func (t *Trace) AppendFrame(f *Frame) error {
	if t == nil {
		return errors.New("Slice is Nil")
	}

	t.frames = append(t.frames, f)

	return nil
}

// FrameType represents an Enum of types of Frames which Phosphor can record
type FrameType int32

const (
	UnknownFrameType = FrameType(0) // No idea...

	// Calls
	Req     = FrameType(1) // Client Request dispatch
	Rsp     = FrameType(2) // Client Response received
	In      = FrameType(3) // Server Request received
	Out     = FrameType(4) // Server Response dispatched
	Timeout = FrameType(5) // Client timed out waiting

	// Developer initiated annotations
	Annotation = FrameType(6)
)

// A Frame represents the smallest individually fired component of a trace
// These can be assembled into spans, and entire traces of a request to our systems
type Frame struct {
	TraceId      string // Global Trace Identifier
	SpanId       string // Identifier for this span, non unique - eg. RPC calls would have 4 frames with this id
	ParentSpanId string // Parent span - eg. nested RPC calls

	Timestamp time.Time     // Timestamp the event occured, can only be compared on the same machine
	Duration  time.Duration // Optional: duration of the event, eg. RPC call

	Hostname    string // Hostname this event originated from
	Origin      string // Fully qualified name of the message origin
	Destination string // Optional: Fully qualified name of the message destination

	FrameType FrameType // The type of Frame

	Payload     string            // The payload, eg. RPC body, or Annotation
	PayloadSize int32             // Bytes of payload
	KeyValue    map[string]string // Key value debug information
}
