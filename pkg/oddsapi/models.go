package oddsapi

import (
	"fmt"
	"time"
)

type Sport struct {
	Key          string `json:"key"`
	Active       bool   `json:"active"`
	Group        string `json:"group"`
	Details      string `json:"details"`
	Title        string `json:"title"`
	HasOutrights bool   `json:"has_outrights"`
}

func (s *Sport) String() string {
	return fmt.Sprintf("key:%s", s.Key)
}

type Odds struct {
	H2H []float64 `json:"h2h"`
}

type Site struct {
	Key            string `json:"site_key"`
	Nice           string `json:"site_nice"`
	LastUpdateUnix int64  `json:"last_update"`
	Odds           Odds   `json:"odds"`
}

type Match struct {
	SportKey         string   `json:"sport_key"`
	SportNice        string   `json:"sport_nice"`
	Teams            []string `json:"teams"`
	CommenceTimeUnix int64    `json:"commence_time"`
	HomeTeam         string   `json:"home_team"`
	Sites            []Site   `json:"sites"`
	SitesCount       int      `json:"sites_count"`
}

func (m *Match) String() string {
	return fmt.Sprintf("sport_key:%s , home_team:%v , commence_time:%s", m.SportKey, m.HomeTeam, time.Unix(m.CommenceTimeUnix, 0))
}
