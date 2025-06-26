"use client"
import { useState } from "react";

export default function Home() {
  const [loginForm, setLogin] = useState(true)

  const loginResolve = async (event) => {
    event.preventDefault()

    const formData = new FormData(event.currentTarget)
    let body = {}

    formData.forEach((val, key) =>{
        body[key] = val
    })

    let result = await fetch("localhost:8080/auth/login", {
        method: 'POST',
        data: JSON.stringify(body)
    })
    console.log(result)
  }

  const registerResolve = async (event) => {
    event.preventDefault()

    const formData = new FormData(event.currentTarget)
    let body = {}

    formData.forEach((val, key) =>{
        body[key] = val
    })

    console.log(body["Avatar"])
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

  if (loginForm) {
    return (
      <div className="h-full w-full flex flex-col justify-center items-center text-white gap-10">
        <section className="w-1/3 h-2/4">
          <form id="login" onSubmit={loginResolve} className="w-full h-full flex flex-col bg-primaryT p-10 justify-between items-center neon-xl rounded-3xl" method="post">
            <div className="flex flex-col h-fit gap-5 w-full justify-between">
              <input name="Mail" type="text" className="bg-primaryT h-11 p-3 neon-sm rounded-xl" placeholder="Mail" required/>
              <input name="Password" type="password" className="bg-primaryT h-11 p-3 neon-sm rounded-xl" placeholder="Password" required/>
            </div>
            <div className="flex flex-col w-full items-center">
              <input type="submit" className="bg-secondary h-13 w-1/2 rounded-xl neon-xl hover:cursor-pointer text-xl duration-100 ease-in-out hover:scale-110" value="Log in"/>
              <a className="hover:cursor-pointer underline active:text-secondary" onClick={() => {
                document.querySelector("#login").reset()
                setLogin(false)
                }}>Or register here</a>
            </div>
          </form>
        </section>
        <div className="bg-primaryT neon-xl rounded-3xl w-1/3 h-1/6 flex flex-row gap-5 p-3 justify-around items-center">
          <img id="logGoogle" src="/googleLogo.svg" className="w-fit h-3/4 bg-white neon-sm p-2 rounded-xl duration-100 ease-in-out hover:scale-110"></img>
          <img id="logDiscord" src="/discord.svg" className="w-fit h-3/4 bg-discord neon-sm p-2 rounded-xl duration-100 ease-in-out hover:scale-110"></img>
          <img id="logGithub" src="/github.svg" className="w-fit h-3/4 bg-black neon-sm p-2 rounded-xl duration-100 ease-in-out hover:scale-110"></img>
        </div>
      </div>
    )
  }
  else if(!loginForm) {
    return (
      <div className="h-full w-full flex flex-col justify-center items-center text-white">
        <section className="w-fit h-5/6">
          <form id="register" onSubmit={registerResolve} encType="multipart/form-data" className="w-full h-full flex flex-col bg-primaryT p-10 justify-between items-center neon-xl rounded-3xl" method="post">
            <div className="flex flex-col h-fit gap-5 justify-between">
              <div className="flex flex-row h-full gap-8">
                <div className="grid grid-cols-2 gap-5 w-1/2 h-fit">
                  <input name="FirstName" type="text" className="bg-primaryT h-10 p-3 neon-sm rounded-xl" placeholder="First Name" required/>
                  <input name="LastName" type="text" className="bg-primaryT h-10 p-3 neon-sm rounded-xl" placeholder="Last Name" required/>
                  <input name="Mail" type="text" className="bg-primaryT h-10 p-3 neon-sm rounded-xl col-span-2" placeholder="Mail" required/>
                  <input name="Password" type="password" className="bg-primaryT h-10 p-3 neon-sm rounded-xl col-span-2" placeholder="Password" required/>
                  <input name="RPassword" type="password" className="bg-primaryT h-10 p-3 neon-sm rounded-xl col-span-2" placeholder="Repeat Password" required/>
                  <div className="col-span-2">
                  <p>Date of Birth:</p>
                    <div className="w-full flex flex-row gap-4">
                      <input name="Day" type="text" className="bg-primaryT h-10 p-3 neon-sm rounded-xl w-1/4" placeholder="Day" required/>
                      <input name="Month" type="text" className="bg-primaryT h-10 p-3 neon-sm rounded-xl w-1/4" placeholder="Month" required/>
                      <input name="Year" type="text" className="bg-primaryT h-10 p-3 neon-sm rounded-xl w-max" placeholder="Year" required/>
                    </div> 
                  </div>   
                </div>
                <div className="grid grid-cols-2 gap-5 w-1/2 auto-rows-max">
                  <input name="Nickname" type="text" className="bg-primaryT h-10 p-3 neon-sm rounded-xl col-span-2" placeholder="Nickname (optional)"/>
                  <textarea name="About" className="bg-primaryT h-10 p-3 neon-sm rounded-xl col-span-2 h-40 resize-none" placeholder="About you (optional)" maxLength={512}></textarea>
                  <div className="col-span-2 grid align-center h-fit">
                    <label htmlFor="avatar" className="bg-primaryT h-fit neon-sm rounded-xl w-full p-2 flex flex-row justify-between" >
                      <div>
                        <input name="Avatar" type="file" id="avatar" className="hidden" onChange={changedFile} accept=".gif,.jpg,.jpeg,.png"/>
                        <p>Avatar chosen (optional): </p><p id="fileName">None</p>
                      </div>
                      <div className="w-25 h-25">
                        <img id="preview" className="w-full h-full rounded-xl hidden"></img>
                      </div>
                    </label>
                    
                  </div>
                </div>
              </div>
            </div>
            <div className="flex flex-col w-full items-center">
              <input type="submit" className="bg-secondary h-13 w-1/2 rounded-xl neon-xl hover:cursor-pointer text-xl duration-100 ease-in-out hover:scale-110" value="Register"/>
              <a className="hover:cursor-pointer underline active:text-secondary" onClick={() => {
                document.querySelector("#register").reset()
                setLogin(true)
                }}>Or log in here</a>
            </div>
          </form>
        </section>
      </div>
      )
  }
}
