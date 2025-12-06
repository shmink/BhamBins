package bins

type Collection struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

type Collections struct {
	Bins []Collection `json:"bins"`
}
