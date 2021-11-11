import { useState, useEffect } from 'react'
import logo from './logo.svg'
import './App.css'

interface IAlert{
  title: string;
  interval_second: number;
  count: number;
}

interface IAlertAPIResponse{
  status: number;
  data: IAlert[];
  msg: string;
}

function App() {
  const [alerts, setAlerts] = useState<IAlert[]>([]);
  useEffect(function(){
    setInterval(function(){
    fetch('http://polka:9000/api/alert')
    .then((x) => x.json())
    .then((j: IAlertAPIResponse) => setAlerts(j.data));
      }, 20000);
    }, [])

  return (
    <ul>
      {
        alerts.map((a) => (
          <li>
            {a.title}:{a.interval_second}
          </li>
        ))
        }
    </ul>
  )
}

export default App
