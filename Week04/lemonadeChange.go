package Week04

func lemonadeChange(bills []int) bool {

	if len(bills) ==  0 || bills[0] != 5 {
		return false
	}
	five := 1
	ten :=0
	for i := 1 ;i < len(bills);i++{
		if bills[i] == 5{
			five++
		} else if bills[i] == 10{
			five--
			ten++
		}else if bills[i] == 20{
			if ten >= 1 && five >= 1{
				ten--
				five--
			}else if ten == 0 && five >= 3{
				five -= 3
			}else{
				return false
			}
		}

	}
	return true
}