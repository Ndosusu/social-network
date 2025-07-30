"use client"

import { createContext, useContext, useState } from "react"

const HomeContext = createContext()

export const parseState = (tabState) => {
    return {
        val: tabState[0],
        set: tabState[1]
    }
}

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

        //represent the list of info messages (WIP)
        infoMessages: parseState(useState([])),
    }

    return (
        <HomeContext.Provider value={States}>
            {children}
        </HomeContext.Provider>
    )
}

export function useHomeContext() {
    return useContext(HomeContext)
}

