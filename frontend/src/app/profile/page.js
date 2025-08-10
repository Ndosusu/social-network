"use client"

import ActionMenu from "../actionMenu"
import { useRouter, useSearchParams } from "next/navigation"
import { CheckApiResponse, CheckLogToken } from "../utils"
import { ProfileProvider, useProfileContext } from "./contextProvider"
import { useEffect, useState } from "react"
import { DEFAULT_SERVER_PATH } from "../page"
import { likePost } from "../home/page"
import { CreateAllInfoMessages } from "../infoMessage"
import DetailPostModal from "./detailPostModal"
import { ImageModal } from "./imageModal"

export default function ProfileContextWrapper() {
    const router = useRouter()
    CheckLogToken(router)

    const params = useSearchParams()
    const id = params.get("id")

    return (
        <ProfileProvider>
            <ProfilePage id={id}/>
        </ProfileProvider>
    )
}

export function ProfilePage({id}) {
    const {
        curProfile,
        extraContent,
        modal,
    } = useProfileContext()

    useEffect(() => {
        fetch(DEFAULT_SERVER_PATH + "profile", {
            method: "POST",
            body: JSON.stringify({
                session_uuid: localStorage.getItem("logToken"),
                user_id: id,
            })
        })
        .catch(error => {
            console.log(error)
            throw new Error(error)
        })

        .then(data => data.json())

        .then(response => {
            if(CheckApiResponse(response)) {
                curProfile.set(response.data.Result)
            }
        })
    }, [])

    return (
        <div className="center w-full h-full text-white overflow-scroll p-10">
            <div className="neon-xl bg-primaryT w-full h-fit p-7 rounded-xl center flex flex-col gap-10">
                <div className="flex flex-col gap-3">
                    <div className="flex flex-row gap-6">
                        <img src={curProfile.val.User.Avatar ? DEFAULT_SERVER_PATH + "data/images/" + curProfile.val.User.Avatar : "defaultAvatar.svg"} className="w-50 h-50 rounded-xl bg-primary" onClick={() => {modal.set("imageModal")}} />
                        <div className="flex flex-col justify-between">
                            {curProfile.val.User.Nickname ? <p className="text-5xl h-fit">Pepiño</p> : null}
                            <p className={curProfile.val.User.Nickname ? "text-3xl h-fit text-gray-400" : "text-5xl h-fit"}>Lotr Taré</p>
                            <p className="text-3xl h-fit">User since : {curProfile.val.User.CreatedDate}</p>
                            <p className="text-3xl h-fit">Born on : {curProfile.val.User.BirthDate}</p>
                        </div>
                    </div>
                    {
                        curProfile.val.User.About
                        ? (
                            <div className="flex flex-col gap-3">
                                <p className="text-3xl px-5 font-bold">About me :</p>
                                <p className="text-2xl px-5 break-words">{curProfile.val.User.About}</p>
                            </div>
                        )

                        : null
                    }
                    <div className="flex flex-row justify-around text-2xl p-10 text-center">
                        <div onClick={() => {extraContent.set("following")}} className="cursor-pointer duration-100 hover:scale-110">
                            <p>Following</p>
                            <p>200</p>
                        </div>
                        <div onClick={() => {extraContent.set("followers")}} className="cursor-pointer duration-100 hover:scale-110">
                            <p>Followers</p>
                            <p>100</p>
                        </div>
                        <div onClick={() => {extraContent.set("posts")}} className="cursor-pointer duration-100 hover:scale-110">
                            <p>Posts</p>
                            <p>15</p>
                        </div>
                    </div>
                </div>
            </div>
            <div className="bg-primaryT w-8/10 h-full mt-10 rounded-xl center relative">
                <div className="w-full h-full z-6 absolute neon-xl rounded-xl pointer-events-none" />
                <CheckState />
            </div>
            {
                modal.val
                ? <CreateModal />
                : null
            }
            <ActionMenu />
            <CreateAllInfoMessages />
        </div>
    )
}

function CreateModal() {
    const {
        modal,
    } = useProfileContext()

    return (
        <div id="modalDiv" className="w-screen h-screen absolute left-0 top-0 z-10">
            <div className="w-full h-full bg-black opacity-80 absolute z-11" onClick={() => {modal.set("")}}/>
            <CheckModalState />
        </div>
    )
}

function CheckModalState() {
    const {
        modal,
    } = useProfileContext()

    switch(modal.val) {
        case "detailModal":{
            return <DetailPostModal />
        }

        case "imageModal":{
            return <ImageModal />
        }
    }
}

