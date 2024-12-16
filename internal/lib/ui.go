package lib

import (
	"bufio"
	"fmt"
	"os"
)

func MustPressEnter() {
	fmt.Print("Press 'Enter' to continue...")
	_, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
	Must(err)
}
