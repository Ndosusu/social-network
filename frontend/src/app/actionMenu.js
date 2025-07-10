"use client"

import { useRouter } from "next/navigation"

export default function ActionMenu() {
    const router = useRouter()

    const moveTo = async (event) => {
        router.push(event.target.getAttribute("target"))
    }

    const logOut = async () => {
        localStorage.removeItem("logToken")
        router.push("/")
    }  

    return (
        <div className="neon-xl h-fit w-12 absolute top-1/3 left-0">
            <button className="neon-sm w-10 h-10" target="home" onClick={moveTo}>H</button> {/*home page*/}
            <button className="neon-sm w-10 h-10" target="profile" onClick={moveTo}>P</button> {/*profile page*/}
            <button className="neon-sm w-10 h-10" target="search" onClick={moveTo}>S</button> {/*search page*/}
            <button className="neon-sm w-10 h-10" target="notifications" onClick={moveTo}>N</button> {/*notif page*/}
            <button className="neon-sm w-10 h-10" target="chats" onClick={moveTo}>C</button> {/*chat page*/}
            <button className="neon-sm w-10 h-10">CM</button> {/*chat modal*/}
            <button className="bg-red-500 w-10 h-10" onClick={logOut}></button>
        </div>
    )
}