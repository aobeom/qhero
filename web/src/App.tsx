import './App.css'
import { useState } from "react"
import "mini.css"

function App() {
  const [mediaURL, setMediaURL] = useState<string>("")
  const [mediaData, setMediaData] = useState<any>([])
  const [mediaMsg, setMediaMsg] = useState<any>({ open: false, msg: "" })
  const [mediaLoading, setMediaLoading] = useState<boolean>(false)

  const urlChange = (event: any) => {
    setMediaURL(event.target.value)
  }

  const GetImgs = () => {
    setMediaMsg({ open: false, msg: "" })
    setMediaLoading(true)

    var regex: RegExp = /http(s)?:\/\/([\w-]+.)+[\w-]+(\/[\w- ./?%&=]*)?/

    if (!regex.test(mediaURL) || mediaURL.indexOf("mdpr.jp") === -1) {
      setMediaLoading(false)
      setMediaMsg({ open: true, msg: "URL Error" })
      return false
    }

    let mediaURLClear: any = mediaURL.split(" ")
    let mediaURLNew: string = mediaURLClear[mediaURLClear.length - 1]
    let apiurl: string = "/api/mdpr?url=" + mediaURLNew

    fetch(apiurl, {
      method: 'GET',
    }).then(res => res.json())
      .then(data => {
        let status: number = data.status
        if (status === 1) {
          setMediaLoading(false)
          setMediaMsg({ open: false, msg: "" })
          setMediaData(data.data)
        } else {
          setMediaLoading(false)
          setMediaMsg({ open: true, msg: data.message })
        }
      })
      .catch(
        () => {
          setMediaLoading(false)
          setMediaMsg({ open: true, msg: "Server Error" })
        }
      )
  }
  return (
    <div className="App">
      <div className="App-header">
        <div>
          <input
            placeholder="mdpr url"
            value={mediaURL}
            onChange={(event) => urlChange(event)}
            className="App-input"
          />
          <button
            className="primary"
            onClick={() => GetImgs()}
          >
            GET
          </button>
        </div>
      </div>
      <div className="App-main">
        {mediaLoading ? <div className="spinner primary"></div> :
          mediaMsg.open ? <button className="secondary">{mediaMsg.msg}</button > :
            mediaData.map((media: any, index: number) => {
              return <div key={index}>
                <img src={media} alt="" className='App-result-img' />
              </div>
            })
        }
      </div>
    </div>
  );
}

export default App
