// import './Row.css'


function Row({ files, setFiles, index }) {
  console.log(files[index])
  return (
    <tr>
        <td>{ index }</td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
    </tr>  
  );
}

export default Row;
