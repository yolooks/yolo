package config

var CONFIG_DEFINE_TPL = `package config

// sign
const (
	SIGN_SALT  = "@mesh"
	SIGN_ERROR = "sign error"
)

// file
const (
	FILE_READ_ERROR   = "read file error: %v"
	FILE_WRITE_ERROR  = "write file error: %v"
	JSON_DECODE_ERROR = "json decode failed: %v"
)
`
