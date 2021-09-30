package imgdatabase

import (
	"fmt"
	"log"

	"github.com/sugarme/tokenizer/pretrained"
)

// Preprocess returns preprocessed text
func Preprocess() []string {
	// comment here
	tk := pretrained.BertBaseUncased()

	sentence := `The quick brown fox jumps over the [L A Z Y D O G]`
	en, err := tk.EncodeSingle(sentence)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tokens: %q\n", en.Tokens)
	return en.Tokens
}
