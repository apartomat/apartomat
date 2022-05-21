import React, { useState } from "react"

import { Box, Heading, Text, Button } from "grommet"
import { Add } from "grommet-icons"

import { useAddContact, ContactType, Contact } from "../useAddContact"
import { useProject, Project  as ProjectType, ProjectFiles, ProjectContacts, ProjectHouses, HouseRooms } from "../useProject"

import AddContact from "../AddContact/AddContact"
import ContactCard from "../ContactCard/ContactCard"
import UpdateContact from "../UpdateContact/UpdateContact"

export default function Contacts({ contacts, projectId, notify }: { contacts: ProjectContacts, projectId: string, notify: (val: { message: string }) => void }) {
    const [showAddContact, setShowAddContact] = useState(false)

    const [updateContact, setUpdateContact] = useState<Contact | undefined>(undefined)

    const [justAdded, setJustAdded] = useState([] as Contact[])

    const [justDeleted, setJustDeleted] = useState([] as string[])

    switch (contacts.list.__typename) {
        case "ProjectContactsList":

            const list = [...contacts.list.items, ...justAdded].filter(contact => !justDeleted.includes(contact.id))

            return (
                <>
                    <Box margin={{top: "small"}}>
                        <Heading level={4} margin={{ bottom: "xsmall"}}>
                            {list.length === 1 ? "Заказчик" : "Заказчики"}
                        </Heading>
                        <Box direction="row" gap="small" wrap>
                            {[...list.map((contact) => {
                                return (
                                    <ContactCard
                                        key={contact.id}
                                        contact={contact}
                                        onDelete={(contact: Contact) => {
                                            setJustDeleted([...justDeleted, contact.id])
                                            notify({ message: "Контакт удален"})
                                        }}
                                        onClickUpdate={(contact: Contact) => {
                                            setUpdateContact(contact)
                                        }}
                                    />
                                )
                            }), <Button key="" icon={<Add/>} label="Добавить" onClick={() => setShowAddContact(true) }/>]}
                        </Box>
                    </Box>

                    {showAddContact ?
                        <AddContact
                            projectId={projectId}
                            setShow={setShowAddContact}
                            onAdd={(contact: Contact) => {
                                setJustAdded([...justAdded, contact])
                                notify({ message: "Контакт добавлен"})
                            }}
                        /> : null}

                    {updateContact ?
                        <UpdateContact
                            contact={updateContact}
                            hide={() => { setUpdateContact(undefined) }}
                            onUpdate={(contact: Contact) => {
                                notify({ message: "Контакт сохранен"})
                            }}
                        /> : null}
                </>
            )
        default:
            return (
                <Box margin={{top: "small"}}>
                    <Text>n/a</Text>
                </Box>
            )
    }
}