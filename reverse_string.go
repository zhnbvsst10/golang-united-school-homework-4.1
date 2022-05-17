package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	output := []rune(input) // convert to rune
    	for i, j := 0, len(output)-1; i < j; i, j = i+1, j-1 {
  
        // swap the letters of the string,
        // like first with last and so on.
        output[i], output[j] = output[j], output[i]
    	}
  
    	// return the reversed string.
   	// return string(rn)
	return string(output)
}
