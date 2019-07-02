package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	bottleCapacityMl  int = 20000
	maxHotCapacityMl  int = 1000
	maxColdCapacityMl int = 2000
	waterOutAmountMl  int = 10000 //デフォルト200ml テストの為10000ml
)

type waterAmount struct {
	currentRefillBottoleAmountMl int
	currentHotTankAmountMl       int
	currentColdTankAmountMl      int
}

//水の準備　フィールドにある変数に水を入れる
func prepareWater(wa *waterAmount) {
	changeBottle(wa)
	wa.currentRefillBottoleAmountMl -= wa.currentHotTankAmountMl
	wa.currentHotTankAmountMl += maxHotCapacityMl
	wa.currentRefillBottoleAmountMl -= wa.currentColdTankAmountMl
	wa.currentColdTankAmountMl += maxColdCapacityMl

}

//文字の入力を促す 入力された文字が指定された文字と一致したらリターン
func scanWithRestrictions(useStrings ...string) (string, bool) {

	scanner := bufio.NewScanner(os.Stdin) //入力待ちの処理
	scanner.Scan()
	var inputStr string
	isError := false
	inputStr = scanner.Text()

	for _, useStr := range useStrings {
		if useStr == inputStr {
			return inputStr, isError
		}
	}
	fmt.Println("指定された文字を入力して下さい")
	isError = true
	return inputStr, isError

}

//ボトルを交換する
func changeBottle(wa *waterAmount) {
	wa.currentRefillBottoleAmountMl = bottleCapacityMl
}

//水を出す
func drainWater(waterTemperatur string, wa *waterAmount) {
	switch waterTemperatur {
	case "1":
		wa.currentColdTankAmountMl -= waterOutAmountMl
		wa.currentRefillBottoleAmountMl -= waterOutAmountMl
		wa.currentColdTankAmountMl += waterOutAmountMl
	case "2":
		wa.currentHotTankAmountMl -= waterOutAmountMl
		wa.currentRefillBottoleAmountMl -= waterOutAmountMl
		wa.currentHotTankAmountMl += waterOutAmountMl
	}

}

//水量チェック　水が空だったら補充を促す
func checkWaterAmountAndRefill(wa *waterAmount) {
	if wa.currentRefillBottoleAmountMl <= 0 {
		for {
			fmt.Println("水を補充して下さい。")
			fmt.Println("補充が完了したら「１」を入力して下さい。")
			useString := []string{"1"}
			_, isError := scanWithRestrictions(useString...)
			if !isError {
				break
			}
		}
		changeBottle(wa)
	}
}

//水量の表示
func printWatereAmount(wa *waterAmount) {
	fmt.Printf("ボトルの残量は%vmlです。\n", wa.currentRefillBottoleAmountMl)
}

func main() {
	var wa = new(waterAmount)
	//水を準備
	prepareWater(wa)
	for {
		checkWaterAmountAndRefill(wa)
		printWatereAmount(wa)
		fmt.Println("冷たい水の場合は[1]を、熱いお湯の場合は[2]を入力して下さい。")
		rawInputStr := []string{"1", "2"}
		inputStr, isError := scanWithRestrictions(rawInputStr...)
		if isError {
			continue
		}

		//水を出す
		drainWater(inputStr, wa)

	}
}
