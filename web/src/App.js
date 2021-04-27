import logo from './logo.svg';
import './App.css';
import { useState } from "react"

function App() {
  const [seed, setSeed] = useState(-1)
  const GetSeed = () => {
    fetch("/api/seed", {
      method: 'GET',
      dataType: 'json'
    }).then(res => res.json())
      .then(data => {
        let status = data.status
        if (status === 1) {
          setSeed(data.data)
        } else {
          setSeed(-1)
        }
      })
      .catch(
        () => {
          setSeed(-1)
        }
      )
  }
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <div><input value={seed} /></div>
        <div>
          <button onClick={() => GetSeed()} >Get</button>
        </div>

      </header>
    </div>
  );
}

export default App;
