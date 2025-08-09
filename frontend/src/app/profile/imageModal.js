"use client"

import { newInfoMessage } from "../infoMessage"
import { DEFAULT_SERVER_PATH } from "../page"
import { CheckApiResponse } from "../utils"
import { useProfileContext } from "./contextProvider"

export function ImageModal() {
    const {
        curProfile,
        modal,
    } = useProfileContext()

    const changedFile = async (event) => {
        const preview = document.querySelectorAll(".preview")
        const fileName = document.querySelector("#fileName")
        const file = event.target.files[0]

        if (file) {
            let reader = new FileReader()
            fileName.textContent = file.name
            reader.onload = (e) => {
                preview.forEach(obj => {
                    obj.setAttribute("src", e.target.result)
                }) 
            };
            reader.readAsDataURL(file);
        } else {
            preview.forEach(obj => {
                obj.setAttribute("src", curProfile.val.User.Avatar ? DEFAULT_SERVER_PATH + "data/images/" + curProfile.val.User.Avatar : "defaultAvatar.svg")
            }) 
            fileName.textContent = "None"
            console.log("no file")
        }
    }

    const imageResolve = (event) => {
        event.preventDefault()

        const data = new FormData(event.currentTarget)
        data.append("session_uuid", localStorage.getItem("logToken"))

        fetch(DEFAULT_SERVER_PATH + "profile/avatar", {
            method: "PUT",
            body: data
        })
        .catch(error => {
            console.log(error)
            throw new Error(error)
        })

        .then(data => data.json())

        .then(response => {
            if(CheckApiResponse(response)) {
                curProfile.val.User.Avatar = response.data.Result
                modal.set("")
            } else {
                newInfoMessage("Failed to modify avatar.")
            }
        })
    }

    return (
        <div className="absolute inset-1/2 -translate-1/2 neon-xl w-2/3 h-fit z-12 bg-primaryT rounded-xl overflow-hidden">
            <form className="w-full h-full flex flex-col items-center p-5 gap-7" encType="multipart/form-data" onSubmit={imageResolve}>
                <p className="text-2xl">Modify avatar :</p>
                <div className="grid align-center h-fit w-9/10">
                    <label htmlFor="avatar" className="bg-primaryT h-fit neon-sm rounded-xl w-full p-2 flex flex-row justify-between" >
                        <div>
                            <input name="avatar" type="file" id="avatar" className="hidden" onChange={changedFile} accept=".gif,.jpg,.jpeg,.png"/>
                            <p>Avatar chosen (optional): </p><p id="fileName">None</p>
                        </div>
                    </label>
                </div>
                <p>Preview :</p>
                <div className="flex flex-row justify-center items-end w-full gap-10">
                    <img className="preview h-10 w-10 rounded-xl bg-primary" src={curProfile.val.User.Avatar ? DEFAULT_SERVER_PATH + "data/images/" + curProfile.val.User.Avatar : "defaultAvatar.svg"} />
                    <img className="preview h-30 w-30 rounded-xl bg-primary" src={curProfile.val.User.Avatar ? DEFAULT_SERVER_PATH + "data/images/" + curProfile.val.User.Avatar : "defaultAvatar.svg"} />
                    <img className="preview h-50 w-50 rounded-xl bg-primary" src={curProfile.val.User.Avatar ? DEFAULT_SERVER_PATH + "data/images/" + curProfile.val.User.Avatar : "defaultAvatar.svg"} />
                </div>
                <input type="submit" value="Confirm" className="justify-self-end py-3 px-7 bg-secondary neon-sm rounded-xl" />
            </form>
        </div>
    )
}