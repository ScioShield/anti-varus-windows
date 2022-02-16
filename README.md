# Anti-varus

The misspelling of anti-virus is deliberate, this is a project to further understand how antivirus software works, not replicate a real AV. I wouldn't want anyone use this software expecting a real antivirus.

## Alternative working AV software

| Windows                                       | Linux                              |
|-----------------------------------------------|------------------------------------|
| [Malwarebytes](https://www.malwarebytes.com/) | [ClamAV](https://www.clamav.net/)  |
| [Sophos](https://home.sophos.com/en-us)       |                                    |
| [Kaspersky](https://www.kaspersky.co.uk/)     |                                    |
| [Bitdefender](https://www.bitdefender.co.uk/) |                                    |

Blog post: TBD 

## Compilation  

You will need Golang installed  
Download the code to your Downloads folder (unzip the code)  
Open Powershell  
Open the Download folder (replace USERNAME with your username)  
<code>cd C:\Users\Incognito\Downloads\anti-varus-windows-master\anti-varus-windows-master</code>  
<code>go build anti-varus-windows.go</code>  
Place the varus.vdb file in the correct location (replace USERNAME with your username)  
<code>mkdir C:\Users\USERNAME\AppData\Local\anti-varus</code>  
<code>cp C:\Users\USERNAME\Downloads\varus.vdb C:\Users\USERNAME\AppData\Local\anti-varus</code>  
Run the software against a test file (the .vdb only has sigs for the vdb file, the unziped eicar file, and the compiled anti-varus-windows.exe)  
```
PS C:\Users\USERNAME\Downloads\anti-varus-windows-master\anti-varus-windows-master> .\anti-varus-windows.exe
Enter what file you would like to scan (full path and filename):  
C:\Users\USERNAME\Downloads\anti-varus-windows-master\anti-varus-windows-master\anti-varus-windows.exe  
Match found!  
Target hex: 78797a  
Target string: xyz
Target file: C:\Users\USERNAME\Downloads\anti-varus-windows-master\anti-varus-windows-master\anti-varus-windows.exe
File SHA256 Hash is: SHA256 hash of C:\Users\USERNAME\Downloads\anti-varus-windows-master\anti-varus-windows-master\anti-varus-windows.exe:
4acd2b62d46a3a4e160a1287ec0eb8591d2ac99a956c0fa00cfc390ef5bba6b6
CertUtil: -hashfile command completed successfully.
```

Do not run anti-varus against large files! 
