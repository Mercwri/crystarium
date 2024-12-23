package reportdata

type ReportQuery struct {
	ReportData struct {
		Report `graphql:"report(code: $code)"`
	}
}

type FightQuery struct {
	ReportData struct {
		Report struct {
			Events Event `graphql:"events(fightIDs: $fightID, filterExpression: $fex, startTime:$timestamp, endTime:999999)"`
		} `graphql:"report(code: $code)"`
	}
}

type Report struct {
	Title  string
	Fights []Fight
	Code   string
}

type Fight struct {
	ID   int
	Kill bool
}

type Event struct {
	Data              []Data `json:"data" scalar:"true"`
	NextPageTimestamp float64
}

type Data struct {
	Timestamp           int     `json:"timestamp"`
	Type                string  `json:"type"`
	SourceID            int     `json:"sourceID,omitempty"`
	TargetID            int     `json:"targetID,omitempty"`
	AbilityGameID       int     `json:"abilityGameID,omitempty"`
	Fight               int     `json:"fight"`
	HitType             int     `json:"hitType,omitempty"`
	Amount              int     `json:"amount,omitempty"`
	UnmitigatedAmount   int     `json:"unmitigatedAmount,omitempty"`
	Multiplier          float64 `json:"multiplier,omitempty"`
	PacketID            int     `json:"packetID,omitempty"`
	Duration            int     `json:"duration,omitempty"`
	Value               int     `json:"value,omitempty"`
	Bars                int     `json:"bars,omitempty"`
	DirectHit           bool    `json:"directHit,omitempty"`
	Melee               bool    `json:"melee,omitempty"`
	ExtraAbilityGameID  int     `json:"extraAbilityGameID,omitempty"`
	Unpaired            bool    `json:"unpaired,omitempty"`
	Overheal            int     `json:"overheal,omitempty"`
	Tick                bool    `json:"tick,omitempty"`
	GaugeID             string  `json:"gaugeID,omitempty"`
	Data1               string  `json:"data1,omitempty"`
	Data2               string  `json:"data2,omitempty"`
	Data3               string  `json:"data3,omitempty"`
	Data4               string  `json:"data4,omitempty"`
	AttackerID          int     `json:"attackerID,omitempty"`
	Absorbed            int     `json:"absorbed,omitempty"`
	Absorb              int     `json:"absorb,omitempty"`
	FinalizedAmount     float64 `json:"finalizedAmount,omitempty"`
	Simulated           bool    `json:"simulated,omitempty"`
	ExpectedAmount      int     `json:"expectedAmount,omitempty"`
	ExpectedCritRate    int     `json:"expectedCritRate,omitempty"`
	ActorPotencyRatio   float64 `json:"actorPotencyRatio,omitempty"`
	GuessAmount         float64 `json:"guessAmount,omitempty"`
	DirectHitPercentage float64 `json:"directHitPercentage,omitempty"`
	BonusPercent        int     `json:"bonusPercent,omitempty"`
	Stack               int     `json:"stack,omitempty"`
	ExtraInfo           int     `json:"extraInfo,omitempty"`
}
