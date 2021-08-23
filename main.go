package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

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
		ID           int    `xml:"internal-id,attr"`
		ProjectName  string `xml:"location>metro>name"`
		BuildingID   int    `xml:"yandex-building-id"`
		BuildingName string `xml:"building-name"`
		SectionName  string `xml:"building-section"`
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

func main() {
	xmlFile, err := os.Open("export_yandex_leningradka_msk.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	data := &Data{}

	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'data' which we defined above
	xml.Unmarshal(byteValue, data)

	var project []Project
	for i := 0; i < len(data.DataOffer); i++ {
		newSection := true
		newBuilding := true
		newProject := true
		var pr Project
		var build Building
		var sect Section
		var lot Lot
		j := 0 // ind project
		k := 0 // ind building
		z := 0 // ind section

		lot.ID = data.DataOffer[i].ID
		lot.Floor = data.DataOffer[i].Floor
		lot.KitchenSquare = data.DataOffer[i].KitchenSquare
		lot.LivingSquare = data.DataOffer[i].LivingSquare
		lot.LotType = data.DataOffer[i].LotType
		lot.Price = data.DataOffer[i].Price
		lot.RoomType = data.DataOffer[i].RoomType
		lot.TotalSquare = data.DataOffer[i].TotalSquare

		sect.Name = data.DataOffer[i].SectionName
		sect.Lot = append(sect.Lot, lot)

		build.ID = data.DataOffer[i].BuildingID
		build.Name = data.DataOffer[i].BuildingName
		build.Section = append(build.Section, sect)

		pr.Name = data.DataOffer[i].ProjectName
		pr.Address = data.DataOffer[i].Address
		pr.Description = data.DataOffer[i].Description
		pr.Building = append(pr.Building, build)

		for j < len(project) {
			if project[j].Name == data.DataOffer[i].ProjectName {
				newProject = false
				break
			}
			j++
		}
		if newProject {
			project = append(project, pr)
		} else {
			for k < len(project[j].Building) {
				if project[j].Building[k].Name == data.DataOffer[i].BuildingName {
					newBuilding = false
					break
				}
				k++
			}
			if newBuilding {
				project[j].Building = append(project[j].Building, build)
			} else {
				for z < len(project[j].Building[k].Section) {
					if project[j].Building[k].Section[z].Name == data.DataOffer[i].SectionName {
						newSection = false
						break
					}
					z++
				}
				if newSection {
					project[j].Building[k].Section = append(project[j].Building[k].Section, sect)
				} else {
					project[j].Building[k].Section[z].Lot = append(project[j].Building[k].Section[z].Lot, lot)
				}
			}
		}
	}
	fmt.Println(len(project))
	fmt.Println(project)
}
