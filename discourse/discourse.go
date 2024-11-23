package discourse

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Discourse struct {
	Url           string
	ApiKey        string
	AdminUserName string
}

const (
	POST = "POST"
	GET  = "GET"
)

func (self Discourse) prepareGetReq(path string, isAdmin bool) *http.Request {
	req, err := http.NewRequest(GET, self.Url+path, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Api-Key", self.ApiKey)

	if isAdmin {
		req.Header.Set("Api-Username", self.AdminUserName)
	}

	return req
}

func (self Discourse) preparePostReq(
	path string,
	payload []byte,
	isAdmin bool,
) *http.Request {
	endpoint := self.Url + path
    req, err := http.NewRequest(POST, endpoint, bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Api-Key", self.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	if isAdmin {
		req.Header.Set("Api-Username", self.AdminUserName)
	}

	return req
}

func (self Discourse) CreateUser(payload *CreateDiscourseUserDto) ([]byte, error) {
	bytes, err := json.Marshal(CreateDiscourseUserBody{
		Email:    payload.Email,
		Username: payload.Username,
		Password: payload.Password,
		Active:   payload.Active,
	})

	if err != nil {
		return nil, err
	}

	req := self.preparePostReq("/users.json", bytes, true)
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
    
    bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
    
    if err != nil {
        return nil, err
    }
    
    return bodyBytes, nil
}

func (self Discourse) GetLatestPosts() ([]byte, error) {
    cache, _ := GetCache("latest-posts")
    
    if cache != "" {
        return []byte(cache), nil
    }

    req := self.prepareGetReq("/posts.json", true)
    client := &http.Client{}

    resp, err := client.Do(req)

    if err != nil {
        return nil, err
    }

    bodyBytes, err := io.ReadAll(resp.Body)
    defer resp.Body.Close()

    if err != nil {
        return nil, err
    }
    
    SetCache("latest-posts", string(bodyBytes))

    return bodyBytes, nil
}
