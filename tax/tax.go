package tax

import (
	"fmt"
	"os"
)

func CalcTax() {
	var lower_bound_standard [1010]int
	var tax_number [13][1010]int
	var tax_value = 0.0

	fmt.Printf("이 간이세액표의 해당 세액은 소득세법에 따른 근로소득공제, 기본공제, 특별소득공제 및 특별세액공제 중 일부, 연금보험료공제, 근로소득세액공제와 해당 세율을 반영하여 계산한 금액임. 이 경우 “특별소득공제 및 특별세액공제 중 일부”는 다음의 계산식에 따라 계산한 금액을 소득공제하여 반영한 것임.\n")
	var real_family_number int
	fmt.Printf("부양가족의 수를 입력해주세요(본인 및 배우자도 각각 1명으로 보아 계산하여야 합니다.) :  ")
	fmt.Scan(&real_family_number)
	fmt.Print("\n")
	var children_number int
	fmt.Printf("7세 이상 20세 이하 자녀의 수를 입력해주세요 : ")
	fmt.Scan(&children_number)
	fmt.Printf("\n")
	var qualified_family_number = real_family_number + children_number

	var monthly_salary int
	var tax_percentage int
	fmt.Printf("비과세 및 자녀 학자금 지원금액을 제외한 월 급여액을 입력해주세요(천원 단위) : ")
	fmt.Scan(&monthly_salary)
	fmt.Printf("\n")
	fmt.Printf("선택하신 세액의 비율(80%%, 100%%, 120%%)을 입력해주세요. 선택하지 아니한 경우 100(%%)를 입력해주세요 : ")
	fmt.Scan(&tax_percentage)
	fmt.Printf("\n")

	read, _ := os.Open("standard.txt")
	for i := 1; i < 654; i++ {
		fmt.Fscan(read, &lower_bound_standard[i])
	}
	read1, _ := os.Open("1person.txt")
	for i := 59; i < 648; i++ {
		fmt.Fscan(read1, &tax_number[1][i])
	}
	read2, _ := os.Open("2person.txt")
	for i := 115; i < 648; i++ {
		fmt.Fscan(read2, &tax_number[2][i])
	}
	read3, _ := os.Open("3person.txt")
	for i := 169; i < 648; i++ {
		fmt.Fscan(read3, &tax_number[3][i])
	}
	read4, _ := os.Open("4person.txt")
	for i := 186; i < 648; i++ {
		fmt.Fscan(read4, &tax_number[4][i])
	}
	read5, _ := os.Open("5person.txt")
	for i := 203; i < 648; i++ {
		fmt.Fscan(read5, &tax_number[5][i])
	}
	read6, _ := os.Open("6person.txt")
	for i := 220; i < 648; i++ {
		fmt.Fscan(read6, &tax_number[6][i])
	}
	read7, _ := os.Open("7person.txt")
	for i := 237; i < 648; i++ {
		fmt.Fscan(read7, &tax_number[7][i])
	}
	read8, _ := os.Open("8person.txt")
	for i := 254; i < 648; i++ {
		fmt.Fscan(read8, &tax_number[8][i])
	}
	read9, _ := os.Open("9person.txt")
	for i := 270; i < 648; i++ {
		fmt.Fscan(read9, &tax_number[9][i])
	}
	read10, _ := os.Open("10person.txt")
	for i := 286; i < 648; i++ {
		fmt.Fscan(read10, &tax_number[10][i])
	}
	read11, _ := os.Open("11person.txt")
	for i := 299; i < 648; i++ {
		fmt.Fscan(read11, &tax_number[11][i])
	}

	var standard_position int
	for i := 1; i < 654; i++ {
		if lower_bound_standard[i] <= monthly_salary {
			standard_position = i
		}
	}

	if qualified_family_number <= 11 {
		if standard_position <= 647 {
			tax_value = float64(tax_number[qualified_family_number][standard_position])
		} else if standard_position == 648 {
			tax_value = float64(tax_number[qualified_family_number][647]) + float64(monthly_salary-10000)*1000*0.98*0.35
		} else if standard_position == 649 {
			tax_value = float64(tax_number[qualified_family_number][647]) + 1372000 + float64(monthly_salary-14000)*1000*0.98*0.38
		} else if standard_position == 650 {
			tax_value = float64(tax_number[qualified_family_number][647]) + 6585600 + float64(monthly_salary-28000)*1000*0.98*0.4
		} else if standard_position == 651 {
			tax_value = float64(tax_number[qualified_family_number][647]) + 7369600 + float64(monthly_salary-30000)*1000*0.4
		} else if standard_position == 652 {
			tax_value = float64(tax_number[qualified_family_number][647]) + 13369600 + float64(monthly_salary-45000)*1000*0.42
		} else {
			tax_value = float64(tax_number[qualified_family_number][647]) + 31009600 + float64(monthly_salary-87000)*1000*0.45
		}
	} else {
		if standard_position <= 647 {
			tax_value = float64(tax_number[10][standard_position])
			tax_value -= float64(tax_number[11][standard_position])
		} else if standard_position == 648 {
			tax_value = float64(tax_number[10][647]) + float64(monthly_salary-10000)*1000*0.98*0.35
			tax_value -= float64(tax_number[11][647]) + float64(monthly_salary-10000)*1000*0.98*0.35
		} else if standard_position == 649 {
			tax_value = float64(tax_number[10][647]) + 1372000 + float64(monthly_salary-14000)*1000*0.98*0.38
			tax_value -= float64(tax_number[11][647]) + 1372000 + float64(monthly_salary-14000)*1000*0.98*0.38
		} else if standard_position == 650 {
			tax_value = float64(tax_number[10][647]) + 6585600 + float64(monthly_salary-28000)*1000*0.98*0.4
			tax_value -= float64(tax_number[11][647]) + 6585600 + float64(monthly_salary-28000)*1000*0.98*0.4
		} else if standard_position == 651 {
			tax_value = float64(tax_number[10][647]) + 7369600 + float64(monthly_salary-30000)*1000*0.4
			tax_value -= float64(tax_number[11][647]) + 7369600 + float64(monthly_salary-30000)*1000*0.4
		} else if standard_position == 652 {
			tax_value = float64(tax_number[10][647]) + 13369600 + float64(monthly_salary-45000)*1000*0.42
			tax_value -= float64(tax_number[11][647]) + 13369600 + float64(monthly_salary-45000)*1000*0.42
		} else {
			tax_value = float64(tax_number[10][647]) + 31009600 + float64(monthly_salary-87000)*1000*0.45
			tax_value -= float64(tax_number[11][647]) + 31009600 + float64(monthly_salary-87000)*1000*0.45
		}
		tax_value *= float64(qualified_family_number - 11)
	}

	fmt.Printf("입력하신 데이터를 바탕으로 계산한 공제대상가족의 수는 %d명 입니다\n", qualified_family_number)
	fmt.Printf("이때의 원천징수세액은 %d원(소득세 %d원, 지방소득세 %d원) 입니다.", int64(tax_value*1.1), int64(tax_value), int64(tax_value*0.1))
}
