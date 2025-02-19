// challenges/packages/begin/main.go
package main

// import the proverbs package
import (
	"github.com/jboursiquot/go-proverbs"
	"fmt"
)
// getRandomProverb returns a random proverb from the proverbs package
func getRandomProverb() string {
	return proverbs.Random().Saying
}

func main() {
	// print the result of calling your getRandomProverb function
	fmt.Println(getRandomProverb())
}
