"use client"

export function newInfoMessage(message, ...classes) {
    const newInfo = document.createElement("div")
    newInfo.classList.add(..."infoMessage rounded-xl neon-sm w-full text-center p-8".split(" ").concat(classes.length > 0 ? classes : "bg-primaryT"))
    newInfo.addEventListener("click", (e) => {e.target.remove()})
    newInfo.addEventListener("animationend", (e) => {e.target.remove()})
    newInfo.textContent = message

    document.getElementById("infoMessagesDiv").prepend(newInfo)
}

export function CreateAllInfoMessages({infoState}) {
    if(!infoState.val) {
        return
    }

    return (
        <div id="infoMessagesDiv" className="absolute top-0 inset-x-1/2 -translate-x-1/2 w-1/4 flex flex-col gap-5 z-90 p-3">
        </div>
    )
}
