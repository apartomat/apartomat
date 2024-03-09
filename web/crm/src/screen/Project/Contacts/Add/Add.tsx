import React, { useEffect, useState, Dispatch, SetStateAction } from "react"

import { Box, Button, Heading, Layer, Text } from "grommet"
import { FormClose } from "grommet-icons"

import Form, { ContactFormData } from "../Form/Form"

import { useAddContact, ContactType, ProjectContact } from "./useAddContact"

export default function Add({
    projectId,
    setShow,
    onAdd,
}: {
    projectId: string
    setShow: Dispatch<SetStateAction<boolean>>
    onAdd: (contact: ProjectContact) => void
}) {
    const [value, setValue] = useState({} as ContactFormData)

    const [add, { data, loading, error }] = useAddContact()

    const handleSubmit = (event: React.FormEvent) => {
        const { fullName } = value

        const details = []

        if (value.phone) {
            details.push({ type: ContactType.Phone, value: value.phone })
        }

        if (value.email) {
            details.push({ type: ContactType.Email, value: value.email })
        }

        if (value.instagram) {
            details.push({ type: ContactType.Instagram, value: value.instagram })
        }

        add(projectId, { fullName, details })

        event.preventDefault()
    }

    useEffect(() => {
        switch (data?.addContact.__typename) {
            case "ContactAdded": {
                const {
                    addContact: { contact },
                } = data
                onAdd(contact)
                setShow(false)
            }
        }
    }, [data, setShow, onAdd]) // todo

    return (
        <Layer onClickOutside={() => setShow(false)} onEsc={() => setShow(false)}>
            <Box pad="medium" gap="medium" width="large">
                <Box direction="row" justify="between" align="center">
                    <Heading level={3} margin="none">
                        Добавить контакт
                    </Heading>
                    <Button icon={<FormClose />} onClick={() => setShow(false)} />
                </Box>
                <Form
                    contact={value}
                    onSet={setValue}
                    onSubmit={handleSubmit}
                    submit={
                        <Box direction="row" justify="between" margin={{ top: "large" }}>
                            <Button
                                type="submit"
                                primary
                                label={loading ? "Сохранение..." : "Сохранить"}
                                disabled={loading}
                            />
                            <Box>
                                <Text color="status-critical">
                                    <ErrorMessage res={data?.addContact} />
                                </Text>
                            </Box>
                            {error && (
                                <Box>
                                    <Text color="status-critical">{error.message}</Text>
                                </Box>
                            )}
                        </Box>
                    }
                />
            </Box>
        </Layer>
    )
}

/* eslint-disable  @typescript-eslint/no-explicit-any */
function ErrorMessage({
    res,
}: {
    res: { __typename: "Forbidden"; message: string } | { __typename: "ServerError"; message: string } | any
}) {
    switch (res?.__typename) {
        case "Forbidden":
            return <>Доступ запрещен</>
        case "ServerError":
            return <>Ошибка сервера</>
    }

    return null
}
