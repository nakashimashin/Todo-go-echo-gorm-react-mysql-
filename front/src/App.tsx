import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { Todo } from './pages/Todo'
import { Auth } from './pages/Auth'

function App() {

  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Todo />} />
          <Route path="/auth" element={<Auth />} />
        </Routes>
      </BrowserRouter>
    </>
  )
}

export default App
