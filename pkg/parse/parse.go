package parse

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func FromXML() []Project {
	xmlFile, err := os.Open("pkg/parse/export_yandex_leningradka_msk.xml")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	data := &Data{}

	err = xml.Unmarshal(byteValue, data)
	if err != nil {
		fmt.Println(err)
	}

	var projects []Project
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

		for j < len(projects) {
			if projects[j].Name == data.DataOffer[i].ProjectName {
				newProject = false
				break
			}
			j++
		}
		if newProject {
			projects = append(projects, pr)
		} else {
			for k < len(projects[j].Building) {
				if projects[j].Building[k].Name == data.DataOffer[i].BuildingName {
					newBuilding = false
					break
				}
				k++
			}
			if newBuilding {
				projects[j].Building = append(projects[j].Building, build)
			} else {
				for z < len(projects[j].Building[k].Section) {
					if projects[j].Building[k].Section[z].Name == data.DataOffer[i].SectionName {
						newSection = false
						break
					}
					z++
				}
				if newSection {
					projects[j].Building[k].Section = append(projects[j].Building[k].Section, sect)
				} else {
					projects[j].Building[k].Section[z].Lot = append(projects[j].Building[k].Section[z].Lot, lot)
				}
			}
		}
	}
	return projects
}
