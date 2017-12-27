package main

import (
	. "github.com/mojocn/recommend-go/collabFilter"
	"fmt"
)

func main() {
	// User product matrix. 0 indicates products not viewed by user.
	// Uses cosine similarity.


	// arguments to **MakeRatingMatrix** are: data, nrows, ncols.
	prefs := MakeRatingMatrix([]float64{
		2, 3, 4, 1, 5,
		3, 0, 3, 3, 0,
		4, 4, 1, 2, 3,
		2, 4, 0, 3, 4,
		3, 1, 3, 0, 4}, 5, 5)

	// Can also load/build matrix from a text file
	//prefs := Load("path/to/file", "separator")


	// product titles <- column titles for prefs matrix
	products := []string{"Spiderman", "Big Momma's House", "Vanilla Sky", "Pacific Rim", "The Mask"}
	// gets recommendations for user 1 (second row) for un-rated products.
	prods, scores, err := GetRecommendations(prefs, 1, products)
	if err != nil {
		fmt.Println("WHAT!?")
	}
	fmt.Printf("\nRecommended Products are: %v, with scores: %v", prods, scores)

	// For a binary matrix, use the getBinaryRecommendations function in the exact same way.
	// Uses Jaccard Similarity to return confidence/probabality of user's recommendations
	binaryPrefs := MakeRatingMatrix([]float64{
		1, 1, 1, 1, 0,
		0, 1, 1, 0, 1,
		1, 1, 1, 1, 1,
		1, 1, 0, 0, 1,
		1, 0, 1, 1, 1}, 5, 5)
	// Returns recommended products for User ID 1 (second row) in descending order, w/ corresponding confidence/probability,
	// and error - if applicable.

	fmt.Println(GetBinaryRecommendations(binaryPrefs, 1, products))
}
