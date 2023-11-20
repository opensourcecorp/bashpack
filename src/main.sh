#!/usr/bin/env bash
set -euo pipefail

# Root directory for the bashpack tree. Can be user-overridden.
BASHPACK_ROOT="${BASHPACK_ROOT:-}"
if [[ "$(id -u)" -eq 0 ]] ; then
  BASHPACK_ROOT="/usr/local/share/bashpack"
else
  BASHPACK_ROOT="${HOME}/.bashpack"
fi
BASHPACK_LIB="${BASHPACK_ROOT}/lib"
export BASHPACK_ROOT
export BASHPACK_LIB
mkdir -p \
  "${BASHPACK_ROOT}" \
  "${BASHPACK_LIB}"

# Set up ezlog, which is ironic considering bashpack can be used to get ezlog itself
[[ -d /tmp/ezlog ]] || {
  printf 'Getting ezlog...\n'
  git clone --depth=1 https://github.com/opensourcecorp/ezlog.git /tmp/ezlog > /dev/null 2>&1
}
# shellcheck disable=SC1091
source /tmp/ezlog/src/main.sh

# _sanitize-pkg-name takes a package name (a URI or something like it) and
# converts it into a canonical filesystem path
_sanitize-pkg-name() {
  if [[ -z "${1:-}" ]] ; then
    log-fatal "Package URI not provided"
  fi
  local pkg="${1:-}"
  pkg="${pkg/http:\/\//}" # http://
  pkg="${pkg/https:\/\//}" # https://
  pkg="${pkg/git:\/\//}" # git://
  pkg="${pkg/git@/}" # git@<path>
  pkg="${pkg/:/\/}" # hopefully the trailing colon in e.g. 'git@github.com:'
  pkg="${pkg/.git/}"

  log-debug "Sanitized '${1}' to '${pkg}'" >&2
  printf '%s' "${pkg}" # return
  return 0
}

# _cache fetches and caches packages
_cache() {
  if [[ -z "${1:-}" ]] ; then
    log-fatal "Package URI not provided"
  fi
  local pkg
  pkg="$(_sanitize-pkg-name "${1:-}")"
  log-debug "Processing cache check for package '${pkg}'" >&2
  local pkg_path="${BASHPACK_LIB}/${pkg}"
  if [[ ! -d "${pkg_path}" ]] ; then
    log-info "Package '${pkg}' does not exist locally; retrieving" >&2
    git -C "${BASHPACK_LIB}" clone "${1}" "${pkg}" > /dev/null 2>&1 # clones to path based on the source path
    log-info "bashpack package '${pkg}' sucessfully cached at ${pkg_path}" >&2
  else
    log-debug "bashpack package '${pkg}' was already cached at ${pkg_path}" >&2
  fi
  return 0
}

# _mainpath retrieves the host path to the backpack package's main file
_mainpath() {
  if [[ -z "${1:-}" ]] ; then
    log-fatal "Package URI not provided"
  fi
  local pkg
  pkg="$(_sanitize-pkg-name "${1}")"
  _cache "${1}"
  local pkg_path="${BASHPACK_LIB}/${pkg}"

  local mainfile=''
  if [[ -f "${pkg_path}/manifest.bashpack" ]] ; then
    log-debug "Found manifest.bashpack at '${pkg_path}/manifest.bashpack', will try to parse" >&2
    local mainfile
    mainfile="${pkg_path}/$(awk -F'=' '/^main/ { gsub(/ /, "") ; print $2 }' "${pkg_path}"/manifest.bashpack)"
  else
    log-warn "No manifest.bashpack found at '${pkg_path}', will use the first 'main.sh' found" >&2
    local mainfile
    mainfile="$(find "${pkg_path}" -type f -name main.sh | sort | head -n1)"
    log-warn "Found first main.sh to be '${mainfile}'" >&2
  fi
  log-debug "Using '${mainfile}' as mainfile" >&2

  printf '%s' "${mainfile}"
  return 0
}

_main() {
  case "${1:-}" in
    run)
      bash "$(_mainpath "${2:-}")"
    ;;
    mainpath)
      _mainpath "${2:-}"
    ;;
    cache)
      _cache "${2:-}"
    ;;
    "")
      log-fatal "Must supply a valid bashpack command"
    ;;
    *)
      log-fatal "Unrecognized command '${1:-}'"
    ;;
  esac
}

if [[ "${BASH_SOURCE[0]}" == "${0}" ]] ; then
  _main "$@"
fi
