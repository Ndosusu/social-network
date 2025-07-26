package types

type User struct {
	Id          int
	Email       string
	Nickname    string
	FirstName   string
	LastName    string
	BirthDate   string
	About       string
	Avatar      *string
	CreatedDate string
	PrivateMode bool
	IsClient    bool
}

type Session struct {
	Id        int
	Uuid      string
	User      *User
	CreatedAt string
	ExpiresAt string
}

type Post struct {
	Id          int
	Author      *User
	Message     string
	Image       *string
	Date        string
	PrivacyMode int
	Group       *Group
}

type PostFeed struct {
	Post         *Post
	Like         *Like
	LikeCount    int
	CommentCount int
	GroupTitle   string
}

type Comment struct {
	Id      int
	Author  *User
	Post    *Post
	Message string
	Image   *string
	Date    string
}

type CommentFeed struct {
	Comment   *Comment
	Like      *Like
	LikeCount int
}

type Like struct {
	Id      int
	Author  *User
	Post    *Post
	Comment *Comment
}

type Notif struct {
	Id           int
	NotifType    int
	Receiver     *User
	Sender       *User
	Group        *Group
	Event        *Event
	DateCreation string
}

type Group struct {
	Id           int
	Admin        *User
	Title        string
	About        string
	DateCreation string
}

type Event struct {
	Id           int
	Group        *Group
	Title        string
	About        string
	DateSchedule string
	DateCreation string
}

type Chat struct {
	Id       int
	Receiver *User
	Sender   *User
	Group    *Group
}

type Log struct {
	Id      int
	Author  *User
	Chat    *Chat
	Message string
	Date    string
}
