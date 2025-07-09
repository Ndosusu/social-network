"use client"

export default function DetailPostModal() {
    return (
        <div id="detailPostModal" postid={0} className="modal neon-xl bg-primaryT h-9/10 w-3/4 absolute z-10 inset-x-1/8 inset-y-1/20 rounded-xl hidden overflow-scroll">
            <div className="w-5/6 rounded-xl neon-sm">
                <div className="w-full postHeader bg-primaryT p-2">
                    <p id="detailAuthor"></p>
                </div>
                <div className="w-full h-fit p-4">
                    <p id="detailMessage"></p>
                </div>
                <div id="detailCommentList">

                </div>
            </div>
        </div>
    )
}

export function CreateCom(data) {
    return (
        <div>
            <p>{data.key}</p>
        </div>
    )
}