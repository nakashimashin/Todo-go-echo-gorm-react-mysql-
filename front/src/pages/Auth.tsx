import React, { useState } from 'react'
import axios from 'axios'
import { useNavigate } from 'react-router-dom'

export const Auth = () => {
  const [email, setEmail] = useState<string>('')
  const [password, setPassword] = useState<string>('')
  const navigate = useNavigate()

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault()
    try {
      const res = await axios.post(
        `${import.meta.env.VITE_REACT_APP_API_URL}/login`,
        {
          email,
          password,
        }
      )
      console.log(res.data)
      localStorage.setItem('token', res.data.token)
      navigate('/todo')
    } catch (error) {
      console.error('Error login', error)
    }
  }
  return (
    <div>
      <form name="login_form" onSubmit={handleLogin}>
        <div className="flex flex-col">
          <label htmlFor="login">Login</label>
          <label htmlFor="password">Password</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="border border-black w-[130px]"
          />
          <label htmlFor="email">Email</label>
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="border border-black w-[130px]"
          />
        </div>
        <button className="mt-3 bg-red-500 hover:bg-red-300 w-[50px] h-[30px] border rounded font-bold text-white text-[15px]">
          ログイン
        </button>
      </form>
    </div>
  )
}
