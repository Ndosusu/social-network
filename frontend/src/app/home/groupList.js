"use client"

import {useState} from "react"

export default function PrivateGroupList() {
    const [groupList, setGroupList] = useState([])

    //fetch groups
    const groups = [
        {
            id: "1",
            name: "thug shaker central",
        },
        {
            id: "2",
            name: "wizards' den",
        },
        {
            id: "3",
            name: "kawaii desu neeeeee",
        }
    ]

    return (
        <div className="w-1/5 h-9/10 neon-xl rounded-xl flex flex-col pointer-events-auto p-2">
            <div className="h-full w-full rounded-xl flex flex-col gap-1 overflow-scroll">
                {groups.map((obj, i) => <CreateGroupRadio group={obj} key={i} />)}
            </div>
        </div>
    )
}

function CreateGroupRadio({group}) {
    return (
        <label htmlFor={"radioGroup"+group.id} className="p-3 hover:bg-hovered rounded-xl">
            <input type="radio" name="group_id" id={"radioGroup"+group.id} value={group.id} form="newPostForm" className="hidden" />
            <p>{group.name}</p>
        </label>
    )
}