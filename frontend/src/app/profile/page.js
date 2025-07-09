"use client"

import ActionMenu from "../actionMenu"
import { CreatePost } from "../home/page"
import DetailPostModal from "../home/detailPostModal"

export default function ProfilePage() {
    const temp = [
        {
            Id:1,
            AuthorId: 3,
            Message: "tkt",
            nbLike: 0,
            nbCom: 1,
        },
        {
            Id:1,
            AuthorId: 3,
            Message: "tkt",
            nbLike: 0,
            nbCom: 1,
        }
    ]

    const hideModal = async () => {
        document.getElementById("modalDiv").classList.add("hidden")
        document.querySelectorAll(".modal").forEach(obj => {
            obj.classList.add("hidden")
        })
    }

    return (
        <div className="center w-full h-full text-white overflow-scroll">
            <div className="neon-xl w-9/10 h-fit p-7 rounded-xl center mt-7 flex flex-col gap-10">
                <div className="flex flex-col gap-3">
                    <div className="flex flex-row gap-6">
                        <img src="discord.svg" className="w-50 h-50 rounded-xl bg-primary"></img>
                        <div className="flex flex-col justify-between">
                            <p className="text-5xl h-fit">Pepiño</p>
                            <p className="text-3xl h-fit text-gray-400">Lotr Taré</p>
                            <p className="text-3xl h-fit">User since : 09/07/2025</p>
                            <p className="text-3xl h-fit">Born on : 09/07/2025</p>
                        </div>
                    </div>
                    <p className="text-3xl px-5">About me :</p>
                    <p className="text-2xl px-5 break-all">blablablablablabalabalabalabalabalabalabalabalabalabalabalabalabalabalabaalabalabalabalanalabalabalabalabalabalaa</p>
                    <div className="flex flex-row justify-around text-2xl p-10 text-center">
                        <div>
                            <p>Following</p>
                            <p>200</p>
                        </div>
                        <div>
                            <p>Followers</p>
                            <p>100</p>
                        </div>
                        <div>
                            <p>Posts</p>
                            <p>15</p>
                        </div>
                    </div>
                </div>
            </div>
            <div className="neon-xl w-8/10 min-h-5 h-fit mt-10 rounded-xl center py-5 flex flex-col items-center gap-8">
                {temp.map((obj, i) => <CreatePost post={obj} key={i} />)}
            </div>
            <ActionMenu />
            <div id="modalDiv" className="w-screen h-screen absolute hidden top-0">
                <div className="w-full h-full bg-black opacity-80 absolute z-5" onClick={hideModal}/>
                <DetailPostModal />
            </div>
        </div>
    )
}