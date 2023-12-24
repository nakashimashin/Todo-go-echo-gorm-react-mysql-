import { Button } from '@/components/ui/button'
import { Task } from '@/types/index'
import axios from 'axios'
import { useEffect, useState } from 'react'
import { GrUpdate } from 'react-icons/gr'

export const Todo = () => {
  useEffect(() => {
    ;(async () => {
      try {
        const res = await axios.get(
          `${import.meta.env.VITE_REACT_APP_API_URL}/tasks`
        )
        console.log(res.data)
        setTasks(res.data)
      } catch (error) {
        console.error('Error fetching tasks', error)
      }
    })()
  }, [])

  const [tasks, setTasks] = useState<Omit<Task, 'created_at' | 'updated_at'>[]>(
    [{ id: 0, title: 'task0' }]
  )
  const [newTask, setNewTask] = useState<
    Omit<Task, 'created_at' | 'updated_at' | 'id'>
  >({ title: '' })
  const [selectedTasks, setSelectedTasks] = useState<Set<number>>(new Set())

  const handleNewTaskChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setNewTask({ title: e.target.value })
  }

  const createTask = async () => {
    try {
      const res = await axios.post(
        `${import.meta.env.VITE_REACT_APP_API_URL}/task`,
        newTask
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
    try {
      for (const taskId of selectedTasks) {
        await axios.delete(
          `${import.meta.env.VITE_REACT_APP_API_URL}/task/${taskId}`
        )
      }
      setTasks(tasks.filter((task) => !selectedTasks.has(task.id)))
      setSelectedTasks(new Set())
    } catch (error) {
      console.error('Error deleting tasks', error)
    }
  }

  return (
    <div className="flex flex-col items-center justify-center mt-[50px]">
      <div className="text-[30px]">Todoリスト</div>
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
      <div className="flex flex-col items-center space-y-3 mt-3">
        {tasks.map((task) => {
          return (
            <div className="flex space-x-2" key={task.id}>
              <input
                type="checkbox"
                onChange={(e) =>
                  handleCheckboxChange(task.id, e.target.checked)
                }
                checked={selectedTasks.has(task.id)}
              />
              <label className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70">
                {task.title}
              </label>
              <GrUpdate className="text-[15px]" />
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
