// import './Row.css'


function Row({ files, setFiles, index }) {
  console.log(files[index])
  return (
    <tr>
        <td>{ files[index].Filetype }</td>
        <td>{ files[index].Filename }</td>
        <td>{ files[index].Description }</td>
        <td>{ files[index].Uploader }</td>
        <td>{ files[index].UnixTimestamp }</td>
        <td>Remove file</td>
    </tr>  
  );
}

export default Row;
