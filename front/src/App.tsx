import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { Todo } from './pages/Todo'
import { Auth } from './pages/Auth'

function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path="/todo" element={<Todo />} />
          <Route path="/" element={<Auth />} />
        </Routes>
      </BrowserRouter>
    </>
  )
}

export default App
