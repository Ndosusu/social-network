"use client"
import { useRouter, useSearchParams } from "next/navigation"
import { CheckLogToken } from "../checkToken"
import { useEffect, useState } from "react"
import ActionMenu from "../actionMenu"
import NewPostModal from "./newPostModal"
import DetailPostModal from "./detailPostModal"

export default function Home() {
    const router = useRouter()
    CheckLogToken(router)
    
    const [loading, setLoading] = useState(true)
    const [posts, setPosts] = useState(null)
    const [curDetail, setDetail] = useState(null)
    const [curModal, setModal]= useState("")
    
    useEffect(() => {
        fetch("http://localhost:8080/feed/global",{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                session_uuid: localStorage.getItem("logToken"),
                limit: 15,
            })
        })
        .catch(error => {
            setLoading(false)
            throw new Error(error)
        })

        .then(data => data.json())

        .then(result => {
            console.log(result)
            if(result.data) {
                setLoading(false)    
                setPosts(result.data.Result)
            } else {
                setLoading(false)
                throw new Error("No data.")
            }
        })
    }, [])

    const CheckState = () => {
        if (loading) {
            return <p>loading...</p>
        }
        if (!posts) {
            return <p>failed to fetch data.</p>
        }
        return posts.map((obj, i) => <CreatePost postFeed={obj} key={i} setDetail={setDetail} setModal={setModal} />)
    }

    const CreateModal = () => {
        return (
            <div id="modalDiv" className="w-screen h-screen absolute ">
                <div className="w-full h-full bg-black opacity-80 absolute z-5" onClick={async () => {setModal("")}}/>
                <CheckModalState />
            </div>
        )
    }

    const CheckModalState = () => {
        switch(curModal) {
            case "newPostModal": {
                return <NewPostModal />
            }

            case "detailModal": {
                return <DetailPostModal post={curDetail} />
            }
        }
    }

    return (
        <div className="text-white h-full w-full grid items-center text-xl">
            <div className="bg-primaryT h-6/4 w-2/3 neon-xl center grid items-center">
                <div className="w-full h-screen overflow-scroll flex flex-col items-center p-4 gap-7">
                    <CheckState />
                </div>
            </div>
            <div className="fixed neon-xl w-1/10 h-fit max-h-5/6 left-5/6 top-1/12 postAction p-7">
                <div className="neon-sm p-5 rounded-xl flex flex-col items-center" onClick={() => {setModal("newPostModal")}}>
                    <img src="/new.svg" className="h-max"></img>
                    <p className="text-sm text-center">New post</p>
                </div>
            </div>
            <ActionMenu />
            {
                curModal != "" 
                ? <CreateModal /> 
                : null
            }
        </div>
    )
}

export function CreatePost(data) {
    const postFeed = data.postFeed
    const post = postFeed.Post
    const setDetail = data.setDetail
    const setModal = data.setModal

    return (
        <div className="w-5/6 rounded-xl neon-sm duration-100 hoverable hover:scale-110" onClick={async () => {
            setDetail(post)
            setModal("detailModal")
        }}>
            <div className="w-full postHeader bg-primaryT p-2">
                {post.AuthorId || "no"}
            </div>
            <div className="w-full h-fit p-4">
                {post.Message || "no"}
            </div>
            <div className="p-3 flex w-full gap-4">
                <div className="min-w-1/10 flex items-center">
                    <img src="/like.svg" className="h-8"></img>
                    <p>{postFeed.LikeCount || "0"}</p>
                </div>
                <div className="min-w-1/10 flex items-center gap-1">
                    <img src="/comment.svg" className="h-8"></img>
                    <p>{postFeed.CommentCount || "0"}</p>
                </div>
            </div>
        </div>
    )
}