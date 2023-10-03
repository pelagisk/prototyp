import React from "react"
import './App.css'
import Table from '../Table/Table'
import Form from '../Form/Form'

function App() {
  
  const [displayForm, setDisplayForm] = React.useState(false)

  const toggleUploadView = () => {
    setDisplayForm(!displayForm);
  }  

  if (displayForm === false) {
    return (
      <div className="App">      
        <Table />
        <button onClick={toggleUploadView.bind(null)}>Upload new file</button>
      </div>
    )
  } else {
    return (
      <div className="App">      
        <Form toggleUploadView={toggleUploadView} />
        <button className="button" onClick={toggleUploadView.bind(null)}>Cancel upload (back to File Archive)</button>
      </div>
    )
  }

}

export default App;
