import { Button } from '@/components/ui/button'
import { Task } from '@/types/index'
import axios from 'axios'
import { useEffect, useState } from 'react'

export const Todo = () => {
  const [tasks, setTasks] = useState<Omit<Task, 'created_at' | 'updated_at'>[]>(
    [{ id: 0, title: 'task0' }]
  )

  const [selectedTasks, setSelectedTasks] = useState<Set<number>>(new Set())

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
      <div className="flex flex-col items-center space-y-3 mt-3">
        {tasks.map((task) => {
          return (
            <div className="flex space-x-2" key={task.id}>
              <label className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70">
                <input
                  type="checkbox"
                  onChange={(e) =>
                    handleCheckboxChange(task.id, e.target.checked)
                  }
                  checked={selectedTasks.has(task.id)}
                />
                {task.title}
              </label>
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
