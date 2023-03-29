import React, { useEffect, useState } from "react"

import { Box, BoxExtendedProps, Button } from "grommet"
import { Add } from "grommet-icons"

import { ProjectHouses } from "../useProject"

import Room from "./Room/Room"
import { Add as AddRoom } from "./Add/Add"
import { Update as UpdateRoom } from "./Update/Update"
import { ProjectScreenHouseRoomFragment as ProjectScreenHouseRoom } from "api/graphql"

export default function Rooms({
    houses,
    onAddRoom,
    onDeleteRoom,
    onUpdateRoom,
    ...boxProps
}: {
    houses: ProjectHouses,
    onAddRoom?: (room: ProjectScreenHouseRoom ) => void,
    onUpdateRoom?: (room: ProjectScreenHouseRoom ) => void,
    onDeleteRoom?: (room: ProjectScreenHouseRoom ) => void,
} & BoxExtendedProps) {
    const [ house, setHouse ] = useState(firstHouse(houses))

    const [ rooms, setRooms ] = useState<ProjectScreenHouseRoom[]>([])

    const [ showAddRoom, setShowAddRoom ] = useState(false)

    const [ updateRoom, setUpdateRoom ] = useState<ProjectScreenHouseRoom | undefined>(undefined)

    useEffect(() => {
        setHouse(firstHouse(houses))
    }, [ houses ])

    useEffect(() => {
        if (house && house.rooms.list.__typename === "HouseRoomsList") {
            setRooms(house.rooms.list.items)
        }
    }, [ house ])


    const handleAddRoom = (room: ProjectScreenHouseRoom) => {
        setShowAddRoom(false)
        onAddRoom && onAddRoom(room)
    }

    const handleDeleteRoom = (room: ProjectScreenHouseRoom) => {
        setRooms(rooms.filter(r => r.id !== room.id))
        onDeleteRoom && onDeleteRoom(room)
    }

    return (
        <Box {...boxProps}>
            <Box direction="row" wrap>
                {rooms.map((room) => {
                    return (
                        <Room
                            key={room.id}
                            room={room}
                            margin={{right: "xsmall", bottom: "small"}}
                            onClickUpdate={(room) => setUpdateRoom(room)}
                            onDelete={handleDeleteRoom}
                        />
                    )
                })}
                <Button key="" icon={<Add/>} label="Добавить" onClick={() => setShowAddRoom(true) } margin={{bottom: "small"}}/>
            </Box>

            {showAddRoom && house &&
                <AddRoom
                    houseId={house?.id}
                    onEsc={() => { setUpdateRoom(undefined) }}
                    onClickClose={() => setShowAddRoom(false) }
                    onAdd={handleAddRoom}
                />
            }

            {updateRoom &&
                <UpdateRoom
                    room={updateRoom}
                    onEsc={() => { setUpdateRoom(undefined) }}
                    onClickClose={() => { setUpdateRoom(undefined) }}
                    onUpdate={(room: ProjectScreenHouseRoom) => {
                        setUpdateRoom(undefined)
                        onUpdateRoom && onUpdateRoom(room)
                    }}
                />
            }
        </Box>
    )
}

function firstHouse(houses: ProjectHouses) {
    switch (houses.list.__typename) {
        case "ProjectHousesList":
            return houses.list.items[0]
        default:
            return undefined
    }
}