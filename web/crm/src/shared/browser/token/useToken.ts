import useLocalStorage from "@d2k/react-localstorage"

export function useToken() {
    return useLocalStorage("token", "")
}
