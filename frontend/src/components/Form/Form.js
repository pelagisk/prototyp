import React from "react"
import axios from "axios"
import config from "../../config"

function Form() {

  return (
    <form>        
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
