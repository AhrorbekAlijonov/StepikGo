package main

import (
	"fmt"
	"time"
)

func main() {
	const unixTime int64 = 1589570165

	var minutlar, sekundlar int
	fmt.Scanf("%d мин. %d сек.", &minutlar, &sekundlar)

	davomEtish := time.Duration(minutlar)*time.Minute + time.Duration(sekundlar)*time.Second

	asosiyVaqt := time.Unix(unixTime, 0).UTC()

	natijaVaqt := asosiyVaqt.Add(davomEtish)

	fmt.Println(natijaVaqt.Format("Mon Jan 2 15:04:05 MST 2006"))
}
