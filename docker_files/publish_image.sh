#!/bin/bash

set -euCo pipefail

function build_path() {
    local -r path="docker_files/$1"

    echo "$path"
}

function check_if_docker_file_exists(){
    local -r docker_file_path="${1}/Dockerfile"

    if [[ ! -f "$docker_file_path" ]] ; then
        echo "Not found: ${docker_file_path}"
        exit 1
    fi
}

function to_image_name() {
     local -r original_name="$1"
     local -r image_name="$(echo "$original_name" | sed -e "s/_/-/g")"

     echo "gcr.io/memoq-backend/$image_name"
}

function fetch_tag() {
    local -r docker_file_path="${1}/Dockerfile"

    echo $(head -n 1 $docker_file_path | sed 's/# version: //g')
}

function build_image(){
    local -r path="$1"
    local -r image_name="$2"
    local -r tag="$3"

    docker build "$path" -t "${image_name}:${tag}"
}

function push_image(){
    local -r image_name="$1"
    local -r tag="$2"

    docker push "${image_name}:${tag}"
}

function main(){
    local -r target="$1"

    local -r path=$(build_path "$target")

    check_if_docker_file_exists "$path"

    local -r image_name=$(to_image_name "$target")

    local -r tag=$(fetch_tag "$path")

    build_image "$path" "$image_name" "$tag"

    push_image "$image_name" "$tag"
}

if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
    echo "Build image and push to gcr."
    echo "Usage: ./docker_files/publish_image.sh go_gclound|datastore_emulator."

    main "$1"
fi
