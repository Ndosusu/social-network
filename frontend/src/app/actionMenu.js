"use client"

import { useRouter } from "next/navigation"

let opened = false

export default function ActionMenu() {
    const router = useRouter()

    const moveTo = async (event) => {
        router.push(event.target.getAttribute("target"))
    }

    const logOut = async () => {
        localStorage.removeItem("logToken")
        router.push("/")
    }

    const showMenu = () => {
        const menu = document.getElementById("ActionMenu")
        const arrow = document.getElementById("arrow")

        menu.classList.add("neon-xl")
        menu.classList.remove("-translate-x-1/1")
        menu.setAttribute("opened", true)
        opened = true
        arrow.setAttribute("src", "leftArrow.svg")
    }

    const hideMenu = () => {
        const menu = document.getElementById("ActionMenu")
        const arrow = document.getElementById("arrow")

        menu.classList.remove("neon-xl")
        menu.classList.add("-translate-x-1/1")
        opened = false
        menu.setAttribute("opened", false)
        arrow.setAttribute("src", "rightArrow.svg")
    }

    let defaultState = "-translate-x-1/1"
    if(opened) {
        defaultState = "neon-xl"
    }

    return (
        <div id="ActionMenu" opened="true" onMouseOut={hideMenu} onMouseOver={showMenu} className={"bg-primaryT h-fit w-fit absolute inset-y-1/2 -translate-y-1/2 left-0 flex flex-col p-4 gap-4 rounded-br-xl rounded-tr-xl duration-500 "+defaultState}>
            <button className="neon-sm w-12 h-12 bg-primaryT rounded-xl" target="home" onClick={moveTo}>H</button>
            <button className="neon-sm w-12 h-12 bg-primaryT rounded-xl" target="profile" onClick={moveTo}>P</button>
            <button className="neon-sm w-12 h-12 bg-primaryT rounded-xl" target="search" onClick={moveTo}>S</button>
            <button className="neon-sm w-12 h-12 bg-primaryT rounded-xl" target="notifications" onClick={moveTo}>N</button>
            <button className="neon-sm w-12 h-12 bg-primaryT rounded-xl" target="chats" onClick={moveTo}>C</button>
            <button className="neon-sm w-12 h-12 bg-primaryT rounded-xl">CM</button>
            <button className="neon-sm bg-red-500 w-12 h-12 rounded-xl" onClick={logOut}></button>
            <div id="ActionOpen" className="bg-secondary neon-sm w-1/3 h-1/5 absolute -right-1/3 rounded-br-xl rounded-tr-xl p-1">
                <img id="arrow" src="rightArrow.svg" className="h-full" />
            </div>
        </div>
    )
}