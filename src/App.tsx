import React from 'react';
import logo from './logo.svg';
import './App.css';
import { useState } from 'react';
import TextField from '@mui/material/TextField';
import Typography from '@mui/material/Typography';




function App() {

  const [count, setCount] = useState(0)
  const[argA,setArgA] = useState(0)
  const[argB,setArgB] = useState(0)

  function runGo() {
    setCount(window.myGolangFunction(argA, argB))

  }


  return (
    <div className='App'>
      <header className='App-header'>
        
        <TextField label="A" variant="outlined" onChange={(e: React.ChangeEvent<HTMLInputElement>) => {setArgA(+e.target.value)}} />
        <TextField label="B" variant="outlined" onChange={(e: React.ChangeEvent<HTMLInputElement>) => {setArgB(+e.target.value)}} />
        <button onClick={runGo}>A+B = </button>
        <Typography variant="h1" component="h2">{count}</Typography>
      </header>

    </div>
  );
}

export default App;
