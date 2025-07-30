"use client"

export let InfoMessages = []

export function newInfoMessage(message, ...classes) {
    const obj = {
        message: message,
        classes: classes,
    }

    InfoMessages.push(obj)
}

export function CreateAllInfoMessages({allInfoMessages}) {
    if(!allInfoMessages) {
        return
    }

    return (
        <div className="absolute top-0 inset-x-1/2 -translate-x-1/2 w-1/4 flex flex-col gap-5 z-90 p-3">
            {allInfoMessages.map((obj, i) => <CreateInfoMessage info={obj} key={i} /> )}
        </div>
    )
}

function CreateInfoMessage({info}) {
    return (
        <div className={"infoMessage rounded-xl neon-sm w-full text-center p-8 " + (info.classes.length > 0 ? info.classes.join(" ") : "bg-primaryT")} onAnimationEnd={(e) => e.target.remove()}>
            <p>{info.message}</p>
        </div>
    )
}