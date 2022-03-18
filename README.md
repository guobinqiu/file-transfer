# Resumable File Upload Demo

### Run

1. Run `go run generateLargeFile` to get a 1GB file for mock purpose.

2. Run `go run fileUploadServer.go` to get the `upload` directory running as a file system.

3. Open `fileUpload.html` in browser to choose files to be uploaded.

### Notes

After getting file upload server up and running 
which start to mornitor changes from `uploads`directory right away through file hook 
written in the `fileUploadServer.go`. You could implement your own business
logic from inside respective file hooks.

### Ouputs

```bash
Guobins-MacBook-Pro:file-transfer guobin$ go run fileUploadServer.go 
  

[tusd] 2021/01/18 10:56:22 event="RequestIncoming" method="OPTIONS" path="" requestId="" 
[tusd] 2021/01/18 10:56:22 event="ResponseOutgoing" status="200" method="OPTIONS" path="" requestId="" 
[tusd] 2021/01/18 10:56:22 event="RequestIncoming" method="POST" path="" requestId="" 
[tusd] 2021/01/18 10:56:22 event="UploadCreated" id="ff81c21e2646fa1007d4f0e283a93b0b" size="1073741824" url="http://localhost:8080/files/ff81c21e2646fa1007d4f0e283a93b0b" 
[tusd] 2021/01/18 10:56:22 event="ResponseOutgoing" status="201" method="POST" path="" requestId="" 
Upload test.txt created
[tusd] 2021/01/18 10:56:22 event="RequestIncoming" method="OPTIONS" path="ff81c21e2646fa1007d4f0e283a93b0b" requestId="" 
[tusd] 2021/01/18 10:56:22 event="ResponseOutgoing" status="200" method="OPTIONS" path="ff81c21e2646fa1007d4f0e283a93b0b" requestId="" 
[tusd] 2021/01/18 10:56:22 event="RequestIncoming" method="PATCH" path="ff81c21e2646fa1007d4f0e283a93b0b" requestId="" 
[tusd] 2021/01/18 10:56:22 event="ChunkWriteStart" id="ff81c21e2646fa1007d4f0e283a93b0b" maxSize="1073741824" offset="0" 
Uploaded 12%
Uploaded 26%
Uploaded 39%
Uploaded 52%
Uploaded 62%
[tusd] 2021/01/18 10:56:28 event="ChunkWriteComplete" id="ff81c21e2646fa1007d4f0e283a93b0b" bytesWritten="708787576" 
Uploaded 66%
[tusd] 2021/01/18 10:56:28 event="ResponseOutgoing" status="204" method="PATCH" path="ff81c21e2646fa1007d4f0e283a93b0b" requestId="" 
[tusd] 2021/01/18 10:57:15 event="RequestIncoming" method="OPTIONS" path="ff81c21e2646fa1007d4f0e283a93b0b" requestId="" 
[tusd] 2021/01/18 10:57:15 event="ResponseOutgoing" status="200" method="OPTIONS" path="ff81c21e2646fa1007d4f0e283a93b0b" requestId="" 
[tusd] 2021/01/18 10:57:15 event="RequestIncoming" method="HEAD" path="ff81c21e2646fa1007d4f0e283a93b0b" requestId="" 
[tusd] 2021/01/18 10:57:15 event="ResponseOutgoing" status="200" method="HEAD" path="ff81c21e2646fa1007d4f0e283a93b0b" requestId="" 
[tusd] 2021/01/18 10:57:15 event="RequestIncoming" method="OPTIONS" path="ff81c21e2646fa1007d4f0e283a93b0b" requestId="" 
[tusd] 2021/01/18 10:57:15 event="ResponseOutgoing" status="200" method="OPTIONS" path="ff81c21e2646fa1007d4f0e283a93b0b" requestId="" 
[tusd] 2021/01/18 10:57:15 event="RequestIncoming" method="PATCH" path="ff81c21e2646fa1007d4f0e283a93b0b" requestId="" 
[tusd] 2021/01/18 10:57:15 event="ChunkWriteStart" id="ff81c21e2646fa1007d4f0e283a93b0b" maxSize="364954248" offset="708787576" 
Uploaded 77%
Uploaded 89%
Uploaded 100%
[tusd] 2021/01/18 10:57:18 event="ChunkWriteComplete" id="ff81c21e2646fa1007d4f0e283a93b0b" bytesWritten="364954248" 
Upload ff81c21e2646fa1007d4f0e283a93b0b finished
[tusd] 2021/01/18 10:57:18 event="ResponseOutgoing" status="204" method="PATCH" path="ff81c21e2646fa1007d4f0e283a93b0b" requestId="" 
```

### Reference

Frontend

> https://transloadit.com/

Backend

> http://tusd.io/

