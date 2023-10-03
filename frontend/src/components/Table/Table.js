import React from 'react'
import axios from 'axios'
import config from '../../config'
import './Table.css';
import Row from '../Row/Row'


function Table() {

  const [files, setFiles] = React.useState(null);

  React.useEffect(() => {
    axios.get(config.URL_OF_API + '/files').then((response) => {
      setFiles(response.data);
    })
    .catch((response) => {
      console.log(response.data)
    })
  }, []);

  if (!files) return null

  const tableRows = [...files.keys()].map(
    index => {
      return (
        <Row key={index} files={files} setFiles={setFiles} index={index} />
      )
    }
  )

  return (
    <table>    
      <thead>
        <tr>
          <td></td>
          <td>Filename</td>
          <td>Description</td>
          <td>Uploader</td>
          <td>Date</td>
          <td></td>
        </tr>      
      </thead>
      <tbody>   
        { tableRows }
      </tbody>       
    </table>
  );
}

export default Table;
