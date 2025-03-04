package respmodel

type NotePathResp struct {
	Value    string         `json:"value"`
	Label    string         `json:"label"`
	Tag      string         `json:"tag"`
	Children []NotePathResp `json:"children"`
}
