package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))

	for index, taxRate := range taxRates {

		doneChans[index] = make(chan bool)

		// now both cmdmanager and filemanager implement the IOManager interface, can easily switch between FILE-MANAGER and CMD-MANAGER
		// cmd := cmdmanager.New()
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index])

		// if err != nil {
		// 	fmt.Println("Couldn't process the job: ", err)
		// }

	}

	// Wait for all channels to complete, This ensures that the program does not exit until all workers complete.
	for _, doneChan := range doneChans {
		<-doneChan //<-doneChan waits until a value is sent into doneChan from another goroutine
	}
}
