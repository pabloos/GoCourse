package greet

// Greet represents the parameters from a greet message
type Greet struct {
	Name     string `json:"name,omitempty"`
	Location string `json:"location,omitempty"`
}
