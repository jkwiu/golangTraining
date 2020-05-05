package dataStruct

func Hash(s string) int {

	h := 0
	A := 256
	B := 3571
	for i := 0; i < len(s); i++ {
		h = (h*A + int(s[i])) % B
	}
	return h
}

type keyValue struct {
	key   string
	value string
}

type Map struct {

	//Hash 충돌을 방지하기 위해 같은 값을 가진 key들을 다시 배열의 배열 속으로
	//3571개의 key array가 있고, 이건 다시 list로 들어가고, list의 값은 key value의 값이 들어간다.
	keyArray [3571][]keyValue
}

//빈 맵을 만들어서 반환
func CreatMap() *Map {
	return &Map{}
}

//key, value 값 추가
func (m *Map) Add(key, value string) {
	h := Hash(key)
	m.keyArray[h] = append(m.keyArray[h], keyValue{key, value})
}

//key, value 값 가져오기
func (m *Map) Get(key string) string {
	h := Hash(key)
	for i := 0; i < len(m.keyArray[h]); i++ {
		if m.keyArray[h][i].key == key {
			return m.keyArray[h][i].value
		}
	}
	return ""
}
