package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Team represents a team in the league.
type Team struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// Match represents a football match.
type Match struct {
	Date   string `json:"date"`
	Team1  Team   `json:"team1"`
	Team2  Team   `json:"team2"`
	Score1 int    `json:"score1"`
	Score2 int    `json:"score2"`
}

// Round represents a matchday with its matches.
type Round struct {
	Name    string  `json:"name"`
	Matches []Match `json:"matches"`
}

// League represents the overall league structure.
type League struct {
	Name   string  `json:"name"`
	Rounds []Round `json:"rounds"`
}

// fetchJSON retrieves JSON data from the given URL.
func fetchJSON(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	return body, nil
}

// parseLeague parses the JSON data into a League struct.
func parseLeague(data []byte) (League, error) {
	var league League
	if err := json.Unmarshal(data, &league); err != nil {
		return League{}, fmt.Errorf("error unmarshaling JSON: %w", err)
	}
	return league, nil
}

// goalsScoredForTeam calculates the total goals scored by the team with the given key.
func goalsScoredForTeam(league League, teamKey string) int {
	totalGoals := 0
	teamKeyLower := strings.ToLower(teamKey)
	for _, round := range league.Rounds {
		for _, match := range round.Matches {
			if strings.ToLower(match.Team1.Key) == teamKeyLower {
				totalGoals += match.Score1
			}
			if strings.ToLower(match.Team2.Key) == teamKeyLower {
				totalGoals += match.Score2
			}
		}
	}
	return totalGoals
}

// run fetches the league JSON from the URL and returns total goals for the specified team.
func run(teamKey string) int {
	url := "The S3 link from the challenge."
	body, err := fetchJSON(url)
	if err != nil {
		log.Fatalf("Error fetching JSON: %v", err)
	}

	league, err := parseLeague(body)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	return goalsScoredForTeam(league, teamKey)
}

func main() {
	team := "mancity"
	goals := run(team)
	fmt.Printf("Total goals scored by %s: %d\n", team, goals)
}
