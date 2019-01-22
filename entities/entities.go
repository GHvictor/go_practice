package entities

type user struct {
	Name string
	email string
}

type Admin struct {
	user
	Right int
}
