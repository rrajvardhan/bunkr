package shared

type Route string

const (
	None      Route = ""
	Home      Route = "Home"
	Host      Route = "host"
	Connect   Route = "connect"
	Client    Route = "client"
	About     Route = "about"
	Dashboard Route = "dashboard"
	Files     Route = "manage files"
	Users     Route = "connected users"
	Upload    Route = "upload files"
)
