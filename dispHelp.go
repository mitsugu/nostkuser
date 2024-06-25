package main

import (
	"fmt"
	"os"
)

// dispHelp
func dispHelp() {
	usageTxt := `Usage :
	nostkuser <hex epub>`
	fmt.Fprintf(os.Stderr, "%s\n", usageTxt)
}

