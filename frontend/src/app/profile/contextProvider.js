"use client"

import { createContext, useContext, useState } from "react"
import { parseState } from "../utils"

const ProfileContext = createContext()

export function ProfileProvider({children}) {
    const States = {
        //user to show the profile of
        curUser: parseState(useState(null)),

        //user's created posts list
        postList: parseState(useState([])),

        //list of users following the current user
        followedList: parseState(useState([])),

        //list of users followed by the current user
        followingList: parseState(useState([])),

        //showed modal
        modal: parseState(useState("")),
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