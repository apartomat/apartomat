import React, { useState, useEffect, useRef } from "react"

import { Box, Button, Drop, Card, CardHeader, CardBody, CardFooter } from "grommet"
import { Instagram, Trash } from "grommet-icons"

import useDeleteContact, { Contact, ContactType } from "../useDeleteContact"

export default function ContactCard(
    { contact, onDelete, onClickUpdate }:
    { contact: Contact , onDelete: (contact: Contact) => void, onClickUpdate: (contact: Contact) => void }
) {
    const ref = useRef(null)

    const [showCard, setShowCard] = useState(false)

    const [showDeleteConfirm, setShowDeleteConfirm] = useState(false)

    const [deleteContact, { data } ] = useDeleteContact()

    const handleDelete = () => {
        setShowDeleteConfirm(true)
    }

    const handleDeleteConfirm = () => {
        deleteContact(contact.id)
    }

    const handleDeleteCancel = () => {
        setShowDeleteConfirm(false)
        setShowCard(false)
    }

    useEffect(() => {
        switch (data?.deleteContact.__typename) {
            case "ContactDeleted":
                setShowCard(false)
                onDelete(contact)
        }
    }, [ data, contact, onDelete ])

    return (
        <Box>
            <Button
                key={contact.id}
                ref={ref}
                primary
                color="light-2"
                label={contact.fullName}
                style={{whiteSpace: "nowrap"}}
                onClick={() => setShowCard(!showCard) }
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
                        <CardHeader pad={{horizontal: "medium", top: "medium"}} style={{fontWeight: "bold"}}>{contact.fullName}</CardHeader>
                        <CardBody pad="medium">
                            {contact.details.filter(c => ![ContactType.Instagram].includes(c.type)).map(c => {
                                return <Box pad={{vertical: "small"}}>{c.value}</Box>
                            })}
                            {contact.details.filter(c => [ContactType.Instagram].includes(c.type)).length > 0
                                ? <Box pad={{vertical: "small"}} direction="row">
                                {contact.details.filter(c => [ContactType.Instagram].includes(c.type)).map(c => {
                                    switch (c.type) {
                                        case ContactType.Instagram:
                                            return <Button icon={<Instagram color="primary"/>} plain href={c.value}/>
                                    }

                                    return null
                                })}
                                </Box>
                                : null
                            }
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
                                    onClickUpdate(contact)
                                }}/>
                            }

                        </CardFooter>
                    </Card>
                </Drop>
            }
        </Box>
    )
}