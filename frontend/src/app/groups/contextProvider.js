"use client"

import { createContext, useEffect, useState } from "react"
import { CheckApiResponse, parseState } from "../utils"
import { DEFAULT_SERVER_PATH } from "../page"

const GroupContext = createContext()

export function GroupProvider({children}) {
    const States = {
        //the list of groups the user is a part of
        groupList: parseState(useState([])),

        //selected group
        curGroup: parseState(useState(null)),

        //chat history of the group
        chatHistory: parseState([]),

        //event history of the group
        eventHistory: parseState([]),

        //post history of the group
        postHistory: parseState([]),
    }

    useEffect(() => {
        fetch(DEFAULT_SERVER_PATH + "groups/list", {
            method: "POST",
            body: JSON.stringify({
                session_uuid: localStorage.getItem("logToken"),
            })
        })
        .catch(error => {
            console.log(error)
            throw new Error(error)
        })

        .then(data => data.json())

        .then(response => {
            console.log(response)
            if(CheckApiResponse(response)) {
                States.groupList.set(response.data.Result)
            }
        })
    }, [])

    return (
        <GroupContext.Provider value={States} >
            {children}
        </GroupContext.Provider>
    )
}