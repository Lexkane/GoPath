package main

import (
	"fmt"
	"github.com/thoas/go-funk"
)
//var (
//	A map[int][]string
//)

func fill() map[int][]string{
	X := make(map[int][]string)
	X[1] = append(X[1], "IvanoFrankivsk", "Striy", "Hodoriv", "Ternopil", "Hmelnitskiy", "Jmerinka", "Slobidka", "Pomichna",
		"Kirovograd", "KriviyRig", "Zaporizzya", "Dnipro", "Konstantinivka")
	X[2] = append(X[1], "IvanoFrankivsk", "Striy", "Hodoriv", "Ternopil", "Hmelnitskiy", "Jmerinka", "Slobidka", "Pomichna",
		"Kirovograd", "KriviyRig", "Zaporizzya", "Dnipro", "Konstantinivka")
	X[12] = append(X[12], "Lviv", "Krasne", "Ternopil", "Hmelnitskiy", "Jmerinka", "Slobidka", "Rozdilna", "Odessa")
	X[13] = append(X[13], "Odessa", "Rozdilna", "Slobidka", "Jmerinka", "Hmelnitskiy", "Ternopil", "Krasne", "Lviv")
	X[17] = append(X[17], "Uzgorod", "Sambir", "Lviv", "Krasne", "Ternopil", "Shepetivka", "Novograd-Volinskiy", "Jitmir", "Fastiv", "Kyiv", "Hrebinka", "Romodan", "Poltava", "Harkiv")
	X[18] = append(X[18], "Harkiv", "Poltava", "Romodan", "Hrebinka", "Kyiv", "Fastiv", "Jitmir", "Novograd-Volinskiy", "Shepetivka", "Ternopil", "Krasne", "Lviv", "Sambir", "Uzgorod")
	X[42] = append(X[42], "Truskavets", "Lviv", "Krasne", "Ternopil", "Shepetivka", "Novograd-Volinskiy", "Jitmir", "Fastiv", "Kyiv", "Hrebinka",
		"Romodan", "Poltava", "Krasnograd", "Dnipro")
	X[43] = append(X[43], "Dnipro", "Krasnograd", "Poltava", "Hrebinka",
		"Romodan", "Kyiv", "Fastiv", "Jitmir", "Novograd-Volinskiy", "Shepetivka", "Ternopil", "Krasne", "Lviv", "Truskavets")
	X[97] = append(X[97], "Kyiv", "Korosten", "Sarni", "Kovel")
	X[110] = append(X[110], "Lviv", "Krasne", "Ternopil", "Hmelnitskiy", "Jmerinka", "Slobidka", "Pomichne", "Kolosivka", "Mikolayiv")
	X[111] = append(X[111], "Mikolayiv", "Kolosivka", "Pomichne", "Slobidka", "Jmerinka", "Hmelnitskiy", "Ternopil", "Krasne", "Lviv")
	X[120] = append(X[120], "Lviv", "Krasne", "Ternopil", "Shepetivka", "Novograd-Volinskiy", "Jitmir", "Fastiv", "Kyiv", "Hrebinka",
		"Romodan", "Poltava", "Krasnograd", "Dnipro", "Zaporizhzia")
	X[121] = append(X[121], "Zaporizhzia", "Dnipro", "Krasnograd", "Poltava", "Hrebinka", "Romodan", "Kyiv", "Fastiv", "Jitmir", "Novograd-Volinskiy", "Shepetivka", "Ternopil", "Krasne", "Lviv")
	X[136] = append(X[136], "Chernivtsi", "Kolomiya", "Ivano-Frankivsk", "Hodoriv", "Ternopil", "Hmelnitskiy", "Jmerinka", "Slobidka", "Rozdilna", "Odessa")
	X[137] = append(X[137], "Odessa", "Rozdilna", "Slobidka", "Jmerinka", "Hmelnitskiy", "Ternopil", "Hodoriv", "Ivano-Frankivsk", "Kolomiya", "Chernivtsi")
	X[142] = append(X[142], "Lviv", "Rava-Ruska", "Chervonograd", "Volodimir-Volinskiy", "Kovel")
	X[143] = append(X[143], "Kovel", "Volodimir-Volinskiy", "Chervonograd", "Rava-Ruska", "Lviv")
	X[606] = append(X[606], "Lviv", "Striy", "Ivano-Frankivsk", "Dilyatin", "Rahiv")
	X[607] = append(X[607], "Rahiv", "Dilyatin", "Ivano-Frankivsk", "Striy", "Lviv")
	X[726] = append(X[726], "Kyiv", "Nigin", "Konotop", "Vorozba", "Sumi", "Trostyanets", "Kharkiv")
	X[743] = append(X[743], "Kyiv", "Fastiv", "Jitomir", "Novograd-Volinskiy", "Shepetivka", "Ternopil", "Krasne", "Lviv")
	X[744] = append(X[744], "Lviv", "Krasne", "Ternopil", "Shepetivka", "Novograd-Volinskiy", "Jitomir", "Fastiv", "Kyiv")
	X[775] = append(X[755], "Kharkiv", "Trostyanets", "Sumi", "Vorozba", "Konotop", "Nigin", "Kyiv")
	X[780] = append(X[780], "Kyiv", "Nigin", "Konotop", "Vorozba", "Sumi")
	X[781] = append(X[781], "Sumi", "Vorozba", "Konotop", "Nigin", "Kyiv")
	X[806] = append(X[806], "Lviv", "Krasne", "Rivne")
	X[807] = append(X[807], "Rivne", "Krasne", "Lviv")
	return X
}

type Result struct {
	RouteId         int
	StartRouteIndex int
	EndRouteIndex   int
}
func FindPath( startStation string, endStation string,X map[int][] string,) (int, []Result) {

	var pathresult [] Result
	var counter int

	for key, value := range X {
		if indexStart := funk.IndexOf(value, startStation); indexStart != -1 {
			if indexEnd := funk.IndexOf(value, endStation); indexEnd != -1 {
				result := Result{key, indexStart, indexEnd}
				pathresult = append(pathresult, result)
				counter++
			}

		}
	}

	return counter, pathresult
}


func main() {
	X := fill()
	fmt.Print("Enter startstation ")
	var startPath string
	var endPath string
	fmt.Scanln(&startPath)
	fmt.Print("Enter endstation ")
	fmt.Scanln(&endPath)

	var pathresult [] Result
	var counter int

	for key, value := range X {
		if indexStart := funk.IndexOf(value, startPath); indexStart != -1 {
			if indexEnd := funk.IndexOf(value, endPath); indexEnd != -1 {
				result := Result{key, indexStart, indexEnd}
				pathresult = append(pathresult, result)
				counter++
			}

		}
	}

	for _, v := range pathresult {
		s1 := fmt.Sprintf("Route ID is %v", v.RouteId)
		s2 := fmt.Sprintf("Route is heading from %s to %s", X[v.RouteId][0], X[v.RouteId][len(X[v.RouteId])-1])
		s3 := fmt.Sprintf("Starting station is  %s ,End station is %s", X[v.RouteId][v.StartRouteIndex], X[v.RouteId][v.EndRouteIndex])
		fmt.Println(s1)
		fmt.Println(s2)
		fmt.Println(s3)
	}
	if (counter == 0) {
		fmt.Print("No matches found")
	}

}