import React from "react"
import axios from "axios"
import config from "../../config"
import './Form.css'

function Form({ toggleUploadView }) {

  // state to specify if file was rejected during upload
  const [isRejected, setIsRejected] = React.useState(false)

  // ref to store the form data
  const refForm = React.createRef()

  // triggered by onSubmit 
  const handleSubmit = event => {    
    event.preventDefault()
    axios.post(
      config.URL_OF_API + '/files',
      new FormData(refForm.current),
      {
        headers: {
          "Content-Type": "multipart/form-data",
        }
      }
    )
    .then((response) => {
      setIsRejected(false)
      // back to the table view and trigger update of state so that new file list is downloaded!
      toggleUploadView()
    })
    .catch((response) => {
      console.log("Upload was rejected")
      console.log(response)
      setIsRejected(true)      
    })
  }

  // displays if the file was rejected
  const displayRejected = () => {
    if (isRejected === true) {
      return (
        <p className="red">The file was rejected! Try again!</p>
      )
    } else {
      return (
        <p></p>
      )
    }
  }
  
  return (
    <form onSubmit={handleSubmit} ref={refForm}>  
      { displayRejected() }
      <br />      
      <input type="file" name="file" placeholder="Filename" />
      <br />
      <input type="text" name="filename" placeholder="Change filename to (leave blank to keep uploaded filename)" />
      <br />
      <input type="text" name="description" placeholder="Description" />
      <br />
      <input type="text" name="uploader" placeholder="Uploader" />
      <br />
      <input className="button" type="submit" value="Upload file" />
    </form>        
  )
}

export default Form
