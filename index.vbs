Set WinScriptHost = CreateObject("WScript.Shell")
WinScriptHost.Run Chr(34) & "c:\wamp\wampmanager.exe" & Chr(34), 0
WinScriptHost.Run Chr(34) & "c:\Projects\Go\src\myapi\myapi.exe" & Chr(34), 0
WinScriptHost.Run Chr(34) & "c:\Projects\Go\src\myapi\browser.bat" & Chr(34), 0
Set WinScriptHost = Nothing