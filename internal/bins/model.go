package bins

type Collection struct {
	Name       string `json:"name"`
	WebDate    string `json:"webDate"`
	ActualDate string `json:"actualDate"`
}

type Collections struct {
	Bins []Collection `json:"bins"`
}
