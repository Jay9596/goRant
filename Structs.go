package goRant

//Rants is struct for recieving Rants response
type Rants struct {
	Success  bool   `json:"success"`
	Error    string `json:"error"`
	Rants    []Rant `json:"rants"`
	Settings string `json:"settings"`
	Set      string `json:"set"`
	Wrw      int    `json:"wrw"`
	News     News   `json:"news"`
}

//Rant is a struct to save Rant data
type Rant struct {
	ID            int      `json:"id"`
	Text          string   `json:"text"`
	Score         int      `json:"score"`
	CreatedTime   int      `json:"created_time"`
	AttachedImage Image    `json:"attached_image"`
	NumComments   int      `json:"num_comments"`
	Tags          []string `json:"tags"`
	VoteState     int      `json:"vote_state"`
	Edited        bool     `json:"edited"`
	UserID        int      `json:"user_id"`
	Username      string   `json:"user_username"`
	UserScore     int      `json:"user_score"`
}

//RantRes is a struct to recieve Rant response
type RantRes struct {
	Success  bool      `json:"success"`
	Error    string    `json:"error"`
	Rant     Rant      `json:"rant"`
	Comments []Comment `json:"comments"`
}

//Comment is a struct to store comment data
type Comment struct {
	ID          int    `json:"id"`
	RantID      int    `json:"rant_id"`
	Body        string `json:"body"`
	Score       int    `json:"score"`
	CreatedTime int    `json:"created_time"`
	UserID      int    `json:"user_id"`
	Username    string `json:"user_username"`
	UserScore   int    `json:"user_score"`
}

//User is a struct to store User data
type User struct {
	Username    string `json:"username"`
	Score       int    `json:"score"`
	About       string `json:"about"`
	Location    string `json:"location"`
	CreatedTime int    `json:"created_time"`
	Skills      string `json:"skills"`
	Github      string `json:"github"`
	Website     string `json:"website"`
	Content     struct {
		Content Content `json:"content"`
		Counts  Counts  `json:"counts"`
	} `json:"content"`
	DPP int `json:"dpp"`
}

//Content is a struct to save Content of a User
// It has rants,upvotes,comments,favoutites of a user
type Content struct {
	Rants     []Rant    `json:"rants"`
	Upvoted   []Rant    `json:"upvoted"`
	Comments  []Comment `json:"comments"`
	Favorites []Rant    `json:"favorites"`
}

//Counts is a struct to save Counts of a User
// It has count(int values) of Rants, Upvotes, Comments, Favourites, and Colabs of a user
type Counts struct {
	Rants      int `json:"rants"`
	Upvotes    int `json:"upvoted"`
	Comments   int `json:"comments"`
	Favourites int `json:"favourites"`
	Colabs     int `json:"colabs"`
}

//Image struct is used to store info about Images
// Any attached images in rants, and/or comments
// It has Height, Width, and URl of the image
type Image struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

//News struct is used to store info about any News from devRant
type News struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	Headline string `json:"headline"`
	Body     string `json:"body"`
	Footer   string `json:"footer"`
	Height   int    `json:"height"`
	Action   string `json:"action"`
}

//UserRes is a struct to recieve User response
type UserRes struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Profile User   `json:"profile"`
}

//SearchRes is a struct to recieve Search response
type SearchRes struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Rants   []Rant `json:"results"`
}

//UserIDRes is a struct to recieve UserId response
//that is used to fetch a User
type UserIDRes struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	UserID  int    `json:"user_id"`
}
