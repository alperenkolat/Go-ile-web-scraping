package catch_err

import (
	"fmt"
)

func Control(err error) {
	if err != nil {
		fmt.Println(err)
	}

}
