package pets

type AddPetCmd struct {
	Name string
	Tags []string
}

type Added struct {
	Name string
	Tags []string
}

type Deleted struct{}
