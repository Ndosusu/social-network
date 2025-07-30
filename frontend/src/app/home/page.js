"use client"
import { useRouter, useSearchParams } from "next/navigation"
import { CheckLogToken } from "../checkToken"
import { useEffect, useState } from "react"
import ActionMenu from "../actionMenu"
import NewPostModal from "./newPostModal"
import DetailPostModal from "./detailPostModal"
import { DEFAULT_SERVER_PATH } from "../page"

export default function Home() {
    const router = useRouter()
    CheckLogToken(router)
    
    const [loading, setLoading] = useState(true)
    const [posts, setPosts] = useState([])
    const [curDetail, setDetail] = useState(null)
    const [curModal, setModal]= useState("")
    const [feed, setFeed] = useState("global")
    
    useEffect(() => {
        setPosts(null)
        setLoading(true)
        fetch(DEFAULT_SERVER_PATH + "feed/" + feed,{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                session_uuid: localStorage.getItem("logToken"),
                limit: 15,
            })
        })
        .catch(error => {
            setLoading(false)
            throw new Error(error)
        })

        .then(data => data.json())

        .then(response => {
            if(response.data) {
                setLoading(false)    
                setPosts(response.data.Result)
            } else {
                setLoading(false)
                throw new Error("No data.")
            }
        })
    }, [feed])

    const getNextPosts = async () => {
        fetch(DEFAULT_SERVER_PATH + "feed/" + feed, {
            method:"POST",
            body: JSON.stringify({
                session_uuid: localStorage.getItem("logToken"),
                limit: 15,
                last_id: posts[posts.length - 1].Post.Id,
            }),
        })
        .catch(error => {
            throw new Error(error)
        })
        .then(data => data.json())
        .then(response => {
            if(response.data.Result)
                setPosts(posts.concat(response.data.Result))
        })
    }

    const CheckState = () => {
        if (loading) {
            //replace later with good div instead of simple text
            return <p>loading...</p>
        }
        if (!posts) {
            //replace later with good div instead of simple text
            return <p>No posts found.</p>
        }

        return posts.map((obj, i) => <CreatePost postFeed={obj} key={i} setDetail={setDetail} setModal={setModal} />)
    }

    const CreateModal = () => {
        return (
            <div id="modalDiv" className="w-screen h-screen absolute ">
                <div className="w-full h-full bg-black opacity-80 absolute z-5" onClick={async () => {setModal("")}}/>
                <CheckModalState />
            </div>
        )
    }

    const CheckModalState = () => {
        switch(curModal) {
            case "newPostModal": {
                return <NewPostModal posts={posts} postsFn={setPosts} modalFn={setModal} />
            }

            case "detailModal": {
                return <DetailPostModal postFeed={curDetail} postsFn={setPosts} modalFn={setModal} postsList={posts} />
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
                            <input id="globalFeed" name="feedRadio" type="radio" defaultChecked onClick={() => {setFeed("global")}} className="hidden" />
                            <p className="w-full">Global</p>
                        </label>
                        <label htmlFor="followFeed" className="neon-sm p-2 w-full h-fit flex flex-row rounded-xl text-center duration-100 hover:scale-110">
                            <input id="followFeed" name="feedRadio" type="radio" onClick={() => {setFeed("follow")}} className="hidden" />
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
                <div className="neon-sm p-5 rounded-xl flex flex-col items-center" onClick={() => {setModal("newPostModal")}}>
                    <img src="/new.svg" className="h-max"></img>
                    <p className="text-sm text-center">New post</p>
                </div>
            </div>
            <ActionMenu />
            {
                curModal != "" 
                ? <CreateModal /> 
                : null
            }
        </div>
    )
}

export function CreatePost(data) {
    const [postFeed, setPostFeed] = useState(data.postFeed)
    const post = postFeed.Post
    const setDetail = data.setDetail
    const setModal = data.setModal

    return (
        <div className="w-5/6 rounded-xl neon-sm duration-100 hoverable hover:scale-110" onClick={async () => {
            setDetail(postFeed)
            setModal("detailModal")
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
                    <input type="button" className="hidden" onClick={() => {likePost(postFeed, setPostFeed)}} />
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

export async function likePost(postFeed, fn) {
    const cloneFeed = structuredClone(postFeed)
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