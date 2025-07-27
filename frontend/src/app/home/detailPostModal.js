"use client"

import { useState } from "react"

export default function DetailPostModal(data) {
    return (
        <div className="modal neon-xl bg-primaryT h-9/10 w-3/5 absolute z-10 inset-1/2 -translate-1/2 rounded-xl overflow-scroll">
            {data.postFeed ? <DetailContent postFeed={data.postFeed} /> : <PostNotFound /> }
        </div>
    )
}

function DetailContent({postFeed}) {
    const [newCom, setNewCom] = useState(false)
    const [liked, setLiked] = useState(postFeed.Like ? true : false)
    const post = postFeed.Post

    const coms = [ //replace with fetch
        {
            Comment: {
                Id: 9,
                Message: "tkt",
                Author: {
                    Id: 4,
                    Nickname: "wiz"
                }
            },
            Like: null,
            LikeCount: 0,
        },
        {
            Comment: {
                Id: 10,
                Message: "inquiète toi",
                Author: {
                    Id: 5,
                    Nickname: "ziw"
                }
            },
            Like: {
                Id: 11,
            },
            LikeCount: 9,
        }
    ]

    return (
        <div className="w-5/6 h-full flex flex-col items-center p-7 gap-5 center">
            <div className="w-full min-h-40 rounded-xl neon-sm bg-primaryT">
                <div className="w-full postHeader bg-primaryT p-2">
                    <p id="detailAuthor">{post.Author.Nickname || "Author not found"}</p>
                </div>
                <div className="w-full h-fit p-4">
                    <p id="detailMessage">{post.Message || "Content not found"}</p>
                </div>
                <label className="min-w-1/10 flex items-center" onClick={(e) => {e.stopPropagation()}}>
                    <input type="button" className="hidden" onClick={() => {likePost(postFeed, setLiked)}} />
                    <img src={liked ? "/likeActive.svg" : "/like.svg"} className={"h-8 "}></img>
                    <p className={liked ? "text-secondary" : null}>{postFeed.LikeCount || "0"}</p>
                </label>
            </div>
            <div className="w-full flex flex-row justify-between">
                <input type="button" value={newCom? "See comments" : "New comment"} className="bg-secondary neon-sm p-3 rounded-xl duration-100" onClick={() => {setNewCom(!newCom)}} />
                <div className="flex flex-row justify-end w-full">
                    <input type="button" value="Delete" className="bg-red-500 neon-sm p-3 rounded-xl duration-100" onClick={() => {deletePost(post)}} />
                </div>
            </div>
            {
                newCom 
                ? <NewComInput />
                : (coms.length > 0 
                    ? <CreateComList comList={coms} />
                    : <NoComs />)
            }
        </div>
    )
}

function deletePost(post) {
    fetch("http://localhost:8080/posts", {
        method: "DELETE",
        body: {
            session_uuid: localStorage.getItem("logToken"),
            post_id: post.Id
        }
    })
    .catch(error => {
        throw new Error(error)
    })

    .then(data => data.json())

    .then(data => {
        if(data.Result != "Ok") {
            throw new Error("Post deletion failed.")
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

function CreateComList({comList}) {
    return (
        <div className="w-5/6 h-fit flex flex-col gap-7 p-2 pt-1">
            <p className="font-bold">{comList.length} Comments :</p>
            {comList.map((obj, i) => <CreateCom comFeed={obj} key={i} />)}
        </div>
    )
}

function NoComs() {
    return (
        <div className="w-5/6 flex flex-col items-center gap-7 p-5">
            <p className="font-bold">No comments</p>
        </div>
    )
}

function CreateCom({comFeed}) {
    const [liked, setLiked] = useState(comFeed.Like ? true : false)
    const com = comFeed.Comment

    return (
         <div className="w-full h-fit rounded-xl neon-sm bg-primaryT">
            <div className="w-full postHeader bg-primaryT p-2">
                <p id="detailAuthor">{com.Author.Nickname}</p>
            </div>
            <div className="w-full h-fit p-4">
                <p id="detailMessage">{com.Message}</p>
            </div>
            <label className="p-3 flex w-full gap-1 w-1/10 items-center" onClick={(e) => {e.stopPropagation()}}>
                <input type="button" className="hidden" onClick={() => {likeCom(comFeed, setLiked)}} />
                <img src={liked ? "/likeActive.svg" : "/like.svg"} className={"h-8 "}></img>
                <p className={liked ? "text-secondary" : null}>{comFeed.LikeCount || "0"}</p>
            </label>
        </div>
    )
}

async function likeCom(comFeed, fn) {
    fetch("http://localhost:8080/likes", comFeed.Like 
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

    .then(data => {
        switch(typeof data.Result) {
            case "string": {
                fn(false)
            }
            case "object": {
                fn(true)
            }
        }
    })
}

function NewComInput() {
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

    const handleForm = async (event) => {
        event.preventDefault()
    }

    return (
        <form className="w-full h-full flex flex-col gap-5" onSubmit={handleForm}>
            <textarea name="Message" className="w-full resize-none neon-sm rounded-xl bg-primaryT h-25 overflow-scroll p-3 flex-grow" placeholder="Write your comment here" maxLength={1024} required />
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
            <input type="submit" value="Send" className="bg-secondary neon-sm rounded-xl px-5 py-3 w-fit self-end" />
        </form>
    )
}