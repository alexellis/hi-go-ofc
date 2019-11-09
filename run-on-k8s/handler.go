package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) string {
	return fmt.Sprintf("I'm running on K8s and have full e2e TLS from you to my Pod %s", string(req))
}
