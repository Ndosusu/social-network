"use client"

import { useEffect, useState } from "react"
import { likePost } from "./page"
import { CreateComList, NoComs, NewComInput } from "./createComList"

export default function DetailPostModal(data) {
    return (
        <div className="modal neon-xl bg-primaryT h-9/10 w-3/5 absolute z-10 inset-1/2 -translate-1/2 rounded-xl overflow-scroll">
            {data.postFeed ? <DetailContent postFeed={data.postFeed} postsFn={data.postsFn} modalFn={data.modalFn} postsList={data.postsList} /> : <PostNotFound /> }
        </div>
    )
}

function DetailContent(data) {
    const [newCom, setNewCom] = useState(false)
    const [coms, setComs] = useState(null)
    const [postFeed, setPostFeed] = useState(data.postFeed)
    const post = postFeed.Post

    useEffect(() => {
        fetch("http://localhost:8080/feed/detail", {
            method: "POST",
            body: JSON.stringify({
                session_uuid: localStorage.getItem("logToken"),
                post_id: post.Id,
            })
        })
        .catch(error => {
            throw new Error(error)
        })

        .then(data => data.json())
        .then(response => {
            setComs(response.data.Result)
        })
    }, [])

    return (
        <div className="w-5/6 h-full flex flex-col items-center p-7 gap-5 center">
            <div className="w-full min-h-40 rounded-xl neon-sm bg-primaryT">
                <div className="w-full postHeader bg-primaryT p-2">
                    <p id="detailAuthor">{post.Author.Nickname || post.Author.FirstName + " " + post.Author.LastName || "Author not found"}</p>
                </div>
                <div className="w-full h-fit p-4">
                    <p id="detailMessage">{post.Message || "Content not found"}</p>
                </div>
                <label className="w-fit flex items-center select-none p-3 gap-1 duration-100 hover:scale-110" onClick={(e) => {e.stopPropagation()}}>
                    <input type="button" className="hidden" onClick={() => {likePost(postFeed, setPostFeed)}} />
                    <img src={postFeed.Like ? "/likeActive.svg" : "/like.svg"} className="h-8"></img>
                    <p className={postFeed.Like ? "text-secondary" : null}>{postFeed.LikeCount || "0"}</p>
                </label>
            </div>
            <div className="w-full flex flex-row justify-between">
                <input type="button" value={newCom? "See comments" : "New comment"} className="bg-secondary neon-sm p-3 rounded-xl duration-100 hover:scale-110" onClick={() => {setNewCom(!newCom)}} />
                <div className="flex flex-row justify-end w-full">
                    <input type="button" value="Delete" className="bg-red-500 neon-sm p-3 rounded-xl duration-100 hover:scale-110" onClick={async () => {deletePost(postFeed, data.modalFn, data.postsFn, data.postsList)}} />
                </div>
            </div>
            {
                newCom 
                ? <NewComInput postFeed={postFeed} comFeedList={coms} comFn={setComs} stateFn={setNewCom} />
                : (coms && coms.length > 0 
                    ? <CreateComList comList={coms} />
                    : <NoComs />)
            }
        </div>
    )
}

function deletePost(postFeed, modalFn, postsFn, postsList) {
    fetch("http://localhost:8080/posts", {
        method: "DELETE",
        body: JSON.stringify({
            session_uuid: localStorage.getItem("logToken"),
            post_id: postFeed.Post.Id,
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
            let index
            postsList.forEach((obj, i) => {
                if(obj.Post.Id == postFeed.Post.Id) {
                    index = i
                }
            })
            postsList.splice(index, 1)
            postsFn(postsList)
            modalFn("")
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