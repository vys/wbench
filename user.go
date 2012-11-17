package main

type User struct {
	Id      uint64
	Name    string
	Xp      uint64
	Level   uint64
	Items   []Item
	Friends []uint64
}

type Item struct {
	Id   uint64
	Type string
	Data string
}

func NewUser() *User {
	it := make([]Item, 20)
	for i, _ := range it {
		it[i].Id = 1000 + uint64(i)
		it[i].Type = "sometype"
		it[i].Data = "some data blah blah blah"
	}
	friends := make([]uint64, 50)
	for i, _ := range friends {
		friends[i] = uint64(i) + 10000000
	}
	user := &User{
		Id:      1292983,
		Name:    "Vinay",
		Xp:      100,
		Level:   200,
		Items:   it,
		Friends: friends,
	}
	return user
}
