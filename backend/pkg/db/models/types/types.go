package types

type User struct {
	Id           int
	Email        string
	Nickname     string
	First_name   string
	Last_name    string
	Birth_date   string
	About        string
	Avatar       string
	Created_date string
	Private_mode bool
}

type Session struct {
	Id        int
	Uuid      string
	UserId    int
	CreatedAt string
	ExpiresAt string
	User      *User
}

type Post struct {
	Id          int
	AuthorId    int
	Message     string
	Image       *string
	Date        string
	PrivacyMode int
	GroupId     int
	Author      *User
}

type PostFeed struct {
	Post         *Post
	LikeCount    int
	CommentCount int
	GroupTitle   string
}

type Comment struct {
	Id       int
	AuthorId int
	PostId   int
	Message  string
	Image    *string
	Date     string
	Author   *User
}

type CommentFeed struct {
	Comment   *Comment
	LikeCount int
}

type Like struct {
	Id        int
	UserId    int
	PostId    int
	CommentId int
	Author    *User
	Post      *Post
	Comment   *Comment
}

type Notif struct {
	Id           int
	NotifType    int
	ReceiverId   int
	SenderId     int
	GroupId      int
	EventId      int
	DateCreation string
	Receiver     *User
	Sender       *User
	Group        *Group
	Event        *Event
}

type Group struct {
	Id           int
	AdminId      int
	Title        string
	About        string
	DateCreation string
	Admin        *User
}

type Event struct {
	Id           int
	GroupId      int
	Title        string
	About        string
	DateSchedule string
	DateCreation string
	Group        *Group
}

type Chat struct {
	Id         int
	ReceiverId int
	SenderId   int
	GroupId    int
	Receiver   *User
	Sender     *User
	Group      *Group
}

type Log struct {
	Id       int
	ChatId   int
	AuthorId int
	Message  string
	Date     string
	Author   *User
}
