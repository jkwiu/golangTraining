package main

import "fmt"

// 아래와 같이 오렌지 잼을 추가했을 뿐인데 작업량이 어마어마하고 소스가 더러워진다.
func main() {
	// 절차적 프로그래밍
	// 1. 빵 두개를 꺼낸다.
	breads := GetBreads(2)
	// jam := &StrawberryJam{}
	jam := &OrangeJam{}
	// 2. 딸기잼 뚜껑을 연다.
	// OpenStrawberryJam(jam)
	OpenOrangeJam(jam)
	// 3. 딸기잼을 한스푼 뜬다.
	// spoon := GetOneSpoon(jam)
	spoon := GetOneOrangeJamSpoon(jam)
	// 4. 딸기잼을 빵에 바른다.
	// PutJamOnBread(breads[0], spoon)
	PutOrangeJamOnBread(breads[0], spoon)
	// 5. 빵을 덮는다.
	sandwitch := MakeSandwitch(breads)
	// 6. 완성
	fmt.Println(sandwitch.val)
}

type Bread struct {
	val string
}

type StrawberryJam struct {
	opened bool
}

type OrangeJam struct {
	opened bool
}

type SpoonOfStrawberry struct {
}

type SpoonOfOrangeJam struct {
}

type Sandwitch struct {
	val string
}

func GetBreads(num int) []*Bread {
	breads := make([]*Bread, num)
	for i := 0; i < num; i++ {
		breads[i] = &Bread{val: "bread"}
	}
	return breads
}

func OpenStrawberryJam(jam *StrawberryJam) {
	jam.opened = true
}

func OpenOrangeJam(jam *OrangeJam) {
	jam.opened = true
}

func GetOneSpoon(_ *StrawberryJam) *SpoonOfStrawberry {
	return &SpoonOfStrawberry{}
}

func GetOneOrangeJamSpoon(_ *OrangeJam) *SpoonOfOrangeJam {
	return &SpoonOfOrangeJam{}
}

func PutJamOnBread(bread *Bread, jam *SpoonOfStrawberry) {
	bread.val += " + Strawberry Jam"
}

func PutOrangeJamOnBread(bread *Bread, jam *SpoonOfOrangeJam) {
	bread.val += " + Orange Jam"
}

func MakeSandwitch(breads []*Bread) *Sandwitch {
	sandwitch := &Sandwitch{}
	for i := 0; i < len(breads); i++ {
		sandwitch.val += breads[i].val + " + "
	}
	return sandwitch
}
