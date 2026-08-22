package dto

type RoleCreateRequestPayload struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type RoleCreateResponsePayload struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type RoleUpadateRequestPayload struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type RoleUpdateResponsePayload struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}