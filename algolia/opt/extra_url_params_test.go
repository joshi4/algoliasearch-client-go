package opt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractExtraURLParams(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected map[string]string
	}{
		{
			opts:     []interface{}{nil},
			expected: map[string]string{},
		},
		{
			opts:     []interface{}{ExtraURLParams(map[string]string{})},
			expected: map[string]string{},
		},
		{
			opts: []interface{}{
				map[string]string{"key1": "value1", "key2": "value2"},
			},
			expected: map[string]string{},
		},
		{
			opts: []interface{}{
				ExtraURLParams(map[string]string{"key1": "value1", "key2": "value2"}),
				map[string]string{"key5": "value5"},
				ExtraURLParams(map[string]string{"key3": "value3", "key4": "value4"}),
			},
			expected: map[string]string{"key1": "value1", "key2": "value2", "key3": "value3", "key4": "value4"},
		},
	} {
		res := ExtractExtraURLParams(c.opts...)
		require.EqualValues(t, c.expected, res)
	}
}