"use client"

import { useState } from "react"
import { DEFAULT_SERVER_PATH } from "../page"
import { newInfoMessage } from "../infoMessage"
import { useHomeContext } from "./contextProvider"
import { CheckApiResponse } from "../utils"

//component that creates the comments feed dynamically
export function CreateComList({comList, setComList}) {
    const {
        curPost,
    } = useHomeContext()

    return (
        <div className="w-5/6 h-fit flex flex-col gap-7 p-2 pt-1">
            <p className="font-bold">{curPost.val.CommentCount} Comments :</p>
            {comList.map((obj, i) => <CreateCom comFeed={obj} comListState={{comList, setComList}} key={i} />)}
        </div>
    )
}

//returns a div to tell the user that there is no comment, CALLED ONLY IF NO COMMENTS
export function NoComs() {
    return (
        <div className="w-5/6 flex flex-col items-center gap-7 p-5">
            <p className="font-bold">No comments</p>
        </div>
    )
}

//create a div for a comment with the appropriate data
function CreateCom(data) {
    //state needed to handle the liked state of the comment
    const [comFeed, setComFeed] = useState(data.comFeed)
    const com = comFeed.Comment
    const {
        comList,
        setComList,
    } = data.comListState
    const {
        curPost,
        feedPosts,
    } = useHomeContext()

    //function to call when user deletes a comment
    const deleteCom = () => {
        //api call to delete in db
        fetch(DEFAULT_SERVER_PATH + "comments", {
            method: "DELETE",
            body: JSON.stringify({
                session_uuid: localStorage.getItem("logToken"),
                comment_id: com.Id,
            })
        })
        .catch(error => {
            console.log(error)
            throw new Error(error)
        })

        //make data readable as json object
        .then(data => data.json())

        .then(response => {
            if(CheckApiResponse(response)){
                comList.splice(comList.indexOf(comFeed), 1)
                setComList(comList)

                const copy = structuredClone(curPost.val)
                const listCopy = structuredClone(feedPosts.val)
                const index = feedPosts.val.indexOf(curPost.val)

                --copy.CommentCount

                listCopy[index] = copy
                feedPosts.set(listCopy)
                curPost.set(copy)

                newInfoMessage("Comment deleted successfully")
            } else {
                newInfoMessage("Failed to delete comment", "bg-red-500")
            }
        })
    }

    return (
         <div className="w-full h-fit rounded-xl neon-sm bg-primaryT box-border">
            <div className="w-full postHeader bg-primaryT p-2 flex flex-row items-center gap-4">
                <img src={com.Author.Avatar ? DEFAULT_SERVER_PATH + "data/images" + com.Author.Avatar : "defaultAvatar.svg"} className="h-10 rounded-xl" />
                <p id="detailAuthor">{com.Author.Nickname || com.Author.FirstName + " " + com.Author.LastName || "Author not found"}</p>
            </div>
            <div className="w-full h-fit p-4">
                <p className="break-all">{com.Message}</p>
            </div>
            {
                com.Image 
                ? <img src={DEFAULT_SERVER_PATH + "data/images/" + com.Image} className="center max-w-full rounded-xl"/>
                : null
            }
            <div className="w-full flex flex-row justify-between">
                <label className="p-3 flex gap-1 w-fit items-center duration-100 hover:scale-110" onClick={(e) => {e.stopPropagation()}}>
                    <input type="button" className="hidden" onClick={() => {likeCom(comFeed, setComFeed)}} />
                    <img src={comFeed.Like ? "/likeActive.svg" : "/like.svg"} className={"h-8 "}></img>
                    <p className={comFeed.Like ? "text-secondary" : null}>{comFeed.LikeCount || "0"}</p>
                </label>
                {
                    com.Author.IsClient

                    ?
                    <label className="bin p-3 flex gap-1 w-fit items-center duration-100 hover:scale-110" onClick={(e) => {e.stopPropagation()}}>
                        <input type="button" className="hidden" onClick={deleteCom} />
                        <img src="/bin.svg" className="h-8" />
                    </label>

                    : null
                }
            </div>
        </div>
    )
}

