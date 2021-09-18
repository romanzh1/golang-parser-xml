package parse

type Project struct {
	ID          int
	Name        string
	Description string
	Address     string
	Building    []Building
}

type Building struct {
	ID      int
	Name    string
	Section []Section
}

type Section struct {
	ID   int
	Name string
	Lot  []Lot
}

type Lot struct {
	ID            int
	Floor         int
	TotalSquare   float32
	LivingSquare  float32
	KitchenSquare float32
	Price         float32
	LotType       string
	RoomType      string
}

type Data struct {
	DataOffer []struct {
		ID            int     `xml:"internal-id,attr"`
		ProjectName   string  `xml:"location>metro>name"`
		BuildingID    int     `xml:"yandex-building-id"`
		BuildingName  string  `xml:"building-name"`
		SectionName   string  `xml:"building-section"`
		Description   string  `xml:"description"`
		Address       string  `xml:"location>address"`
		Floor         int     `xml:"floor"`
		TotalSquare   float32 `xml:"area>value"`
		LivingSquare  float32 `xml:"living-space>value"`
		KitchenSquare float32 `xml:"kitchen-space>value"`
		Price         float32 `xml:"price>value"`
		LotType       string  `xml:"category"`
		RoomType      string  `xml:"property-type"`
	} `xml:"offer"`
}
