package scandaloriantypes

import "strconv"

type Scan interface {
	SetDefaults(scan *Scan)
}

//ScanRequest object instructing system on how to scan.
type ScanRequest struct {
	Address         string           `json:"address,omitempty"`
	Host            string           `json:"host,omitempty"`
	PortScan        *PortScan        `json:"port_scan,omitempty"`
	ApplicationScan *ApplicationScan `json:"application_scan,omitempty"`
}

//Scan structure to send to message queue for scanning
type ScanMetaData struct {
	IP        string `json:"ip"`
	ScanID    string `json:"scan_id"`
	RequestID string `json:"request_id"`
	Stream    string `json:"-"`
}

//Top level object to define a port scan
type PortScan struct {
	ScanMetaData
	Run         bool     `json:"run"`
	PPS         int      `json:"pps,omitempty"` //Set rate limiter value
	Ports       []string `json:"ports,omitempty"`
	TopTen      bool     `json:"top_ten,omitempty"`
	TopHundred  bool     `json:"top_hundred,omitempty"`
	TopThousand bool     `json:"top_thousand,omitempty"`
	//Errors      []string `json:"errors,omitempty"`
}

func (ps *PortScan) SetDefaults(scan *ScanMetaData) {
	ps.IP = scan.IP
	ps.ScanID = scan.ScanID
	ps.RequestID = scan.RequestID
	if ps.PPS == 0 {
		ps.PPS = 4000
	}
	if len(ps.Ports) == 0 {
		/*
			Not Yet Implemented
			if ps.TopTen {

			}
			if ps.TopHundred{}
			if ps.TopThousand{}
		*/
		for i := 0; i <= 65535; i++ {
			ps.Ports = append(ps.Ports, strconv.Itoa(i))
		}
	}
}

//Top Level Object to define application level scanning
type ApplicationScan struct {
	ScanMetaData
	Run         bool      `json:"run"`
	Flags       BaseFlags `json:"flags"`
	UDPFlags    UDPFlags  `json:"udp_flags,omitempty"`
	HttpOptions *Http     `json:"http_options,omitempty"`
}

func (ps *ApplicationScan) SetDefaults(scan *ScanMetaData) {
	ps.IP = scan.IP
	ps.ScanID = scan.ScanID
	ps.RequestID = scan.RequestID
}
