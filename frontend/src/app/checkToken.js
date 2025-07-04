"use client"

import { useEffect, useState } from "react"

export function CheckLogToken(router) {
    useEffect(() => {
        let token = localStorage.getItem("logToken")
        if(!token) {
            router.push("/")
        }
    })
}