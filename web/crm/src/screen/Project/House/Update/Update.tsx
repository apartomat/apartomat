import React, { useEffect, useState } from "react"

import { ProjectScreenHouse } from "../../useProject"
import useUpdateHouse, { House as HouseType } from "./useUpdateHouse"

import { Box, Button, Heading, Layer } from "grommet"
import { FormClose } from "grommet-icons"

import { Form } from "../Form/Form"

export function Update({
    house,
    onUpdate,
    onEsc,
    onClickClose,
}: {
    house: ProjectScreenHouse
    onUpdate?: (house: HouseType) => void
    onEsc?: () => void
    onClickClose?: () => void
}) {
    const [errorMessage, setErrorMessage] = useState<string | undefined>()

    const [value, setValue] = useState({
        city: house.city,
        address: house.address,
        housingComplex: house.housingComplex,
    })

    const [updateHouse, { data, error, loading }] = useUpdateHouse()

    const handleSubmit = (event: React.FormEvent) => {
        updateHouse(house.id, { ...value })
        event.preventDefault()
    }

    useEffect(() => {
        if (data?.updateHouse) {
            switch (data.updateHouse.__typename) {
                case "HouseUpdated":
                    onUpdate && onUpdate(data?.updateHouse.house)
                    break
                case "Forbidden":
                    setErrorMessage("Доступ запрещен")
                    break
                case "NotFound":
                    setErrorMessage("Неизвестный адрес")
                    break
                default:
                    setErrorMessage("Ошибка сервера")
            }
        }
    }, [data, onUpdate])

    useEffect(() => {
        setErrorMessage(error ? "Ошибка сервера" : undefined)
    }, [error])

    return (
        <Layer onEsc={onEsc}>
            <Box pad="medium" gap="medium" width="medium">
                <Box direction="row" justify="between" align="center">
                    <Heading level={2} margin="none">
                        Адрес
                    </Heading>
                    <Button icon={<FormClose />} onClick={onClickClose} />
                </Box>

                {errorMessage && (
                    <Box
                        background={{ color: "status-critical", opacity: "strong" }}
                        round="medium"
                        pad={{ vertical: "small", horizontal: "medium" }}
                        margin={{ top: "small" }}
                    >
                        {errorMessage}
                    </Box>
                )}

                <Form
                    value={value}
                    onChange={setValue}
                    onSubmit={handleSubmit}
                    submit={
                        <Box>
                            <Box direction="row" justify="between" margin={{ top: "small" }}>
                                <Button
                                    type="submit"
                                    primary
                                    label={loading ? "Сохранение..." : "Сохранить"}
                                    disabled={loading}
                                />
                            </Box>
                        </Box>
                    }
                />
            </Box>
        </Layer>
    )
}
