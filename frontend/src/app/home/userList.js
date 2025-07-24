"use client"

import { useState } from "react"

export default function PrivateUserList() {
    const [userList, setUsers] = useState([])

    //fetch users
    const tempUser = [
        {
            id:"5",
            name:"wiz",
        },
        {
            id:"8",
            name:"ziw",
        },
        {
            id:"9",
            name:"ziw",
        },
        {
            id:"10",
            name:"ziw",
        },
        {
            id:"11",
            name:"ziw",
        },
        {
            id:"12",
            name:"ziw",
        },
        {
            id:"13",
            name:"ziw",
        },
        {
            id:"14",
            name:"ziw",
        },
        {
            id:"15",
            name:"ziw",
        },
        {
            id:"16",
            name:"ziw",
        },
        {
            id:"17",
            name:"ziw",
        },
        {
            id:"18",
            name:"ziw",
        },
        {
            id:"19",
            name:"ziw",
        }
    ]

    return (
        <div className="w-1/5 h-9/10 neon-xl rounded-xl flex flex-col pointer-events-auto p-2">
            <div className="max-h-1/2 h-fit flex flex-col">
                <p>Selected users :</p>
                <div id="selectedList" className="overflow-scroll rounded-xl">
                    
                </div>
            </div>
            <div className="min-h-1/2 flex flex-col">
                <p>Available users :</p>
                <div id="userList" className="overflow-scroll rounded-xl">
                    {tempUser.map((obj, i) => <CreateUserCheckbox user={obj} key={i} />)}
                </div>
            </div>
        </div>
    )
}

function CreateUserCheckbox({user}) {
        return (
            <label id="userCheckBox" htmlFor={"checkbox"+user.id} className="w-full flex flex-row p-3 hover:bg-hovered text-xl items-center gap-3 select-none">
                <input name="followers_id" form="newPostForm" value={user.id} id={"checkbox"+user.id} type="checkbox" className="hidden" onClick={async (event) => {
                    const div = event.target
                    if(!div.checked) {
                        document.getElementById("userList").append(div.parentNode)
                    } else {
                        document.getElementById("selectedList").append(div.parentNode)
                    }
                }} />
                <img src="discord.svg" className="bg-discord h-10 rounded-xl" />
                <p>{user.name}</p>
            </label>
        )
    }