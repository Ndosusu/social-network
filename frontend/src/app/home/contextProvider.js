"use client"

import { createContext, useContext, useState } from "react"

const HomeContext = createContext()

export function HomeProvider({children}) {
    const parseState = (tabState) => {
        return {
            val: tabState[0],
            set: tabState[1]
        }
    }

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

        //represent the list of info messages (WIP)
        infoMessages: parseState(useState([]))
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

