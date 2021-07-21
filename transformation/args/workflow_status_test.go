package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWorkflowStatusArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {

		mapExpected := map[string]interface{}{}
		expected, _ := json.Marshal(mapExpected)

		args := NewWorkflowStatusArgs()

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("with options", func(t *testing.T) {

		jobID := "10"
		mapExpected := map[string]interface{}{
			"job_id": jobID,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewWorkflowStatusArgs().SetJobID(jobID)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

}
