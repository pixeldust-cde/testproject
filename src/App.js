import React from 'react';
import './App.css';

function App() {
  return (
    <div className="App">
      <h4>FiboServe 1.0</h4>
      <form>
      <p>Please enter the iteration desired.</p>
      <input type="text" placeholder="iteration"></input>
      <input type="submit" value="Submit"></input>
      </form>
    </div>
  );
}

export default App;
