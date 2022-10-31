package main

import (
	"fmt"

	charger "github.com/onyor/gomodule/Charger"
	logger "github.com/onyor/gomodule/Logger"
	summer "github.com/onyor/gomodule/Summer"
)

func main() {
	fmt.Println(charger.Carp(5, 3))
	fmt.Println(logger.Summer(5, 3))
	summer.Log("angara messi")

	// git config \
	// --global \
	// url."git@github.com".insteadOf \
	// "https://github.com"

	// git config \
	// --global \
	// url."https://${onyor}:${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}@github.com".insteadOf \
	// "https://github.com"
}

// git remote set-url origin https://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@github.com/onyor/gomodule.git
// git remote add origin
