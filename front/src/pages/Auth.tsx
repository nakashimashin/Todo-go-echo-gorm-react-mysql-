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
    <div className="flex flex-col items-center justify-center">
      <form
        name="login_form"
        onSubmit={handleLogin}
        className="mt-[30px] flex items-center flex-col"
      >
        <div className="flex flex-col">
          <label htmlFor="login" className="flex justify-center text-[40px]">
            Login
          </label>
          <label htmlFor="password">Password</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="border border-black w-[250px] h-[40px]"
          />
          <label htmlFor="email">Email</label>
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="border border-black w-[250px] h-[40px]"
          />
        </div>
        <button className="mt-3 bg-red-500 hover:bg-red-300 w-[100px] h-[50px] border rounded font-bold text-white text-[20px]">
          ログイン
        </button>
      </form>
    </div>
  )
}
