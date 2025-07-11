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

    const actionClicked = async () => {
        const div = document.getElementById("ActionOpen")
        const menu = document.getElementById("ActionMenu")
        const arrow = document.getElementById("arrow")

        if(div.getAttribute("opened") == "true") {
            menu.classList.remove("neon-xl")
            menu.classList.add("-translate-x-1/1")
            div.setAttribute("opened", false)
            arrow.setAttribute("src", "rightArrow.svg")
        } else {
            menu.classList.add("neon-xl")
            menu.classList.remove("-translate-x-1/1")
            div.setAttribute("opened", true)
            arrow.setAttribute("src", "leftArrow.svg")
        }
    }

    return (
        <div id="ActionMenu" className="bg-primaryT h-fit w-fit absolute inset-y-1/2 -translate-y-1/2 left-0 -translate-x-1/1 flex flex-col p-4 gap-4 rounded-br-xl rounded-tr-xl duration-500">
            <button className="neon-sm w-12 h-12 bg-primaryT rounded-xl" target="home" onClick={moveTo}>H</button>
            <button className="neon-sm w-12 h-12 bg-primaryT rounded-xl" target="profile" onClick={moveTo}>P</button>
            <button className="neon-sm w-12 h-12 bg-primaryT rounded-xl" target="search" onClick={moveTo}>S</button>
            <button className="neon-sm w-12 h-12 bg-primaryT rounded-xl" target="notifications" onClick={moveTo}>N</button>
            <button className="neon-sm w-12 h-12 bg-primaryT rounded-xl" target="chats" onClick={moveTo}>C</button>
            <button className="neon-sm w-12 h-12 bg-primaryT rounded-xl">CM</button>
            <button className="neon-sm bg-red-500 w-12 h-12 rounded-xl" onClick={logOut}></button>
            <div id="ActionOpen" className="bg-secondary neon-sm w-1/3 h-1/5 absolute -right-1/3 rounded-br-xl rounded-tr-xl p-1" onClick={actionClicked}>
                <img id="arrow" src="rightArrow.svg" className="h-full" />
            </div>
        </div>
    )
}