package characterdata

type Query struct {
	CharacterData struct {
		Character `graphql:"character(name: $name, serverSlug: $server, serverRegion: $region)"`
	}
}

type Character struct {
	ID           int
	Name         string
	LodestoneID  int          `graphql:"lodestoneID"`
	ZoneRankings ZoneRankings `scalar:"true" graphql:"zoneRankings(zoneID: $zoneID)"`
}

type ZoneRankings struct {
	Difficulty int       `json:"difficulty"`
	Metric     string    `json:"metric"`
	Rankings   []Ranking `json:"rankings"`
}

type Ranking struct {
	Encounter struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"encounter"`
	RankPercent   float32
	MedianPercent float32
	LockedIn      bool
	TotalKills    int64
	FastestKill   int64
	Spec          string
	BestSpec      string
}
