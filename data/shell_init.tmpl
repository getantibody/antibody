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
			#while read bundle; do
			#	source "$bundle"
			#done < <(--antibody-wrapper "$@")

			which -s "parallel" && p=parallel || p=xargs
			--antibody-wrapper "$@" | $p --verbose --verbose ls -l

			//source < <(--antibody-wrapper "$@"| xargs --verbose cat)
			;;

		# Perfom a simple health check on the current env this hook
		# is running on. Check if we should replace ourselves with the
		# latest version from the main binary
		check_hook|update_hook)
			local antibody_version="$(--antibody-wrapper version)"
            echo "Antibody version: $antibody_version"
			echo "Shell hook version: $version"

			if [[ "$antibody_version" != "$version" ]]; then
				echo "--"
				echo "Versions are not out of sync."
				echo "Replacing myself with the latest copy via:"
                echo "  antibody shell_init"
				-antibody-shell-hook shell_init
			fi
			;;

		# These actually return what should be executed in the shell
		apply|shell_init)
			source < <(--antibody-wrapper "$@")
			;;

		# Anything else just send it on up as is.
		*)  command antibody "$@" ;;
	esac
}

alias antibody=-antibody-shell-hook

