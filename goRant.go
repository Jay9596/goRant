//Package goRant implements a API-wrapper
//for the public API of http://www.devrant.io
package goRant

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

//All Constants are written in uppercase, and have no need to be exported
const (
	API      = "https://www.devrant.io/api"
	VERSION  = 3
	RANTS    = "%s/devrant/rants?sort=%s&limit=%d&skip=%d&app=%d"
	RANT     = "%s/devrant/rants/%d?app=%d"
	USERID   = "%s/get-user-id?username=%s&app=%d"
	USER     = "%s/users/%d?app=%d"
	SEARCH   = "%s/devrant/search?term=%s&app=%d"
	SURPRISE = "%s/devrant/rants/surprise?app=%d"
	WEEKLY   = "%s/devrant/weekly-rants?app=%d"
	COLLABS  = "%s/devrant/collabs?app=%d"
	STORIES  = "%s/devrant/story-rants?app=%d"
)

//Client in an empty interface
// It is has all functionalitios of goRant Wrapper
type Client struct {
}

//Rants return an array of Rants from devRant.io
//
//It has 3 parameters: Sort, Limit, Skip
// Sort can take 3 values, "algo", "recent", "top"
// Limit is an int value, used to specify number of rants to get
// Skip is an int value, used to specify number of rants to skip
// Example:
// If you eant the recent 20 rants: Rants("recent",20,0)
// If you want the next 20 rants: Rants("recent",20,20)
func (c *Client) Rants(sort string, limit int, skip int) ([]Rant, error) {
	if sort == "" {
		sort = "algo"
	} else if (sort != "algo") && (sort != "recent") && (sort != "top") {
		return nil, errors.New("Invalid sort method")
	}
	if limit <= 0 {
		limit = 50
	}

	url := fmt.Sprintf(RANTS, API, sort, limit, skip, VERSION)
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln("error", err)
	}
	var data Rants
	json.NewDecoder(res.Body).Decode(&data)
	if !data.Success && data.Error != "" {
		return nil, errors.New(data.Error)
	}
	return data.Rants, nil

}

//GetRant returns a singel Rant
// Just pass a rant-id as parameter
func (c *Client) GetRant(id int) (Rant, []Comment, error) {
	url := fmt.Sprintf(RANT, API, id, VERSION)
	res, err := http.Get(url)
	if err != nil {
		return Rant{}, nil, err
	}
	var data RantRes
	json.NewDecoder(res.Body).Decode(&data)
	if !data.Success && data.Error != "" {
		return Rant{}, nil, errors.New(data.Error)
	}
	return data.Rant, data.Comments, nil
}

//Profile returns a User profile info
// It takes a username of type string as input
func (c *Client) Profile(username string) (User, error) {
	id, err := getUserID(username)
	if err != nil {
		return User{}, err
	}
	//fmt.Println("id: ", id)
	url := fmt.Sprintf(USER, API, id, VERSION)
	res, err := http.Get(url)
	if err != nil {
		return User{}, err
	}
	var data UserRes
	json.NewDecoder(res.Body).Decode(&data)
	if !data.Success && data.Error != "" {
		return User{}, errors.New(data.Error)
	}
	return data.Profile, nil
}

//Used to get userID from a given name
// Only used internally for Profile func
func getUserID(username string) (int, error) {
	url := fmt.Sprintf(USERID, API, username, VERSION)
	res, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	var userID UserIDRes
	json.NewDecoder(res.Body).Decode(&userID)
	if !userID.Success && userID.Error != "" {
		return 0, errors.New(userID.Error)
	}
	return userID.UserID, nil
}

//Search returns array of Rants as the search result
func (c *Client) Search(query string) ([]Rant, error) {
	url := fmt.Sprintf(SEARCH, API, query, VERSION)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var data SearchRes
	json.NewDecoder(res.Body).Decode(&data)
	if !data.Success && data.Error != "" {
		return nil, errors.New(data.Error)
	}
	return data.Rants, nil
}

//Surprise returns a random Rant
func (c *Client) Surprise() (Rant, []Comment, error) {
	url := fmt.Sprintf(SURPRISE, API, VERSION)
	res, err := http.Get(url)
	if err != nil {
		return Rant{}, nil, nil
	}
	var data RantRes
	json.NewDecoder(res.Body).Decode(&data)
	if !data.Success && data.Error != "" {
		return Rant{}, nil, errors.New(data.Error)
	}
	return data.Rant, data.Comments, nil
}

//WeeklyRant return an array of Rants from weekly rants tab of devRant
func (c *Client) WeeklyRant() ([]Rant, error) {
	url := fmt.Sprintf(WEEKLY, API, VERSION)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var data Rants
	json.NewDecoder(res.Body).Decode(&data)
	if !data.Success && data.Error != "" {
		return nil, errors.New(data.Error)
	}
	return data.Rants, nil
}

//Collabs return an array of Rants from the collabs tob of devRant
func (c *Client) Collabs() ([]Rant, error) {
	url := fmt.Sprintf(COLLABS, API, VERSION)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var data Rants
	json.NewDecoder(res.Body).Decode(&data)
	if !data.Success && data.Error != "" {
		return nil, errors.New(data.Error)
	}
	return data.Rants, nil
}

//Stories return an array of Rants from stories tab of devRant
func (c *Client) Stories() ([]Rant, error) {
	url := fmt.Sprintf(STORIES, API, VERSION)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var data Rants
	json.NewDecoder(res.Body).Decode(&data)
	if !data.Success && data.Error != "" {
		return nil, errors.New(data.Error)
	}
	return data.Rants, nil
}

//New returns a new Client
func New() *Client {
	return new(Client)
}
