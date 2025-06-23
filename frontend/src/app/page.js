"use client"
import { useState } from "react";

export default function Home() {
  const [loginForm, setLogin] = useState(true)
  if (loginForm) {
    return (
      <div className="h-full w-full flex flex-col justify-center items-center text-white gap-10">
        <section className="w-1/3 h-2/4">
          <form id="login" className="w-full h-full flex flex-col bg-primary p-10 justify-between items-center neon-xl rounded-3xl" method="post">
            <div className="flex flex-col h-fit gap-5 w-full justify-between">
              <input name="Mail" type="text" className="bg-primary h-11 p-3 neon-sm rounded-xl" placeholder="Mail"/>
              <input name="Password" type="password" className="bg-primary h-11 p-3 neon-sm rounded-xl" placeholder="Password"/>
            </div>
            <div className="flex flex-col w-full items-center">
              <input type="submit" className="bg-secondary h-13 w-1/2 rounded-xl neon-xl hover:cursor-pointer" value="Log in"/>
              <a className="hover:cursor-pointer" onClick={() => {
                document.querySelector("#login").reset()
                setLogin(false)
                }}>Or register here</a>
            </div>
          </form>
        </section>
        <div className="bg-primary neon-xl rounded-3xl w-1/3 h-1/6">

        </div>
      </div>
    )
  }
  else if(!loginForm) {
    return (
      <div className="h-full w-full flex flex-col justify-center items-center text-white">
        <section className="w-2/5 h-5/6">
          <form id="register" className="w-full h-full flex flex-col bg-primary p-10 justify-between items-center neon-xl rounded-3xl" method="post">
            <div className="flex flex-col h-fit gap-5 w-full justify-between">
              <input name="FirstName" type="text" className="bg-primary h-10 p-3 neon-sm rounded-xl" placeholder="First Name"/>
              <input name="LastName" type="text" className="bg-primary h-10 p-3 neon-sm rounded-xl" placeholder="Last Name"/>
              <input name="Nickname" type="text" className="bg-primary h-10 p-3 neon-sm rounded-xl" placeholder="Nickname"/>
              <input name="Mail" type="text" className="bg-primary h-10 p-3 neon-sm rounded-xl" placeholder="Mail"/>
              <input name="Password" type="text" className="bg-primary h-10 p-3 neon-sm rounded-xl" placeholder="Password"/>
              <input name="RPassword" type="text" className="bg-primary h-10 p-3 neon-sm rounded-xl" placeholder="Repeat Password"/>
              <div>
                <p>Date of Birth:</p>
                <div className="w-full flex flex-row gap-4">
                  <input name="Day" type="text" className="bg-primary h-10 p-3 neon-sm rounded-xl w-1/4" placeholder="Day" />
                  <input name="Month" type="text" className="bg-primary h-10 p-3 neon-sm rounded-xl w-1/4" placeholder="Month"/>
                  <input name="Year" type="text" className="bg-primary h-10 p-3 neon-sm rounded-xl w-1/4" placeholder="Year"/>
                </div>          
              </div>
            </div>
            <div className="flex flex-col w-full items-center">
              <input type="submit" className="bg-secondary h-13 w-1/2 rounded-xl neon-xl hover:cursor-pointer" value="Register"/>
              <a className="hover:cursor-pointer" onClick={() => {
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
