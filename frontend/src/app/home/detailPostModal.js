"use client"

import { useEffect, useState } from "react"
import { likePost } from "./page"
import CreateComList, { NoComs, NewComInput } from "./createComList"
import { DEFAULT_SERVER_PATH } from "../page"

export default function DetailPostModal(data) {
    const [newCom, setNewCom] = useState(false)
    const [coms, setComs] = useState(null)
    const [postFeed, setPostFeed] = useState(data.postFeed)
    const post = postFeed.Post

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
            setComs(response.data.Result)
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
                    ? <CreateComList comList={coms} postFeed={postFeed} />
                    : <NoComs />)
            }
        </div>
        )
    }

    const getNextComs= async () => {
        if(!coms) {
            return
        }
        
        fetch(DEFAULT_SERVER_PATH + "feed/detail", {
            method: "POST",
            body: JSON.stringify({
                session_uuid: localStorage.getItem("logToken"),
                post_id: post.Id,
                limit: 15,
                last_id: coms[coms.length - 1].Comment.Id,
            })
        })
        .catch(error => {
            throw new Error(error)
        })
        .then(data => data.json())
        .then(response => {
            if(response.data.Result)
                setComs(coms.concat(response.data.Result))
        })
    }

    return (
        <div className="modal neon-xl bg-primaryT h-9/10 w-3/5 absolute z-10 inset-1/2 -translate-1/2 rounded-xl overflow-scroll" onScroll={(e) => {
            //Check if user scrolled to the bottom, 1 is needed as a safety because scrollHeight and clientHeight are rounded numbers but not scrollTop
            if(e.target.scrollHeight - e.target.clientHeight - e.target.scrollTop <= 1 && !newCom) 
                getNextComs()
        }}>
            {data.postFeed ? <DetailContent /> : <PostNotFound /> }
        </div>
    )
}

function deletePost(postFeed, modalFn, postsFn, postsList) {
    fetch(DEFAULT_SERVER_PATH + "posts", {
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