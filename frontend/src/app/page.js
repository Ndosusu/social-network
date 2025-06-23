export default function Home() {
  return (
    <div className="h-full w-full flex justify-center items-center">
      <section id="login" className="w-1/3">
        <form className="grid bg-white gap-5">
          <input type="text" className="bg-red-300"></input>
          <input type="password" className="bg-red-300"></input>
          <input type="submit" className="bg-red-300"></input>
        </form>
      </section>
      <section id="register" className="">

      </section>
    </div>
  );
}
