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
                    <div className="w-full h-fit p-4">
                        text
                    </div>
                    <div className="p-3 flex w-full gap-4">
                        <div className="w-1/10 flex items-center">
                            <img src="/like.svg" className="h-8"></img>
                            <p>113k</p>
                        </div>
                        <div className="w-1/10 flex items-center gap-1">
                            <img src="/comment.svg" className="h-8"></img>
                            <p>189</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}