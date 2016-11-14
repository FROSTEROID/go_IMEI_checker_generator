package main

import "fmt" 		// to output some text
import "os"			// to get the cmd args
import "time"		// to seed the random tree
import "math/rand"	// too grow the random tree
//import "strconv"	// extracting numbers out of chars here!

const IMEI_BASE_DIGITS_COUNT int = 14 // The number of digits without the last - the control one.
const ASCII_ZERO_NUMBER uint8 = 48

func main() {
	var argc int = len(os.Args)
	if argc == 1 { // Nothing given, working as a generator:
		var sum int = 0 								// the control sum of digits
		var toAdd int = 0
		var digits [IMEI_BASE_DIGITS_COUNT+1]int 		// darray for dadigits
		
		lolrndsrc := rand.NewSource(time.Now().UnixNano())	// creating a randomizer instance initiated with TIME!
		lolrnd := rand.New(lolrndsrc)
		
		for i := 0; i<IMEI_BASE_DIGITS_COUNT; i++{		// generating all the base digits
			digits[i] = lolrnd.Intn(10)					// with lolrnd.Intn
			toAdd = digits[i]
			if (i+1) % 2 == 0 { 						// special proc for every 2nd one
				toAdd *= 2 				
				if (toAdd >= 10)  { toAdd = (toAdd %10)+1 }
			}
			sum += toAdd								// and summarizing
			fmt.Printf("%d",digits[i])					// and even printing them here!
		}
		var ctrlDigit int = (sum * 9) % 10 				// calculating the control digit
		digits[IMEI_BASE_DIGITS_COUNT] = ctrlDigit		// adding to darray
		fmt.Printf("%d",ctrlDigit)						// and even printing it here!

	}else { // Got some argument(s), gonna try to check for IMEI(s).
		var sum int = 0 									// the control sum of digits
		var toAdd uint8 = 0
		var onlyDigits bool = true
		args := os.Args										// put args to a pretty arr
		for i:=1; i<argc; i++{								// all of them will be checked
			sum = 0	
			onlyDigits = true
			if len(args[i]) == IMEI_BASE_DIGITS_COUNT+1 {	// firstly checking length
				for l:=0; l<IMEI_BASE_DIGITS_COUNT; l++{	// now walking through characters
					if( ( args[i][l] >= ASCII_ZERO_NUMBER ) && ( args[i][l] <= 58 ) ) {// if this is a digit
						digit := args[i][l] - ASCII_ZERO_NUMBER//convert the ascii char to a digit
						toAdd = digit
						if (l+1) % 2 == 0 { 		// special proc for every 2nd one
							toAdd *= 2
							if (toAdd >= 10)  { toAdd = (toAdd %10)+1 }
						}
						sum += int(toAdd) 						// sum with the sum
						
					}else {onlyDigits = false; break;}		// if it wasn't a digit, stop this check
				}
				if(onlyDigits){								// if only found digits
					var ctrlDigit int = (sum * 9) % 10		// calculating the control digit
					// Finally if the the last character is the ctrlDigit:
					if((args[i][IMEI_BASE_DIGITS_COUNT] - ASCII_ZERO_NUMBER) == uint8(ctrlDigit)){ 
						fmt.Printf("%s is TRULY an IMEI!\n", args[i])
					}else{								// if the ctrlDigit is not zbs...
						fmt.Printf("%s has WRONG CONTROL DIGIT! Must have %d!\n", args[i], ctrlDigit)
					}
				}else{									// if containing non-digits...
					fmt.Printf("%s isn't an IMEI as it HAS NON-DIGITS!\n", args[i])
				}
			}else{										// if not fitting by length...
				fmt.Printf("%s isn't an IMEI even by it's LENGTH. Must be %d digits!\n", 
						args[i], IMEI_BASE_DIGITS_COUNT+1)
				continue
			}
		}
	}
	
}