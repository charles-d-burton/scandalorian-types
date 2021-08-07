package scandaloriantypes

//ScanRequest object instructing system on how to scan.
type ScanRequest struct {
	Address string `json:"address,omitempty"`
	Host    string `json:"host,omitempty"`
	PPS     int    `json:"pps,omitempty"`

	ScanTypes []string    `json:"scan_types,omitempty"`
	Options   ScanOptions `json:"scan_options,omitempty"`
}

//Scan structure to send to message queue for scanning
type Scan struct {
	IP        string   `json:"ip"`
	ScanID    string   `json:"scan_id"`
	RequestID string   `json:"request_id"`
	Ports     []string `json:"ports,omitempty"`
	Stream    string   `json:"-"`
}

//ScanOptions optional parameters to set for a scan
type ScanOptions struct {
	TopTen      bool                   `json:"top_ten,omitempty"`
	TopHundred  bool                   `json:"top_hundred,omitempty"`
	TopThousand bool                   `json:"top_thousand,omitempty"`
	PPS         int                    `json:"pps,omitempty"` //Set rate limiter value
	AppOptions  ApplicationScanOptions `json:"application_scan_options,omitempty"`
}

type ApplicationScanOptions struct {
	Flags       BaseFlags `json:"flags"`
	UDPFlags    UDPFlags  `json:"udp_flags,omitempty"`
	HttpOptions Http      `json:"http_options,omitempty"`
}
