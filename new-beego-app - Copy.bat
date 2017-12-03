@echo Set value of APP same as your app folder

set APP=myapi
set GOPATH=C:\Projects\Go
//set BEE=%GOPATH%\bin\bee
//%BEE% new %APP%
//cd %APP%

taskkill /F /IM chrome.exe /T
taskkill /F /IM chromeprocess /T
tskill chrome

echo cmd /c start /min "c:\Projects\Go\bin" myapi.exe  > run.bat
//echo "C:\Program Files (x86)\Mozilla Firefox\firefox.exe"  -kiosk -fullscreen http://localhost/manojvijayan/angular_hotel >> run.bat
echo  exit >> run.bat
start run.bat

"C:\Program Files (x86)\Google\Chrome\Application\chrome.exe" --incognito -kiosk -fullscreen http://localhost/manojvijayan/angular_hotel
 
//"C:\Program Files (x86)\Mozilla Firefox\firefox.exe"  -kiosk -fullscreen http://localhost/manojvijayan/angular_hotel 
//&& exit
//exit /b
cls


//pause
//start http://localhost/manojvijayan/angular_hotel
//"C:\Program Files (x86)\Google\Chrome\Application\chrome.exe" -kiosk -fullscreen http://localhost/manojvijayan/angular_hotel

