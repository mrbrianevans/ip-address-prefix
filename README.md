# IP Address prefix finder
This small Go application gets the first three subdivisions of an IPv4 address. 

For example, 34.240.50.53 would output 34.240.50 

This is useful in web analytics as a user might get reassigned a network IP address, and the last subdivision changes, but you want to associate the tracking of the new IP address with the old. 

## Usage
To run this application, download the .exe file and navigate to its location in command line using `cd`

Type `ipaddressprefix` followed by a list of IP address and press enter to run.
The IP addresses can be in quotation marks `"`, but they don't have to be. They can be comma seperated `, ` or just seperated by a single blank space ` `

Example:

```cmd
>> ipaddressprefix 69.160.160.133 130.43.180.51 92.40.200.148 5.151.93.193 34.240.50.53
```
The output will be: first repeating the inputs, then each prefix on a new line
```cmd
>>[69.160.160.133, 130.43.180.51, 92.40.200.148, 5.151.93.193]
>>69.160.160
>>130.43.180
>>92.40.200
>>5.151.93
```

## System requirements
This application is written in Go, and the source code is available, but you do not need to know Go or have Go installed to run it. Any Windows machine can run this in command line. 


Finally, if you have any requests or suggestions on improving the tool (for example reading IP addresses from a file) please let me know.
