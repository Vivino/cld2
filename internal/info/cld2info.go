package info

type Languages interface {
	GetEstimates() []Estimate
	GetTextBytes() int
	GetReliable() bool
}

type Estimate interface {
	GetLanguage() uint16
	GetPercent() int
	GetNormScore() float64
}

// Languages are probable languages of the supplied text
type LanguagesImpl struct {
	Estimates []Estimate // Possible languages returned in order of confidence
	TextBytes int        // the amount of non-tag/letters-only text found
	Reliable  bool       // Does CLD2 see the result as reliable?
}

func (l LanguagesImpl) GetEstimates() []Estimate {
	return l.Estimates
}

func (l LanguagesImpl) GetTextBytes() int {
	return l.TextBytes
}

func (l LanguagesImpl) GetReliable() bool {
	return l.Reliable
}

var _ Languages = LanguagesImpl{}

// Single Language estimate
type EstimateImpl struct {
	Language uint16
	Percent  int // text percentage 0..100 of the top 3 languages.

	// NormScore is internal language scores as a ratio to normal score for real text in that language.
	// Scores close to 1.0 indicate normal text, while scores far away
	// from 1.0 indicate badly-skewed text or gibberish.
	NormScore float64
}

func (e EstimateImpl) GetLanguage() uint16 {
	return e.Language
}

func (e EstimateImpl) GetPercent() int {
	return e.Percent
}

func (e EstimateImpl) GetNormScore() float64 {
	return e.NormScore
}

var _ Estimate = EstimateImpl{}
