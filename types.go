package azuredevopssdk


type Project struct {
	Id				string	`json:"id"`
	Name    		string	`json:"name"`
	Description 	string	`json:"description"`
	Visibility 		string	`json:"visibility"`
	Capabilities	Capabilities `json:"capabilities,omitempty"`
}

type Capabilities struct {
	Versioncontrol Versioncontrol `json:"versioncontrol"`
	ProcessTemplate ProcessTemplate `json:"processTemplate"`
}

type Versioncontrol struct {
	SourceControlType string `json:"sourceControlType"`
}

type ProcessTemplate struct {
	TemplateTypeId	string `json:"templateTypeId"`
}

type ResponseProject struct {
	Id			string	`json:"id"`
	Status    	string	`json:"status"`
	Url 		string	`json:"url"`
}


type ResponseOperation struct {
	Id			string	`json:"id"`
	Status    	string	`json:"status"`
	Url 		string	`json:"url"`
}