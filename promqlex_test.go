package promqlex

import (
	"errors"
	"gopkg.in/yaml.v3"
	"io"
	"strings"
	"testing"
)

func TestPromQLExParser_ValidPromQLIsValidPromQLEx(t *testing.T) {
	yamlDecoder := yaml.NewDecoder(strings.NewReader(promqlExamplesYaml))
	for {
		var example Examples
		err := yamlDecoder.Decode(&example)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		t.Run(example.Source, func(t *testing.T) {
			for _, query := range example.Queries {
				t.Run(query, func(t *testing.T) {
					//if strings.HasPrefix(query, "sum by") {
					//	runtime.Breakpoint()
					//}
					_, _, _ = parsePromQLEx(t, query)
				})
			}
		})
	}
}
