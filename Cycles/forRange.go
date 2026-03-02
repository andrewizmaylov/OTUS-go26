package cycles
import {fmt }
func handle(input) {
	for i, v := range input {
		fmt.Printf("%d: %s\n", i, input(v))
	}
}
