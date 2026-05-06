@echo off
chcp 65001 >nul
title 跨平台备份工具编译器

echo ================================================
echo    备份工具 - 全平台编译脚本 v2.0
echo ================================================
echo.

:: 检查 Go 是否安装
where go >nul 2>nul
if errorlevel 1 (
    echo ❌ 错误: 未找到 Go 编译器
    echo 请从 https://golang.org/dl/ 下载并安装 Go
    pause
    exit /b 1
)

:: 显示 Go 版本
echo [系统信息]
echo Go 版本: 
go version
echo.

:: 创建输出目录
set OUTPUT_DIR=build
if not exist "%OUTPUT_DIR%" mkdir "%OUTPUT_DIR%"
if not exist "%OUTPUT_DIR%\windows" mkdir "%OUTPUT_DIR%\windows"
if not exist "%OUTPUT_DIR%\linux" mkdir "%OUTPUT_DIR%\linux"
if not exist "%OUTPUT_DIR%\macos" mkdir "%OUTPUT_DIR%\macos"

set APP_NAME=backup-tool
set VERSION=2.0.0

echo [编译开始]
echo 程序: %APP_NAME%
echo 版本: %VERSION%
echo 输出目录: %OUTPUT_DIR%\
echo.

:: ===================== Windows =====================
echo ========== Windows 平台 ==========
echo.

:: Windows 64位
echo [1/6] Windows 64位 (amd64)...
set GOOS=windows
set GOARCH=amd64
go build -ldflags "-s -w" -o "%OUTPUT_DIR%\windows\%APP_NAME%-windows-amd64.exe"
if errorlevel 1 (
    echo ❌ 编译失败
) else (
    echo ✅ 成功: %APP_NAME%-windows-amd64.exe (%.0f KB)
    for %%F in ("%OUTPUT_DIR%\windows\%APP_NAME%-windows-amd64.exe") do echo    大小: %%~zF 字节
)

:: Windows 32位
echo.
echo [2/6] Windows 32位 (386)...
set GOOS=windows
set GOARCH=386
go build -ldflags "-s -w" -o "%OUTPUT_DIR%\windows\%APP_NAME%-windows-386.exe"
if errorlevel 1 (
    echo ⚠️  编译失败 (可能缺少32位工具链)
) else (
    echo ✅ 成功: %APP_NAME%-windows-386.exe
)

:: ===================== Linux =====================
echo.
echo ========== Linux 平台 ==========
echo.

:: Linux 64位
echo [3/6] Linux 64位 (amd64)...
set GOOS=linux
set GOARCH=amd64
go build -ldflags "-s -w" -o "%OUTPUT_DIR%\linux\%APP_NAME%-linux-amd64"
if errorlevel 1 (
    echo ⚠️  编译失败
) else (
    echo ✅ 成功: %APP_NAME%-linux-amd64
    echo    在Linux上使用: chmod +x %APP_NAME%-linux-amd64
)

:: Linux ARM64 (树莓派等)
echo.
echo [4/6] Linux ARM64 (arm64)...
set GOOS=linux
set GOARCH=arm64
go build -ldflags "-s -w" -o "%OUTPUT_DIR%\linux\%APP_NAME%-linux-arm64"
if errorlevel 1 (
    echo ⚠️  编译失败
) else (
    echo ✅ 成功: %APP_NAME%-linux-arm64
)

:: ===================== macOS =====================
echo.
echo ========== macOS 平台 ==========
echo.

:: macOS Intel
echo [5/6] macOS Intel (amd64)...
set GOOS=darwin
set GOARCH=amd64
go build -ldflags "-s -w" -o "%OUTPUT_DIR%\macos\%APP_NAME%-macos-amd64"
if errorlevel 1 (
    echo ⚠️  编译失败
) else (
    echo ✅ 成功: %APP_NAME%-macos-amd64
    echo    在macOS上使用: chmod +x %APP_NAME%-macos-amd64
)

:: macOS Apple Silicon
echo.
echo [6/6] macOS Apple Silicon (arm64)...
set GOOS=darwin
set GOARCH=arm64
go build -ldflags "-s -w" -o "%OUTPUT_DIR%\macos\%APP_NAME%-macos-arm64"
if errorlevel 1 (
    echo ⚠️  编译失败
) else (
    echo ✅ 成功: %APP_NAME%-macos-arm64
)

:: ===================== 验证文件 =====================
echo.
echo ========== 创建验证文件 ==========
echo.

