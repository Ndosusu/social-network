"use client"

import { useEffect, useState } from "react"
import { likePost } from "./page"
import { CreateComList, NoComs, NewComInput } from "./createComList"
import { DEFAULT_SERVER_PATH } from "../page"
import { parseState, useHomeContext } from "./contextProvider"

export default function DetailPostModal() {
    const {
        curPost,
    } = useHomeContext()
    const [commentList, setComList] = useState(null)
    const [commentInput, setComInput] = useState(false)
    const post = curPost.val.Post

    useEffect(() => {
        fetch(DEFAULT_SERVER_PATH + "feed/detail", {
            method: "POST",
            body: JSON.stringify({
                session_uuid: localStorage.getItem("logToken"),
                post_id: post.Id,
                limit: 15,
            })
        })
        .catch(error => {
            throw new Error(error)
        })

        .then(data => data.json())
        .then(response => {
            setComList(response.data.Result)
        })
    }, [])

    const DetailContent = () => {
        return (
            <div className="w-5/6 h-fit flex flex-col items-center p-7 gap-5 center box-content">
                <div className="w-full h-fit rounded-xl neon-sm bg-primaryT flex flex-col">
                    <div className="w-full postHeader bg-primaryT p-2 flex flex-row items-center gap-4">
                        <img src={post.Author.Avatar ? DEFAULT_SERVER_PATH + "data/images" + post.Author.Avatar : "defaultAvatar.svg"} className="h-10 rounded-xl" />
                        <p id="detailAuthor">{post.Author.Nickname || post.Author.FirstName + " " + post.Author.LastName || "Author not found"}</p>
                    </div>
                    <div className="w-full h-fit p-4">
                        <p className="break-all">{post.Message || "Content not found"}</p>
                    </div>
                    <label className="w-fit flex items-center select-none p-3 gap-1 duration-100 hover:scale-110" onClick={(e) => {e.stopPropagation()}}>
                        <input type="button" className="hidden" onClick={() => {likePost(curPost.val, curPost.set)}} />
                        <img src={curPost.val.Like ? "/likeActive.svg" : "/like.svg"} className="h-8"></img>
                        <p className={curPost.val.Like ? "text-secondary" : null}>{curPost.val.LikeCount || "0"}</p>
                    </label>
                </div>
                <div className="w-full flex flex-row justify-between">
                    <input type="button" value={commentInput? "See comments" : "New comment"} className="bg-secondary neon-sm p-3 rounded-xl duration-100 hover:scale-110" onClick={() => {setComInput(!commentInput)}} />
                    <div className="flex flex-row justify-end w-full">
                        <input type="button" value="Delete" className="bg-red-500 neon-sm p-3 rounded-xl duration-100 hover:scale-110" onClick={async () => {deletePost()}} />
                    </div>
                </div>
                {
                    commentInput
                    ? <NewComInput commentList={{val: commentList, set: setComList}} commentInput={{val: commentInput,set: setComInput}} />
                    : (commentList && commentList.length > 0 
                        ? <CreateComList commentList={{val: commentList, set:setComList}} />
                        : <NoComs />)
                }
            </div>
        )
    }

    const getNextComs= async () => {
        if(!commentList) {
            return
        }
        
        fetch(DEFAULT_SERVER_PATH + "feed/detail", {
            method: "POST",
            body: JSON.stringify({
                session_uuid: localStorage.getItem("logToken"),
                post_id: post.Id,
                limit: 15,
                last_id: commentList[commentList.length - 1].Comment.Id,
            })
        })
        .catch(error => {
            throw new Error(error)
        })
        .then(data => data.json())
        .then(response => {
            if(response.data.Result)
                setComList(commentList.concat(response.data.Result))
        })
    }

    return (
        <div className="modal neon-xl bg-primaryT h-9/10 w-3/5 absolute z-12 inset-1/2 -translate-1/2 rounded-xl overflow-scroll" onScroll={(e) => {
            //Check if user scrolled to the bottom, 1 is needed as a safety because scrollHeight and clientHeight are rounded numbers but not scrollTop
            if(e.target.scrollHeight - e.target.clientHeight - e.target.scrollTop <= 1 && !commentInput) 
                getNextComs()
        }}>
            {curPost.val ? <DetailContent /> : <PostNotFound /> }
        </div>
    )
}

function deletePost() {
    const {
        curPost,
        modal,
        feedPosts,
    } = useHomeContext()

    fetch(DEFAULT_SERVER_PATH + "posts", {
        method: "DELETE",
        body: JSON.stringify({
            session_uuid: localStorage.getItem("logToken"),
            post_id: curPost.val.Post.Id,
        })
    })
    .catch(error => {
        throw new Error(error)
    })

    .then(data => data.json())

    .then(response => {
        if(response.data.Result != "Ok") {
            throw new Error("Post deletion failed.")
        } else {
            feedPosts.val.splice(feedPosts.val.indexOf(curPost.val), 1)
            feedPosts.set(postsList)
            modal.set("")
        }
    })
}

function PostNotFound() {
    return (
        <div>
            <p>Post not found :c</p>
        </div>
    )
}