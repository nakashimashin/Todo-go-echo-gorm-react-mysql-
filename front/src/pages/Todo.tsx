import { Button } from '@/components/ui/button'
import { Task } from '@/types/index'
import axios from 'axios'
import React, { useEffect, useState } from 'react'
import { GrUpdate } from 'react-icons/gr'
import { useNavigate } from 'react-router-dom'

export const Todo = () => {
  useEffect(() => {
    ;(async () => {
      try {
        const token = localStorage.getItem('token')
        const res = await axios.get(
          `${import.meta.env.VITE_REACT_APP_API_URL}/api/tasks`,
          {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          }
        )
        console.log(res.data)
        setTasks(res.data)
      } catch (error) {
        console.error('Error fetching tasks', error)
      }
    })()
  }, [])

  const navigate = useNavigate()
  const [tasks, setTasks] = useState<Omit<Task, 'created_at' | 'updated_at'>[]>(
    [{ id: 0, title: 'task0' }]
  )
  const [newTask, setNewTask] = useState<
    Omit<Task, 'created_at' | 'updated_at' | 'id'>
  >({ title: '' })
  const [selectedTasks, setSelectedTasks] = useState<Set<number>>(new Set())
  const [editingTaskId, setEditingTaskId] = useState<number | null>(null)
  const [editingTitle, setEditingTitle] = useState('')

  const handleNewTaskChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setNewTask({ title: e.target.value })
  }

  const createTask = async () => {
    const token = localStorage.getItem('token')
    try {
      const res = await axios.post(
        `${import.meta.env.VITE_REACT_APP_API_URL}/api/task`,
        newTask,
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      )
      console.log(res.data)
      setTasks([...tasks, res.data])
      setNewTask({ title: '' })
    } catch (error) {
      console.error('Error creating task', error)
    }
  }

  const handleCheckboxChange = (taskId: number, isChecked: boolean) => {
    setSelectedTasks((prevSelectedTasks) => {
      const newSelectedTasks = new Set(prevSelectedTasks)
      if (isChecked) {
        newSelectedTasks.add(taskId)
      } else {
        newSelectedTasks.delete(taskId)
      }
      return newSelectedTasks
    })
  }

  const deleteSelectedTasks = async () => {
    const token = localStorage.getItem('token')
    try {
      for (const taskId of selectedTasks) {
        await axios.delete(
          `${import.meta.env.VITE_REACT_APP_API_URL}/api/task/${taskId}`,
          {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          }
        )
      }
      setTasks(tasks.filter((task) => !selectedTasks.has(task.id)))
      setSelectedTasks(new Set())
    } catch (error) {
      console.error('Error deleting tasks', error)
    }
  }

  const handleEditClick = (task: Omit<Task, 'created_at' | 'updated_at'>) => {
    setEditingTaskId(task.id)
    setEditingTitle(task.title)
  }

  const handleEditingChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setEditingTitle(e.target.value)
  }

  const saveEdit = async (taskId: number) => {
    const token = localStorage.getItem('token')
    try {
      const res = await axios.put(
        `${import.meta.env.VITE_REACT_APP_API_URL}/api/task/${taskId}`,
        {
          title: editingTitle,
        },
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      )
      setTasks(tasks.map((task) => (task.id === taskId ? res.data : task)))
      setEditingTaskId(null)
      console.log(res.data)
    } catch (error) {
      console.error('Error updating task', error)
    }
  }

  const handleLogout = () => {
    localStorage.removeItem('token')
    navigate('/')
  }

  return (
    <div className="flex flex-col items-center justify-center mt-[50px]">
      <div className="text-[30px]">Todoリスト</div>
      <button
        onClick={handleLogout}
        className="mt-3 bg-blue-500 hover:bg-blue-300 w-[80px] h-[30px] border rounded ml-3 font-bold text-white text-[15px]"
      >
        ログアウト
      </button>
      <div>
        <input
          type="text"
          placeholder="タスクを追加"
          className="border border-black"
          value={newTask.title}
          onChange={handleNewTaskChange}
        />
        <button
          onClick={createTask}
          className="mt-3 bg-red-500 hover:bg-red-300 w-[50px] h-[30px] border rounded ml-3 font-bold text-white text-[15px]"
        >
          追加
        </button>
      </div>
      <div className="flex flex-col justify-start space-y-3 mt-3">
        {tasks.map((task) => {
          return (
            <div className="flex space-x-2" key={task.id}>
              <input
                type="checkbox"
                onChange={(e) =>
                  handleCheckboxChange(task.id, e.target.checked)
                }
                checked={selectedTasks.has(task.id)}
                className="cursor-pointer"
              />
              {editingTaskId === task.id ? (
                <input
                  type="text"
                  value={editingTitle}
                  onChange={handleEditingChange}
                  onKeyDown={(e) => {
                    if (e.key === 'Enter') {
                      saveEdit(task.id)
                    }
                  }}
                  className="border border-black"
                />
              ) : (
                <label className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70">
                  {task.title}
                </label>
              )}
              <GrUpdate
                onClick={() => handleEditClick(task)}
                className="text-[15px] cursor-pointer"
              />
            </div>
          )
        })}
      </div>
      <Button className="mt-3" onClick={deleteSelectedTasks}>
        削除
      </Button>
    </div>
  )
}
