package gotools

import (
	"encoding/json"
	"os"
)

func PrettyPrint(v any) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "\t")
	return enc.Encode(v)
}
