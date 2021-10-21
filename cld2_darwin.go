//+build darwin,go1.8
//+build amd64 arm64

package cld2

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"plugin"
	"reflect"
	"strings"
)

//go:generate go build -buildmode=plugin -o lib/cld2go.so github.com/Vivino/cld2/internal/plugin

func init() {
	LoadPlugin = func(paths ...string) error {
		if Enabled {
			return nil
		}
		p, err := findPlugin(paths...)
		if err != nil {
			return err
		}
		return usePlugin(p)
	}
}

func usePlugin(p *plugin.Plugin) error {
	f, err := p.Lookup("PluginDetect")
	if err != nil {
		return err
	}
	var ok bool
	Detect, ok = f.(func(text string) string)
	if !ok {
		return fmt.Errorf("PluginDetect: wrong signature: %T", f)
	}

	f, err = p.Lookup("PluginDetectLang")
	if err != nil {
		return err
	}
	dl, ok := f.(func(text string) uint16)
	if !ok {
		return fmt.Errorf("PluginDetectLang: wrong signature: %T", f)
	}
	DetectLang = func(text string) Language {
		return Language(dl(text))
	}

	f, err = p.Lookup("PluginDetectThree")
	if err != nil {
		return err
	}
	dt, ok := f.(func(text string) interface{})
	if !ok {
		return fmt.Errorf("PluginDetectThree: wrong signature: %T", f)
	}
	DetectThree = func(text string) Languages {
		l, err := parseLanguage(dt(text))
		if err != nil {
			log.Println(err)
			return Languages{}
		}
		return l
	}
	Enabled = true
	return nil
}

func findPlugin(paths ...string) (*plugin.Plugin, error) {
	if len(paths) == 0 {
		return nil, fmt.Errorf("no paths specified for opening cld2go.so")
	}
	var errors bytes.Buffer
	for i, path := range paths {
		abspath, err := getAbsoluteCLD2Path(path)
		if err != nil {
			errors.WriteString(err.Error())
			continue
		}
		p, err := plugin.Open(abspath)
		if err == nil {
			return p, nil
		}
		errors.WriteString(fmt.Sprintf("\n\terror opening file (path: %s), error: (%v)", paths[i], err))
	}
	return nil, fmt.Errorf("could not open cld2go.so: %s", errors.String())
}

func getAbsoluteCLD2Path(path string) (string, error) {
	if !strings.Contains(path, "cld2go.so") {
		abspath, err := filepath.Abs(filepath.Join(path, "cld2go.so"))
		if err != nil {
			return "", fmt.Errorf("\n\tcould not create absolute path for (%s): %v", path, err)
		}
		return abspath, nil
	}
	return path, nil
}

func parseLanguage(v interface{}) (Languages, error) {
	val := reflect.ValueOf(v)

	reliable := val.FieldByName("Reliable")
	if reliable.Kind() != reflect.Bool {
		return Languages{}, NewParseError("Reliable", reflect.Bool, reliable.Kind())
	}

	textBytes := val.FieldByName("TextBytes")
	if textBytes.Kind() != reflect.Int {
		return Languages{}, NewParseError("TextBytes", reflect.Int, textBytes.Kind())
	}

	estimates, err := parseEstimates(val.FieldByName("Estimates"))
	if err != nil {
		return Languages{}, err
	}

	return Languages{
		Estimates: estimates,
		TextBytes: int(textBytes.Int()),
		Reliable:  reliable.Bool(),
	}, nil
}

func parseEstimates(value reflect.Value) ([]Estimate, error) {
	if value.Kind() != reflect.Array && value.Kind() != reflect.Slice {
		return nil, NewParseError("Estimates", reflect.Slice, value.Kind())
	}
	estimates := []Estimate{}
	for i := 0; i < value.Len(); i++ {
		estimate, err := parseEstimate(value.Index(i))
		if err != nil {
			return nil, err
		}
		if estimate.Language != UNKNOWN_LANGUAGE {
			estimates = append(estimates, estimate)
		}
	}
	return estimates, nil
}

func parseEstimate(value reflect.Value) (Estimate, error) {
	language := value.FieldByName("Language")
	if language.Kind() != reflect.Uint16 {
		return Estimate{}, NewParseError("Language", reflect.Uint16, language.Kind())
	}
	percent := value.FieldByName("Percent")
	if percent.Kind() != reflect.Int {
		return Estimate{}, NewParseError("Percent", reflect.Int, percent.Kind())
	}
	normScore := value.FieldByName("NormScore")
	if normScore.Kind() != reflect.Float64 {
		return Estimate{}, NewParseError("NormScore", reflect.Float64, normScore.Kind())
	}
	return Estimate{
		Language:  Language(language.Uint()),
		Percent:   int(percent.Int()),
		NormScore: normScore.Float(),
	}, nil
}

// ErrCLD2ParseError represents an error of unexpected output from cld2
var ErrCLD2ParseError = errors.New("could not parse cld2 output")

// ParseError represents the details of a parsing error
type ParseError struct {
	err     error
	details string
}

// Error returns the details of the parsing error in the form of string
func (err ParseError) Error() string {
	return fmt.Sprintf("%v: %v", err.err, err.details)
}

// Unwrap returns the inner error of the ParseError structure
func (err ParseError) Unwrap() error {
	return err.err
}

// NewParseError instatiates and returns a new ParseError
func NewParseError(property string, expected, actual reflect.Kind) error {
	return ParseError{
		err: ErrCLD2ParseError,
		details: fmt.Sprintf("unexpected type for property %v: (expected: %v) (actual: %v)",
			property, expected, actual),
	}
}
