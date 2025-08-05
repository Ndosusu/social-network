"use client"
import { useRouter } from "next/navigation"
import { CheckApiResponse, CheckLogToken } from "../utils"
import {  useState } from "react"
import ActionMenu from "../actionMenu"
import NewPostModal from "./newPostModal"
import DetailPostModal from "./detailPostModal"
import { DEFAULT_SERVER_PATH } from "../page"
import { CreateAllInfoMessages, newInfoMessage } from "../infoMessage"
import { HomeProvider, useHomeContext } from "./contextProvider"

//needed to give the context to the whole Home page
export default function HomeContextWrapper() {
    return (
        <HomeProvider>
            <Home />
        </HomeProvider>
    )
}

//main page component
export function Home() {
    //create router to redirect and check if user allowed to access website
    const router = useRouter()
    CheckLogToken(router)

    console.log("Rendering HomePage...")

    const {
        feedLoading,
        currentFeed,
        feedPosts,
        modal,
    } = useHomeContext()

    //called when reaching the end of the post list to get the next ones
    const getNextPosts = () => {
        const lastId = feedPosts.val[feedPosts.val.length - 1].Post.Id

        console.log("Fetching posts after #" + lastId )

        //api call to get the post from the correct feed and give the id of the last post
        fetch(DEFAULT_SERVER_PATH + "feed/" + currentFeed.val, {
            method:"POST",
            body: JSON.stringify({
                session_uuid: localStorage.getItem("logToken"),
                limit: 15,
                last_id: lastId,
            }),
        })
        .catch(error => {
            throw new Error(error)
        })

        //make data readable as json object
        .then(data => data.json())

        //add to post list
        .then(response => {
            if(response.data.Result)
                feedPosts.set(feedPosts.val.concat(response.data.Result))
        })
    }

    //check the state of the feed and show the according div
    const CheckState = () => {
        if (feedLoading.val) {
            //replace later with good div instead of simple text
            return <p>loading...</p>
        }
        if (!feedPosts.val) {
            //replace later with good div instead of simple text
            return <p>No posts found.</p>
        }
        return feedPosts.val.map((obj, i) => <CreatePost postFeed={obj} key={i}/>)
    }

    //component to handle modals
    const CreateModal = () => {
        return (
            <div id="modalDiv" className="w-screen h-screen absolute z-10">
                <div className="w-full h-full bg-black opacity-80 absolute z-11" onClick={() => {modal.set("")}}/>
                <CheckModalState />
            </div>
        )
    }

    //check which modal to show
    const CheckModalState = () => {
        switch(modal.val) {
            case "newPostModal": {
                return <NewPostModal />
            }

            case "detailModal": {
                return <DetailPostModal />
            }
        }
    }

    return (
        <div className="text-white h-full w-full grid items-center text-xl">
            <div className="bg-primaryT h-6/4 w-2/3 center grid items-center relative">
                <div className="w-full h-full neon-xl absolute z-6 pointer-events-none" />
                <div className="w-full h-screen overflow-hidden flex flex-col">
                    <div className="h-fit w-full flex flex-row justify-around p-3 px-10 gap-7 bg-primary">
                        <label htmlFor="globalFeed" className="neon-sm p-2 w-full h-fit flex flex-row rounded-xl text-center duration-100 hover:scale-110">
                            <input id="globalFeed" name="feedRadio" type="radio" defaultChecked onClick={() => {currentFeed.set("global")}} className="hidden" />
                            <p className="w-full">Global</p>
                        </label>
                        <label htmlFor="followFeed" className="neon-sm p-2 w-full h-fit flex flex-row rounded-xl text-center duration-100 hover:scale-110">
                            <input id="followFeed" name="feedRadio" type="radio" onClick={() => {currentFeed.set("follow")}} className="hidden" />
                            <p className="w-full">Followed</p>
                        </label>
                    </div>
                    <div className="flex flex-col w-full flex-grow overflow-hidden items-center gap-7 relative">
                        <div className="absolute h-full w-full pointer-events-none fade z-5" />
                        <div className="h-full p-4 py-8 w-full flex flex-col items-center overflow-scroll gap-7" onScroll={(e) => {
                            //Check if user scrolled to the bottom, 1 is needed as a safety because scrollHeight and clientHeight are rounded numbers but not scrollTop
                            if(e.target.scrollHeight - e.target.clientHeight - e.target.scrollTop <= 1) 
                                getNextPosts()
                        }}>
                            <CheckState />
                        </div>
                    </div>
                </div>
            </div>
            <div className="fixed neon-xl w-1/10 h-fit max-h-5/6 left-5/6 top-1/12 postAction p-7 z-7">
                <div className="neon-sm p-5 rounded-xl flex flex-col items-center" onClick={() => {modal.set("newPostModal")}}>
                    <img src="/new.svg" className="h-max"></img>
                    <p className="text-sm text-center">New post</p>
                </div>
            </div>
            <CreateAllInfoMessages />
            <ActionMenu />
            {
                modal.val != "" 
                ? <CreateModal /> 
                : null
            }
        </div>
    )
}

