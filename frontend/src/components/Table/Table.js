// import './Table.css';
import Row from '../Row/Row';


function Table() {

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
        <Row />
      </tbody>       
    </table>
  );
}

export default Table;
