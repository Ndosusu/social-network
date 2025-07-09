"use client"
import { useRouter } from "next/navigation"
import { CheckLogToken } from "../checkToken"
import { useEffect, useState } from "react"
import ActionMenu from "../actionMenu"
import NewPostModal from "./newPostModal"
import DetailPostModal, { CreateCom } from "./detailPostModal"

export default function Home() {
    const router = useRouter()
    const [loading, setLoading] = useState(true)
    const [posts, setPosts] = useState(null)

    CheckLogToken(router)
    
    useEffect(() => {
        fetch("http://localhost:8080/posts?limit=15",{
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            }})
            .catch(error => {
                setLoading(false)
                throw new Error(error)
            })

            .then(data => data.json())

            .then(data => {
            if(data.success) {
                setLoading(false)    
                setPosts(data.data)
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
        return posts.map((obj, i) => <CreatePost post={obj} key={i} />)
    }

    const hideModal = async () => {
        document.getElementById("modalDiv").classList.add("hidden")
        document.querySelectorAll(".modal").forEach(obj => {
            obj.classList.add("hidden")
        })
    }

    return (
        <div className="text-white h-full w-full grid items-center">
            <div className="bg-primaryT h-6/4 w-2/3 neon-xl center grid items-center">
                <div className="w-full h-screen overflow-scroll flex flex-col items-center p-4 gap-7">
                    <CheckState />
                </div>
            </div>
            <div className="fixed neon-xl w-1/10 h-fit max-h-5/6 left-5/6 top-1/12 postAction p-7">
                <div className="neon-sm p-5 rounded-xl flex flex-col items-center" onClick={async () => {showModal("newPostModal")}}>
                    <img src="/new.svg" className="h-max"></img>
                    <p className="text-sm text-center">New post</p>
                </div>
            </div>
            <ActionMenu />
            <div id="modalDiv" className="w-screen h-screen absolute hidden ">
                <div className="w-full h-full bg-black opacity-80 absolute z-5" onClick={hideModal}/>
                <NewPostModal />
                <DetailPostModal />
            </div>
        </div>
    )
}

const showModal = (modalId) => {
    document.getElementById("modalDiv").classList.remove("hidden")
    document.getElementById(modalId).classList.remove("hidden")
}

function CreatePost(data) {
    const post = data.post
    const updateModal = () => {
        document.getElementById("detailAuthor").textContent = post.AuthorId
        document.getElementById("detailMessage").textContent = post.Message
        fetch("http://localhost:8080/comments?post_id="+post.Id, {
            method: "GET"
        })
        .catch(error => {
            throw new Error("Failed to fetch comments.")
        })
        .then(data => {
            data.json()
        })
        .then(data => {
            if(data.success) {
                document.getElementById("detailCommentList").innerHTML = data.data.map((obj, i) => <CreateCom com={obj} key={i} />)
            } else {
                throw new Error("No data.")
            }
        })
    }

    return (
        <div className="w-5/6 rounded-xl neon-sm" onClick={async () => {
            showModal("detailPostModal")
            updateModal()
        }}>
            <div className="w-full postHeader bg-primaryT p-2">
                {post.AuthorId}
            </div>
            <div className="w-full h-fit p-4">
                {post.Message}
            </div>
            <div className="p-3 flex w-full gap-4">
                <div className="w-1/10 flex items-center">
                    <img src="/like.svg" className="h-8"></img>
                    <p>{post.nbLike}</p>
                </div>
                <div className="w-1/10 flex items-center gap-1">
                    <img src="/comment.svg" className="h-8"></img>
                    <p>{post.nbCom}</p>
                </div>
            </div>
        </div>
    )
}