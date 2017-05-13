// Package cld2 implements language detection using the
// Compact Language Detector.
//
// This package includes the relevant sources from the cld2
// project, so it doesn't require any external dependencies.
// For more information about CLD2, see https://github.com/CLD2Owners/cld2
package cld2

import "github.com/klauspost/cld2/internal/info"

// Detect returns the language code for detected language
// in the given text.
var Detect func (text string) string

// DetectLang returns the language code for detected language
// in the given text.
// ENGLISH is returned if the language cannot be detected.
var DetectLang func(text string) Language

// DetectThree returns up to three language guesses.
// Extended languages are enabled.
// Unknown languages are removed from the resultset.
var DetectThree func(text string) Languages

var Enabled bool

func infoToLanguages(in info.Languages) Languages {
	res :=  Languages {
		TextBytes: in.TextBytes,
		Reliable: in.Reliable,
	}
	for _, est := range in.Estimates {
		if Language(est.Language) == UNKNOWN_LANGUAGE {
			continue
		}
		res.Estimates = append(res.Estimates, Estimate{Language: Language(est.Language), NormScore: est.NormScore, Percent:est.Percent })
	}
	return res
}