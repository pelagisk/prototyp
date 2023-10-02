curl -X POST http://localhost:8080/files \
    -H "Content-Type: multipart/form-data" \
    -F "description=Some description here" \
    -F "uploader=Axel Gagge" \
    -F "file=@testfiles/PLC.jpg"; 
echo "";
curl -X GET http://localhost:8080/files;
echo "";
curl -X GET http://localhost:8080/files/1;
echo "";
curl -X DELETE http://localhost:8080/files/1;