// import './Row.css'
import FileIcon from '../FileIcon/FileIcon'
import UploadDate from '../UploadDate/UploadDate'

function Row({ files, setFiles, index }) {
  const file = files[index]
  return (
    <tr>
        <td><FileIcon mime={file.Mime} /></td>
        <td>{ file.Filename }</td>
        <td>{ file.Description }</td>
        <td>{ file.Uploader }</td>
        <td><UploadDate unixTimestamp={file.UnixTimestamp} /></td>
        <td>Remove file</td>
    </tr>  
  );
}

export default Row;
