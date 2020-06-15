package greet

type Greet struct {
	Name     string `json:"name,omitempty"`
	Location string `json:"location,omitempty"`
}
