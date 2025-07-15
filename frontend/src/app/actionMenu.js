"use client"

import { useRouter } from "next/navigation"
import { useState } from "react"

let opened = false

export default function ActionMenu() {
    const router = useRouter()
    const [chatModals, setChatModal] = useState([])

    const moveTo = async (event) => {
        router.push(event.target.getAttribute("target"))
    }

    const logOut = async () => {
        localStorage.removeItem("logToken")
        router.push("/")
    }

    const showMenu = async () => {
        const menu = document.getElementById("ActionMenu")
        const arrow = document.getElementById("arrow")

        menu.classList.add("neon-xl")
        menu.classList.remove("-translate-x-1/1")
        menu.setAttribute("opened", true)
        opened = true
        arrow.setAttribute("src", "leftArrow.svg")
    }

    const hideMenu = async () => {
        const menu = document.getElementById("ActionMenu")
        const arrow = document.getElementById("arrow")

        menu.classList.remove("neon-xl")
        menu.classList.add("-translate-x-1/1")
        opened = false
        menu.setAttribute("opened", false)
        arrow.setAttribute("src", "rightArrow.svg")
    }

    const newChatModal = async () => {
        const mainContainer = document.getElementById("mainContainer")
        const modal = document.createElement("div")

        let selDiv = null
        let posX = 0
        let posY = 0

        modal.className = "absolute resize bg-white h-50 w-50 top-0 left-0 z-40"

        const setDrag = async (event) => {
            event.preventDefault()
            selDiv = event.target
            posX = event.clientX
            posY = event.clientY

            window.addEventListener("mousemove", followMouse)
        }

        const dragEnd = async (event) => {
            event.preventDefault()
            selDiv = null
            window.removeEventListener("mousemove", followMouse)
        }

        const followMouse = (event) => {
            if(selDiv) {
                let offsetX = posX - selDiv.style.left.replace("px", "")
                let offsetY = posY - selDiv.style.top.replace("px", "")
                posX = event.clientX
                posY = event.clientY

                selDiv.style.left = posX - offsetX +"px"
                selDiv.style.top = posY - offsetY +"px"

                console.log(offsetX, offsetY)
            }
        }

        modal.onmousedown = setDrag
        modal.onmouseup = dragEnd

        mainContainer.append(modal)
        setChatModal(chatModals.concat(modal))
        console.log(chatModals)
    }

    let defaultState = "-translate-x-1/1"
    if(opened) {
        defaultState = "neon-xl"
    }

    return (
        <div id="ActionMenu" opened="true" onMouseOut={hideMenu} onMouseOver={showMenu} className={"bg-primaryT h-fit w-fit absolute inset-y-1/2 -translate-y-1/2 left-0 flex flex-col p-4 gap-4 rounded-br-xl rounded-tr-xl duration-500 z-50 "+defaultState}>
            <button className="neon-sm w-12 h-12 bg-primaryT rounded-xl" target="home" onClick={moveTo}>H</button>
            <button className="neon-sm w-12 h-12 bg-primaryT rounded-xl" target="profile" onClick={moveTo}>P</button>
            <button className="neon-sm w-12 h-12 bg-primaryT rounded-xl" target="search" onClick={moveTo}>S</button>
            <button className="neon-sm w-12 h-12 bg-primaryT rounded-xl" target="notifications" onClick={moveTo}>N</button>
            <button className="neon-sm w-12 h-12 bg-primaryT rounded-xl" target="chats" onClick={moveTo}>C</button>
            <button className="neon-sm w-12 h-12 bg-primaryT rounded-xl" onClick={newChatModal}>CM</button>
            <button className="neon-sm bg-red-500 w-12 h-12 rounded-xl" onClick={logOut}></button>
            <div id="ActionOpen" className="bg-secondary neon-sm w-1/3 h-1/5 absolute -right-1/3 rounded-br-xl rounded-tr-xl p-1">
                <img id="arrow" src="rightArrow.svg" className="h-full" />
            </div>
        </div>
    )
}