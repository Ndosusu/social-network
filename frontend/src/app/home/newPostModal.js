"use client"

import { useState } from "react"
import PrivateUserList from "./userList"

export default function NewPostModal({posts, postsFn, modalFn}) {
    const [privacyState, setPrivacy] = useState(1)

    const newPostResolve = async (event) => {
        event.preventDefault()

        const formData = new FormData(event.currentTarget)
        formData.append("session_uuid", localStorage.getItem("logToken"))

        fetch("http://localhost:8080/posts", {
            method: "POST",
            body: formData,
        })
        .catch(error => {
            console.log(error)
            throw new Error(error)
        })
        .then(data => data.json())
        .then(response => {
            modalFn("")
            postsFn([{
                CommentCount: 0,
                LikeCount: 0,
                GroupTitle: "",
                Post: response.data.Result,
            }].concat(posts))
        })
    }

    const changedFile = async (event) => {
        const preview = document.querySelector("#preview")
        const fileName = document.querySelector("#fileName")
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

    return (
        <div className="modal absolute w-full h-full z-10 inset-1/2 -translate-1/2 flex items-center justify-center pointer-events-none gap-7">
            <div className="neon-xl bg-primaryT h-9/10 w-3/5 rounded-xl pointer-events-auto">
                <form id="newPostForm" encType="multipart/form-data" className="w-full h-full flex flex-col justify-between items-center p-7 gap-7" onSubmit={newPostResolve}>
                    <div className="w-full h-3/4 flex flex-col items-center gap-7 flex-grow">
                        <textarea name="message" className="w-full h-full bg-primaryT neon-sm resize-none rounded-xl flex-grow p-3" placeholder="Content"></textarea>
                        <label htmlFor="postImage" className="bg-primaryT h-fit neon-sm rounded-xl w-8/10 p-2 flex flex-row justify-between" >
                            <div>
                                <input name="image" type="file" id="postImage" className="hidden" onChange={changedFile} accept=".gif,.jpg,.jpeg,.png"/>
                                <p>Chosen file (optional): </p><p id="fileName" className="fileName">None</p>
                            </div>
                            <div className="w-25 h-25">
                                <img id="preview" className="preview w-full h-full rounded-xl hidden"></img>
                            </div>
                        </label>
                        <CreateOptionList opts={["Public", "Followers only", "Private"]} fn={setPrivacy} />
                    </div>
                    <input type="submit" className="neon-sm text-2xl rounded-xl px-10 py-5 bg-secondary duration-100 hover:cursor-pointer hover:scale-110" value="Post"></input>
                </form>
            </div>
            <CheckPrivacyState privacy={privacyState} />
        </div>
    )
}

function CreateOptionList({opts, fn}) {
    return (
        <div className="w-full h-fit flex justify-center gap-10 text-center">
            {opts.map((obj, i) => <CreateOption text={obj} val={i+1} key={i} fn={fn} />)} 
        </div>
    )
    //val +1 to account for the fact that the list begins at index 0 whereas the starting index in the db is 1
}

function CreateOption({text, val, fn}) {
    return (
        <label htmlFor={"radio"+val} className="neon-sm rounded-xl p-2 w-2/10 duration-100 hover:scale-110">
            <input type="radio" name="privacy_mode" id={"radio"+val} value={val} defaultChecked={val == 1 ? true : false} onClick={() => {fn(val)}} className="hidden" />
            <p>{text}</p>
        </label>
    )
}

function CheckPrivacyState({privacy}) {
    //check if privacy is equal to 3, 3 is the value given if the user chose to make the post on a whitelist
    if (privacy == 3) 
        return <PrivateUserList />
}