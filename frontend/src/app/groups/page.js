"use client"

import { useRouter } from "next/navigation"
import { CheckLogToken } from "../utils"
import ActionMenu from "../actionMenu"
import { GroupProvider } from "./contextProvider"

export default function GroupContextWrapper() {
    return (
        <GroupProvider>
            <GroupPage/>
        </GroupProvider>
    )
}

export function GroupPage() {
    const router = useRouter()
    CheckLogToken(router)

    console.log("Rendering groups page...")

    return (
         <div className="w-full h-full flex flex-row gap-9 p-5 text-white">
            <div className="neon-xl bg-primaryT w-1/5 rounded-xl flex flex-col overflow-scroll p-4 gap-5">
                
            </div>
            <div className="neon-xl bg-primaryT h-full w-4/5 rounded-xl flex flex-col p-4 gap-4">
                <div className="h-full w-full flex flex-col gap-5 overflow-scroll rounded-xl p-4">
                    
                </div>
                <textarea id="chatInput" className="neon-sm bg-primaryT w-full rounded-xl text-xl p-3 h-15 break-all resize-none hidden" />
            </div>
            <ActionMenu />
        </div>
    )
}