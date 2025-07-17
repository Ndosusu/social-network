"use client"

import { useState } from "react"
import ActionMenu from "../actionMenu"
import { CheckLogToken } from "../checkToken"
import { useRouter } from "next/navigation"

export default function ChatPage() {
    const router = useRouter()
    CheckLogToken(router)

    const userList = [
        {
            uuid:0,
            nickname: "wiz",
            status: true,
        },
        {
            uuid:1,
            nickname: "ziw",
            status: false,
        }
    ]

    const [chatList, setChat] = useState(null)

    return (
        <div className="w-full h-full flex flex-row gap-9 p-5 text-white">
            <div className="neon-xl bg-primaryT w-1/5 rounded-xl flex flex-col overflow-scroll p-4 gap-5">
                {userList.map((obj, i) => <CreateUser user={obj} key={i} onclick={setChat} />)}
            </div>
            <div className="neon-xl bg-primaryT h-full w-4/5 rounded-xl flex flex-col p-4 gap-4">
                <div className="h-full w-full flex flex-col gap-5 overflow-scroll rounded-xl p-4">
                    <CreateMessageHistory chatList={chatList} />
                </div>
                <textarea id="chatInput" className="neon-sm bg-primaryT w-full rounded-xl text-xl p-3 h-15 break-all resize-none hidden" />
            </div>
            <ActionMenu />
        </div>
    )
}

//Change message history creation to be on userlist object clicked
function CreateUser(data) {
    const chatList = [ //temporary chat list, to be replaced with a fetch
        {
            message: "okok",
            author: "ziw",
        },
        {
            message: "okok2",
            author: "ziw",
        },
        {
            message: "nonon",
            author: "wiz",
        },
        {
            message: "nonon",
            author: "wiz",
        },
        {
            message: "nonon",
            author: "wiz",
        },
        {
            message: "nonon",
            author: "wiz",
        },
        {
            message: "nonon",
            author: "wiz",
        },
        {
            message: "nonon",
            author: "wiz",
        },
        {
            message: "nonon",
            author: "wiz",
        },
        {
            message: "nonon",
            author: "wiz",
        },
        {
            message: "nonon",
            author: "wiz",
        },
        {
            message: "nonon",
            author: "wiz",
        },
        {
            message: "nonon",
            author: "wiz",
        },
        {
            message: "nonon",
            author: "wiz",
        },
        {
            message: "nonon",
            author: "wiz",
        },
        {
            message: "nonon",
            author: "wiz",
        },
        {
            message: "nonon",
            author: "wiz",
        },
        {
            message: "nonon",
            author: "wiz",
        },
        {
            message: "nonon",
            author: "wiz",
        }
    ]

    const user = data.user

    const userClicked = (chatList) => {
        const chatInput = document.getElementById("chatInput")
        chatInput.classList.remove("hidden")
        chatInput.setAttribute("placeholder", "Send a message to "+ user.nickname)
        data.onclick(chatList)
    }

    return (
        <div className="w-full h-fit text-xl flex gap-3 items-center" onClick={async () => {userClicked(chatList)}} >
            <img className="h-12 w-12 bg-discord rounded-xl" src="discord.svg" />
            <p className="h-fit">{user.nickname}</p>
        </div>
    )
}

//Make chat history fusion back to back messages from the same person
function CreateMessageHistory(data) {
    const chatList = data.chatList
    if (!chatList) {
        return <p className="text-2xl">No chat selected</p>
    } else if (chatList.length == 0) {
        return <p className="text-2xl">Start the conversation !</p>
    } else {
        return (
            chatList.map((obj, i) => <CreateMessage chat={obj} key={i} />)
        )
    }
}

function CreateMessage(data) {
    const chat = data.chat

    return (
        <div className="w-full h-fit text-xl flex flex-col">
            <div className="flex flex-row gap-3 items-center">
                <img src="discord.svg" className="bg-discord h-10 w-10 rounded-xl" />
                <p className="font-bold h-fit">{chat.author}</p>
            </div>
            <p>{chat.message}</p>
        </div>
    )
}