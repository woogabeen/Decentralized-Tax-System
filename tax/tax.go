package tax

import (
	"fmt"
	"os"
)

var lowerBoundStansard [1010]int
var taxNumber [13][1010]int
var TaxVal = 0.0
var standardPosition int

func CalcTax() {
	fmt.Printf("세금 거래를 시작하시겠습니까?  (예 / 아니오)\n")
	var whetherStart string
	fmt.Scan(&whetherStart)
	if whetherStart == "예" {
		fmt.Printf("이 간이세액표의 해당 세액은 소득세법에 따른 근로소득공제, 기본공제, 특별소득공제 및 특별세액공제 중 일부, 연금보험료공제, 근로소득세액공제와 해당 세율을 반영하여 계산한 금액임. 이 경우 “특별소득공제 및 특별세액공제 중 일부”는 다음의 계산식에 따라 계산한 금액을 소득공제하여 반영한 것임.\n\n")

		fmt.Printf("부양가족의 수를 입력해주세요(본인 및 배우자도 각각 1명으로 보아 계산하여야 합니다.) : ")
		var realFamilyNum int
		fmt.Scan(&realFamilyNum)
		fmt.Print("\n")

		fmt.Printf("7세 이상 20세 이하 자녀의 수를 입력해주세요 : ")
		var childrenNum int
		fmt.Scan(&childrenNum)
		fmt.Printf("\n")

		fmt.Printf("비과세 및 자녀 학자금 지원금액을 제외한 월 급여액을 입력해주세요(천원 단위) : ")
		var monthlySalary int
		fmt.Scan(&monthlySalary)
		fmt.Printf("\n")

		fmt.Printf("선택하신 세액의 비율(80%%, 100%%, 120%%)을 입력해주세요. 선택하지 아니한 경우 100(%%)를 입력해주세요 : ")
		var taxRatio int
		fmt.Scan(&taxRatio)
		fmt.Printf("\n")

		var qualFamilyNum = realFamilyNum + childrenNum

		read, _ := os.Open("tax/standard.txt")
		for i := 1; i < 654; i++ {
			fmt.Fscan(read, &lowerBoundStansard[i])
		}
		read1, _ := os.Open("tax/1person.txt")
		for i := 59; i < 648; i++ {
			fmt.Fscan(read1, &taxNumber[1][i])
		}
		read2, _ := os.Open("tax/2person.txt")
		for i := 115; i < 648; i++ {
			fmt.Fscan(read2, &taxNumber[2][i])
		}
		read3, _ := os.Open("tax/3person.txt")
		for i := 169; i < 648; i++ {
			fmt.Fscan(read3, &taxNumber[3][i])
		}
		read4, _ := os.Open("tax/4person.txt")
		for i := 186; i < 648; i++ {
			fmt.Fscan(read4, &taxNumber[4][i])
		}
		read5, _ := os.Open("tax/5person.txt")
		for i := 203; i < 648; i++ {
			fmt.Fscan(read5, &taxNumber[5][i])
		}
		read6, _ := os.Open("tax/6person.txt")
		for i := 220; i < 648; i++ {
			fmt.Fscan(read6, &taxNumber[6][i])
		}
		read7, _ := os.Open("tax/7person.txt")
		for i := 237; i < 648; i++ {
			fmt.Fscan(read7, &taxNumber[7][i])
		}
		read8, _ := os.Open("tax/8person.txt")
		for i := 254; i < 648; i++ {
			fmt.Fscan(read8, &taxNumber[8][i])
		}
		read9, _ := os.Open("tax/9person.txt")
		for i := 270; i < 648; i++ {
			fmt.Fscan(read9, &taxNumber[9][i])
		}
		read10, _ := os.Open("tax/10person.txt")
		for i := 286; i < 648; i++ {
			fmt.Fscan(read10, &taxNumber[10][i])
		}
		read11, _ := os.Open("tax/11person.txt")
		for i := 299; i < 648; i++ {
			fmt.Fscan(read11, &taxNumber[11][i])
		}

		for i := 1; i < 654; i++ {
			if lowerBoundStansard[i] <= monthlySalary {
				standardPosition = i
			}
		}

		if qualFamilyNum <= 11 {
			if standardPosition <= 647 {
				TaxVal = float64(taxNumber[qualFamilyNum][standardPosition])
			} else if standardPosition == 648 {
				TaxVal = float64(taxNumber[qualFamilyNum][647]) + float64(monthlySalary-10000)*1000*0.98*0.35
			} else if standardPosition == 649 {
				TaxVal = float64(taxNumber[qualFamilyNum][647]) + 1372000 + float64(monthlySalary-14000)*1000*0.98*0.38
			} else if standardPosition == 650 {
				TaxVal = float64(taxNumber[qualFamilyNum][647]) + 6585600 + float64(monthlySalary-28000)*1000*0.98*0.4
			} else if standardPosition == 651 {
				TaxVal = float64(taxNumber[qualFamilyNum][647]) + 7369600 + float64(monthlySalary-30000)*1000*0.4
			} else if standardPosition == 652 {
				TaxVal = float64(taxNumber[qualFamilyNum][647]) + 13369600 + float64(monthlySalary-45000)*1000*0.42
			} else {
				TaxVal = float64(taxNumber[qualFamilyNum][647]) + 31009600 + float64(monthlySalary-87000)*1000*0.45
			}
		} else {
			if standardPosition <= 647 {
				TaxVal = float64(taxNumber[10][standardPosition])
				TaxVal -= float64(taxNumber[11][standardPosition])
			} else if standardPosition == 648 {
				TaxVal = float64(taxNumber[10][647]) + float64(monthlySalary-10000)*1000*0.98*0.35
				TaxVal -= float64(taxNumber[11][647]) + float64(monthlySalary-10000)*1000*0.98*0.35
			} else if standardPosition == 649 {
				TaxVal = float64(taxNumber[10][647]) + 1372000 + float64(monthlySalary-14000)*1000*0.98*0.38
				TaxVal -= float64(taxNumber[11][647]) + 1372000 + float64(monthlySalary-14000)*1000*0.98*0.38
			} else if standardPosition == 650 {
				TaxVal = float64(taxNumber[10][647]) + 6585600 + float64(monthlySalary-28000)*1000*0.98*0.4
				TaxVal -= float64(taxNumber[11][647]) + 6585600 + float64(monthlySalary-28000)*1000*0.98*0.4
			} else if standardPosition == 651 {
				TaxVal = float64(taxNumber[10][647]) + 7369600 + float64(monthlySalary-30000)*1000*0.4
				TaxVal -= float64(taxNumber[11][647]) + 7369600 + float64(monthlySalary-30000)*1000*0.4
			} else if standardPosition == 652 {
				TaxVal = float64(taxNumber[10][647]) + 13369600 + float64(monthlySalary-45000)*1000*0.42
				TaxVal -= float64(taxNumber[11][647]) + 13369600 + float64(monthlySalary-45000)*1000*0.42
			} else {
				TaxVal = float64(taxNumber[10][647]) + 31009600 + float64(monthlySalary-87000)*1000*0.45
				TaxVal -= float64(taxNumber[11][647]) + 31009600 + float64(monthlySalary-87000)*1000*0.45
			}
			TaxVal *= float64(qualFamilyNum - 11)
		}

		fmt.Printf("입력하신 데이터를 바탕으로 계산한 공제대상가족의 수는 %d명 입니다.\n", qualFamilyNum)
		fmt.Printf("이때의 원천징수세액은 %d원(소득세 %d원, 지방소득세 %d원) 입니다.\n\n", int64(TaxVal*1.1), int64(TaxVal), int64(TaxVal*0.1))

		fmt.Printf("거래를 진행 중 입니다. \n")
		var whetherEnd string
		fmt.Scan(&whetherEnd)
	}
}
