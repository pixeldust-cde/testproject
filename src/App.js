import React, {useState} from 'react';
import './App.css';

function App() {
  const [iteration, setIteration] = useState(0)
  const [message, setResponseMessage] = useState('?')
  const baseUrl = 'http://127.0.0.1:8081/fib/v1/iterate'

  function handleSubmit(e){
    e.preventDefault();
    setResponseMessage("...calculating...")
    console.log(iteration);
    const headers = {'iteration': iteration}
    fetch(baseUrl, { headers: headers})
    .then(response => response.text())
    .then(data => setResponseMessage(data))
    .catch(err => console.log(err));
  }

  const handleChange = e => {
    setIteration(e.target.value);
  }

  return (
    <div className="App">
      <h4>FiboServe 1.0</h4>
      <form onSubmit={handleSubmit}>
      <p>Please enter the x for your f(x).</p>
      <input id="iteration" type="text" onChange={handleChange}></input>
      <input type="submit" value="Submit"></input>
      <p>f(x) = {message}</p>
      </form>
    </div>
  );
}

export default App;
