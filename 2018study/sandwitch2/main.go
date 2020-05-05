package main

import "fmt"

//오곡빵, 통밀빵 -> 빵 의 관계
type Breads interface {
	GetOneBread() OneBread
}

//빵을 String으로 연결
type OneBread interface {
	String() string
}

//허니잼, 개꿀잼을 -> 잼 의 관계
type Jams interface {
	GetOneSpoon() SpoonOfJam
}

//잼을 String으로 연결
type SpoonOfJam interface {
	String() string
}

//오곡빵
type OgokBread struct {
}

//오곡빵 하나
type OneOgokBread struct {
}

//오곡빵 하나를 집음
func (o *OgokBread) GetOneBread() OneBread {
	return &OneOgokBread{}
}

//오곡빵에 String값 입력
func (o *OneOgokBread) String() string {
	return "오곡빵"
}

//통밀빵
type TongmilBread struct {
}

type OneTongmilBread struct {
}

func (o *TongmilBread) GetOneBread() OneBread {
	return &OneTongmilBread{}
}

func (o *OneTongmilBread) String() string {
	return "통밀빵"
}

//허니잼
type HoneyJam struct {
}

type SpoonOfHoneyJam struct {
}

func (h *HoneyJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfHoneyJam{}
}

func (s *SpoonOfHoneyJam) String() string {
	return "허니잼 "
}

//개꿀잼
type DogHoneyJam struct {
}

type SpoonDogHoneyJam struct {
}

func (h *DogHoneyJam) GetOneSpoon() SpoonOfJam {
	return &SpoonDogHoneyJam{}
}

func (s *SpoonDogHoneyJam) String() string {
	return "개꿀잼 "
}

//샌드위치 제작
func MakeSandwitch(bread Breads, jam Jams) string {
	sandwitch := jam.GetOneSpoon().String() + bread.GetOneBread().String()
	return sandwitch
}

//시식
func EatSandwitch(sandwitch string) {
	hangul := HasConsonantSuffix(sandwitch)
	var toEatSandwitch string
	if hangul {
		toEatSandwitch = sandwitch + "을 먹었다"
	} else {
		toEatSandwitch = sandwitch + "를 먹었다"
	}
	fmt.Println(toEatSandwitch)
}

//한글 받침 판독기
var (
	start = rune(44032)
	end   = rune(55204)
)

func HasConsonantSuffix(s string) bool {
	numEnds := 28
	result := false
	for _, r := range s {
		if start <= r && r < end {
			index := int(r - start)
			result = index%numEnds != 0
		}
	}
	return result
}

func main() {
	bread := &OgokBread{}
	jam := &HoneyJam{}
	bread2 := &TongmilBread{}
	jam2 := &DogHoneyJam{}

	sandwitch := MakeSandwitch(bread, jam)
	sandwitch2 := MakeSandwitch(bread2, jam)
	sandwitch3 := MakeSandwitch(bread, jam2)
	sandwitch4 := MakeSandwitch(bread2, jam2)
	drink := "음료수"

	fmt.Printf("%v\n%v\n%v\n%v\n", sandwitch, sandwitch2, sandwitch3, sandwitch4)

	EatSandwitch(sandwitch)
	EatSandwitch(sandwitch2)
	EatSandwitch(sandwitch3)
	EatSandwitch(sandwitch4)
	EatSandwitch(drink)

}
