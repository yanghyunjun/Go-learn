package main

import (
	"fmt"
	"strings"
)

func lengthAndUpper(name string) (leng int, upper string) {

	defer fmt.Println(`length and upper done`)

	leng = len(name)
	upper = strings.ToUpper(name)
	return

} /*Return two variables and Execute defer after function ends*/

func AddFunc(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num

	}
	return total
} /* for syntex and Index occurs when using range*/

func KorAgeCheck(age int) bool {
	if korAge := age + 2; korAge < 20 {
		return false
	}
	return true
} /* if syntex and swich syntex samething*/

type User struct {
	name         string
	age          int
	favoritebook []string
} /*structs*/

func main() {
	// length, toUpper := lengthAndUpper("hyunjun")
	// fmt.Println(length, toUpper)
	/*func lengthAndUpper*/

	// result := AddFunc(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// fmt.Println(result)
	/*func AddFunc*/

	// fmt.Println(KorAgeCheck(18))
	// fmt.Println(KorAgeCheck(16))
	/*func LoginCheck*/

	// a := 2
	// b := &a
	// c := &a
	// fmt.Println(a, *b, c)
	/*pointer "&" this memoryAddress , "*" this look at the memory address.*/

	// names := []string{"yang","hong","kim"}
	// names = append(names,"park")
	/*Do not use push() , use to appand(arr,value) */

	// user := map[string]string{"name": "hyunjun", "age": "27"}
	// for key, value := range user {
	// 	fmt.Println(key, value)
	// }
	/*map syntex*/

	// favoritebook := []string{"javascript", "react"}
	// userinfo := User{name: "hyunjun", age: 27, favoritebook: favoritebook}
	// fmt.Println(userinfo)
	/*structs syntex*/
}
