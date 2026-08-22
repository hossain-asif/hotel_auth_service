package dto

type PermissionCreateRequestPayload struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Resource    string  `json:"resource"`
	Action      string  `json:"action"`
}

type PermissionCreateResponsePayload struct {
	Name        string  `json:"name"`
}

type PermissionUpadateRequestPayload struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Resource    string  `json:"resource"`
	Action      string  `json:"action"`
}

type PermissionUpdateResponsePayload struct {
	Name        string  `json:"name"`
}