export const Auth = () => {
  return (
    <div>
      <form name="login_form">
        <div className="flex flex-col">
          <label htmlFor="login">Login</label>
          <label htmlFor="password">Password</label>
          <input type="password" className="border border-black w-[130px]" />
          <label htmlFor="email">Email</label>
          <input type="email" className="border border-black w-[130px]" />
        </div>
        <button className="mt-3 bg-red-500 hover:bg-red-300 w-[50px] h-[30px] border rounded font-bold text-white text-[15px]">
          ログイン
        </button>
      </form>
    </div>
  )
}
