package structs

type Menu struct {
	C struct {
		Coffee int `json:"Coffee"`
	} `json:"C"`
	D struct {
		Dosa int `json:"Dosa"`
	} `json:"D"`
	T struct {
		TomatoSoup int `json:"Tomato Soup"`
	} `json:"T"`
	I struct {
		Idli int `json:"Idli"`
	} `json:"I"`
	V struct {
		Vada int `json:"Vada"`
	} `json:"V"`
	B struct {
		BhatureChhole int `json:"Bhature&Chhole"`
	} `json:"B"`
	P struct {
		PaneerPakoda int `json:"Paneer Pakoda"`
	} `json:"P"`
	M struct {
		Manchurian int `json:"Manchurian"`
	} `json:"M"`
	H struct {
		HakkaNoodle int `json:"Hakka Noodle"`
	} `json:"H"`
	F struct {
		FrenchFries int `json:"French Fries"`
	} `json:"F"`
	J struct {
		Jalebi int `json:"Jalebi"`
	} `json:"J"`
	L struct {
		Lemonade int `json:"Lemonade"`
	} `json:"L"`
	S struct {
		SpringRoll int `json:"spring Roll"`
	} `json:"S"`
}
