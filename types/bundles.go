package types

type BundlePost struct {
	Name      string `json:"name"`
	Notes     string `json:"notes,omitempty"`
	OwnerID   int    `json:"owner_id,omitempty"`
	OwnerType string `json:"owner_type,omitempty"`
}

type Bundle struct {
	Cover          *Photo        `json:"bundle_cover,omitempty"`
	Items          []BundledItem `json:"bundled_items,omitempty"`
	ItemsCount     int           `json:"bundled_items_count,omitempty"`
	CreatedAt      Date          `json:"created_at"`
	FavoritesCount int           `json:"favorites_count,omitempty"`
	FirstPhoto     *Photo        `json:"first_photo,omitempty"`
	ID             int           `json:"id"`
	Name           string        `json:"name"`
	Notes          string        `json:"notes,omitempty"`
	Permalink      string        `json:"permalink"`
	UpdatedAt      Date          `json:"updated_at"`
	Owner          *UserSmall    `json:"user"`
}

type BundledItem struct {
	//BOZO
}

type Photo struct {
	//BOZO
}

type UserSmall struct {
	//BOZO
}
