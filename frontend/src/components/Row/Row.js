import axios from 'axios'
import config from '../../config'
import './Row.css'
import FileIcon from '../FileIcon/FileIcon'
import UploadDate from '../UploadDate/UploadDate'

function Row({ files, setFiles, index }) {
  const file = files[index]

  // creates a downloading anonymous function for Axios, for the given filename
  // see https://stackoverflow.com/questions/41938718/how-to-download-files-using-axios
  const fileDownloader = (filename) => {
    return (response) => {
        // create file link in browser's memory
        const href = URL.createObjectURL(response.data);    
        // create "a" HTML element with href to file & click
        const link = document.createElement('a');
        link.href = href;
        link.setAttribute('download', filename); //or any other extension
        document.body.appendChild(link);
        link.click();    
        // clean up "a" element & remove ObjectURL
        document.body.removeChild(link);
        URL.revokeObjectURL(href);
    }
  }

  const invokeDownload = () => {
    axios({
      url: config.URL_OF_API + '/files/' + file.ID,
      method: 'GET',
      responseType: 'blob',
    })
    .then(fileDownloader(file.Filename))
    .catch((response) => {
      console.log(response.data)
    });
  }

  const invokeDelete = () => {
    // first delete file from store and db
    axios.delete(config.URL_OF_API + '/files/' + file.ID)
    .then((response) => {
      console.log(response)
      // then find the file to be deleted
      const matches = files.reduce((arr, e, i) => {
        if (e.ID === file.ID) { arr.push(i) }
        return arr
      }, [])
      console.log(matches)
      // if exactly one match is found, remove it from files and update state
      if (matches.length === 1) {
        files.splice(index, 1)
        setFiles([...files])
        console.log("Deleted file!")
      } else {
        console.log("Error in deleting from prop 'files'. Number of matches found: ", matches.length)
      }      
    })
    .catch((response) => {
      console.log(response.data)
    })
  }

  return (
    <tr key={file.ID} className={'row-' + (index % 2)}>
        <td><FileIcon mime={file.Mime} /></td>
        <td><button className="link-button" onClick={() => invokeDownload()}>{file.Filename}</button></td>
        <td>{ file.Description }</td>
        <td>{ file.Uploader }</td>
        <td><UploadDate unixTimestamp={file.UnixTimestamp} /></td>
        <td><button className="link-button" onClick={() => invokeDelete()}>Delete</button></td>
    </tr>  
  );
}

export default Row;
