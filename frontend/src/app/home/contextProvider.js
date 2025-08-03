"use client"

import { createContext, useContext, useEffect, useState } from "react"
import { DEFAULT_SERVER_PATH } from "../page"

//create context hook
const HomeContext = createContext()

//small function to make state creation faster
export const parseState = (tabState) => {
    return {
        val: tabState[0],
        set: tabState[1]
    }
}

//context provider component, gives access to the context to all child component as well as giving the context it's value
export function HomeProvider({children}) {

    const States = {
        //sets feed loading state
        feedLoading: parseState(useState(true)),

        //dictates which feed to fetch
        currentFeed: parseState(useState("global")),

        //represent the current's feed post list
        feedPosts: parseState(useState([])),

        //represent the post to show the details of
        curPost: parseState(useState(null)),

        //gives the name of the modal to show
        modal: parseState(useState("")),

        //determines whether to show to new comment form or the comment list
        commentInput: parseState(useState(false)),

        //represents the comments list
        commentList: parseState(useState(null)),
    }

    //reset post feed and refetch it
    useEffect(() => {
        console.log("Fetching posts for feed " + States.currentFeed.val + "...")
        States.feedPosts.set(null)
        States.feedLoading.set(true)
        //api call to get the correct post feed
        fetch(DEFAULT_SERVER_PATH + "feed/" + States.currentFeed.val,{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                session_uuid: localStorage.getItem("logToken"),
                limit: 15,
            })
        })
        .catch(error => {
            States.feedLoading.set(false)
            throw new Error(error)
        })

        //make data readable as json object
        .then(data => data.json())

        //if posts returned, set post list
        .then(response => {
            States.feedLoading.set(false)
            if(response.data) {  
                States.feedPosts.set(response.data.Result)
            } else {
                throw new Error("No data.")
            }
        })
    }, [States.currentFeed.val])

    //api call to get all comments of the given post
    useEffect(() => {
        if (!States.curPost.val) {
            return
        }
        console.log("Fetching comments for post #"+States.curPost.val.Post.Id)

        fetch(DEFAULT_SERVER_PATH + "feed/detail", {
            method: "POST",
            body: JSON.stringify({
                session_uuid: localStorage.getItem("logToken"),
                post_id: States.curPost.val.Post.Id,
                limit: 15,
            })
        })
        .catch(error => {
            throw new Error(error)
        })

        //make data readable as json object
        .then(data => data.json())

        //update comment list
        .then(response => {
            // setComList(response.data.Result)
            States.commentList.set(response.data.Result)
        })
    }, [States.curPost.val])

    return (
        <HomeContext.Provider value={States}>
            {children}
        </HomeContext.Provider>
    )
}

//get the context
export function useHomeContext() {
    return useContext(HomeContext)
}

