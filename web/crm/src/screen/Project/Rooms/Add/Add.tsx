import React, { useEffect, useState } from "react"

import { Button, Box, Heading, Layer, LayerExtendedProps, Text } from "grommet"
import { FormClose } from "grommet-icons"

import Form, { Value as FormValue } from "../Form/Form"
import useAddRoom, { ProjectScreenHouseRoomFragment as ProjectScreenHouseRoom } from "./useAddRoom"

export function Add({
    houseId,
    onClickClose,
    onAdd,
    ...layerProps
}: {
    houseId: string,
    onClickClose?: () => void,
    onAdd?: (room: ProjectScreenHouseRoom) => void,
} & LayerExtendedProps) {
    const [ value, setValue ] = useState({name: "", square: "", level: "" } as FormValue)

    const [ addRoom, { data, loading, error } ] = useAddRoom()

    const handleSubmit = (event: React.FormEvent) => {
        const { name, square } = value

        addRoom(houseId, { name, square: parseFloat(square.replace(",", ".")) })

        event.preventDefault()
    }

    useEffect(() => {
        switch (data?.addRoom.__typename) {
            case "RoomAdded":
                onAdd && onAdd(data.addRoom.room)
        }
    }, [ data, onAdd ])

    return (
        <Layer {...layerProps}>
            <Box pad="medium" gap="medium" width="large">
                <Box direction="row" justify="between" align="center">
                    <Heading level={2} margin="none">Комната</Heading>
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
                                label="Добавить"
                                disabled={loading}
                            />
                            {error && <Box><Text color="status-critical">{error.message}</Text></Box>}
                            <Box><Text color="status-critical"><ErrorMessage res={data?.addRoom}/></Text></Box>
                        </Box>
                    }
                />
            </Box>
        </Layer>
    )
}

/* eslint-disable  @typescript-eslint/no-explicit-any */
function ErrorMessage({res}: {res: { __typename: "NotFound", message: string } |  { __typename: "Forbidden", message: string } | { __typename: "ServerError", message: string } | any }) {
    switch (res?.__typename) {
        case "NotFound":
            return <>Не найдено</>
        case "Forbidden":
            return <>Доступ запрещен</>
        case "ServerError":
            return <>Ошибка сервера</>
    }

    return null
}

export default  Add
