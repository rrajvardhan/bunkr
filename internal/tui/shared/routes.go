package shared

type Route string

const (
	None      Route = "none"
	Dashboard Route = "dashboard"
	Host      Route = "host"
	Auth      Route = "auth"
)
