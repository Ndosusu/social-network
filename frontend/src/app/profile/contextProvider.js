"use client"

import { createContext, useContext, useEffect, useState } from "react"
import { CheckApiResponse, parseState } from "../utils"
import { DEFAULT_SERVER_PATH } from "../page"

const ProfileContext = createContext()

export function ProfileProvider({children, id}) {
    const States = {
        //user to show the profile of
        curProfile: parseState(useState(null)),

        //user's created posts list
        postList: parseState(useState([])),

        //list of users following the current user
        followedList: parseState(useState([])),

        //list of users followed by the current user
        followingList: parseState(useState([])),

        //showed modal
        modal: parseState(useState("")),

        //determines the content of the box under the main profile div
        extraContent: parseState(useState("posts")),

        //determines which feed to fetch
        curFeed: parseState(useState("profile")),

        //selected post
        curPost: parseState(useState(null)),

        //list of comments for the selected post
        commentList: parseState(useState([])),

        //boolean to determine whether to show the new comment form
        commentInput: parseState(useState(false)),
    }

    //fetch given profile
    useEffect(() => {
        fetch(DEFAULT_SERVER_PATH + "profile", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                session_uuid: localStorage.getItem("logToken"),
                user_id: id,
            })
        })
        .catch(error => {
            console.log(error)
            throw new Error(error)
        })

        .then(data => data.json())

        .then(response => {
            if(CheckApiResponse(response)) {
                States.curProfile.set(response.data.Result)
            }
        })
    }, [])

    //fill postList
    useEffect(() => {
        if(!States.curProfile.val) {
            return
        }

        //api call to get all post of curUser
        fetch(DEFAULT_SERVER_PATH + "feed/" + States.curFeed.val, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                session_uuid: localStorage.getItem("logToken"),
                user_id: States.curProfile.val.User.Id,
                limit: 15,
            })
        })
        .catch(error => {
            console.log(error)
            States.postList.set([])
            throw new Error(error)
        })

        .then(data => data.json())

        .then(response => {
            if(CheckApiResponse(response)) {
                States.postList.set(response.data.Result)
            }
        })
    }, [States.curProfile.val, States.curFeed.val])

    //get list of people that follow you
    useEffect(() => {
        
    }, [States.curProfile.val])

    //get list of poeple the user follow
    useEffect(() => {

    }, [States.curProfile.val])

    return (
        <ProfileContext.Provider value={States}>
            {children}
        </ProfileContext.Provider>
    )
}

export function useProfileContext() {
    return useContext(ProfileContext)
}