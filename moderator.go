package geddit


type Child struct {
	Name string `json:"name"`
}

type Moderator struct {
	Children []Child `json:"children"`
}
