#!/bin/bash
appName=gtbx-server

function menu() {
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gtbx-server_amd64
  echo "1. 编译 Android"
  echo "2. 编译 Raspberry"
  echo "3. 编译 Ubuntu"
  echo "4. 编译 MacOS"
  echo "5. 编译 Linux"
  echo "6. 编译 Windows"
  echo "请输入编号:"
  read index

  case "$index" in
  [1]) (cmake_jni) ;;
  [2]) (cmake_raspberry) ;;
  [3]) (cmake_ubuntu) ;;
  [4]) (cmake_macos) ;;
  [5]) (cmake_linux) ;;
  *) echo "exit" ;;
  esac

  CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o ${BUILD_NAME}

  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BUILD_NAME}_amd64
}

#function buildAll1() {
#  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ../*.go
#  CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ../*.go
#  CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build *.go
#  CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build *.go
#  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build *.go
#}

function buildApp() {
  rm -rf ${appName}*
#  linux=('386' 'amd64' 'arm' 'arm64' 'mips' 'mips64' 'mips64le' 'mipsle' 'riscv64')
  linux=('386' 'amd64' 'arm' 'arm64')
  for model in ${linux[@]}; do
    echo "正在编译 linux ${model}"
    CGO_ENABLED=0 GOOS=linux GOARCH=${model} go build -o ${appName}_linux_${model} ../*.go
#    CGO_ENABLED=0 GOOS=linux GOARCH=${model} go build -o ${appName}_linux_${model} ../*.go
    tar -zcf ${appName}_linux_${model}.tar.gz ${appName}_linux_${model}
    rm -rf ${appName}_linux_${model}
  done

  windows=('386' 'amd64' 'arm' 'arm64')
  for model in ${windows[@]}; do
    echo "正在编译 windows ${model}"
#    CGO_ENABLED=0 GOOS=windows GOARCH=${model} go build -o ${appName}_windows_${model}.exe ../*.go
    CGO_ENABLED=0 GOOS=windows GOARCH=${model} go build -o ${appName}_windows_${model}.exe ../*.go
    zip -rq ${appName}_windows_${model}.zip ${appName}_windows_${model}.exe
    rm -rf ${appName}_windows_${model}.exe
  done

  darwin=('arm64' 'amd64')
  for model in ${darwin[@]}; do
    echo "正在编译 darwin ${model}"
#    CGO_ENABLED=0 GOOS=darwin GOARCH=${model} go build -o ${appName}_darwin_${model} ../*.go
    CGO_ENABLED=0 GOOS=darwin GOARCH=${model} go build -o ${appName}_darwin_${model} ../*.go
    tar -zcf ${appName}_darwin_${model}.tar.gz ${appName}_darwin_${model}
    rm -rf ${appName}_darwin_${model}
  done

  freebsd=('386' 'amd64')
  for model in ${freebsd[@]}; do
    echo "正在编译 darwin ${model}"
#    CGO_ENABLED=0 GOOS=freebsd GOARCH=${model} go build -o ${appName}_freebsd_${model} ../*.go
    CGO_ENABLED=0 GOOS=freebsd GOARCH=${model} go build -o ${appName}_freebsd_${model} ../*.go
    tar -zcf ${appName}_freebsd_${model}.tar.gz ${appName}_freebsd_${model}
    rm -rf ${appName}_freebsd_${model}
  done
}

#build
function test001() {
  os_all='linux windows darwin freebsd'
  #  os_all=('386' 'amd64' 'arm' 'arm64' 'mips' 'mips64' 'mips64le' 'mipsle' 'riscv64')
  arch_all='386 amd64 arm arm64 mips64 mips64le mips mipsle riscv64'
  for os in $os_all; do
    frp_dir_name="${os}"
    echo $frp_dir_name
  done
}

function test002() {
  ostypes=$(go tool dist list)
  for i in $(go tool dist list); do
    echo "--->$i"
  done
}

function buildAll() {
  rm -rf ${appName}*
  ostypes=$(go tool dist list)
  echo $ostypes
  for i in $(go tool dist list); do
    array=($(echo $i | tr '/' ' '))
    echo "编译 ${appName}_${array[1]}_${array[2]}"
    if [ "x${array[1]}" = x"windows" ]; then
      CGO_ENABLED=0 GOOS=${array[1]} GOARCH=${array[2]} go build -o ${appName}_${array[1]}_${array[2]}.exe ../*.go
      zip -rq ${appName}_${array[1]}_${array[2]}.zip ${appName}_${array[1]}_${array[2]}.exe
      rm -rf ${appName}_${array[1]}_${array[2]}.exe
    else
      CGO_ENABLED=0 GOOS=${array[1]} GOARCH=${array[2]} go build -o ${appName}_${array[1]}_${array[2]} ../*.go
      tar -zcf ${appName}_${array[1]}_${array[2]}.tar.gz ${appName}_${array[1]}_${array[2]}
      rm -rf ${appName}_${array[1]}_${array[2]}
    fi
  done

}

function main() {
    dirName='cmd'
    rm -rf out
    mkdir out
    cd out
    buildApp
    rm -rf $dirName
    mkdir $dirName
    mv *zip *.gz $dirName
    zip -rq ${dirName}.zip ${dirName}
}

main