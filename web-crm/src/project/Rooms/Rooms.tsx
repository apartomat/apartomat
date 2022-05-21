import React from "react"

import { Box, Heading, Text } from "grommet"

import { ProjectHouses, HouseRooms } from "../useProject"

export default function Rooms({ houses }: { houses: ProjectHouses }) {
    const house = firstHouse(houses)

    return (
        <Box margin={{top: "small"}}>
            <Heading level={4} margin={{ bottom: "xsmall"}}>Комплектация</Heading>
            <Box height="36px" justify="center">
                {
                    (house ? <RoomsText rooms={house.rooms}/> : <Text>n/a</Text>)
                }
            </Box>
        </Box>
    )
}

function RoomsText({ rooms }: { rooms: HouseRooms }) {
    switch (rooms.list.__typename) {
        case "HouseRoomsList":
            if (rooms.list.items.length === 0) {
                return (
                    <Text>n/a</Text>
                )
            }

            return (
                <Text>{rooms.list.items.length} помещений, {rooms.list.items.reduce((acc, room) => {
                    return acc + (room.square || 0)
                }, 0)} м<sup>2</sup></Text>
            )
        default:
            return (
                <Text>n/a</Text>
            )
    }
}

function firstHouse(houses: ProjectHouses) {
    switch (houses.list.__typename) {
        case "ProjectHousesList":
            return houses.list.items[0]
        default:
            return undefined
    }
}