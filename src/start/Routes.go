package start

import (
	"fmt"
	"github.com/thoas/go-funk"
)

type Result struct {
	RouteId         int
	StartRouteIndex int
	EndRouteIndex   int
}

func FindPath(startStation string, endStation string) (int, []Result) {

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
func InterpretResult(result []Result){
	for _,v:=range result {
		s1 := fmt.Sprintf("Route ID is %v",X[v.RouteId])
		s2:=fmt.Sprintf("Route is heading from %s to %s",X[v.RouteId][0],X[v.RouteId][len(X)])
		s3:=fmt.Sprintf("Starting station is  %s ,End station is %s",X[v.RouteId][v.StartRouteIndex],X[v.RouteId][v.EndRouteIndex])
		fmt.Println(s1)
		fmt.Println(s2)
		fmt.Println(s3)
	}
}