:: 创建 README
(
echo # %APP_NAME% v%VERSION% - 跨平台备份工具
echo.
echo ## 已编译的平台版本
echo.
echo ### Windows
echo - %APP_NAME%-windows-amd64.exe (64位)
echo - %APP_NAME%-windows-386.exe (32位)
echo.
echo ### Linux
echo - %APP_NAME%-linux-amd64 (64位)
echo - %APP_NAME%-linux-arm64 (ARM64/树莓派)
echo.
echo ### macOS
echo - %APP_NAME%-macos-amd64 (Intel芯片)
echo - %APP_NAME%-macos-arm64 (Apple Silicon)
echo.
echo ## 使用方法
echo.
echo ### Windows
echo 直接运行对应的exe文件
echo.
echo ### Linux/macOS
echo 1. 添加执行权限: chmod +x %APP_NAME%-linux-amd64
echo 2. 运行: ./%APP_NAME%-linux-amd64 "源目录" "目标目录"
echo.
echo ## 功能特性
echo - ✅ 增量备份
echo - ✅ 时间戳文件列表
echo - ✅ 多平台支持
echo - ✅ 简单易用
) > "%OUTPUT_DIR%\README.txt"

echo ✅ 创建: README.txt

:: 创建校验文件
echo.
echo [生成文件校验]
cd "%OUTPUT_DIR%"
(
echo 文件校验和 (SHA256)
echo ====================
echo.
echo [Windows]
for %%f in (windows\*.exe) do (
    echo %%f
    certutil -hashfile "%%f" SHA256 | findstr /v "certutil Hash"
    echo.
)

echo [Linux]
for %%f in (linux\*) do (
    echo %%f
    certutil -hashfile "%%f" SHA256 | findstr /v "certutil Hash"
    echo.
)

echo [macOS]
for %%f in (macos\*) do (
    echo %%f
    certutil -hashfile "%%f" SHA256 | findstr /v "certutil Hash"
    echo.
)
) > checksums.txt 2>nul

echo ✅ 创建: checksums.txt

cd ..

:: ===================== 打包 =====================
echo.
echo ========== 创建发布包 ==========
echo.

:: 检查7zip是否可用
where 7z >nul 2>nul
if %errorlevel% equ 0 (
    echo [创建ZIP压缩包]
    7z a -tzip "%OUTPUT_DIR%\%APP_NAME%-v%VERSION%-windows.zip" "%OUTPUT_DIR%\windows\*" "%OUTPUT_DIR%\README.txt" "%OUTPUT_DIR%\checksums.txt" >nul
    if exist "%OUTPUT_DIR%\%APP_NAME%-v%VERSION%-windows.zip" (
        echo ✅ Windows包: %APP_NAME%-v%VERSION%-windows.zip
    )
    
    7z a -tzip "%OUTPUT_DIR%\%APP_NAME%-v%VERSION%-linux.zip" "%OUTPUT_DIR%\linux\*" "%OUTPUT_DIR%\README.txt" "%OUTPUT_DIR%\checksums.txt" >nul
    if exist "%OUTPUT_DIR%\%APP_NAME%-v%VERSION%-linux.zip" (
        echo ✅ Linux包: %APP_NAME%-v%VERSION%-linux.zip
    )
    
    7z a -tzip "%OUTPUT_DIR%\%APP_NAME%-v%VERSION%-macos.zip" "%OUTPUT_DIR%\macos\*" "%OUTPUT_DIR%\README.txt" "%OUTPUT_DIR%\checksums.txt" >nul
    if exist "%OUTPUT_DIR%\%APP_NAME%-v%VERSION%-macos.zip" (
        echo ✅ macOS包: %APP_NAME%-v%VERSION%-macos.zip
    )
) else (
    echo ⚠️  跳过ZIP压缩 (需要安装7zip)
    echo 您可以从 build\ 目录手动获取文件
)

:: ===================== 最终输出 =====================
echo.
echo ================================================
echo 🎉 编译完成！
echo ================================================
echo.
echo 输出目录结构:
echo %OUTPUT_DIR%\
echo ├── windows\
echo │   ├── %APP_NAME%-windows-amd64.exe
echo │   └── %APP_NAME%-windows-386.exe
echo ├── linux\
echo │   ├── %APP_NAME%-linux-amd64
echo │   └── %APP_NAME%-linux-arm64
echo ├── macos\
echo │   ├── %APP_NAME%-macos-amd64
echo │   └── %APP_NAME%-macos-arm64
echo ├── README.txt
echo ├── checksums.txt
echo └── *.zip (发布包)
echo.
echo 验证编译结果:
echo 1. Windows测试: build\windows\%APP_NAME%-windows-amd64.exe --help
echo 2. 查看文件列表: dir /s build\
echo.
echo ================================================
pause