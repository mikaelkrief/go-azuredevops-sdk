package myservice
 
const baseURL string = "https://www.myservice.com/v1"
 
type Client struct {
	Username string
	Password string
}
 
func NewBasicAuthClient(username, password string) *Client {
	return &Client{
		Username: username,
		Password: password,
	}
}

type Todo struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}
 
func (s *Client) AddTodo(todo *Todo) error {
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
}


func (s *Client) doRequest(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(s.Username, s.Password)
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
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}



func (s *Client) GetTodo(id int) (*Todo, error) {
	url := fmt.Sprintf(baseURL+"/%s/todos/%d", s.Username, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data Todo
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}