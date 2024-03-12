import React, { PropsWithChildren, useEffect, useState } from "react"
import {
    DndContext,
    KeyboardSensor,
    PointerSensor,
    useSensor,
    useSensors,
    pointerWithin,
    DragEndEvent,
} from "@dnd-kit/core"
import {
    arrayMove,
    SortableContext,
    sortableKeyboardCoordinates,
    rectSortingStrategy,
    useSortable,
} from "@dnd-kit/sortable"

import { Box, BoxExtendedProps, Button } from "grommet"
import { Add } from "grommet-icons"

import { ProjectHouses } from "../useProject"

import Room from "./Room/Room"
import { Add as AddRoom } from "./Add/Add"
import { Update as UpdateRoom } from "./Update/Update"
import { ProjectScreenHouseRoomFragment as ProjectScreenHouseRoom } from "api/graphql"

import useMoveRoomToPosition from "pages/Project/Rooms/useMoveRoomToPosition"
import { CSS } from "@dnd-kit/utilities"

export default function Rooms({
    houses,
    onAddRoom,
    onDeleteRoom,
    onUpdateRoom,
    ...boxProps
}: {
    houses: ProjectHouses
    onAddRoom?: (room: ProjectScreenHouseRoom) => void
    onUpdateRoom?: (room: ProjectScreenHouseRoom) => void
    onDeleteRoom?: (room: ProjectScreenHouseRoom) => void
} & BoxExtendedProps) {
    useEffect(() => {
        setHouse(firstHouse(houses))
    }, [houses])

    const [house, setHouse] = useState(firstHouse(houses))

    useEffect(() => {
        if (house && house.rooms.list.__typename === "HouseRoomsList") {
            setRooms(house.rooms.list.items)
        }
    }, [house])

    const [rooms, setRooms] = useState<ProjectScreenHouseRoom[]>([])

    const [showAddRoom, setShowAddRoom] = useState(false)

    const [updateRoom, setUpdateRoom] = useState<ProjectScreenHouseRoom | undefined>(undefined)

    const [moveRoomToPosition] = useMoveRoomToPosition()

    const handleAddRoom = (room: ProjectScreenHouseRoom) => {
        setShowAddRoom(false)
        onAddRoom && onAddRoom(room)
    }

    const handleDeleteRoom = (room: ProjectScreenHouseRoom) => {
        setRooms(rooms.filter((r) => r.id !== room.id))
        onDeleteRoom && onDeleteRoom(room)
    }

    const sensors = useSensors(
        useSensor(PointerSensor, {
            activationConstraint: {
                distance: 8,
            },
        }),
        useSensor(KeyboardSensor, {
            coordinateGetter: sortableKeyboardCoordinates,
        })
    )

    function handleDragEnd(event: DragEndEvent) {
        const { active, over } = event

        if (!over) {
            return
        }

        if (active.id !== over.id) {
            setRooms(() => {
                const oldIndex = rooms.findIndex(({ id }) => id === active.id)
                const newIndex = rooms.findIndex(({ id }) => id === over.id)

                const room = rooms.find(({ id }) => id === active.id)

                if (room) {
                    moveRoomToPosition(room.id, newIndex + 1)
                }

                return arrayMove(rooms, oldIndex, newIndex)
            })
        }
    }

    return (
        <Box {...boxProps}>
            <Box></Box>
            <Box direction="row" wrap>
                <DndContext sensors={sensors} collisionDetection={pointerWithin} onDragEnd={handleDragEnd}>
                    <SortableContext items={rooms} strategy={rectSortingStrategy}>
                        {rooms.map((room) => {
                            return (
                                <SortableItem key={room.id} id={room.id}>
                                    <Room
                                        key={room.id}
                                        room={room}
                                        margin={{
                                            right: "xsmall",
                                            bottom: "small",
                                        }}
                                        onClickUpdate={(room) => setUpdateRoom(room)}
                                        onDelete={handleDeleteRoom}
                                    />
                                </SortableItem>
                            )
                        })}
                    </SortableContext>
                </DndContext>
                <Button
                    icon={<Add />}
                    label="Добавить"
                    onClick={() => setShowAddRoom(true)}
                    margin={{ bottom: "small" }}
                />
            </Box>

            {showAddRoom && house && house.__typename === "House" && (
                <AddRoom
                    houseId={house.id}
                    onEsc={() => {
                        setUpdateRoom(undefined)
                    }}
                    onClickClose={() => setShowAddRoom(false)}
                    onAdd={handleAddRoom}
                />
            )}

            {updateRoom && (
                <UpdateRoom
                    room={updateRoom}
                    onEsc={() => {
                        setUpdateRoom(undefined)
                    }}
                    onClickClose={() => {
                        setUpdateRoom(undefined)
                    }}
                    onUpdate={(room: ProjectScreenHouseRoom) => {
                        setUpdateRoom(undefined)
                        onUpdateRoom && onUpdateRoom(room)
                    }}
                />
            )}
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

function SortableItem({ id, children }: PropsWithChildren<{ id: string }>) {
    const { attributes, listeners, setNodeRef, transform, transition } = useSortable({ id })

    const style = {
        transform: CSS.Transform.toString(transform ? { x: transform.x, y: transform.y, scaleX: 1, scaleY: 1 } : null),
        transition,
    }

    return (
        <div ref={setNodeRef} style={style} {...attributes} {...listeners}>
            {children}
        </div>
    )
}
