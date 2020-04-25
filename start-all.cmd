start cmd /c "bin\webdav.exe -addr 127.0.0.1:3001 -root .data/1"
start cmd /c "bin\webdav.exe -addr 127.0.0.1:3002 -root .data/2"
start cmd /c "bin\webdav.exe -addr 127.0.0.1:3003 -root .data/3"
start cmd /c "bin\mkv.exe -volumes localhost:3001,localhost:3002,localhost:3003 -db .data/indexdb server"
