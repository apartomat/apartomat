import { Box, BoxExtendedProps, CheckBox } from "grommet"
import { VisualizationsScreenHouseRoomFragment } from "../api/useVisualizations"
import { useSearchParamsRoomsFilter } from "./useSearchParamsRoomsFilter"

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
