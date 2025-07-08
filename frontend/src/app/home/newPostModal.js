"use client"

export default function NewPostModal() {
    return (
        <div id="newPostModal" className="w-screen h-screen absolute hidden ">
            <div className="w-full h-full bg-black opacity-80 absolute z-5" onClick={async () => {document.getElementById("newPostModal").classList.add("hidden")}}>
                
            </div>
            <div className="neon-xl bg-primaryT h-9/10 w-3/4 absolute z-10 inset-x-1/8 inset-y-1/20 rounded-xl">
                
            </div>
        </div>
    )
}