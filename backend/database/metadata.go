package database

// how metadata is stored in datebase
// integer UnixTimestamp represents time in unix format
type Metadata struct {
	ID            int64
	Filename      string
	Mime          string
	Description   string
	Uploader      string
	UnixTimestamp int64
}
