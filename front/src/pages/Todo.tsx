import { Checkbox } from '@/components/ui/checkbox'
import { Task } from '@/types/index'

const tasks : Task[] = [
  {
    id: 1,
    title: "task1",
    created_at: new Date("2021-10-01"),
    updated_at: new Date("2021-10-01"),
  },
  {
    id: 2,
    title: "task2",
    created_at: new Date("2021-10-01"),
    updated_at: new Date("2021-10-01"),
  },
  {
    id: 3,
    title: "task3",
    created_at: new Date("2021-10-01"),
    updated_at: new Date("2021-10-01"),
  },
]


export const Todo = () => {

  return (
    <div className='flex flex-col items-center justify-center mt-[50px]'>
      <div className='text-[30px]'>Todoリスト</div>
      <div className='flex flex-col items-center space-y-3 mt-3'>
        {tasks.map((task: Task) => {
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
