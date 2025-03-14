package maps

import "fmt"

func main() {

	myWebsites := map[string]string{ // map[key]value
		"SiteName": "GolangCode.com",
		"SiteURL":  "https://golangcode.com",
	}

	// add to map
	myWebsites["merdeka"] = "https://merdeka.com"

	// delete from map
	delete(myWebsites, "merdeka")

	fmt.Println("Site Name: " + myWebsites["SiteURL"])

}
