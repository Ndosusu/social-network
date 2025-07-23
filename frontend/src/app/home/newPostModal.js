"use client"

import { useState } from "react"

export default function NewPostModal() {
    const [privacyState, setPrivacy] = useState(1)

    const newPostResolve = async (event) => {
        event.preventDefault()

        const formData = new FormData(event.currentTarget)

        fetch("http://localhost:8080/posts", {
            method: "POST",
            headers: {
                'Content-Type': 'application/json',
            },
            body: formData,
        })
        .catch(error => {
            throw new Error(error)
        })
        .then(data => {
            data.body.getReader().read().then((done, value) => {
                console.log(value)
            })
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
                <form id="newPostForm" className="w-full h-full flex flex-col justify-between items-center p-7 gap-7" onSubmit={newPostResolve}>
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
                        <CreateOptionList opts={["Public", "Followers only", "Private", "Group"]} fn={setPrivacy} />
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
        <div className="w-full h-fit flex justify-between gap-5 text-center">
            {opts.map((obj, i) => <CreateOption text={obj} val={i+1} key={i} fn={fn} />)} 
        </div>
    )
    //val +1 to account for the fact that the list begins at index 0 whereas the starting index in the db is 1
}

function CreateOption({text, val, fn}) {
    return (
        <label htmlFor={"radio"+val} className="neon-sm rounded-xl p-2 w-2/10">
            <input type="radio" name="radioPrivacy" id={"radio"+val} value={val} defaultChecked={val == 1 ? true : false} onClick={() => {fn(val)}} className="hidden" />
            <p>{text}</p>
        </label>
    )
}

function CheckPrivacyState({privacy}) {
    switch (privacy) {
        case 3: {
            return <PrivateUserList />
        }

        case 4: {
            return <PrivateGroupList />
        }
    }
}

function PrivateUserList() {
    const [selectedUsers, setSelectedUsers] = useState([])
    const [userList, setUsers] = useState([])

    //fetch users
    const tempUser = [
        {
            id:"5",
            name:"wiz",
        },
        {
            id:"8",
            name:"ziw",
        },
        {
            id:"8",
            name:"ziw",
        },
        {
            id:"8",
            name:"ziw",
        },
        {
            id:"8",
            name:"ziw",
        }
    ]

    const selectUser = (event) => {
        const div = event.target
        div.setAttribute("selected", "true")
    }

    const CreateUserCheckbox = ({user, init = false}) => {
        return (
            <label id="userCheckBox" className="w-full flex flex-row p-3 hover:bg-hovered text-xl items-center gap-3">
                <input form="newPostForm" name="checkboxUser" id={"checkbox"+user.id} type="checkbox" className="hidden" defaultChecked={init ? true : false} />
                <img src="discord.svg" className="bg-discord h-10 rounded-xl" />
                <p>{user.name}</p>
            </label>
        )
    }

    return (
        <div className="w-1/5 h-9/10 neon-xl rounded-xl flex flex-col pointer-events-auto p-2">
            <p>Selected users :</p>
            {selectedUsers.map((obj, i) => <CreateUserCheckbox user={obj} key={i} init={true} />)}
            <p>Available users :</p>
            {tempUser.map((obj, i) => <CreateUserCheckbox user={obj} key={i} />)}
        </div>
    )
}

function PrivateGroupList() {
    const [selectedGroup, setGroup] = useState(null)

    //fetch groups

    return (
        <div className="w-1/5 h-9/10 neon-xl rounded-xl flex flex-col pointer-events-auto">

        </div>
    )
}