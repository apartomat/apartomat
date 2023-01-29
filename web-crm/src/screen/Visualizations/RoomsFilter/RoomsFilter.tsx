import React, { useState, useEffect } from "react"

import { VisualizationsScreenHouseRoomFragment } from "api/types"

import { Box, BoxExtendedProps, CheckBox } from "grommet"

export default function RoomsFilter({
    rooms,
    onSelectRooms,
    ...props
}: {
    onSelectRooms?: (id: string[]) => void,
    rooms: VisualizationsScreenHouseRoomFragment[]
} & BoxExtendedProps) {
    const [ roomsChecked, setRoomsChecked ] = useState<string[]>([])
    const [ metaKey, setMetaKey ] = useState<boolean>(false)

    useEffect(() => {
        onSelectRooms && onSelectRooms(roomsChecked)
    }, [ roomsChecked ])

    const handleKeyDown = (event: KeyboardEvent) => {
        if (event.key === "Meta") {
            setMetaKey(true)
        }
    }

    const handleKeyUp = (event: KeyboardEvent) => {
        if (event.key === "Meta") {
            setMetaKey(false)
        }
    }

    useEffect(() => {
        window.addEventListener("keydown", handleKeyDown)
        window.addEventListener("keyup", handleKeyUp)

        return () => {
            window.removeEventListener("keydown", handleKeyDown)
            window.addEventListener("keyup", handleKeyUp)
        }
    })

    return (
        <Box direction="row" wrap {...props}>
            {rooms.map((room) => {
                return (
                    <CheckBox
                        key={room.id}
                        onChange={({ target: { checked }}: React.ChangeEvent<HTMLInputElement>): void => {
                            if (metaKey) {
                                if (checked) {
                                    setRoomsChecked([...roomsChecked, room.id])
                                } else {
                                    setRoomsChecked(roomsChecked.filter(item => item !== room.id))
                                }
                            } else {
                                if (!checked && roomsChecked.length === 1) {
                                    setRoomsChecked([])
                                } else {
                                    setRoomsChecked([room.id])
                                }
                            }
                        }}
                        checked={roomsChecked.includes(room.id)}
                    >
                        {({ checked }: { checked: boolean }) => (
                            <Box
                                pad={{horizontal: "small", vertical: "xsmall"}}
                                background={checked ? "brand" : "light-1"}
                                round="medium"
                            >{room.name}</Box>
                        )}
                    </CheckBox>
                )
            })}
        </Box>
    )
}