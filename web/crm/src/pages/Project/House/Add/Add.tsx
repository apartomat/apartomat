import React, { useEffect, useState } from "react"

import useAddHouse, { House as HouseType } from "./useAddHouse"

import { Box, Button, Heading, Layer } from "grommet"
import { FormClose } from "grommet-icons"

import { Form, Value as FormValue } from "../Form/Form"

export function Add({
    projectId,
    onAdd,
    onEsc,
    onClickClose,
}: {
    projectId: string
    onAdd?: (house: HouseType) => void
    onEsc?: () => void
    onClickClose?: () => void
}) {
    const [errorMessage, setErrorMessage] = useState<string | undefined>()

    const [value, setValue] = useState({ city: "", address: "", housingComplex: "" } as FormValue)

    const [addHouse, { data, error, loading }] = useAddHouse()

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault()
        await addHouse(projectId, { ...value })
    }

    useEffect(() => {
        if (data?.addHouse) {
            switch (data.addHouse.__typename) {
                case "HouseAdded": {
                    const {
                        addHouse: { house },
                    } = data
                    onAdd && onAdd(house)
                    break
                }

                case "NotFound":
                    setErrorMessage("Неизвестный проект")
                    break
                case "Forbidden":
                    setErrorMessage("Доступ запрещен")
                    break
                default:
                    setErrorMessage("Ошибка сервера")
            }
        }

        if (error) {
            setErrorMessage("Ошибка сервера")
        }
    }, [data, error, onAdd])

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

                <Box>
                    <Form
                        value={value}
                        onChange={setValue}
                        onSubmit={handleSubmit}
                        submit={
                            <Box>
                                <Box direction="row" justify="between" margin={{ top: "large" }}>
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
            </Box>
        </Layer>
    )
}
