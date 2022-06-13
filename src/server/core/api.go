package core

type StandardResponse struct {
	APIVersion string            `json:"apiVersion"`
	Context    string            `json:"context"`
	Method     string            `json:"method"`
	Params     map[string]string `json:"params,omitempty"`
	Error      any               `json:"error,omitempty"`
	Data       any               `json:"data"`
}

type ShortResponse struct {
	Message string `json:"message"`
}
