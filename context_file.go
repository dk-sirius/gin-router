package xrouter

type AttachFile struct {
	file string
	data []byte
}

func NewAttachFile(file string, data []byte) *AttachFile {
	return &AttachFile{
		file: file,
		data: data,
	}
}

func (a *AttachFile) GetHeaderValue() string {
	return "attachment; filename=" + a.file
}

