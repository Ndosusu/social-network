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