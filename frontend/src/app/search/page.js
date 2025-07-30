"use client"

import { useRouter } from "next/navigation"
import ActionMenu from "../actionMenu"
import { CheckLogToken } from "../checkToken"

export default function SearchPage() {
    const router = useRouter()
    CheckLogToken(router)

    const result = [
        {
            nickname: "wiz",
            status: true,
        },
        {
            nickname: "ziw",
            status: false,
        }
    ]

    return (
        <div className="center w-9/10 h-full text-white flex flex-col gap-10 justify-around p-7">
            <input type="text" placeholder="Search..." className="w-full h-fit neon-xl bg-primaryT rounded-xl p-3 text-2xl" />
            <div className="neon-xl bg-primaryT w-full h-full rounded-xl text-2xl flex flex-col gap-4 p-6">
                {result.map((obj, i) => <CreateSearchResult user={obj} key={i} />)}
            </div>
            <ActionMenu />
        </div>
    )
}

function CreateSearchResult(data) {
    const user = data.user
    return (
        <div className="w-full h-fit">
            <div className="flex flex-row w-full gap-5">
                <img src="discord.svg" className="bg-discord h-15 w-15 rounded-xl" />
                <div className="flex flex-row justify-between items-center w-full">
                    <p>{user.nickname}</p>
                    <CheckStatus status={user.status} />
                </div>
            </div>
        </div>
    )
}

function CheckStatus(data) {
    const status = data.status
    if (status) {
        return (
            <div className="bg-green-500 h-5 w-5 rounded-full mr-5"/>
        )
    } else {
        return (
            <div className="bg-gray-500 h-5 w-5 rounded-full mr-5"/>
        )
    }
}