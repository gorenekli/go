I will demonstrate here a simple server and client side of a Project developed in language “go”

<h2>The problem is:</h2>

There is a simple implementation of file server based on http.FileServer handle ( https://pkg.go.dev/net/http#example-FileServer ).

The server instance is running on top of simple file folder which doesn’t have nested subfolders.
Please implement client which downloads files using this server.
You should download a file containing char 'A' on earlier position than other files.
In case several files have the 'A' char on the same the earliest position you should download all of them.
Each TCP connection is limited by speed. The total bandwidth is unlimited.
You can use any disk space for temporary files.
The goal is to minimize execution time and data size to be transferred.
 
======================================================

Example

If the folder contains the following files on server:

'file1' with contents: "---A---"

'file2' with contents: "--A------"  

'file3' with contents: "------------"

'file4' with contents: "==A=========="

then 'file2' and 'file4' should be downloaded

<h2>For the solution:</h2>
We should consider the following conditions:

•	The directory is empty

•	None of the filenames include the letter “A”

•	“a” letter is also considered

•	Some filenames length is few or very long (filename with one letter or 50 letters)

After you run the server side code:

You can access the related pages as:

Localhost:8090

Localhost:8090/hello

Localhost:8090/headers

Localhost:8090/docs

Localhost:8090/root

/hello and /headers call the function hello and headers respectively.

/docs and /root list the files under /users/share/doc and / respectively.


On the client side: 

Reads the content through http://localhost:8090/docs/

Append all the filenames to the sfile[] list.

The first and last lines which consist of \<pre> and \</pre> won’t be appended.

With the use of index() method the leftmost indexes of letter “A” in each filenames are appended to winner[]. -1 is appended if there is no letter “A” in a filename.

The filenames who has the lowest indexes (and also >-1) are appended to dlfile[]. Finally filenames in dlfile[] are downloaded to the default directory.
