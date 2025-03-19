package main

import (
	"fmt"

	"example.com/price-calculator/cmdmanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// now both cmdmanager and filemanager implement the IOManager interface, can easily switch between FILE-MANAGER and CMD-MANAGER
		cmd := cmdmanager.New()
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(cmd, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Couldn't process the job: ", err)
		}

	}
}