//component to create a single post with the correct data
function CreatePost(data) {
    const {
        curPost,
        feedPosts,
        modal,
    } = useHomeContext()
    const [postFeed, setPostFeed] = useState(data.postFeed)
    const post = postFeed.Post

    const updatePostList = (post) => {
        const index = feedPosts.val.indexOf(postFeed)
        feedPosts.val[index] = post
    }
    
    return (
        <div className="w-5/6 rounded-xl neon-sm duration-100 hoverable hover:scale-110" onClick={() => {
            curPost.set(postFeed)
            modal.set("detailModal")
        }}>
            <div className="w-full postHeader bg-primaryT p-2 flex flex-row gap-4 items-center">
                <img src={post.Author.Avatar ? DEFAULT_SERVER_PATH + "data/images" + post.Author.Avatar : "defaultAvatar.svg"} className="h-10 rounded-xl" />
                {post.Author.Nickname || post.Author.FirstName + " " + post.Author.LastName || "Author not found"}
            </div>
            <div className="w-full h-fit p-4 flex flex-row justify-between gap-3">
                <p className="break-all">{post.Message || "Content not found"}</p>
                {
                    post.Image
                    ? <img src={ DEFAULT_SERVER_PATH + "data/images/" + post.Image} className="h-25 max-w-1/3" />
                    : null
                }
            </div>
            <div className="p-3 flex w-full gap-4">
                <label className="min-w-1/10 flex items-center duration-100 hover:scale-110" onClick={(e) => {e.stopPropagation()}}>
                    <input type="button" className="hidden" onClick={() => {likePost(postFeed, (post) => {
                            updatePostList(post)
                            setPostFeed(post)
                        })
                    }} />
                    <img src={postFeed.Like ? "/likeActive.svg" : "/like.svg"} className={"h-8 "}></img>
                    <p className={postFeed.Like ? "text-secondary" : null}>{postFeed.LikeCount || "0"}</p>
                </label>
                <div className="min-w-1/10 flex items-center gap-1">
                    <img src="/comment.svg" className="h-8"></img>
                    <p>{postFeed.CommentCount || "0"}</p>
                </div>
            </div>
        </div>
    )
}

//called when user likes a post
export async function likePost(postFeed, fn) {
    //make a clone so that the update state function works later
    const cloneFeed = structuredClone(postFeed)

    console.log("Changing like state of post #" + postFeed.Post.Id + ", becoming " + (postFeed.Like ? "Unliked" : "Liked"))
    
    //api call with a ternary to either delete the like or add it
    fetch(DEFAULT_SERVER_PATH + "likes",
        postFeed.Like 
        ? {
            method: "DELETE",
            body: JSON.stringify({
                session_uuid: localStorage.getItem("logToken"),
                like_id: postFeed.Like.Id,
            })
        }
        : {
            method: "POST",
            body: JSON.stringify({
                session_uuid: localStorage.getItem("logToken"),
                post_id: postFeed.Post.Id,
            })
        }
    )

    .catch(error => {
        throw new Error(error)
    })

    //make data readable as json object
    .then(data => data.json())

    //api returns "Ok" as a string if delete or the like as an object, act accordingly
    .then(response => {
        if(CheckApiResponse(response)) {
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
            newInfoMessage("Failed to like Post", "bg-red-500")
        }
    })
}