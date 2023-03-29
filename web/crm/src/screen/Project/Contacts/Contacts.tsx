import React, { useState } from "react"

import { Box, BoxExtendedProps, Text, Button } from "grommet"
import { Add } from "grommet-icons"

import { ProjectContact } from "./Add/useAddContact"
import { ProjectContacts } from "../useProject"

import Contact from "./Contact/Contact"
import AddContact from "./Add/Add"
import UpdateContact from "./Update/Update"

export default function Contacts({
    contacts,
    projectId,
    onAdd,
    onUpdate,
    onDelete,
    notify,
    ...boxProps
}: {
    contacts: ProjectContacts,
    projectId: string,
    notify: (val: { message: string }) => void,
    onAdd?: (house: ProjectContact) => void,
    onUpdate?: (house: ProjectContact) => void,
    onDelete?: (house: ProjectContact) => void
} & BoxExtendedProps) {
    const [showAddContact, setShowAddContact] = useState(false)

    const [updateContact, setUpdateContact] = useState<ProjectContact | undefined>(undefined)

    const [justAdded, setJustAdded] = useState([] as ProjectContact[])

    const [justDeleted, setJustDeleted] = useState([] as string[])

    switch (contacts.list.__typename) {
        case "ProjectContactsList":
            const list = [...contacts.list.items, ...justAdded].filter(contact => !justDeleted.includes(contact.id))

            return (
                <Box {...boxProps}>
                    <Box direction="row" wrap>
                        {[...list.map((contact) => {
                            return (
                                <Contact
                                    key={contact.id}
                                    contact={contact}
                                    onDelete={(contact: ProjectContact) => {
                                        setJustDeleted([...justDeleted, contact.id])
                                        notify({ message: "Контакт удален"})
                                    }}
                                    onClickUpdate={(contact: ProjectContact) => {
                                        setUpdateContact(contact)
                                    }}
                                    width={{min:"xsmall"}}
                                    overflow="hidden"
                                    margin={{right: "xsmall", bottom: "small"}}
                                />
                            )
                        }), <Button key="" icon={<Add/>} label="Добавить" onClick={() => setShowAddContact(true) } margin={{bottom: "small"}}/>]}
                    </Box>

                    {showAddContact ?
                        <AddContact
                            projectId={projectId}
                            setShow={setShowAddContact}
                            onAdd={(contact: ProjectContact) => {
                                setJustAdded([...justAdded, contact])
                                notify({ message: "Контакт добавлен"})
                            }}
                        /> : null}

                    {updateContact ?
                        <UpdateContact
                            contact={updateContact}
                            hide={() => { setUpdateContact(undefined) }}
                            onUpdate={(contact: ProjectContact) => {
                                notify({ message: "Контакт сохранен"})
                            }}
                        /> : null}
                </Box>
            )
        default:
            return (
                <Box margin={{top: "small"}}>
                    <Text>n/a</Text>
                </Box>
            )
    }
}