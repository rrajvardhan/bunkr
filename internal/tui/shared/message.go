package shared

type NavigateMsg struct {
	Target View
}

type SetPasswordMsg struct {
	Password string
}

type SetServerMsg struct {
	Url      string
	Password string
}

type ResetTo struct {
	Target View
}

type UploadSuccessMsg struct {
	FileName string
}

type UploadErrorMsg struct {
	Err error
}

type AuthErrorMsg struct {
	Err error
}
