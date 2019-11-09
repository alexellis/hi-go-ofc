package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) string {
	return fmt.Sprintf("e2e tls on K8s, great eh: %s", string(req))
}
