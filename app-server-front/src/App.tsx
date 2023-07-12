import { useEffect, useState } from 'react'

type Topic = {
  id: number
  title: string
}

function App() {
  const [topics, setTopics] = useState<Topic[]>([])

  useEffect(()=>{
    setTopics([
      {
        id: 1,
        title: "Reactの使い方"
      },
      {
        id: 2,
        title: "Viteの使い方"
      }
    ])
  },[])

  return (
    <>
      <div>
        <h1>掲示板アプリ</h1>
        <h2>話題一覧</h2>
        <ul>
          {topics.map((topic) => (
            <li key={topic.id}>{topic.title}</li>
          ))}
        </ul>
      </div>
    </>
  )
}

export default App
