package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoughRequest(t *testing.T) {
	// Initialize struct
	request := DoughRequest{
		DoughBallAmount: 4,
		Hydration:       70,
		DoughBallWeight: 270,
	}

	// Marshal struct to JSON
	jsonData, err := json.Marshal(request)
	if err != nil {
		t.Errorf("Failed to marshal struct to JSON: %v", err)
	}

	// Assert the JSON string
	assert.Equal(t, `{"doughBallAmount":4,"hydration":70,"doughBallWeight":270}`, string(jsonData))

	// Unmarshal JSON to struct
	var unmarshalledRequest DoughRequest
	err = json.Unmarshal(jsonData, &unmarshalledRequest)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON to struct: %v", err)
	}

	// Assert the unmarshalled struct
	assert.Equal(t, request, unmarshalledRequest)
}
