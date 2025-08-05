"use client"

import { createContext, useContext, useState } from "react"
import { parseState } from "../utils"

const ProfileContext = createContext()

export function ProfileProvider({children}) {
    const tempUser = {
        User: {
            Id: 1,
            FirstName: "Lotr",
            LastName: "Taré",
            Nickname: "Pepiño",
            PrivateMode: 1,
            BirthDate: "24/02/2004",
            CreatedDate: "05/08/2025",
            isClient: true,
            About: "balabalabalabelebelebeleaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa aaaaaaaaaaaa aaaaaaaaaa aaaaaaaaaaaaa aaaaaaaaaaaaaa aaaaaaaaaaaaaa aaaaaaaaaaaa aaaaaaaaaaaa aaaaaaaaaaaa",
        },
        IsFollower: false,
        IsFollowed: false,
    }

    const tempPostList = [
        {
            Post: {
                Author: {
                    Avatar: "temp.png",
                    Nickname: "OtrPepiño",
                },
                Message: "test post"
            },
            LikeCount: 0,
            CommentCount: 0,            
        },
        {
            Post: {
                Author: {
                    Avatar: "temp.png",
                    Nickname: "OtrPepiño",
                },
                Message: "test post"
            },
            LikeCount: 0,
            CommentCount: 0,            
        },
        {
            Post: {
                Author: {
                    Avatar: "temp.png",
                    Nickname: "OtrPepiño",
                },
                Message: "test post"
            },
            LikeCount: 0,
            CommentCount: 0,            
        },
        {
            Post: {
                Author: {
                    Avatar: "temp.png",
                    Nickname: "OtrPepiño",
                },
                Message: "test post"
            },
            LikeCount: 0,
            CommentCount: 0,            
        },
        {
            Post: {
                Author: {
                    Avatar: "temp.png",
                    Nickname: "OtrPepiño",
                },
                Message: "test post"
            },
            LikeCount: 0,
            CommentCount: 0,            
        }
    ]

    const tempUserList = [
        {
            Id: 1,
            Nickname: "EnkorPepiño",
            Avatar: "temp.png"
        },
        {
            Id: 2,
            Nickname: "EnkorPepiño",
            Avatar: "temp.png"
        },
        {
            Id: 3,
            Nickname: "EnkorPepiño",
            Avatar: "temp.png"
        },
        {
            Id: 4,
            Nickname: "EnkorPepiño",
            Avatar: "temp.png"
        },
        {
            Id: 1,
            Nickname: "EnkorPepiño",
            Avatar: "temp.png"
        },
        {
            Id: 1,
            Nickname: "EnkorPepiño",
            Avatar: "temp.png"
        },
        {
            Id: 1,
            Nickname: "EnkorPepiño",
            Avatar: "temp.png"
        },
        {
            Id: 1,
            Nickname: "EnkorPepiño",
            Avatar: "temp.png"
        },
        {
            Id: 1,
            Nickname: "EnkorPepiño",
            Avatar: "temp.png"
        },
        {
            Id: 1,
            Nickname: "EnkorPepiño",
            Avatar: "temp.png"
        },
        {
            Id: 1,
            Nickname: "EnkorPepiño",
            Avatar: "temp.png"
        },
        {
            Id: 1,
            Nickname: "EnkorPepiño",
            Avatar: "temp.png"
        },
        {
            Id: 1,
            Nickname: "EnkorPepiño",
            Avatar: "temp.png"
        }
    ]

    const States = {
        //user to show the profile of
        curProfile: parseState(useState(tempUser)),

        //user's created posts list
        postList: parseState(useState(tempPostList)),

        //list of users following the current user
        followedList: parseState(useState(tempUserList)),

        //list of users followed by the current user
        followingList: parseState(useState(tempUserList)),

        //showed modal
        modal: parseState(useState("")),

        //determines the content of the box under the main profile div
        extraContent: parseState(useState("posts")),

        //determines which feed to fetch
        curFeed: parseState(useState("profile")),

        //selected post
        curPost: parseState(useState(null)),
    }

    return (
        <ProfileContext.Provider value={States}>
            {children}
        </ProfileContext.Provider>
    )
}

export function useProfileContext() {
    return useContext(ProfileContext)
}