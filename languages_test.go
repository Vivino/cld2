package cld2

import "testing"

func TestLanguageFromCode(t *testing.T) {
	l := LanguageFromCode("da")
	if l != DANISH {
		t.Error("want 'da' code to return Danish, got %s", l.String())
	}

	l = LanguageFromCode("something")
	if l != UNKNOWN_LANGUAGE {
		t.Error("want 'something' code to return Unknown Language, got %s", l.String())
	}

	l = LanguageFromCode("un")
	if l != UNKNOWN_LANGUAGE {
		t.Error("want 'something' code to return Unknown Language, got %s", l.String())
	}
}

func TestLanguageFromID(t *testing.T) {
	l := NewLanguage(uint16(THAI))
	if l != THAI {
		t.Errorf("want Thai, got %v", l)
	}

	l = NewLanguage(65000)
	if l != UNKNOWN_LANGUAGE {
		t.Errorf("want Unknown Language, got %v", l)
	}

	l = NewLanguage(0)
	if l != ENGLISH {
		t.Errorf("want English, got %v", l)
	}

	l = NewLanguage(uint16(NUM_LANGUAGES))
	if l != UNKNOWN_LANGUAGE {
		t.Errorf("want Unknown Language, got %v", l)
	}
}

func TestLoadPlugin(t *testing.T) {
	if !Enabled {
		t.Skip("want to start enabled")
		return
	}
	err := LoadPlugin("./lib/notfound.so")
	if err != nil {
		// We should never get error if enabled.
		t.Error(err)
	}
	Enabled = false
	err = LoadPlugin("lib/notfound.so")
	if err == nil && err != ErrNoPlugins {
		// We should get an error
		t.Error(err)
	}
	err = LoadPlugin("lib/cld2go.so")
	if err != nil && err != ErrNoPlugins {
		// We should get an error
		t.Fatal(err)
	}
	if !Enabled {
		t.Error("We should have been enabled now.")
	}
}
