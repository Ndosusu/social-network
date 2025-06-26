export default function Home() {
    return (
        <div className="text-white h-full w-full grid items-center">
            <PostList/>
        </div>
    )
}

function PostList() {
    let temp = [{
        id: 1,
        user: "wiz",
        content: "text",
        nbLike: "132k",
        nbCom: "123"
    },{
        id: 2,
        user: "ziw",
        content: "abcd",
        nbLike: "1.8k",
        nbCom: "90"
    }]

    return(
        <div className="bg-primaryT h-5/4 w-2/3 neon-xl center grid items-center">
            <div className="w-full h-screen overflow-scroll flex flex-col items-center p-4 gap-7">
                {temp.map((obj, index) => (<CreatePost post={obj} key={obj.id}/>))}
            </div>
        </div>
    )
}

function CreatePost(data) {
    const post = data.post
    return (
        <div className="w-5/6 rounded-xl neon-sm">
            <div className="w-full postHeader bg-primaryT p-2">
                {post.user}
            </div>
            <div className="w-full h-fit p-4">
                {post.content}
            </div>
            <div className="p-3 flex w-full gap-4">
                <div className="w-1/10 flex items-center">
                    <img src="/like.svg" className="h-8"></img>
                    <p>{post.nbLike}</p>
                </div>
                <div className="w-1/10 flex items-center gap-1">
                    <img src="/comment.svg" className="h-8"></img>
                    <p>{post.nbCom}</p>
                </div>
            </div>
        </div>
    )
}