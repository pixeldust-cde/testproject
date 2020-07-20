import React, {useState} from 'react';
import axios from 'axios';
import './App.css';

function App() {
  const [iteration, setIteration] = useState(0)
  const [message, setResponseMessage] = useState('')
  
  const handleSubmit = e => {
    e.preventDefault();
    axios.get('localhost:8081/fib/v1/iterate').then(res =>{
      console.log(res.data);
      setResponseMessage(res.data);
    }).catch(err=>{
      console.log(err);
    })
  }

  const handleChange = e => {
    setIteration(e.target.value);
  }

  return (
    <div className="App">
      <h4>FiboServe 1.0</h4>
      <p>{message}</p>
      <form onSubmit={handleSubmit}>
      <p>Please enter the x for your f(x).</p>
      <input id="iteration" type="text" onChange={handleChange}></input>
      <input type="submit" value="Submit"></input>
      </form>
    </div>
  );
}

export default App;
