package azuredevopssdk
 
import "fmt"
import "net/http"
import "encoding/json"
import "io/ioutil"
import "encoding/base64"
import "time"
import "bytes"

const baseURL string = "https://dev.azure.com/"
 
type Client struct {
	client *http.Client
	organization string
	encToken string
}



func NewClientWith(organization string, token string) (*Client, error) {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	
	return &Client{netClient, organization, basicAuth(":"+token)}, nil
}



func basicAuth(token string) string {
	auth := token
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func (s *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization","Basic " + s.encToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// fmt.Println(resp.StatusCode)
	if 200 != resp.StatusCode {
		if  resp.StatusCode == 203 {
			return nil, fmt.Errorf("%s", "BAD TOKEN")
		}else{
			return nil, fmt.Errorf("%s", body)
		}
	}
	return body, nil
}

func (s *Client) CreateProject(project Project) (string, error) {
	var a [1]interface{}
	a[0] = project
	projectBytes, _ := json.Marshal(a)
	projectReader := bytes.NewReader(projectBytes)
	url := fmt.Sprintf(baseURL+"%s/_apis/projects?api-version=5.0-preview.3",s.organization)
	fmt.Println(url)
	req, err := http.NewRequest("POST", url, projectReader)
	if err != nil {
		return "", err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return "", err
	}

	var resp RespProject
	json.Unmarshal(bytes, &resp)
	return resp.Id, nil
}

/*
func (s *Client) GetBuildDefinitions(project string) ([]byte, error) {
	url := fmt.Sprintf(baseURL+"%s/%s/_apis/build/definitions?api-version=5.0-preview.7", s.organization, project)
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
*/

/*func (s *Client) AddTodo(todo *Todo) error {
	url := fmt.Sprintf(baseURL+"/%s/todos", s.Username)
	fmt.Println(url)
	j, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	_, err = s.doRequest(req)
	return err
}*/