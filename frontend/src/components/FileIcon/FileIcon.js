function FileIcon({ mime }) {

    // map MIME to suffix

    var fileTypesFromMime = {
			"application/pdf": "pdf",
			"image/jpeg":      "jpg",
			"application/xml": "xml",
			"text/xml":        "xml",
		}

    // verified during upload on backend
    if (!(mime in fileTypesFromMime)) {
      console.log("File type " + mime + " is not allowed!")
      return null
    }
    var fileType = fileTypesFromMime[mime]
    var iconPath = fileType + "-icon.png"
  
    return (
      <img src={iconPath} alt={fileType} width="32" />
    )
}
  
export default FileIcon
  

