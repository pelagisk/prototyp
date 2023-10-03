// import './Row.css'


function Row({ files, setFiles, index }) {
  const file = files[index]
  return (
    <tr>
        <td>{ file.Filetype }</td>
        <td>{ file.Filename }</td>
        <td>{ file.Description }</td>
        <td>{ file.Uploader }</td>
        <td>{ file.UnixTimestamp }</td>
        <td>Remove file</td>
    </tr>  
  );
}

export default Row;
