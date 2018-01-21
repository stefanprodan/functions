package function

import (
	"fmt"
	"os"
)

func Handle(req []byte) []byte {
	response := fmt.Sprintf("Echo:\n%s\nEnvironment:\n", string(req))
	for _, e := range os.Environ() {
		response += fmt.Sprintf("%s\n", e)
	}

	return []byte(response)
}
