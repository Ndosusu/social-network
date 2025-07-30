"use client"

import { useState } from "react"
import { DEFAULT_SERVER_PATH } from "../page"
import { newInfoMessage } from "../infoMessage"

export default function CreateComList({comList, postFeed}) {
    return (
        <div className="w-5/6 h-fit flex flex-col gap-7 p-2 pt-1">
            <p className="font-bold">{postFeed.CommentCount} Comments :</p>
            {comList.map((obj, i) => <CreateCom comFeed={obj} key={i} />)}
        </div>
    )
}

export function NoComs() {
    return (
        <div className="w-5/6 flex flex-col items-center gap-7 p-5">
            <p className="font-bold">No comments</p>
        </div>
    )
}

function CreateCom(data) {
    const [comFeed, setComFeed] = useState(data.comFeed)
    const com = comFeed.Comment

    return (
         <div className="w-full h-fit rounded-xl neon-sm bg-primaryT box-border">
            <div className="w-full postHeader bg-primaryT p-2 flex flex-row items-center gap-4">
                <img src={com.Author.Avatar ? DEFAULT_SERVER_PATH + "data/images" + com.Author.Avatar : "defaultAvatar.svg"} className="h-10 rounded-xl" />
                <p id="detailAuthor">{com.Author.Nickname || com.Author.FirstName + " " + com.Author.LastName || "Author not found"}</p>
            </div>
            <div className="w-full h-fit p-4">
                <p className="break-all">{com.Message}</p>
            </div>
            <label className="p-3 flex gap-1 w-fit items-center duration-100 hover:scale-110" onClick={(e) => {e.stopPropagation()}}>
                <input type="button" className="hidden" onClick={() => {likeCom(comFeed, setComFeed)}} />
                <img src={comFeed.Like ? "/likeActive.svg" : "/like.svg"} className={"h-8 "}></img>
                <p className={comFeed.Like ? "text-secondary" : null}>{comFeed.LikeCount || "0"}</p>
            </label>
        </div>
    )
}

async function likeCom(comFeed, fn) {
    const cloneFeed = structuredClone(comFeed)
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

    .then(data => data.json())

    .then(response => {
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
    })
}

export function NewComInput({postFeed, comFeedList, comFn, stateFn, infosFn}) {
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

    const newComResolve = async (event) => {
        event.preventDefault()

        const formData = new FormData(event.currentTarget)
        formData.append("session_uuid", localStorage.getItem("logToken"))
        formData.append("post_id", postFeed.Post.Id)

        fetch(DEFAULT_SERVER_PATH + "comments", {
            method: "POST",
            body: formData,
        })
        .catch(error => {
            throw new Error(error)
        })

        .then(data => data.json())
        .then(response => {
            const obj = {
                Like: null,
                LikeCount: null,
                Comment: response.data.Result,
            }
            infosFn(newInfoMessage("Comment created successfully"))
            if(comFeedList){
                comFn([obj].concat(comFeedList))
            } else {
                comFn([obj])
            }           
            stateFn(false)
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