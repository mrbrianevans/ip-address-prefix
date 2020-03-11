# IP Address prefix finder
This small Go application gets the first three subdivisions of an IP address. <br>
For example, 34.240.50.53 would output 34.240.50 <br>
This is useful in web analytics as a user might get reassigned a network IP address, and the last subdivision changes, but you want to associate the tracking of the new IP address with the old. 

<br>
<br>
To run this application, download the .exe file and navigate to its location in command line using `cd`.<br>
Type `ipaddressprefix` followed by a list of IP address in the form "ip", "ip2", "ip3" etc and press enter to run.
Example:<br>
`>> ipaddressprefix "69.160.160.133", "130.43.180.51", "92.40.200.148", "5.151.93.193", "34.240.50.53"`<br>
The output will be: first repeating the inputs, then each prefix on a new line. For the example it would look like this:<br>
`>>[69.160.160.133, 130.43.180.51, 92.40.200.148, 5.151.93.193]`<br>
`>>69.160.160`<br>
`>>130.43.180`<br>
`>>92.40.200`<br>
`>>5.151.93` <br><br>
This application is written in Go, and the source code is available, but you do not need to know Go or have Go installed to run it. Any Windows machine can run this in command line. <br><br>
Finally, if you have any requests or suggestions on improving the tool (for example reading IP addresses from a file) please let me know.