//function to call when user likes a comment
function likeCom(comFeed, fn) {
    //clone the original structure  in order not to modify the original (important to make the set state function work)
    const cloneFeed = structuredClone(comFeed)

    //api call, ternary operator to switch between deleting the like or add it
    fetch(DEFAULT_SERVER_PATH + "likes", comFeed.Like 
        ? {
            method: "DELETE",
            body: JSON.stringify({
                session_uuid: localStorage.getItem("logToken"),
                like_id: comFeed.Like.Id,
            })
        }
        : {
            method: "POST",
            body: JSON.stringify({
                session_uuid: localStorage.getItem("logToken"),
                comment_id: comFeed.Comment.Id,
            })
        }
    )
    .catch(error => {
        throw new Error(error)
    })

    //make data readable as json object
    .then(data => data.json())

    //if action was delete, Result is "ok" as string, else Result is the new like as an object. act accordingly
    .then(response => {
        if(CheckApiResponse(response)){
            switch(typeof response.data.Result) {
                case "string": {
                    cloneFeed.Like = null
                    --cloneFeed.LikeCount
                    fn(cloneFeed)
                    break
                }
                case "object": {
                    cloneFeed.Like = response.data.Result
                    ++cloneFeed.LikeCount
                    fn(cloneFeed)
                    break
                }
            }
        } else {
            newInfoMessage("Failed to like comment", "bg-red-500")
        }
    })
}

//Component called when user wants to create a new comment
export function NewComInput() {
    const {
        curPost,
        feedPosts,
        commentList,
        commentInput,
    } = useHomeContext()

    //called when the user chooses a file to update the preview
    const changedFile = async (event) => {
        const preview = document.querySelector("#previewCom")
        const fileName = document.querySelector("#fileNameCom")
        const file = event.target.files[0]

        if (file) {
            let reader = new FileReader()
            preview.classList.remove("hidden")
            fileName.textContent = file.name
            reader.onload = (e) => {
                preview.setAttribute("src", e.target.result)
            };
            reader.readAsDataURL(file);
        } else {
            preview.setAttribute("src", "")
            preview.classList.add("hidden")
            fileName.textContent = "None"
            console.log("no file")
        }
    }

    //called when the new comment form is submitted
    const newComResolve = (event) => {
        event.preventDefault()

        //create the object to send to the back and adding necessary data not given by the form
        const formData = new FormData(event.currentTarget)
        formData.append("session_uuid", localStorage.getItem("logToken"))
        formData.append("post_id", curPost.val.Post.Id)

        //api request to add the comment
        fetch(DEFAULT_SERVER_PATH + "comments", {
            method: "POST",
            body: formData,
        })
        .catch(error => {
            throw new Error(error)
        })

        //make data readable as json object
        .then(data => data.json())

        //if everything went well, Result is the new comment object. create an empty CommentFeed object and add it to the list
        .then(response => {
            if(CheckApiResponse(response)) {
                const obj = {
                    Like: null,
                    LikeCount: null,
                    Comment: response.data.Result,
                }
                // infosFn(newInfoMessage("Comment created successfully"))
                if(commentList.val){
                    commentList.set([obj].concat(commentList.val))
                } else {
                    commentList.set([obj])
                }

                const copy = structuredClone(curPost.val)
                const listCopy = structuredClone(feedPosts.val)
                const index = feedPosts.val.indexOf(curPost.val)

                ++copy.CommentCount

                listCopy[index] = copy
                feedPosts.set(listCopy)
                curPost.set(copy)

                newInfoMessage("Comment created successfuly")
                commentInput.set(false)
            } else {
                newInfoMessage("Failed to create comment", "bg-red-500")
            }
        })
    }

    return (
        <form className="w-full h-full flex flex-col gap-5" encType="multipart/form-data" onSubmit={newComResolve}>
            <textarea name="message" className="w-full resize-none neon-sm rounded-xl bg-primaryT h-25 overflow-scroll p-3 flex-grow" placeholder="Write your comment here" maxLength={1024} required />
            <div className="col-span-2 grid align-center h-fit">
                <label htmlFor="file" className="bg-primaryT h-fit neon-sm rounded-xl w-full p-2 flex flex-row justify-between" >
                    <div>
                        <input name="File" type="file" id="file" className="hidden" onChange={changedFile} accept=".gif,.jpg,.jpeg,.png"/>
                        <p>File chosen (optional): </p><p id="fileNameCom" className="fileName">None</p>
                    </div>
                    <div className="w-25 h-25">
                        <img id="previewCom" className="preview w-full h-full rounded-xl hidden"></img>
                    </div>
                </label>
            </div>
            <input type="submit" value="Send" className="bg-secondary neon-sm rounded-xl px-5 py-3 w-fit self-end duration-100 hover:scale-110" />
        </form>
    )
}