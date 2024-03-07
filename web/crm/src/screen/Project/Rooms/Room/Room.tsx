import React, { useRef, useState, useEffect } from "react"

import { Box, BoxExtendedProps, Button, Card, CardBody, CardFooter, CardHeader, Drop } from "grommet"
import { Trash } from "grommet-icons"
import { ProjectScreenHouseRoomFragment as HouseRoom } from "api/graphql"
import useDeleteRoom from "./useDeleteRoom"

export default function Room({
    room,
    onClickUpdate,
    onDelete,
    ...boxProps
}: {
    room: HouseRoom,
    onClickUpdate?: (room: HouseRoom) => void,
    onDelete?: (room: HouseRoom) => void
} & BoxExtendedProps) {
    const ref = useRef(null)

    const [showCard, setShowCard] = useState(false)

    const [showDeleteConfirm, setShowDeleteConfirm] = useState(false)

    const [deleteRoom, { data } ] = useDeleteRoom()

    const handleDelete = () => {
        setShowDeleteConfirm(true)
    }

    const handleDeleteConfirm = () => {
        deleteRoom(room.id)
    }

    const handleDeleteCancel = () => {
        setShowDeleteConfirm(false)
        setShowCard(false)
    }

    useEffect(() => {
        switch (data?.deleteRoom.__typename) {
            case "RoomDeleted":
                setShowCard(false)
                onDelete && onDelete(room)
        }
    }, [ data, room, onDelete ])

    return (
        <Box {...boxProps}>
            <Button
                key={room.id}
                ref={ref}
                primary
                color="light-2"
                label={room.name}
                onClick={() => setShowCard(!showCard) }
                style={{ whiteSpace: "nowrap" }}
            />
            {ref.current && showCard &&
                <Drop
                    target={ref.current}
                    align={{left: "right"}}
                    plain
                    onEsc={() => setShowCard(false) }
                    onClickOutside={() => setShowCard(false) }
                >
                    <Card width="medium" background="white" margin="small">
                        <CardHeader pad={{horizontal: "medium", top: "medium"}} style={{fontWeight: "bold"}}>{room.name}</CardHeader>
                        <CardBody pad="medium">
                            {room.square && <>{room.square}&nbsp;м²</>}
                        </CardBody>
                        <CardFooter pad={{horizontal: "small"}} background="light-1" height="xxsmall">
                            {showDeleteConfirm
                                ?
                                    <Box direction="row" gap="small">
                                        <Button primary label="Удалить" size="small" onClick={handleDeleteConfirm}/>
                                        <Button label="Отмена" size="small" onClick={handleDeleteCancel}/>
                                    </Box>
                                : (
                                    <Button icon={<Trash/>} onClick={handleDelete}/>
                                )
                            }

                            {!showDeleteConfirm &&
                                <Button label="Редактировать" size="small" primary onClick={() => {
                                    setShowCard(false)
                                    onClickUpdate && onClickUpdate(room)
                                }}/>
                            }

                        </CardFooter>
                    </Card>
                </Drop>
            }
        </Box>
    )
}