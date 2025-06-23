export default function Home() {
  return (
    <div className="h-full w-full flex justify-center items-center text-white">
      <section id="login" className="w-1/3 h-2/4">
        <form className="w-full h-full flex flex-col bg-primary p-10 justify-between items-center neon-xl rounded-3xl" method="post">
          <div className="flex flex-col h-fit gap-5 w-full justify-between">
            <input type="text" className="bg-primary h-11 p-2 neon-sm rounded-xl" placeholder="Mail"/>
            <input type="password" className="bg-primary h-11 p-2 neon-sm rounded-xl" placeholder="Password"/>
          </div>
          <input type="submit" className="bg-secondary h-13 w-1/2 rounded-xl neon-xl" value="Log in"/>
        </form>
      </section>

      <section id="register" className="">

      </section>
    </div>
  );
}
