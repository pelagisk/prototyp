function ShowDate({ UnixTimestamp }) {
  
    var date = new Date(UnixTimestamp * 1000);
    var year = date.getFullYear();
    var month = String(date.getMonth() + 1).padStart(2, '0');
    var day = String(date.getDay() + 1).padStart(2, '0');
    var formattedDate = year + '-' + month + '-' + day;
    
      return (
          <p>{formattedDate}</p>
      );
  }
    
  export default ShowDate;