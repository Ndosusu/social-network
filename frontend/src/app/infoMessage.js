"use client"

export function newInfoMessage(tabInfo , message, ...classes) {
    const obj = {
        message: message,
        classes: classes,
    }
    console.log(tabInfo)

    tabInfo ? tabInfo.push(obj) : tabInfo = [obj]
    
    return tabInfo
}

export function CreateAllInfoMessages({tabInfo}) {
    if(!tabInfo) {
        return
    }

    return (
        <div className="absolute top-0 inset-x-1/2 -translate-x-1/2 w-1/4 flex flex-col gap-5 z-90 p-3">
            {tabInfo.map((obj, i) => <CreateInfoMessage info={obj} tabInfo={tabInfo} key={i} /> )}
        </div>
    )
}

function CreateInfoMessage({info, tabInfo}) {
    return (
        <div className={"infoMessage rounded-xl neon-sm w-full text-center p-8 " + (info.classes.length > 0 ? info.classes.join(" ") : "bg-primaryT")} onAnimationEnd={(e) => {
            e.target.remove()
            tabInfo.shift()
        }}>
            <p>{info.message}</p>
        </div>
    )
}