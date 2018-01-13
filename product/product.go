package product

type Product struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	ShortDesc string `json:"short_desc"`
	Desc      string `json:"desc"`
}
