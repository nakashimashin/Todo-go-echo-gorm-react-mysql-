'use client'

import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
import * as z from 'zod'

import { Button } from '@/components/ui/button'
import { Checkbox } from '@radix-ui/react-checkbox'
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form"
import { toast } from "@/components/ui/use-toast"

const tasks = [
  {
    id: "1",
    task: "task1",
  },
  {
    id: "2",
    task: "task2",
  },
  {
    id: "3",
    task: "task3",
  },
] as const

const FormSchema = z.object({
  tasks: z.array(z.string()).refine((value) => value.some((task) => task), {
    message: "タスクを選択してください",
  }),
})

export const Todo = () => {
  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      tasks: ["1", "task0"],
    },
  })

  function onSubmit(data: z.infer<typeof FormSchema>) {
    toast({
      title: "You submitted the following tasks:",
      description: (
        <pre className='mt-2 w-[340px] rounded-md bg-slate-950 p-4'>
          <code className='text-white'>{JSON.stringify(data, null, 2)}</code>
        </pre>
      )
    })
  }

  return (
    <Form {...form}>
      <form className='space-y-8'>
        <FormField 
          control={form.control}
          name  = "tasks"
          render={() => (
            <FormItem>
              <div className='mb-4'>
                <FormLabel className='text-base'>Todoリスト</FormLabel>
                <FormDescription>
                  ここにタスクを追加してください。
                </FormDescription>
              </div>
              {tasks.map((task) => (
                <FormField 
                  key={task.id}
                  control={form.control}
                  name='tasks'
                  render={({ field }) => {
                    return (
                      <FormItem
                        key={task.id}
                        className='flex flex-row items-start space-x-3 space-y-0'
                      >
                        <FormControl>
                          <Checkbox
                            checked={field.value?.includes(task.id)}
                            onCheckedChange={(checked) => {
                              return checked
                                ? field.onChange([...field.value, task.id])
                                : field.onChange(
                                    field.value?.filter(
                                      (value) => value !== task.id
                                    )
                                )                             
                            }}
                           />
                        </FormControl>
                        <FormLabel className='font-normal'>
                          {task.task}
                        </FormLabel>
                      </FormItem>
                    )
                  }}
                 />
              ))}
              <FormMessage />
            </FormItem>
          )}
         />
         <Button type='submit'>Submit</Button>
      </form>
    </Form>
  )
}
