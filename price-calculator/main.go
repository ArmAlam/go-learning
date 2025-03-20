package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {

		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)

		// now both cmdmanager and filemanager implement the IOManager interface, can easily switch between FILE-MANAGER and CMD-MANAGER
		// cmd := cmdmanager.New()
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index])

		// if err != nil {
		// 	fmt.Println("Couldn't process the job: ", err)
		// }

	}

	for index := range taxRates {
		// define different cases for different channels, allows us to wait for one channel to emit value,
		// will not wait for other channels, ignore other case if one satisfies
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)

			}

		case <-doneChans[index]:
			fmt.Println("Done")
		}

	}

	// all the code below is replaced by the code above
	// for _, errorChan := range errorChans {
	// 	<-errorChan
	// }

	// // Wait for all channels to complete, This ensures that the program does not exit until all workers complete.
	// for _, doneChan := range doneChans {
	// 	<-doneChan //<-doneChan waits until a value is sent into doneChan from another goroutine
	// }
}
