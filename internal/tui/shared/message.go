package shared

type NavigateMsg struct {
	Target Route
}

type SetPasswordMsg struct {
	Password string
}

type ResetTo struct {
	Target Route
}

type UploadSuccessMsg struct {
	FileName string
}

type UploadErrorMsg struct {
	Err error
}
