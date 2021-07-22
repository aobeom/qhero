import './App.css'
import { useState } from "react"

function App() {
  const [mediaURL, setMediaURL] = useState("")
  const [mediaData, setMediaData] = useState([])

  const urlChange = (event) => {
    setMediaURL(event.target.value)
  }

  const GetImgs = () => {
    var regex = /http(s)?:\/\/([\w-]+.)+[\w-]+(\/[\w- ./?%&=]*)?/

    if (!regex.test(mediaURL) || mediaURL.indexOf("mdpr.jp") === -1) {
      alert("URL Error")
      return false
    }

    let mediaURLClear = mediaURL.split(" ")
    let mediaURLNew = mediaURLClear[mediaURLClear.length - 1]
    let apiurl = "/api/mdpr?url=" + mediaURLNew

    fetch(apiurl, {
      method: 'GET',
      dataType: 'json'
    }).then(res => res.json())
      .then(data => {
        let status = data.status
        if (status === 1) {
          setMediaData(data.data)
        } else {
          alert(data.message)
        }
      })
      .catch(
        () => {
          alert("Server Error")
        }
      )
  }
  return (
    <div className="App">
      <header className="App-header">
        <div>
          <input value={mediaURL} onChange={(event) => urlChange(event)} />
        </div>
        <div>
          <button onClick={() => GetImgs()} >Get</button>
        </div>
      </header>
      <main className="App-main">
        {mediaData.map((media, index) => {
          return <div key={index}>
            <img src={media} alt="" className='App-result-img' />
          </div>
        })}
      </main>
    </div>
  );
}

export default App
