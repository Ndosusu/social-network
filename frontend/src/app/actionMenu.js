"use client"

import { useRouter } from "next/navigation"
import { useState } from "react"

let opened = false
let chatModals = []

export default function ActionMenu() {
    const router = useRouter()

    const [chatModalsNb, setChatModalNb] = useState(chatModals.length)

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
        const obj = {
            x: null,
            y: null,
            with: null,
        }

        setChatModalNb(chatModals.push(obj))
    }

    let defaultState = "-translate-x-1/1"
    if(opened) {
        defaultState = "neon-xl"
    }

    return (
        <div>
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
            <div>
                {chatModals.map((obj, i) => <CreateChatModal obj={obj} id={i} key={i} fn={setChatModalNb} />)}
            </div>
        </div>
    )
}

function CreateChatModal(data) {
    const obj = data.obj
    const key = data.id
    const fn = data.fn

    let selDiv = null
    let posX = 0
    let posY = 0
    let initialState = " inset-1/2 -translate-1/2"
    if(obj.x && obj.y) {
        initialState = ""
    }

    const setDrag = async (event) => {
        document.getSelection().removeAllRanges()
        selDiv = document.getElementById("chatModal"+key)
        posX = event.clientX
        posY = event.clientY

        if(!selDiv.style.left || !selDiv.style.top) {
            let rect = selDiv.getBoundingClientRect()
            selDiv.classList.remove("inset-1/2", "-translate-1/2")

            selDiv.style.left = posX - (posX - rect.left) + "px"
            selDiv.style.top = posY - (posY - rect.top) + "px"

            chatModals[key].x = selDiv.style.left
            chatModals[key].y = selDiv.style.top
        }

        window.addEventListener("mousemove", followMouse)
    }

    const dragEnd = async (event) => {
        event.preventDefault()
        selDiv = null
        window.removeEventListener("mousemove", followMouse)
    }

    const followMouse = (event) => {
        event.preventDefault()
        document.getSelection().removeAllRanges()
        if(selDiv) {
            let offsetX = posX - selDiv.style.left.replace("px", "")
            let offsetY = posY - selDiv.style.top.replace("px", "")

            posX = event.clientX
            posY = event.clientY

            selDiv.style.left = posX - offsetX +"px"
            selDiv.style.top = posY - offsetY +"px"

            chatModals[key].x = selDiv.style.left
            chatModals[key].y = selDiv.style.top
        }
    }

    const removeModal = () => {
        chatModals.splice(key, 1)
        fn(chatModals.length)
    }

    return (
        <div id={"chatModal"+key} className={"chatModal absolute bg-primaryT h-90 w-75 z-40 rounded-xl overflow-scroll flex flex-col neon-sm"+initialState} key={key} style={{top: obj.y, left: obj.x}}>
            <div className="bg-secondary w-full h-fit text-xl neon-sm p-1 flex flex-row justify-between items-center" onMouseDown={setDrag} onMouseUp={dragEnd} >
                <p className="w-fit cursor-default">User123</p>
                <img src="cross.svg" className="h-5" onClick={() => {removeModal(key)}} />
            </div>
            <div className="flex-grow"> {/* rajouter le contenu dynamique */}

            </div>
            <form id={"chatInput"+key} className="p-3 hidden">
                <textarea id="chatInput" className="neon-sm bg-primaryT w-full rounded-xl text-xl p-1 h-10 break-normal resize-none" placeholder="Send a message to User123" />
            </form>
        </div>
    )
}