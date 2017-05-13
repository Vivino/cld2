//+build cld2_disable !cgo

package cld2

// Detect returns the language code for detected language
// in the given text.
func init () {
	Detect = func(text string) string { return UNKNOWN_LANGUAGE.Code() }
	DetectLang = func(text string) Language { return UNKNOWN_LANGUAGE }
	DetectThree = func(text string) Languages { return nil }
}