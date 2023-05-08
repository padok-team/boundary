// Code generated by "make api"; DO NOT EDIT.
package sessionrecordings

import (
	"time"
)

type ConnectionRecording struct {
	Id                string              `json:"id,omitempty"`
	ConnectionId      string              `json:"connection_id,omitempty"`
	BytesUp           uint64              `json:"bytes_up,omitempty"`
	BytesDown         uint64              `json:"bytes_down,omitempty"`
	StartTime         time.Time           `json:"start_time,omitempty"`
	EndTime           time.Time           `json:"end_time,omitempty"`
	Duration          time.Duration       `json:"duration,omitempty"`
	MimeTypes         []string            `json:"mime_types,omitempty"`
	ChannelRecordings []*ChannelRecording `json:"channel_recordings,omitempty"`
}
