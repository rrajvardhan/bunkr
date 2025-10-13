package shared

type View string

const (
	None      View = ""
	Home      View = "Home"
	Host      View = "host"
	Connect   View = "connect"
	Client    View = "client"
	About     View = "about"
	Dashboard View = "dashboard"
	Files     View = "manage files"
	Upload    View = "upload files"
)
