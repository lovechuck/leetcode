#!/bin/bash
# 193. 有效电话号码
isvalid() {
  regx="^([0-9]{3}-|\([0-9]{3}\) )[0-9]{3}-[0-9]{4}$"
  if [[ $1 =~ $regx ]]
  then
      echo "$1"
  fi
}

while read -r line
do
  isvalid "$line"
done < file.txt

