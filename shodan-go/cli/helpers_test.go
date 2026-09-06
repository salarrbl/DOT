package cli

import "testing"

func TestEscapeData(t *testing.T) {
	got := EscapeData("a\nb\tc")
	if got != "a\\nb\\tc" {
		t.Fatalf("%q", got)
	}
}

func TestMatchFilters(t *testing.T) {
	b := map[string]any{"port": float64(80), "org": "Google"}
	if !MatchFilters(b, []string{"org:Google"}) {
		t.Fatal("expected match")
	}
	if MatchFilters(b, []string{"org:Microsoft"}) {
		t.Fatal("expected no match")
	}
}

func TestHumanizeAPIPlan(t *testing.T) {
	if HumanizeAPIPlan("corp") != "Corporate API" {
		t.Fatal(HumanizeAPIPlan("corp"))
	}
}

func TestGetBannerFieldNested(t *testing.T) {
	b := map[string]any{"location": map[string]any{"city": "Vilnius"}}
	if GetBannerField(b, "location.city") != "Vilnius" {
		t.Fatal(GetBannerField(b, "location.city"))
	}
}
