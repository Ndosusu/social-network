export default function Home() {
    return (
        <div className="text-white h-full w-full grid items-center">
            <PostList/>
        </div>
    )
}

function PostList() {
    return(
        <div className="bg-primaryT h-5/4 w-2/3 neon-xl center grid items-center">
            <div className="w-full h-screen overflow-scroll flex flex-col items-center p-4 gap-7">
                <div className="w-5/6 rounded-xl neon-sm">
                    <div className="w-full postHeader bg-primaryT p-2">
                        wiz
                    </div>
                    <div className="w-full h-fit p-2">
                        text
                    </div>
                    <img src="/comment.svg"></img>
                    <img src="/like.svg"></img>
                </div>
            </div>
        </div>
    )
}