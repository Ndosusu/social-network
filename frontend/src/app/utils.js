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