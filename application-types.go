package scandaloriantypes

// BaseFlags contains the options that every flags type must embed
type BaseFlags struct {
	Port           uint   `json:"port"`
	Name           string `json:"name"`
	Timeout        int    `json:"timeout"`
	Trigger        string `json:"trigger"`
	BytesReadLimit int    `json:"byte_read_limit"`
}

type UDPFlags struct {
	LocalPort    uint   `json:"local_port"`
	LocalAddress string `json:"local_address"`
}

type Bacnet struct {
}

type Banner struct {
}

type Dnp3 struct {
}

type Fox struct {
}

type Http struct {
}

type Imap struct {
}

type Ipp struct {
}

type Jarm struct {
}

type Modbus struct {
}

type Mongo struct {
}

type MSSQL struct {
}

type Ntp struct {
}

type Oracle struct {
}

type Pop3 struct {
}

type Postgres struct {
}

type Redis struct {
}

type Siemens struct {
}

type SMB struct {
}

type SMTP struct {
}

type SSH struct {
}

type Telnet struct {
}

type TLS struct {
}
