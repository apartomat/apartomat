import useLocalStorage from "@d2k/react-localstorage"

function useToken() {
    return useLocalStorage("token", "")
}

export default useToken
