package info

// Languages are probable languages of the supplied text
type Languages struct {
	Estimates []Estimate // Possible languages returned in order of confidence
	TextBytes int        // the amount of non-tag/letters-only text found
	Reliable  bool       // Does CLD2 see the result as reliable?
}

// Single Language estimate
type Estimate struct {
	Language uint16
	Percent  int // text percentage 0..100 of the top 3 languages.

	// NormScore is internal language scores as a ratio to normal score for real text in that language.
	// Scores close to 1.0 indicate normal text, while scores far away
	// from 1.0 indicate badly-skewed text or gibberish.
	NormScore float64
}
