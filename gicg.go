package main

import "fmt" 		// to output some text
import "os"			// to get the cmd args
import "time"		// to seed the random tree
import "math/rand"	// too grow the random tree

const IMEI_BASE_DIGITS_COUNT int = 16 // The number of digits without the last - the control one.

func main() {
	if len(os.Args)==1 { // Nothing given, working as a generator:
		var sum int = 0 										// the control sum of digits
		var digits [IMEI_BASE_DIGITS_COUNT+1]int 		// darray for dadigits
		
		lolrndsrc := rand.NewSource(time.Now().UnixNano())	// creating a randomizer instance initiated with TIME!
		lolrnd := rand.New(lolrndsrc)
		
		for i := 0; i<=IMEI_BASE_DIGITS_COUNT; i++{		// generating all the base digits
			digits[i] = lolrnd.Intn(10)					// with lolrnd.Intn
			sum += digits[i]							// and summarizing
			if i % 2 == 0 { sum += digits[i] }			// twice for every 2nd ones
			fmt.Printf("%d",digits[i])					// and even printing them it here!
		}
		var ctrlDigit int = (sum * 9) % 10 				// calculating the control digit
		digits[IMEI_BASE_DIGITS_COUNT] = ctrlDigit		// adding to darray

	}else { // Got some argument(s), gonna try to check for IMEI(s).
		
	}
	
}