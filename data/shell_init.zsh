#!/bin/zsh

function --antibody-wrapper() {
	__ANTIBODY_WRAPPED=true command antibody "$@"
}

function -antibody-shell-hook() {
	local version="{{.Version}}"
	case "$1" in
		# These should be updated to just return the source ala below
		# provided that __ANTIBODY_WRAPPED==true .
		# Currently we just emulate that.
		# These return a newline separated list of files to source.
		bundle|update)
            while read bundle; do
                echo "bundle: $bundle"
                source "$bundle"
            done < <(--antibody-wrapper "$@")
			;;

		# Perfom a simple health check on the current env this hook
		# is running on. Check if we should replace ourselves with the
		# latest version from the main binary
        check_hook|update_hook)
			local antibody_version="$(--antibody-wrapper -v | sed -e 's/^.* version //')"
            echo "Antibody version: $antibody_version"
			echo "Shell hook version: $version"

			if [[ "$antibody_version" != "$version" \
                  || "$1" == "update_hook" ]]; then
				echo "--"
				echo "Versions are not out of sync."
				echo "Replacing myself with the latest copy via:"
                echo "  antibody shell"
				-antibody-shell-hook shell
                return
			fi
			;;

		# These actually return what should be executed in the shell
		apply|shell)
			source <(--antibody-wrapper "$@")
			;;

		# Anything else just send it on up as is.
		*)  command antibody "$@" ;;
	esac
}

alias antibody=-antibody-shell-hook

