"use client"

export default function NewPostModal() {
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
        <div id="newPostModal" className="modal neon-xl bg-primaryT h-9/10 w-3/5 absolute z-10 inset-1/2 -translate-1/2 rounded-xl hidden">
            <form className="w-full h-full flex flex-col justify-between items-center p-7 gap-5" onSubmit={newPostResolve}>
                <div className="w-full h-3/4 flex flex-col items-center gap-7">
                    <textarea name="message" className="w-full h-full neon-sm resize-none rounded-xl flex-grow p-3" placeholder="Content"></textarea>
                    <label htmlFor="postImage" className="bg-primaryT h-fit neon-sm rounded-xl w-8/10 p-2 flex flex-row justify-between" >
                        <div>
                            <input name="image" type="file" id="postImage" className="hidden" onChange={changedFile} accept=".gif,.jpg,.jpeg,.png"/>
                            <p>Chosen file (optional): </p><p id="fileName">None</p>
                        </div>
                        <div className="w-25 h-25">
                            <img id="preview" className="w-full h-full rounded-xl hidden"></img>
                        </div>
                    </label>
                </div>
                <input type="submit" className="neon-sm text-2xl rounded-xl px-10 py-5 bg-secondary duration-100 hover:cursor-pointer hover:scale-110" value="Post"></input>
            </form>
        </div>
    )
}