package models

// Location struct
type Location struct {
	BaseModel
	GoogleMaps  string `json:"google_maps,omitempty"`
	Address     string `json:"address,omitempty"`
	Subdistrict string `json:"subdistrict,omitempty"`
	District    string `json:"district,omitempty"`
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
}

// TableName Location
func (Location) TableName() string {
	return "location"
}
