"use client"
import { useRouter, useSearchParams } from "next/navigation"
import { CheckLogToken } from "../checkToken"
import { useEffect, useState } from "react"
import ActionMenu from "../actionMenu"
import NewPostModal from "./newPostModal"
import DetailPostModal from "./detailPostModal"

export default function Home() {
    const router = useRouter()
    CheckLogToken(router)
    
    const [loading, setLoading] = useState(true)
    const [posts, setPosts] = useState(null)
    const [curDetail, setDetail] = useState(null)
    const [curModal, setModal]= useState("")
    const [feed, setFeed] = useState("global")
    const [groupList, setGroupList] = useState(null)
    const [groupId, setGroup] = useState(null)
    
    useEffect(() => {
        if(feed == "group" && !groupId) {
            return
        }
        
        setPosts(null)
        setLoading(true)
        fetch("http://localhost:8080/feed/"+feed,{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                session_uuid: localStorage.getItem("logToken"),
                limit: 15,
                group_id: feed == "group" ? groupId : null,
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
    }, [feed, groupId])

    useEffect(() => {
        //fetch group list
        const temp = [
            {
                id: 1,
                name: "thug shaker central",
            },
            {
                id: 2,
                name: "wizards' den",
            },
            {
                id: 3,
                name: "kawaii desu neeeeee",
            },
            {
                id: 4,
                name: "Itadakimasuuuuuuuuu",
            }
        ]
        setGroupList(temp)
    }, [])

    const CheckState = () => {
        if (loading) {
            //replace later with good div instead of simple text
            return <p>loading...</p>
        }
        if (!groupId && feed == "group") {
            return groupList.map((obj, i) => <p key={i} onClick={() => {setGroup(obj.id)}}>{obj.name}</p>)
        }
        if (!posts) {
            //replace later with good div instead of simple text
            return <p>failed to fetch data.</p>
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
                return <NewPostModal posts={posts} postsFn={setPosts} modalFn={setModal} groupList={groupList} />
            }

            case "detailModal": {
                return <DetailPostModal postFeed={curDetail} />
            }
        }
    }

    return (
        <div className="text-white h-full w-full grid items-center text-xl">
            <div className="bg-primaryT h-6/4 w-2/3 neon-xl center grid items-center">
                <div className="w-full h-screen overflow-scroll flex flex-col ">
                    <div className="h-fit w-full flex flex-row justify-around p-3 px-10 gap-7">
                        <label htmlFor="globalFeed" className="neon-sm p-2 w-full h-fit flex flex-row rounded-xl text-center duration-100 hover:scale-110">
                            <input id="globalFeed" type="button" onClick={() => {setFeed("global")}} className="hidden" />
                            <p className="w-full">Global</p>
                        </label>
                        <label htmlFor="followFeed" className="neon-sm p-2 w-full h-fit flex flex-row rounded-xl text-center duration-100 hover:scale-110">
                            <input id="followFeed" type="button" onClick={() => {setFeed("follow")}} className="hidden" />
                            <p className="w-full">Followed</p>
                        </label>
                        <label htmlFor="groupFeed" className="neon-sm p-2 w-full h-fit flex flex-row rounded-xl text-center duration-100 hover:scale-110">
                            <input id="groupFeed" type="button" onClick={() => {setGroup(null); setFeed("group")}} className="hidden" />
                            <p className="w-full">Groups</p>
                        </label>
                    </div>
                    <div className="flex flex-col w-full overflow-scroll items-center p-4 gap-7">
                        <CheckState />
                    </div>
                </div>
            </div>
            <div className="fixed neon-xl w-1/10 h-fit max-h-5/6 left-5/6 top-1/12 postAction p-7">
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
    const postFeed = data.postFeed
    const post = postFeed.Post
    console.log(postFeed)
    const setDetail = data.setDetail
    const setModal = data.setModal

    const [liked, setLiked] = useState(postFeed.Like ? true : false)

    return (
        <div className="w-5/6 rounded-xl neon-sm duration-100 hoverable hover:scale-110" onClick={async () => {
            setDetail(postFeed)
            setModal("detailModal")
        }}>
            <div className="w-full postHeader bg-primaryT p-2">
                {post.Author.Nickname || "Author not found"}
            </div>
            <div className="w-full h-fit p-4">
                {post.Message || "Content not found"}
            </div>
            <div className="p-3 flex w-full gap-4">
                <label className="min-w-1/10 flex items-center" onClick={(e) => {e.stopPropagation()}}>
                    <input type="button" className="hidden" onClick={() => {likePost(postFeed, setLiked)}} />
                    <img src={liked ? "/likeActive.svg" : "/like.svg"} className={"h-8 "}></img>
                    <p className={liked ? "text-secondary" : null}>{postFeed.LikeCount || "0"}</p>
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
    fetch("http://localhost:8080/likes",
        postFeed.Like 
        ? {
            method: "DELETE",
            body: JSON.stringify({
                session_uuid: localStorage.getItem("logToken"),
                like_id: postFeed.Like.Id,
            }),
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
        console.log(response)
        switch(typeof response.Data.Result) {
            case "string": {
                fn(false)
            }
            case "object": {
                fn(true)
            }
        }
    })
}