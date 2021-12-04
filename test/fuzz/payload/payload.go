package payload

import (
	"encoding/json"

	"github.com/sigstore/sigstore/pkg/signature/payload"
)

// FuzzMarshalUnMarshalCosign fuzzing function test cosign payload.
func FuzzMarshalUnMarshalCosign(data []byte) int {
	var imgPayload payload.Cosign
	if err := json.Unmarshal([]byte(data), &imgPayload); err != nil {
		return 0
	}
	if _, err := json.Marshal(imgPayload); err != nil {
		panic(err)
	}
	return 1
}
