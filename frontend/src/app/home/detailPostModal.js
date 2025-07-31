"use client"

import { likePost } from "./page"
import { CreateComList, NoComs, NewComInput } from "./createComList"
import { DEFAULT_SERVER_PATH } from "../page"
import { useHomeContext } from "./contextProvider"

export default function DetailPostModal() {
    const {
        curPost,
        commentInput,
        commentList,
    } = useHomeContext()

    const post = curPost.val.Post

    //called after checking if the post exist to show it
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
                    <input type="button" value={commentInput.val? "See comments" : "New comment"} className="bg-secondary neon-sm p-3 rounded-xl duration-100 hover:scale-110" onClick={() => {commentInput.set(!commentInput.val)}} />
                    <div className="flex flex-row justify-end w-full">
                        <input type="button" value="Delete" className="bg-red-500 neon-sm p-3 rounded-xl duration-100 hover:scale-110" onClick={() => {deletePost()}} />
                    </div>
                </div>
                {
                    commentInput.val
                    ? <NewComInput />
                    : (commentList.val && commentList.val.length > 0 
                        ? <CreateComList />
                        : <NoComs />)
                }
            </div>
        )
    }

    //function called when reaching the end of the comment list to get the next ones
    const getNextComs = () => {
        if(!commentList.val) {
            return
        }

        const lastId = commentList.val[commentList.val.length - 1].Comment.Id
        console.log("Fetching comments after #" + lastId + "...")
        
        //api call to get the next comments, it works by giving the id of the last comment you have, gives the next comments that are older than the given one
        fetch(DEFAULT_SERVER_PATH + "feed/detail", {
            method: "POST",
            body: JSON.stringify({
                session_uuid: localStorage.getItem("logToken"),
                post_id: post.Id,
                limit: 15,
                last_id: lastId,
            })
        })
        .catch(error => {
            throw new Error(error)
        })

        //make data readable as json object
        .then(data => data.json())

        //update comment list with new comments
        .then(response => {
            if(response.data.Result) {
                commentList.set(commentList.val.concat(response.data.Result))
            }
        })
    }

    return (
        <div className="modal neon-xl bg-primaryT h-9/10 w-3/5 absolute z-12 inset-1/2 -translate-1/2 rounded-xl overflow-scroll" onScroll={(e) => {
            console.log("scroll")
            console.log(e.target.scrollHeight - e.target.clientHeight - e.target.scrollTop)
            //Check if user scrolled to the bottom, 1 is needed as a safety because scrollHeight and clientHeight are rounded numbers but not scrollTop
            if(e.target.scrollHeight - e.target.clientHeight - e.target.scrollTop <= 1 && !commentInput.val) 
                getNextComs()
        }}>
            {curPost.val ? <DetailContent /> : <PostNotFound /> }
        </div>
    )
}

//called when you try to delete a post
function deletePost() {
    const {
        curPost,
        modal,
        feedPosts,
    } = useHomeContext()

    //api call to delete given post
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

    //make data readable as json object
    .then(data => data.json())

    //api returns "Ok" as string if succesful, act accordingly
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

//called if the selected post is not found
function PostNotFound() {
    return (
        <div>
            <p>Post not found :c</p>
        </div>
    )
}