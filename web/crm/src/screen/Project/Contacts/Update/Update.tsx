import React, { useEffect, useState } from "react"

import { Box, Button, Heading, Layer, Text } from "grommet"
import { FormClose } from "grommet-icons"

import { Form, ContactFormData } from "../Form/Form"

import { useUpdateContact, Contact, ContactType } from "./useUpdateContact"

export default function Update(
    { contact, hide, onUpdate }:
    {
        contact: Contact,
        hide: () => void,
        onUpdate: (contact: Contact) => void
    }) {

    const [ value, setValue ] = useState({
        fullName: contact.fullName,
        phone: contact.details.filter(val => val.type === ContactType.Phone)[0]?.value,
        email: contact.details.filter(val => val.type === ContactType.Email)[0]?.value,
        instagram: contact.details.filter(val => val.type === ContactType.Instagram)[0]?.value
    } as ContactFormData)

    const [ update, { data, loading, error } ] = useUpdateContact()

    const handleSubmit = (event: React.FormEvent) => {
        const { fullName } = value

        const details = []

        if (value.phone) {
            details.push({type: ContactType.Phone, value: value.phone})
        }

        if (value.email) {
            details.push({type: ContactType.Email, value: value.email})
        }

        if (value.instagram) {
            details.push({type: ContactType.Instagram, value: value.instagram})
        }

        update(contact.id, { fullName, details })

        event.preventDefault()
    }

    useEffect(() => {
        switch (data?.updateContact.__typename) {
            case "ContactUpdated":
                hide()
                onUpdate(data.updateContact.contact)
        }
    }, [ data, hide, onUpdate ])

    return (
        <Layer
            onClickOutside={hide}
            onEsc={hide}
        >
                {error && <Box><Text>{error.message}</Text></Box>}

                <Box pad="medium" gap="medium" width={{min: "500px"}}>
                    <Box direction="row"justify="between"align="center">
                        <Heading level={3} margin="none">Изменить контакт</Heading>
                        <Button icon={ <FormClose/> } onClick={hide}/>
                    </Box>
                    <Form
                        contact={value}
                        onSet={setValue}
                        onSubmit={handleSubmit}
                        submit={
                            <Box direction="row" justify="between" margin={{top: "large"}}>
                                <Button
                                    type="submit"
                                    primary
                                    label={loading ? 'Сохранение...' : 'Сохранить' }
                                    disabled={loading}
                                />
                                <Box><Text color="status-critical"><ErrorMessage res={data?.updateContact}/></Text></Box>
                            </Box>
                        }
                    />
                </Box>
        </Layer>
    )
}

/* eslint-disable  @typescript-eslint/no-explicit-any */
function ErrorMessage({res}: { res: { __typename: "Forbidden", message: string } | { __typename: "ServerError", message: string } | any}) {
    switch (res?.__typename) {
        case "Forbidden":
            return <>Доступ запрещен</>
        case "ServerError":
            return <>Ошибка сервера</>
    }

    return null
}
