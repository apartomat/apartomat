import React, { useEffect, useState } from "react"

import { Box, Button, Heading, Layer, LayerExtendedProps, Text } from "grommet"
import { FormClose } from "grommet-icons"

import Form, { Value as FormValue } from "../Form/Form"

import { useUpdateRoom, ProjectScreenHouseRoomFragment as ProjectScreenHouseRoom } from "./useUpdateRoom"

export function Update({
    room,
    onUpdate,
    onClickClose,
    ...layerProps
}: {
    room: ProjectScreenHouseRoom,
    onUpdate?: (room: ProjectScreenHouseRoom) => void,
    onClickClose?: () => void,
} & LayerExtendedProps) {
    const [ value, setValue ] = useState({
        name: room.name,
        square: room.square?.toString().replace(".", ","),
        level: room.level?.toString(),
    } as FormValue)

    const [ updateRoom, { data, loading, error } ] = useUpdateRoom()

    const handleSubmit = (event: React.FormEvent) => {
        const { name, square } = value

        updateRoom(room.id, { name, square: square ? parseFloat(square.replace(",", ".")) : undefined })

        event.preventDefault()
    }

    useEffect(() => {
        switch (data?.updateRoom.__typename) {
            case "RoomUpdated":
                onUpdate && onUpdate(data.updateRoom.room)
        }
    }, [ data, onUpdate ])

    return (
        <Layer {...layerProps}>
                <Box pad="medium" gap="medium" width={{min: "500px"}}>
                    <Box direction="row"justify="between"align="center">
                        <Heading level={3} margin="none">Комната</Heading>
                        <Button icon={ <FormClose/> } onClick={onClickClose}/>
                    </Box>
                    <Form
                        value={value}
                        onChange={setValue}
                        onSubmit={handleSubmit}
                        submit={
                            <Box direction="row" justify="between" margin={{top: "large"}}>
                                <Button
                                    type="submit"
                                    primary
                                    label={loading ? 'Сохранение...' : 'Сохранить' }
                                    disabled={loading}
                                />
                                {error && <Box><Text>{error.message}</Text></Box>}
                                <Box><Text color="status-critical"><ErrorMessage res={data?.updateRoom}/></Text></Box>
                            </Box>
                        }
                    />
                </Box>
        </Layer>
    )
}

function ErrorMessage({res}: { res: { __typename: "Forbidden", message: string } | { __typename: "ServerError", message: string } | any}) {
    switch (res?.__typename) {
        case "Forbidden":
            return <>Доступ запрещен</>
        case "ServerError":
            return <>Ошибка сервера</>
    }

    return null
}

export default Update