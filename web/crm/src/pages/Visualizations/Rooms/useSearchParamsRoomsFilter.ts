import { useSearchParams } from "react-router-dom"
import { useCallback } from "react"

export function useSearchParamsRoomsFilter() {
    const [searchParams, setSearchParams] = useSearchParams()

    const filter = searchParams.getAll("room")

    const setFilter = useCallback((rooms: string[]) => {
        setSearchParams((params: URLSearchParams) => {
            params.delete("room")
            rooms.forEach((room) => params.append("room", room))
            return params
        })
    }, [])

    return [filter, setFilter]
}
