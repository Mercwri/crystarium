package events

type CombatantInfo struct {
	Timestamp int    `json:"timestamp"`
	Type      string `json:"type"`
	Fight     int    `json:"fight"`
	SourceID  int    `json:"sourceID"`
	Gear      []any  `json:"gear"`
	Auras     []struct {
		Source  int    `json:"source"`
		Ability int    `json:"ability"`
		Stacks  int    `json:"stacks"`
		Icon    string `json:"icon"`
		Name    string `json:"name"`
	} `json:"auras"`
	Level              int     `json:"level"`
	SimulatedCrit      float64 `json:"simulatedCrit"`
	SimulatedDirectHit string  `json:"simulatedDirectHit"`
}
