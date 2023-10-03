import React from "react"
import axios from "axios"
import config from "../../config"
import './Form.css'

function Form({ toggleUploadView }) {

  // ref to store the form dataw
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
      // back to the table view and trigger update of state so that new file list is downloaded!
      toggleUploadView()
    })
    .catch((response) => {
      console.log(response)
    })
  }
  
  return (
    <form onSubmit={handleSubmit} ref={refForm}>        
      <input type="file" name="file" placeholder="Filename" />
      <br />
      <input type="text" name="description" placeholder="Description" />
      <br />
      <input type="text" name="uploader" placeholder="Uploader" />
      <br />
      <input className="button" type="submit" value="Upload file" />
    </form>        
  )
}

export default Form;
