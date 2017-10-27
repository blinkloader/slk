#!/bin/sh

install () {

	set -eu

	UNAME=$(uname)
	if [ "$UNAME" != "Linux" -a "$UNAME" != "Darwin" -a "$UNAME" != "OpenBSD" ] ; then
		echo "Sorry, OS not supported: ${UNAME}."
		exit 1
	fi

	if [ "$UNAME" = "Darwin" ] ; then
		OSX_ARCH=$(uname -m)
		if [ "${OSX_ARCH}" = "x86_64" ] ; then
			PLATFORM="darwin_amd64"
		else
			echo "Sorry, architecture not supported: ${OSX_ARCH}."
			exit 1
		fi
	elif [ "$UNAME" = "Linux" ] ; then
		LINUX_ARCH=$(uname -m)
		if [ "${LINUX_ARCH}" = "i686" ] ; then
			PLATFORM="linux_386"
		elif [ "${LINUX_ARCH}" = "x86_64" ] ; then
			PLATFORM="linux_amd64"
		else
			echo "Sorry, architecture not supported: ${LINUX_ARCH}."
			exit 1
		fi
	fi

	LATEST=$(curl -s https://api.github.com/repos/yarikbratashchuk/slk/tags | grep name  | head -n 1 | sed 's/[," ]//g' | cut -d ':' -f 2)
	URL="https://github.com/yarikbratashchuk/slk/releases/download/$LATEST/slk_"$LATEST"_$PLATFORM.tar.gz"
	DEST=/tmp/slk.tar.gz

	if [ -z $LATEST ] ; then
		echo "Error requesting."
		exit 1
	else
		echo "Downloading slk from https://github.com/yarikbratashchuk/slk/releases/download/$LATEST/slk_"$LATEST"_$PLATFORM.tar.gz to $DEST"
		if curl -sL https://github.com/yarikbratashchuk/slk/releases/download/$LATEST/slk_"$LATEST"_$PLATFORM.tar.gz -o $DEST; then
			cd /tmp/
			tar xzf ./slk.tar.gz
			mv ./slk_"$LATEST"_$PLATFORM/slk /usr/local/bin/slk
			mv ./slk_"$LATEST"_$PLATFORM/slkd /usr/local/bin/slkd
			chmod +x /usr/local/bin/slk /usr/local/bin/slkd
			echo "slk installation was successful"
		else
			echo "Installation failed. You may need elevated permissions."
		fi
	fi
}

install
