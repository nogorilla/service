package checkapp

import "encoding/json"

// Info represents information about the service.
type Info struct {
	Status     string `json:"status,omitempty"`
	Build      string `json:"build,omitempty"`
	Host       string `json:"host,omitempty"`
	Name       string `json:"name,omitempty"`
	PodIP      string `json:"podIP,omitempty"`
	Node       string `json:"node,omitempty"`
	Namespace  string `json:"namespace,omitempty"`
	GOMAXPROCS int    `json:"GOMAXPROCS,omitempty"`
}

// Encode implments the encoder interface.
func (app Info) Encode() ([]byte, error) {
	return json.Marshal(app)
}
