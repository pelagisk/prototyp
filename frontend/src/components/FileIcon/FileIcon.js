function FileIcon({ mime }) {

    // TODO refactor, this is ugly
    var fileType = 'none'
    switch (mime) {
      case 'application/pdf': 
        fileType = 'pdf'
        break
      case 'application/xml': 
        fileType = 'xml'
        break
      case 'image/jpeg':
        fileType = 'jpg'
        break
      default:
        console.log("File type " + mime + " is not allowed!")
        return null
    }
    var iconPath = fileType + "-icon.png"
  
    return (
        <img src={iconPath} alt={fileType} width="32" />
    )
}
  
export default FileIcon
  

