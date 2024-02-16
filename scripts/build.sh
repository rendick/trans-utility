#!/bin/bash

cd ..

PS3='Please select an option: '
types=("Configuration (!)" "Building")

BUILD_DIR=$(pwd)

FILE="./settings/env.go"

BOLD="\e[1m"
CLEAR="\e[0m"

configuration () {
    if [[ $1 == "Configuration (!)" ]] ; then
        if [ ! -f "$FILE" ]; then
            echo "package settings" > "$FILE"
            echo "" >> "$FILE"
            echo "const (" >> "$FILE"
            echo "    API_KEY = \"YOUR_API_KEY\"" >> "$FILE"
            echo ")" >> "$FILE"
            echo "File $FILE created with content."
        else
            echo -e "File $FILE already exists. Skipping creation.\n"
        fi

        echo "${FILE}:"

        cat $FILE

        echo -e "\nPaste your API key into the API_KEY field in the ${FILE}."
    elif [[ $1 == "Building" ]] ; then
        go build -o transutil
		mv transutil ./build/
		echo -e "\nThe binary file was build into this folder: ${BOLD}${BUILD_DIR}/build/${CLEAR}"
		sudo ./build/transutil
		sudo chmod 777 /var/log/transutil/
		sudo chmod 777 /var/log/transutil/transutil.log
    else
        echo 'Invalid installation type!'
    fi
}

echo

select type in "${types[@]}" ; do
    if [[ -n $type ]] ; then
        break
    else
        echo 'Please select a valid option!'
    fi
done

configuration "$type"
