package database

// how metadata is stored in datebase
type Metadata struct {
	ID            int64
	Filename      string
	Mime          string
	Description   string
	Uploader      string
	UnixTimestamp int64
}
