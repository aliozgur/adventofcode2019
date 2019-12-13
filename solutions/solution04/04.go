package solution04


import (
	"adventofcode/utils"
	"fmt"

)

func Run(){
	start := 372304
	end := 847060
	count1 := 0
	count2 := 0

	for i:=start;i<=end;i++{
		if !hasDecDigit(i) {
			if hasSameAdjStep1(i){
				count1++
			}

			if hasSameAdjStep2(i){
				count2++
			}
		}
	}

	fmt.Println("Step-1 Count:",count1)
	fmt.Println("Step-2 Count:",count2)
}

func hasSameAdjStep1(value int) (result bool){
	str := fmt.Sprintf("%d",value)
	var cnt = len(str)


	for i := 0; i < cnt;i++ {
		if i+1 == cnt{
			break
		}
		if string(str[i]) == string(str[i+1]){
			result = true
			break
		}
	}
	return
}

func hasSameAdjStep2(value int) (result bool){
	str := fmt.Sprintf("%d",value)
	var cnt = len(str)
	r := make(map[string]int)

	for i := 0; i < cnt;i++ {
		if i+1 == cnt{
			break
		}
		if string(str[i]) == string(str[i+1]){
			val, ok := r[string(str[i])]
			if !ok{
				r[string(str[i])] = 1
			}else{
				r[string(str[i])] = val+1
			}
		}
	}
	for _,value := range r {
		if value == 1 {
			result = true
			break
		}
	}
	return
}

func hasDecDigit(value int) (result bool){
	str := fmt.Sprintf("%d",value)
	var cnt = len(str)
	for i := 0; i < cnt;i++ {
		if i+1 == cnt{
			break
		}

		v1 := utils.ToInt(string(str[i]))
		v2 := utils.ToInt(string(str[i+1]))
		if v2 < v1{
			result = true
			break
		}
	}
	return
}

