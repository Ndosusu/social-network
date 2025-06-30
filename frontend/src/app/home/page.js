export default function Home() {
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
    },{
        id: 2,
        user: "ziw",
        content: "abcd",
        nbLike: "1.8k",
        nbCom: "90"
    },{
        id: 2,
        user: "ziw",
        content: "abcd",
        nbLike: "1.8k",
        nbCom: "90"
    },{
        id: 2,
        user: "ziw",
        content: "abcd",
        nbLike: "1.8k",
        nbCom: "90"
    },{
        id: 2,
        user: "ziw",
        content: "abcd",
        nbLike: "1.8k",
        nbCom: "90"
    },{
        id: 2,
        user: "ziw",
        content: "abcd",
        nbLike: "1.8k",
        nbCom: "90"
    },{
        id: 2,
        user: "ziw",
        content: "abcd",
        nbLike: "1.8k",
        nbCom: "90"
    },{
        id: 2,
        user: "ziw",
        content: "abcd",
        nbLike: "1.8k",
        nbCom: "90"
    },{
        id: 2,
        user: "ziw",
        content: "abcd",
        nbLike: "1.8k",
        nbCom: "90"
    },{
        id: 2,
        user: "ziw",
        content: "abcd",
        nbLike: "1.8k",
        nbCom: "90"
    },{
        id: 2,
        user: "ziw",
        content: "abcd",
        nbLike: "1.8k",
        nbCom: "90"
    },{
        id: 2,
        user: "ziw",
        content: "abcd",
        nbLike: "1.8k",
        nbCom: "90"
    },{
        id: 2,
        user: "ziw",
        content: "abcd",
        nbLike: "1.8k",
        nbCom: "90"
    },{
        id: 2,
        user: "ziw",
        content: "abcd",
        nbLike: "1.8k",
        nbCom: "90"
    },{
        id: 2,
        user: "ziw",
        content: "abcd",
        nbLike: "1.8k",
        nbCom: "90"
    },{
        id: 2,
        user: "ziw",
        content: "abcd",
        nbLike: "1.8k",
        nbCom: "90"
    },{
        id: 2,
        user: "ziw",
        content: "abcd",
        nbLike: "1.8k",
        nbCom: "90"
    },{
        id: 2,
        user: "ziw",
        content: "abcd",
        nbLike: "1.8k",
        nbCom: "90"
    },{
        id: 2,
        user: "ziw",
        content: "abcd",
        nbLike: "1.8k",
        nbCom: "90"
    },{
        id: 2,
        user: "ziw",
        content: "abcd",
        nbLike: "1.8k",
        nbCom: "90"
    },{
        id: 2,
        user: "ziw",
        content: "abcd",
        nbLike: "1.8k",
        nbCom: "90"
    },{
        id: 2,
        user: "ziw",
        content: "abcd",
        nbLike: "1.8k",
        nbCom: "90"
    },{
        id: 2,
        user: "ziw",
        content: "abcd",
        nbLike: "1.8k",
        nbCom: "90"
    },{
        id: 2,
        user: "ziw",
        content: "abcd",
        nbLike: "1.8k",
        nbCom: "90"
    }]

    return (
        <div className="text-white h-full w-full grid items-center">
            <div className="bg-primaryT h-6/4 w-2/3 neon-xl center grid items-center">
                <div className="w-full h-screen overflow-scroll flex flex-col items-center p-4 gap-7">
                    {temp.map((obj, index) => (<CreatePost post={obj} key={index}/>))}
                </div>
            </div>
            <div className="fixed neon-xl w-1/10 h-fit max-h-5/6 left-5/6 top-1/12 postAction p-7">
                <div className="neon-sm p-5 rounded-xl flex flex-col items-center">
                    <img src="/new.svg" className="h-max"></img>
                    <p className="text-sm text-center">New post</p>
                </div>
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