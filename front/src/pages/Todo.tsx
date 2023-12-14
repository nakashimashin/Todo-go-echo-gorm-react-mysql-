import { Checkbox } from '@/components/ui/checkbox'
import { Task } from '@/types/index'
import axios from 'axios'
import { useEffect, useState } from 'react'


export const Todo = () => {
  const [tasks, setTasks] = useState<Omit<Task, 'created_at' | 'updated_at'>[]>([{id: 0, title: 'task0'}])

  useEffect(() => {
    (
      async () => {
        try {
          const res = await axios.get('http://localhost:8080/tasks');
          console.log(res.data)
          setTasks(res.data)
        } catch (error) {
          console.error('Error fetching tasks', error)
        }
      }
    )()
  }, [])

  return (
    <div className='flex flex-col items-center justify-center mt-[50px]'>
      <div className='text-[30px]'>Todoリスト</div>
      <div className='flex flex-col items-center space-y-3 mt-3'>
        {tasks.map(task => {
          return (
            <div className='flex space-x-2' key={task.id}>
              <Checkbox />
              <label 
                className='text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70'
                >
                  {task.title}
              </label>
            </div>
          )
        })}
      </div>
    </div>
  )
}
