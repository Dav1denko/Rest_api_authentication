package restapiauthentication

type Info struct {
	GUID         int    `json:"GUID"`
	Token        string `json:"Token"`
	RefreshToken string `json:"RefreshToken"`
}
