import { Box, BoxExtendedProps, CheckBox } from "grommet"
import { VisualizationsScreenHouseRoomFragment } from "../api/useVisualizations"
import { useSearchParams } from "react-router-dom"
import { useCallback } from "react"

export function Rooms({
    rooms,
    ...boxProps
}: {
    rooms?: VisualizationsScreenHouseRoomFragment[]
} & BoxExtendedProps) {
    const [roomsFilter, setRoomsFilter] = useSearchParamsRoomsFilter()

    return (
        <Box direction="row" wrap gap="small" {...boxProps}>
            {rooms?.map(({ id, name }) => {
                return (
                    <CheckBox
                        key={id}
                        onChange={() => {
                            if (roomsFilter.includes(id)) {
                                setRoomsFilter(roomsFilter.filter((it) => it !== id))
                            } else {
                                setRoomsFilter(roomsFilter.concat([id]))
                            }
                        }}
                        checked={roomsFilter.includes(id)}
                    >
                        {({ checked }: { checked: boolean }) => (
                            <Box
                                pad={{ horizontal: "small", vertical: "xsmall" }}
                                background={checked ? "brand" : "light-1"}
                                round="medium"
                            >
                                {name}
                            </Box>
                        )}
                    </CheckBox>
                )
            })}
        </Box>
    )
}

function useSearchParamsRoomsFilter() {
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
