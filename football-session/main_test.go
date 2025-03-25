package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestFetchJSON verifies that fetchJSON returns the expected data.
func TestFetchJSON(t *testing.T) {
	expected := `{"name": "Test League", "rounds": []}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(expected))
	}))
	defer ts.Close()

	data, err := fetchJSON(ts.URL)
	if err != nil {
		t.Fatalf("fetchJSON error: %v", err)
	}
	if string(data) != expected {
		t.Errorf("expected %q, got %q", expected, string(data))
	}
}

// TestParseLeague checks that parseLeague correctly converts JSON data into a League.
func TestParseLeague(t *testing.T) {
	jsonStr := `{
		"name": "Test League",
		"rounds": [{
			"name": "Round 1",
			"matches": [{
				"date": "2020-01-01",
				"team1": {"key": "t1", "name": "Team One", "code": "T1"},
				"team2": {"key": "t2", "name": "Team Two", "code": "T2"},
				"score1": 2,
				"score2": 1
			}]
		}]
	}`
	league, err := parseLeague([]byte(jsonStr))
	if err != nil {
		t.Fatalf("parseLeague error: %v", err)
	}
	if league.Name != "Test League" {
		t.Errorf("expected league name 'Test League', got %q", league.Name)
	}
	if len(league.Rounds) != 1 {
		t.Errorf("expected 1 round, got %d", len(league.Rounds))
	}
}

// TestGoalsScoredForTeam verifies that goalsScoredForTeam returns the correct totals
func TestGoalsScoredForTeam(t *testing.T) {
	league := League{
		Name: "Test League",
		Rounds: []Round{
			{
				Name: "Round 1",
				Matches: []Match{
					{
						Date:   "2020-01-01",
						Team1:  Team{Key: "t1", Name: "Team One", Code: "T1"},
						Team2:  Team{Key: "t2", Name: "Team Two", Code: "T2"},
						Score1: 2,
						Score2: 1,
					},
					{
						Date:   "2020-01-02",
						Team1:  Team{Key: "t2", Name: "Team Two", Code: "T2"},
						Team2:  Team{Key: "t1", Name: "Team One", Code: "T1"},
						Score1: 3,
						Score2: 0,
					},
				},
			},
		},
	}
	goalsT1 := goalsScoredForTeam(league, "t1")
	if goalsT1 != 2 {
		t.Errorf("expected t1 goals: 2, got %d", goalsT1)
	}
	goalsT2 := goalsScoredForTeam(league, "t2")
	if goalsT2 != 4 {
		t.Errorf("expected t2 goals: 4, got %d", goalsT2)
	}
	goalsCase := goalsScoredForTeam(league, "T1")
	if goalsCase != 2 {
		t.Errorf("expected 'T1' goals: 2, got %d", goalsCase)
	}
}

// runWithURL is a helper for testing run using a custom URL and *testing.T.
func runWithURL(t *testing.T, url string, teamKey string) int {
	body, err := fetchJSON(url)
	if err != nil {
		t.Fatalf("fetchJSON error: %v", err)
	}
	league, err := parseLeague(body)
	if err != nil {
		t.Fatalf("parseLeague error: %v", err)
	}
	return goalsScoredForTeam(league, teamKey)
}

// TestRun verifies the end-to-end flow using a test HTTP server.
func TestRun(t *testing.T) {
	testJSON := `{
		"name": "Test League",
		"rounds": [{
			"name": "Round 1",
			"matches": [{
				"date": "2020-01-01",
				"team1": {"key": "t1", "name": "Team One", "code": "T1"},
				"team2": {"key": "t2", "name": "Team Two", "code": "T2"},
				"score1": 2,
				"score2": 1
			}]
		}]
	}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(testJSON))
	}))
	defer ts.Close()

	goals := runWithURL(t, ts.URL, "t2")
	if goals != 1 {
		t.Errorf("expected t2 goals: 1, got %d", goals)
	}
}
