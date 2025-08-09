package models_user

import (
	"fmt"
	"social-network/pkg/db/models"
	"social-network/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

func (db *UserDB) InsertUser(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			email : string,
			first_name : string,
			last_name : string,
			password : string
			date_birth : string,
			nickname : string,
			avatar : string
			about : string,
		}
	*/

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(obj["password"].(string)), 12)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	stmt := "INSERT INTO users (email, password, first_name, last_name, date_birth, avatar, nick_name, about, date_creation, private_mode) VALUES (?,?,?,?,?,?,?,?,?,?);"
	result, err := db.Conn.Exec(stmt, obj["email"], passwordHash, obj["first_name"], obj["last_name"], obj["date_birth"], obj["avatar"], obj["nickname"], obj["about"], utils.GetCurrentTime(), false)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	newUserId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	response, err := db.InsertSession(int(newUserId))
	if err != nil {
		fmt.Println("Failed to create session for new user")
		return nil, err
	}

	return response, nil
}

func (db *UserDB) SelectUserById(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			user_id : int,
			client_id : int,
		}
	*/

	stmt := `SELECT
				u.id,
				u.first_name,
				u.last_name,
				COALESCE(u.nick_name, ''),
				COALESCE(u.avatar, ''),
				u.private_mode,
				CASE 
					WHEN s.user_id = u.id OR u.private_mode = 0 OR fr.user_from IS NOT NULL 
					THEN u.date_birth 
					ELSE '' 
				END,
				CASE 
					WHEN s.user_id = u.id OR u.private_mode = 0 OR fr.user_from IS NOT NULL 
					THEN u.about 
					ELSE '' 
				END,
				CASE 
					WHEN s.user_id = u.id OR u.private_mode = 0 OR fr.user_from IS NOT NULL 
					THEN u.date_creation 
					ELSE '' 
				END,
				CASE 
					WHEN s.user_id = u.id 
					THEN 1
					ELSE 0 
				END AS is_client,
				CASE 
					WHEN s.user_id != u.id AND fr.user_from IS NOT NULL 
					THEN 1 
					ELSE 0 
				END as is_follower,
				CASE 
					WHEN s.user_id != u.id AND fr2.user_to IS NOT NULL 
					THEN 1 
					ELSE 0 
				END as is_followed,
				(SELECT COUNT(*) FROM follow_rel WHERE user_to = u.id) AS followers_count,
            	(SELECT COUNT(*) FROM follow_rel WHERE user_from = u.id) AS following_count
			FROM users u
			LEFT JOIN sessions s ON s.user_id = ?
			LEFT JOIN follow_rel fr ON fr.user_from = s.user_id AND fr.user_to = u.id
			LEFT JOIN follow_rel fr2 ON fr2.user_from = u.id AND fr2.user_to = s.user_id
			WHERE u.id = ?;`

	result := db.Conn.QueryRow(stmt, obj["client_id"], obj["user_id"])

	var userID, followerCount, followingCount int
	var firstName, lastName, nickname, avatar, dateBirth, about, dateCreation string
	var privateMode, isClient, isFollower, isFollowed bool

	err := result.Scan(&userID, &firstName, &lastName, &nickname, &avatar, &privateMode, &dateBirth, &about, &dateCreation, &isClient, &isFollower, &isFollowed, &followerCount, &followingCount)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	type userProfile struct {
		User           models.User
		IsFollower     bool
		IsFollowed     bool
		FollowersCount int
		FollowingCount int
	}
	profile := userProfile{
		User: models.User{
			Id:          userID,
			FirstName:   firstName,
			LastName:    lastName,
			Nickname:    nickname,
			PrivateMode: privateMode,
			BirthDate:   dateBirth,
			About:       about,
			CreatedDate: dateCreation,
			IsClient:    isClient,
		},
		IsFollower:     isFollower,
		IsFollowed:     isFollowed,
		FollowersCount: followerCount,
		FollowingCount: followingCount,
	}

	if avatar != "" {
		profile.User.Avatar = &avatar
	}

	return &models.Response{Result: profile}, nil
}