function CheckState() {
    const {
        extraContent,
        postList,
        followedList,
        followingList,
    } = useProfileContext()

    let content

    switch(extraContent.val) {
        case "posts": {
            return <CreatePostList postList={postList} />
        }
        
        case "followers": {
            content = followedList.val.map((obj, i) => <CreateUser user={obj} key={i} />)
            break
        }

        case "following": {
            content = followingList.val.map((obj, i) => <CreateUser user={obj} key={i} />)
            break
        }

        default: {
            content = <p>State error</p>
            break
        }
    }

    return (
        <div className="w-full h-full rounded-xl overflow-hidden relative">
            <div className="absolute w-full h-full fade z-5 pointer-events-none" />
            <div className="w-full h-full overflow-scroll flex flex-col items-center gap-3 p-7">
                {content}
            </div>
        </div>
    )
}

function CreateUser(data) {
    const router = useRouter()
    const user = data.user

    return (
        <div className="w-3/4 h-fit flex flex-row items-center gap-3 p-4 rounded-xl hover:bg-hovered cursor-pointer" onClick={() => {router.push("/profile?id=" + user.Id)}}>
            <img src={user.Avatar ? DEFAULT_SERVER_PATH + "data/images/" + user.Avatar : "defaultAvatar.svg"} className="h-12 bg-primary rounded-xl" />
            <p>{user.Nickname || user.FirstName + " " + user.LastName || "User not found"}</p>
        </div>
    )
}

function CreatePostList({postList}) {
    const {
        curFeed,
    } = useProfileContext()

    return (
        <div className="w-full h-full flex flex-col rounded-xl overflow-hidden">
            <div className="flex flex-row justify-around p-10 gap-10 bg-primary">
                <label htmlFor="globalFeed" className="neon-sm p-2 w-full h-fit flex flex-row rounded-xl text-center duration-100 hover:scale-110">
                    <input id="globalFeed" name="feedRadio" type="radio" defaultChecked onClick={() => {curFeed.set("profile")}} className="hidden" />
                    <p className="w-full">Created</p>
                </label>
                <label htmlFor="followFeed" className="neon-sm p-2 w-full h-fit flex flex-row rounded-xl text-center duration-100 hover:scale-110">
                    <input id="followFeed" name="feedRadio" type="radio" onClick={() => {curFeed.set("liked")}} className="hidden" />
                    <p className="w-full">Liked</p>
                </label>
            </div>
            <div className="w-full h-full overflow-hidden relative">
                <div className="absolute top-0 left-0 w-full h-full fade pointer-events-none z-5" />
                <div className="w-full h-full overflow-scroll flex flex-col gap-8 p-7 items-center">
                    {postList.val.length > 0 ? postList.val.map((obj, i) => <CreatePost postFeed={obj} key={i} />) : <p>No posts found.</p>}
                </div>
            </div>
        </div>
    )
}

function CreatePost(data) {
    const {
        curPost,
        postList,
        modal,
    } = useProfileContext()
    const [postFeed, setPostFeed] = useState(data.postFeed)
    const post = postFeed.Post

    const updatePostList = (post) => {
        const index = postList.val.indexOf(postFeed)
        postList.val[index] = post
    }
    
    return (
        <div className="w-5/6 rounded-xl neon-sm duration-100 hoverable hover:scale-110" onClick={() => {
            curPost.set(postFeed)
            modal.set("detailModal")
        }}>
            <div className="w-full postHeader bg-primaryT p-2 flex flex-row gap-4 items-center">
                <img src={post.Author.Avatar ? DEFAULT_SERVER_PATH + "data/images/" + post.Author.Avatar : "defaultAvatar.svg"} className="h-10 rounded-xl" />
                {post.Author.Nickname || post.Author.FirstName + " " + post.Author.LastName || "Author not found"}
            </div>
            <div className="w-full h-fit p-4 flex flex-row justify-between gap-3">
                <p className="break-all">{post.Message || "Content not found"}</p>
                {
                    post.Image
                    ? <img src={ DEFAULT_SERVER_PATH + "data/images/" + post.Image} className="h-25 max-w-1/3" />
                    : null
                }
            </div>
            <div className="p-3 flex w-full gap-4">
                <label className="min-w-1/10 flex items-center duration-100 hover:scale-110" onClick={(e) => {e.stopPropagation()}}>
                    <input type="button" className="hidden" onClick={() => {likePost(postFeed, (post) => {
                            updatePostList(post)
                            setPostFeed(post)
                        })
                    }} />
                    <img src={postFeed.Like ? "/likeActive.svg" : "/like.svg"} className={"h-8 "}></img>
                    <p className={postFeed.Like ? "text-secondary" : null}>{postFeed.LikeCount || "0"}</p>
                </label>
                <div className="min-w-1/10 flex items-center gap-1">
                    <img src="/comment.svg" className="h-8"></img>
                    <p>{postFeed.CommentCount || "0"}</p>
                </div>
            </div>
        </div>
    )
}