package sshconnect

// SSH connect target
type Target struct {
	IP       string // remote ip address
	Port     int    // remote port
	Username string // username
	Password string // password
	Result   string // sync result
}
