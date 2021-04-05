package stats

import (
	"fmt"
	"github.com/Romiz01/bank1/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       2,
			Amount:   53_00,
			Category: "Book",
			Status:   types.StatusOk,
		},
		{
			ID:       1,
			Amount:   51_00,
			Category: "dog",
			Status:   types.StatusOk,
		},
		{
			ID:       3,
			Amount:   52_00,
			Category: "Cat",
			Status:   types.StatusFail,
		},
	}

	fmt.Println(Avg(payments))

	//Output: 5200
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       2,
			Amount:   53_00,
			Category: "Merve",
			Status:   types.StatusOk,
		},
		{
			ID:       1,
			Amount:   51_00,
			Category: "BBQ",
			Status:   types.StatusOk,
		},
		{
			ID:       3,
			Amount:   52_00,
			Category: "auto",
			Status:   types.StatusFail,
		},
	}

	inCategory := types.Category("auto")
	totalInCategory := TotalInCategory(payments, inCategory)
	fmt.Println(totalInCategory)
	//Output:  1000000

}
