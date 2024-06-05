#!/bin/bash
appversion=0.0.0

function getversion() {
  appversion=$(cat version.txt)
  if [ "$appversion" = "" ]; then
    appversion="0.0.0"
    echo $appversion
  else
    v3=$(echo $appversion | awk -F'.' '{print($3);}')
    v2=$(echo $appversion | awk -F'.' '{print($2);}')
    v1=$(echo $appversion | awk -F'.' '{print($1);}')
    if [[ $(expr $v3 \>= 9) == 1 ]]; then
      v3=0
      if [[ $(expr $v2 \>= 9) == 1 ]]; then
        v2=0
        v1=$(expr $v1 + 1)
      else
        v2=$(expr $v2 + 1)
      fi
    else
      v3=$(expr $v3 + 1)
    fi
    ver="$v1.$v2.$v3"
    echo $ver
  fi
}

function todir() {
  pwd
}

function pull() {
  todir
  echo "git pull"
  git pull
}

function forcepull() {
  todir
  echo "git fetch --all && git reset --hard origin/master && git pull"
  git fetch --all && git reset --hard origin/master && git pull
}
#  shellcheck disable=SC2120
function gitpush() {
  commit=""
  if [ ! -n "$1" ]; then
    commit="$(date '+%Y-%m-%d %H:%M:%S') by ${USER}"
  else
    commit="$1 by ${USER}"
  fi

  echo $commit
  git add .
  git commit -m "$commit"
  #  git push -u origin main
  git push
}

function gittag() {
  appversion=$(getversion)
  echo "当前版本：$appversion"
  git add .
  git commit -m "release v$appversion"
  git tag -a v${appversion} -m "release v$appversion"
  git push origin v${appversion}
  echo $appversion >version.txt
}
function m() {
    echo "1. 强制更新"
    echo "2. 普通更新"
    echo "3. 提交项目"
    echo "4. 提交tag"
    echo "请输入编号:"
    read index

    case "$index" in
    [1]) (forcepull);;
    [2]) (pull);;
    [3]) (gitpush);;
    [4]) (gittag);;
    *) echo "exit" ;;
  esac
}

function bootstrap() {
    case $1 in
    pull) (pull) ;;
    m) (m) ;;
      -f) (forcepull) ;;
       *) ( gitpush $1)  ;;
    esac
}


bootstrap m
