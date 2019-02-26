package azuredevopssdk

type Processes struct {
	Count        int        `json:"count"`
	ProcesseList []Processe `json:"value,omitempty"`
}

type Processe struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsDefault   bool   `json:"isDefault,omitempty"`
	Type        string `json:"type,omitempty"`
}
