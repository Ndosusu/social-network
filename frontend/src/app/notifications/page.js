"use client"

import ActionMenu from "../actionMenu"

export default function NotificationPage() {

    const notifs = [ //replace with a fetch
        {
            message: "aaaa"
        },
        {
            message: "bbbbb"
        },
        {
            message: "bbbbb"
        },
        {
            message: "bbbbb"
        },
        {
            message: "bbbbb"
        },
        {
            message: "bbbbb"
        },
        {
            message: "bbbbb"
        },
        {
            message: "bbbbb"
        },
        {
            message: "bbbbb"
        },
        {
            message: "bbbbb"
        },
        {
            message: "bbbbb"
        },
        {
            message: "bbbbb"
        },
        {
            message: "bbbbb"
        },
        {
            message: "bbbbb"
        },
        {
            message: "bbbbb"
        },
        {
            message: "bbbbb"
        },
        {
            message: "bbbbb"
        },
        {
            message: "bbbbb"
        },
        {
            message: "bbbbb"
        },
        {
            message: "bbbbb2"
        }
    ]

    return (
        <div className="w-screen h-screen text-white flex justify-center items-center">
            <div className="neon-xl bg-primaryT w-9/10 h-9/10 rounded-xl text-xl p-5 overflow-hidden flex flex-col gap-4">
                <p className="font-bold">Notifications :</p>
                <div className="flex flex-col w-full h-full p-3 overflow-scroll gap-4">
                    {notifs.map((obj, i) => <CreateNotif notif={obj} key={i} />)}
                </div>
            </div>
            <ActionMenu />
        </div>
    )
}

function CreateNotif(data) { //Check notif type and act accordingly
    const notif = data.notif

    return (
        <div className="neon-sm bg-primaryT p-4 rounded-xl">
            <p>{notif.message}</p>
        </div>
    )
}