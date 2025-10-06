"use client"

import { useEffect } from "react"

export function CheckLogToken(router) {
    useEffect(() => {
        let token = localStorage.getItem("logToken")
        if(!token) {
            router.push("/")
        }
    })
}

export function CheckApiResponse(response) {
    if(response && response.data && response.data.Result) {
        return true
    }
    console.log("ERROR RESPONSE : ", response)
    return false
}

//small function to make state creation faster
export function parseState(tabState) {
    return {
        val: tabState[0],
        set: tabState[1]
    }
}

export function formatDate(date) {
    return date.split(" ")[0].split("-").reverse().join("/")
}