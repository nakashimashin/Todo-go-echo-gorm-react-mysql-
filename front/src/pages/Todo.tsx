'use client'

import { Checkbox } from '@/components/ui/checkbox'

const tasks = [
  {
    id: 1,
    task: "task1",
  },
  {
    id: 2,
    task: "task2",
  },
  {
    id: 3,
    task: "task3",
  },
] as const


export const Todo = () => {

  return (
    <div className='flex flex-col items-center justify-center mt-[50px]'>
      <div className='text-[30px]'>Todoリスト</div>
      <div className='flex flex-col items-center space-y-3 mt-3'>
        <div className='flex space-x-2'>
          <Checkbox id='terms' />
          <label 
            htmlFor="terms"
            className='text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70'
            >
              Accept terms and conditions
          </label>
        </div>
        <div className='flex space-x-2'>
          <Checkbox id='terms' />
          <label 
            htmlFor="terms"
            className='text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70'
            >
              Accept terms and conditions
          </label>
        </div>
      </div>
    </div>
  )
}
