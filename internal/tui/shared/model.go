package shared

type ServerState struct {
	Password string
	Running  bool
}

type Info struct {
	Width  int
	Height int
}

var Term = Info{
	Width:  80,
	Height: 24,
}

var SState = ServerState{
	Password: "",
	Running:  false,
}
