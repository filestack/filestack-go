package url

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParamValueToString(t *testing.T) {

	t.Run("string", func(t *testing.T) {
		assert.Equal(t, "abc", paramValueToString("abc"))
	})

	t.Run("int", func(t *testing.T) {
		assert.Equal(t, "123", paramValueToString(123))
	})

	t.Run("float32", func(t *testing.T) {
		assert.Equal(t, "1.23", paramValueToString(float32(1.23)))
	})

	t.Run("bool", func(t *testing.T) {
		assert.Equal(t, "true", paramValueToString(true))
		assert.Equal(t, "false", paramValueToString(false))
	})

	t.Run("[]int", func(t *testing.T) {
		assert.Equal(t, "[1,2,3]", paramValueToString([]int{1, 2, 3}))
	})

}

func TestTask_AsString(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {
		task := NewTask("store", map[string]interface{}{})
		expected := "store"
		assert.Equal(t, expected, task.AsString())
	})

	t.Run("resize", func(t *testing.T) {
		task := NewTask("resize", map[string]interface{}{
			"width":  "640",
			"height": "480",
		})
		expected := "resize=height:480,width:640"
		assert.Equal(t, expected, task.AsString())
	})

	t.Run("shadow", func(t *testing.T) {
		task := NewTask("shadow", map[string]interface{}{
			"opacity": "50",
			"blur":    "5",
			"vector":  []string{"5", "10"},
		})
		expected := "shadow=blur:5,opacity:50,vector:[5,10]"
		assert.Equal(t, expected, task.AsString())
	})

	t.Run("shadow", func(t *testing.T) {
		task := NewTask("shadow", map[string]interface{}{
			"opacity": "50",
			"blur":    "5",
			"vector":  []string{"5", "10"},
		})
		expected := "shadow=blur:5,opacity:50,vector:[5,10]"
		assert.Equal(t, expected, task.AsString())
	})

	t.Run("partialPixelate", func(t *testing.T) {
		task := NewTask("partialPixelate", map[string]interface{}{
			"objects": [][]int{
				[]int{1, 2, 3, 4},
				[]int{5, 6, 7, 8},
			},
		})
		expected := "partialPixelate=objects:[[1,2,3,4],[5,6,7,8]]"
		assert.Equal(t, expected, task.AsString())
	})
}
