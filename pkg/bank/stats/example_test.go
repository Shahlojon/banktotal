package stats

import (
	"github.com/Shahlojon/bank/pkg/bank/types"
	"fmt"
)

func ExampleAvg(){
	payments := []types.Payment{
	  {
		ID:2,
		Amount:53_00,
		Category: "Cat",
	  },
	  {
		ID:1,
		Amount:53_00,
		Category: "Cat",
	  },
	  {
		ID:3,
		Amount:55_00,
		Category: "Cat",
	  },
	}
  
	fmt.Println(Avg(payments))
  
	//Output: 5366
}