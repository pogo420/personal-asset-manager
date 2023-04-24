package ds

import "time"

type RedeemType int

const (
	L  RedeemType = 0
	NL RedeemType = 1
)

/*
 * ADT for Asset Hierarchies
 */
type AssetBase struct {
	Id          string
	Name        string
	Description string
	Value       float64
}

type LongTermAsset struct {
	AssetBase
	Reedem RedeemType
}

type FDAsset struct {
	AssetBase
	AccountNumber string
}

type RDAsset struct {
	AssetBase
	AccountNumber    string
	StartDate        time.Time
	EndDate          time.Time
	AccumulatedValue float64
}
