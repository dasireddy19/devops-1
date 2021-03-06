package models

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

// DebugVirtualServiceCapture debug virtual service capture
// swagger:model DebugVirtualServiceCapture
type DebugVirtualServiceCapture struct {

	// Number of minutes to capture packets. Use 0 to capture until manually stopped. Special values are 0 - 'infinite'.
	Duration *int32 `json:"duration,omitempty"`

	// Enable SSL session key capture. Field introduced in 18.2.3.
	EnableSslSessionKeyCapture *bool `json:"enable_ssl_session_key_capture,omitempty"`

	// Total number of packets to capture.
	NumPkts *int32 `json:"num_pkts,omitempty"`

	// Enable PcapNg for packet capture. Field introduced in 18.2.5.
	PcapNg *bool `json:"pcap_ng,omitempty"`

	// Number of bytes of each packet to capture. Use 0 to capture the entire packet. Allowed values are 64-1514. Special values are 0 - 'full capture'.
	PktSize *int32 `json:"pkt_size,omitempty"`
}
