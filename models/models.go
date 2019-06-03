package models

type Peer struct {
	ID       string                 `json:"peerID"`
	RawMap   map[string]interface{} `json:"profile"`
	LastPing int64                  `json:"lastPing"`
}

type Listing struct {
	AcceptedCurrencies []string      `json:"acceptedCurrencies"`
	AverageRating      int64         `json:"averageRating"`
	Categories         []string      `json:"categories"`
	CoinType           string        `json:"coinType"`
	ContractType       string        `json:"contractType"`
	Description        string        `json:"description"`
	Hash               string        `json:"hash"`
	Language           string        `json:"language"`
	Location           Location      `json:"location"`
	Moderators         []interface{} `json:"moderators"`
	Nsfw               bool          `json:"nsfw"`
	Price              Price         `json:"price"`
	RatingCount        int64         `json:"ratingCount"`
	Slug               string        `json:"slug"`
	Thumbnail          Thumbnail     `json:"thumbnail"`
	Title              string        `json:"title"`
	PeerSlug           string        `json:"peerSlug"`
	ParentPeer         string        `json:"parentPeer"`
}

type Location struct {
	AddressOne string `json:"addressOne"`
	AddressTwo string `json:"addressTwo"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
	PlusCode   string `json:"plusCode"`
	State      string `json:"state"`
	ZipCode    string `json:"zipCode"`
}

type Price struct {
	Amount       int64  `json:"amount"`
	CurrencyCode string `json:"currencyCode"`
	Modifier     int64  `json:"modifier"`
}

type Thumbnail struct {
	Medium string `json:"medium"`
	Small  string `json:"small"`
	Tiny   string `json:"tiny"`
}

type AdvancedSearchQuery struct {
	Query      string        `json:"query"`
	Filters    []string      `json:"filters"`
	Limit      int           `json:"limit"`
	Start      int           `json:"start"`
	Transforms []interface{} `json:"transforms"`
	Sort       string        `json:"sort"`
}
